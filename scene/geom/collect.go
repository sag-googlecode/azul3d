package geom

import (
	"azul3d.org/v1/math"
	"azul3d.org/v1/scene"
	"azul3d.org/v1/scene/texture"
)

// Collect 'collects' the specified node and all children nodes below it, by
// taking the meshes of all of them, putting the vertices, vertice colors, etc,
// into a single mesh and single node.
//
// Returned is the amount of nodes collected, and the node which represents all
// of the nodes collected.
func Collect(root *scene.Node) (amount int, collected *scene.Node) {
	collected = scene.New(root.Name() + "-collected")
	combined := new(Mesh)
	Add(collected, combined)

	addIndexed := func(index uint32, mesh *Mesh, mat math.Mat4) {
		v := mesh.Vertices[index]
		vect := math.Vec3{float64(v.X), float64(v.Y), float64(v.Z)}
		vect = vect.TransformMat4(mat)
		vtx := Vertex{float32(vect.X), float32(vect.Y), float32(vect.Z)}
		combined.Vertices = append(combined.Vertices, vtx)

		if len(mesh.Normals) > 0 {
			combined.Normals = append(combined.Normals, mesh.Normals[index])
		}

		if len(mesh.Tangents) > 0 {
			combined.Tangents = append(combined.Tangents, mesh.Tangents[index])
		}

		if len(mesh.Bitangents) > 0 {
			combined.Bitangents = append(combined.Bitangents, mesh.Bitangents[index])
		}

		if len(mesh.Colors) > 0 {
			combined.Colors = append(combined.Colors, mesh.Colors[index])
		}

		if len(mesh.BoneWeights) > 0 {
			combined.BoneWeights = append(combined.BoneWeights, mesh.BoneWeights[index])
		}

		for tci, texCoords := range mesh.TextureCoords {
			if len(combined.TextureCoords) > tci {
				combined.TextureCoords[tci] = append(combined.TextureCoords[tci], texCoords[index])
			} else {
				combined.TextureCoords = append(combined.TextureCoords, []texture.Coord{
					texCoords[index],
				})
			}
		}
	}

	root.Traverse(func(index int, n *scene.Node) bool {
		amount++
		for layer, tex := range texture.Textures(n) {
			texture.Set(collected, layer, tex)
		}

		for _, mesh := range Meshes(n) {
			mesh.RLock()

			mat := n.RelativeTransform(root).Mat4()

			if len(mesh.Indices) > 0 {
				for _, index := range mesh.Indices {
					addIndexed(index, mesh, mat)
				}

			} else {
				for _, v := range mesh.Vertices {
					vect := math.Vec3{float64(v.X), float64(v.Y), float64(v.Z)}
					vect = vect.TransformMat4(mat)
					vtx := Vertex{float32(vect.X), float32(vect.Y), float32(vect.Z)}
					combined.Vertices = append(combined.Vertices, vtx)
				}

				for _, v := range mesh.Normals {
					combined.Normals = append(combined.Normals, v)
				}

				for _, v := range mesh.Tangents {
					combined.Tangents = append(combined.Tangents, v)
				}

				for _, v := range mesh.Bitangents {
					combined.Bitangents = append(combined.Bitangents, v)
				}

				for _, v := range mesh.Colors {
					combined.Colors = append(combined.Colors, v)
				}

				for _, v := range mesh.BoneWeights {
					combined.BoneWeights = append(combined.BoneWeights, v)
				}

				for index, texCoords := range mesh.TextureCoords {
					if len(combined.TextureCoords) > index {
						for _, v := range texCoords {
							combined.TextureCoords[index] = append(combined.TextureCoords[index], v)
						}
					} else {
						combined.TextureCoords = append(combined.TextureCoords, texCoords)
					}
				}
			}

			mesh.RUnlock()
		}
		return true
	})

	return amount, collected
}
