// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"azul3d.org/v1/math"
	"azul3d.org/v1/scene"
	"azul3d.org/v1/scene/bucket"
	"azul3d.org/v1/scene/camera"
	"azul3d.org/v1/scene/geom"
	"azul3d.org/v1/scene/shader"
	"azul3d.org/v1/scene/texture"
	"azul3d.org/v1/scene/transparency"
	"sort"
)

type camNodeWrapper struct {
	n *scene.Node
}

func (c camNodeWrapper) Node() *scene.Node {
	return c.n
}

type renderNode struct {
	node          *scene.Node
	camera        camNodeWrapper
	bucket        *bucket.Type
	meshes        []*geom.Mesh
	transparency  transparency.ModeType
	textures      []texture.Pair
	inputTextures []int32
	shader        *shader.Shader
}

// implements bucket.Node interface
func (r renderNode) Node() *scene.Node {
	return r.node
}

type sortedNodes struct {
	slice    []renderNode
	released bool
}

func (s *sortedNodes) Len() int      { return len(s.slice) }
func (s *sortedNodes) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }
func (s *sortedNodes) Less(i, j int) bool {
	a := s.slice[i]
	b := s.slice[j]

	// If the nodes are not inside the same bucket, then we need to sort based
	// off the bucket's sort value.
	if a.bucket != b.bucket {
		return a.bucket.Sort() < b.bucket.Sort()
	}

	// Lastly, we let the active bucket of the two nodes sort however it wishes.
	return a.bucket.Less(a.camera, a, b)
}

type renderRegion struct {
	*camera.Region
	nodes *sortedNodes
}

type sortedRegions []renderRegion

func (s sortedRegions) Len() int      { return len(s) }
func (s sortedRegions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortedRegions) Less(i, j int) bool {
	return s[i].Sort() < s[j].Sort()
}

func (r *Renderer) releaseRegionBuf(buf sortedRegions) {
	// Release each node buffer in-use by the region buffer.
	for _, rr := range buf {
		r.releaseNodeBuf(rr.nodes)
	}

	// Release the actual region buffer.
	select {
	case r.regionBuffers <- buf[:0]:
		// We were able to put it away for re-use later on.
	default:
		// We couldn't put it away for re-use, we must already have the max
		// amount of buffers laying around for re-use.
	}
}

func (r *Renderer) claimRegionBuf() sortedRegions {
	// Claim an existing buffer for re-use or return nil to suggest one should
	// be allocated.
	select {
	case buf := <-r.regionBuffers:
		// We are able to re-use this buffer.
		return buf
	default:
		// We don't have a buffer for re-use at this time.
		return nil
	}
}

func (r *Renderer) releaseNodeBuf(buf *sortedNodes) {
	if buf.released {
		// Buffer is already released and awaiting re-use.
		return
	}
	buf.released = true

	for _, rn := range buf.slice {
		r.releaseInputTexturesBuf(rn.inputTextures)
	}

	// Reslice the buffer.
	buf.slice = buf.slice[:0]

	select {
	case r.nodeBuffers <- buf:
		// We were able to put it away for re-use later on.
	default:
		// We couldn't put it away for re-use, we must already have the max
		// amount of buffers laying around for re-use.
	}
}

func (r *Renderer) claimNodeBuf() *sortedNodes {
	// Claim an existing buffer for re-use or return nil to suggest one should
	// be allocated.
	select {
	case buf := <-r.nodeBuffers:
		// We are able to re-use this buffer.
		buf.released = false
		return buf
	default:
		// We don't have a buffer for re-use at this time.
		return new(sortedNodes)
	}
}

func (r *Renderer) claimInputTexturesBuf() []int32 {
	select {
	case buf := <-r.inputTexturesBuffers:
		return buf
	default:
		return nil
	}
}

func (r *Renderer) releaseInputTexturesBuf(buf []int32) {
	select {
	case r.inputTexturesBuffers <- buf[:0]:
		// We were able to put it away for re-use later on.
	default:
		// We couldn't put it away for re-use, we must already have the max
		// amount of buffers laying around for re-use.
	}
}

