package geom

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/texture"
)

func Collect(root *scene.Node) (amount int, collected *scene.Node) {
	collected = scene.New(root.Name() + "-collected")
	combined := new(Mesh)
	Add(collected, combined)

	addIndexed := func(index uint32, mesh *Mesh, mat *math.Mat4) {
		v := mesh.Vertices[index]
		vect := math.Vector3(v.X, v.Y, v.Z)
		vect = vect.TransformMat4(mat)
		combined.Vertices = append(combined.Vertices, Vertex{vect.X, vect.Y, vect.Z})

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
					vect := math.Vector3(v.X, v.Y, v.Z)
					vect = vect.TransformMat4(mat)
					combined.Vertices = append(combined.Vertices, Vertex{vect.X, vect.Y, vect.Z})
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