func (r *Renderer) collectNodes(root, camNode *scene.Node, lensProjection, camInverse math.Mat4) (nodeBuf *sortedNodes) {
	// Find an existing node buffer for re-use if we have one.
	nodeBuf = r.claimNodeBuf()

	// Check if the camera has a scene node and it is visible.
	camScene := camera.Scene(camNode)
	if camScene == nil || (camScene.Hidden() && !camScene.ShownThrough()) {
		return
	}

	// Traverse the camera's scene and store render nodes inside the node
	// buffer.
	camScene.Traverse(func(index int, n *scene.Node) bool {
		if n.Hidden() && !n.ShownThrough() {
			// Node is not visible, continue traversal.
			return true
		}

		// Acquire mesh list, remove ones that are hidden, contain no vertices.
		meshes := geom.Meshes(n)
		for i := 0; i < len(meshes); i++ {
			m := meshes[i]
			m.RLock()
			if m.Hidden || len(m.Vertices) == 0 {
				// Mesh is hidden or has no vertices, remove it from the slice.
				meshes = append(meshes[:i], meshes[i+1:]...)
			}
			m.RUnlock()
		}

		if len(meshes) == 0 {
			// Node has no visible meshes or ones with vertices, continue
			// traversal.
			return true
		}

		rn := renderNode{
			node:         n,
			camera:       camNodeWrapper{camNode},
			bucket:       bucket.Active(n),
			meshes:       meshes,
			transparency: transparency.ActiveMode(n),
			textures:     texture.Sorted(n),
			shader:       shader.Active(n),
		}

		// Create a input texture ID' slice large enough to hold them all.
		rn.inputTextures = r.claimInputTexturesBuf()
		if rn.inputTextures == nil {
			rn.inputTextures = make([]int32, len(rn.textures))[:0]
		}

		// Add a shader input for the binary transparency mode.
		switch rn.transparency {
		case transparency.Multisample:
			if !r.haveMSTransparency {
				// Fallback to binary transparency.
				shader.SetInput(rn.node, "BinaryTransparency", true)
			} else {
				shader.SetInput(rn.node, "BinaryTransparency", false)
			}

		case transparency.Binary:
			shader.SetInput(rn.node, "BinaryTransparency", true)

		default:
			shader.SetInput(rn.node, "BinaryTransparency", false)
		}

		// Set default shader on node if it has none.
		if rn.shader == nil {
			rn.shader = defaultShader
		}

		// Calculate a world transformation matrix.
		worldMat := n.RelativeTransform(root).Mat4()

		// Apply camera transformation.
		modelView := worldMat.Mul(camInverse)

		// Apply the ZUpRight -> YUpRight coordinate system conversion.
		modelView = modelView.Mul(coordSysConversion)

		// Calculate a model view * projection (MVP) matrix.
		modelViewProjection := modelView.Mul(lensProjection)

		shader.SetInput(n, "Projection", shader.ConvertMat4(lensProjection))
		shader.SetInput(n, "ModelView", shader.ConvertMat4(modelView))
		shader.SetInput(n, "ModelViewProjection", shader.ConvertMat4(modelViewProjection))

		// Append the node to the buffer.
		nodeBuf.slice = append(nodeBuf.slice, rn)

		// Continue traversal
		return true
	})
	return
}

func (r *Renderer) collect(root *scene.Node) sortedRegions {
	// Find an existing region buffer for re-use if we have one.
	regionBuf := r.claimRegionBuf()

	// We start by traversing to find each camera who is visible, and has at
	// least one camera region.
	root.Traverse(func(i int, n *scene.Node) bool {
		if n.Hidden() && !n.ShownThrough() {
			// Node is hidden, continue traversal.
			return true
		}

		if !camera.Is(n) {
			// Node is not a camera node, continue traversal.
			return true
		}

		regions := camera.Regions(n)
		if len(regions) == 0 {
			// Camera does not have any regions to render, continue traversal.
			return true
		}

		// Find lens projection
		lensProjection := camera.Lens(n).Projection()
		camWorldTransform := n.RelativeTransform(root)
		camInverse, _ := camWorldTransform.Mat4().Inverse()

		// Collect the nodes in the camera's scene and sort them.
		nodesBuf := r.collectNodes(root, n, lensProjection, camInverse)
		sort.Sort(nodesBuf)

		// For every region the camera has, store it in the list of regions.
		for _, r := range regions {
			regionBuf = append(regionBuf, renderRegion{r, nodesBuf})
		}

		// Continue traversal.
		return true
	})

	// Sort the regions.
	sort.Sort(regionBuf)

	return regionBuf
}
