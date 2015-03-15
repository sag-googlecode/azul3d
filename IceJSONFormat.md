# WIP #
This page is a work-in-progress. It may be completely inaccurate.

  * Determine how Light should work.
  * Determine how Camera should work with FOV/Ortho lens.

# Ice JSON Format #

This page serves to document the Ice JSON file format used by Azul3D for loading and storing scenes, models, etc on disk. A better overview of Ice can be found in it's package documentation azul3d.org/v1/ice, this page merely serves to document the JSON format used by Ice.

# Names #
All names are case-sensitive and start with an uppercase letter.

# Scene #
This is the literal initial JSON object {} in the file (i.e. top-level).
| [Name](#Names.md)    | Type                             | Optional? | Description |
|:---------------------|:---------------------------------|:----------|:------------|
| `Version`    | Number                           | No        | Describes the file format version. Only 1 is a valid version at this time. |
| `Props`      | Object                           | Yes       | User-defined properties for the scene. |
| `Cameras`    | Object of [Camera](#Camera.md)       | Yes       | Cameras by their unique user-defined names. |
| `Lights`     | Object of [Light](#Light.md)         | Yes       | Lights by their unique user-defined names. |
| `Meshes`     | Object of [Mesh](#Mesh.md)           | Yes       | Meshes by their unique user-defined names. |
| `Textures`   | Object of [Texture](#Texture.md)     | Yes       | Textures by their unique user-defined names. |
| `Objects`    | Object of [Object](#Object.md)       | Yes       | Objects by their unique user-defined names. |
| `Transforms` | Object of [Transform](#Transform.md) | Yes       | Transforms by their unique user-defined names. |
| `Shaders`    | Object of [Shader](#Shader.md)       | Yes       | Shaders by their unique user-defined names. |

# Camera #
A single camera within a scene.
| [Name](#Names.md) | Type   | Optional? | Description |
|:------------------|:-------|:----------|:------------|
| `Props`   | Object | Yes       | User-defined properties for the camera. |

# Light #
A single light source within a scene.
| [Name](#Names.md) | Type   | Optional? | Description |
|:------------------|:-------|:----------|:------------|
| `Props`   | Object | Yes       | User-defined properties for the light. |

# Mesh #
A single mesh with vertices, etc for use by an object.
| [Name](#Names.md)        | Type | Optional? | Description |
|:-------------------------|:-----|:----------|:------------|
| `KeepDataOnLoad` | Bool | Yes       | Whether or not the data slices of the mesh should be kept after the mesh is loaded by the graphics hardware. |
| `Dynamic`        | Bool | Yes       | Hint to the renderer on how the object may be used/manipulated. Does not affect actual usage, it just functions as a performance hint. |
| `AABB`           | [AABB](#AABB.md) | Yes | The axis-aligned bounding box of the mesh. |
| `Indices` | Array of whole numbers | Yes       | Indices into each of the other mesh data arrays. If this array is present the mesh is considered to be indexed and all of the indices in this array must be valid indices into each other data slice of the mesh. |
| `Vertices` | Array of array of three numbers | No       | Vertices which make up the 3D model in `Vertices[vertex][x/y/z]` format. |
| `Colors` | Array of array of four numbers | Yes       | Per-vertex colors of the mesh in `Colors[vertex][r/g/b/a]` format. |
| `TexCoords` | Array of array of array of two numbers | Yes       | List of texture coordinate sets with u/v pairs per-vertex in `TexCoords[set][vertex][u/v]` format. |

# Texture #
A single texture which can be applied to meshes for rendering a graphics object.
| [Name](#Names.md)        | Type   | Optional? | Description |
|:-------------------------|:-------|:----------|:------------|
| `KeepDataOnLoad` | Bool   | Yes       | Whether or not the source image of the texture should be kept after the texture is loaded by the graphics hardware. |
| `Source`         | String | No        | The filepath string to the source image for this texture. |
| `Format`         | String | Yes       | The format to use for storing this texture on the GPU, which may result in lossy conversions (e.g. RGB would lose the alpha channel, etc). Must be one of the [TexFormat](#TexFormat.md) strings. |
| `WrapU`          | String | Yes       | The U wrap mode of this texture. Must be one of the [TexWrap](#TexWrap.md) strings. |
| `WrapV`          | String | Yes       | The V wrap mode of this texture. Must be one of the [TexWrap](#TexWrap.md) strings. |
| `BorderColor`    | Array of four numbers | Yes        | The color of the border when a wrap mode is set to `"BorderColor"`. |
| `MinFilter`      | String | Yes       | The texture filter used for minification of the texture. Must be a valid [TexFilter](#TexFilter.md) string. |
| `MagFilter`      | String | Yes       | The texture filter used for magnification of the texture. Must be a valid [TexFilter](#TexFilter.md) string. |

# Object #
A single object/model within a scene.
| [Name](#Names.md)       | Type            | Optional? | Description |
|:------------------------|:----------------|:----------|:------------|
| `Props`         | Object          | Yes       | User-defined properties for the object. |
| `OcclusionTest` | Bool            | Yes       | Whether or not the object should have occlusion testing enabled. |
| `State`         | [State](#State.md)  | Yes       | The graphics state of which the object should render with. |
| `Meshes`        | Array of String | No        | The names of the meshes associated with this object. Those names must also be present in the [Scene](#Scene.md). |
| `Textures`      | Array of String | Yes       | The names of the textures associated with this object. Those names must also be present in the [Scene](#Scene.md). |
| `Transform`     | String          | Yes       | The name of the transform of this object. The name must also be present in the [Scene](#Scene.md) |
| `Shader`        | String          | Yes       | The name of the shader of this object. The name must also be present in the [Scene](#Scene.md) |

# Transform #
A single transformation to space, used to transform objects in 3D space. If Quaternion rotation is present it is always used over Euler rotation. If an arbitrary 4x4 matrix is present it is always used over the other components.
| [Name](#Names.md)   | Type                   | Optional? | Description |
|:--------------------|:-----------------------|:----------|:------------|
| `Pos`       | Array of three Number  | Yes       | The position of the transform relative to it's parent. |
| `Rot`       | Array of three Number  | Yes       | The Euler rotation of the transform relative to it's parent around x, y, and z axis. |
| `Quat`      | Array of four Number   | Yes       | The Quaternion rotation of the transform relative to it's parent. |
| `Scale`     | Array of three Number  | Yes       | The scale of the transform relative to it's parent along the x, y, and z axis. |
| `Shear`     | Array of three Number  | Yes       | The shear of the transform relative to it's parent along the x, y, and z axis. |
| `Mat4`      | Array of 16 Number     | Yes       | An arbitrary 4x4 transformation matrix to define the transform relative to it's parent (disregarding all other members). |
| `Parent`    | String                 | Yes       | The unique name of the parent transform of this one. The name must also be present in the [Scene](#Scene.md). |

# Shader #
A single shader, used to shade objects when rendering them. Although all components are technically optional, for a shader to be useful in any way it must have both the GLSL vertex and fragment sources defined  either as literal inline sources (GLSLVert, GLSLFrag) or have it's GLSLSources filepaths defined with at least one .vert and one .frag file.
| [Name](#Names.md)         | Type                 | Optional? | Description |
|:--------------------------|:---------------------|:----------|:------------|
| `KeepDataOnLoad`  | Bool                 | Yes       | Whether or not the data source of the shader should be kept in memory after it being loaded by the graphics hardware. |
| `GLSLSources`     | Array of String      | Yes       | A list of filepath's to the GLSL shader sources, of types who are determine by the extensions .vert and .frag. |
| `GLSLVert`        | Array of String      | Yes       | A list of literal GLSL vertex shader source code lines. |
| `GLSLFrag`        | Array of String      | Yes       | A list of literal GLSL fragment shader source code lines. |
| `Inputs`          | Object of JSON types | Yes       | A map of inputs to be fed into the shader program while rendering. Usefulness of the type depends on the renderer implementation (e.g. float64 probably always works, map probably won't, etc). |

# State #
A single graphics state used to render a graphics object.
| [Name](#Names.md)       | Type                     | Optional? | Description |
|:------------------------|:-------------------------|:----------|:------------|
| `AlphaMode`     | String                       | Yes       | Describes the alpha transparency mode which determines how transparent parts of the object are to be rendered. If present it must have a value of `"BlendedAlpha"`, `"BinaryAlpha"`, or `"AlphaToCoverage"`. |
| `Blend`         | [BlendState](#BlendState.md)     | Yes       | Represents how blending between existing (source) and new (destination) pixels in the color buffer occurs when `AlphaMode == "BlendedAlpha"` |
| `WriteRed`      | Bool                         | Yes      | Determines if red values should be written to the color buffer or not when rendering the object. |
| `WriteGreen`    | Bool                         | Yes      | Determines if green values should be written to the color buffer or not when rendering the object. |
| `WriteBlue`     | Bool                         | Yes      | Determines if blue values should be written to the color buffer or not when rendering the object. |
| `WriteAlpha`    | Bool                         | Yes      | Determines if alpha values should be written to the color buffer or not when rendering the object. |
| `Dithering`     | Bool                         | Yes      | Determines if dithering should be used when rendering the object. |
| `DepthTest`     | Bool                         | Yes      | Determines if depth-testing should be used when rendering the object. |
| `DepthWrite`    | Bool                         | Yes      | Determines if depth-writing should be used when rendering the object. |
| `DepthCmp`      | String                       | Yes      | The comparison operator to use for depth testing against existing pixels in the depth buffer. Must be one of the [Cmp](#Cmp.md) strings. |
| `StencilTest`   | Bool                         | Yes      | Determines if stencil-testing should be enabled when rendering the object. |
| `FaceCulling`   | String                       | Yes      | Whether or not (and how) face culling should occur when rendering the object. Must be `"BackFaceCulling"`, `"FrontFaceCulling"`, or `"NoFaceCulling"`     |
| `StencilFront`  | [StencilState](#StencilState.md) | Yes       | The stencil state for front facing pixels, when `StencilTest == true`. |
| `StencilBack`  | [StencilState](#StencilState.md) | Yes       | The stencil state for back facing pixels, when `StencilTest == true`. |

# `BlendState` #
A single blend state used to render a graphics object.
| [Name](#Names.md)  | Type                     | Optional? | Description |
|:-------------------|:-------------------------|:----------|:------------|
| `Color`    | Array of four numbers    | Yes       | Values in the range of 0-1 describing the constant RGBA color to use for blending. |
| `SrcRGB`   | String                   | Yes       | The blend operand to use for the source RGB components. Must be one of the [BlendOp](#BlendOp.md) Strings. |
| `DstRGB`   | String                   | Yes       | The blend operand to use for the destination RGB components. Must be one of the [BlendOp](#BlendOp.md) Strings. |
| `SrcAlpha` | String                   | Yes       | The blend operand to use for the source alpha component. Must be one of the [BlendOp](#BlendOp.md) Strings. |
| `DstAlpha` | String                   | Yes       | The blend operand to use for the destination alpha component. Must be one of the [BlendOp](#BlendOp.md) Strings. |
| `RGBEq`    | String                   | Yes       | The blending equation to use for the RGB components. Must be one of the [BlendEq](#BlendEq.md) Strings. |
| `AlphaEq`  | String                   | Yes       | The blending equation to use for the ralpha component. Must be one of the [BlendEq](#BlendEq.md) Strings. |

# `StencilState` #
A single stencil state used to render front or back pixels of a graphics object.
| [Name](#Names.md)   | Type                   | Optional? | Description |
|:--------------------|:-----------------------|:----------|:------------|
| `WriteMask` | Number                 | Yes       | A bitmask that will be AND'd with each pixel to be written to the stencil buffer, e.g. 0xFFFF would allow writing to the full range of every pixel in the stencil buffer when rendering the object. |
| `ReadMask`  | Number                 | Yes       | A bitmask that will be AND'd with each pixel to be read/compared to the existing value in the stencil buffer, e.g. 0xFFFF would disable the use of the mask altogether. |
| `Reference` | Number                 | Yes       | The reference value that will be used to compare existing values in the stencil buffer against, e.g. `if Reference == 2 && Func == "GreaterOrEqual"`, then any value below 2 would not be affected. |
| `Fail`      | [StencilOp](#StencilOp.md) | Yes       | Fail specifies what stencil operation should occur when the stencil test fails. |
| `DepthFail` | [StencilOp](#StencilOp.md) | Yes       | `DepthFail` specifies what stencil operation should occur when the stencil test passes but the depth test fails. |
| `DepthPass` | [StencilOp](#StencilOp.md) | Yes       | `DepthPass` specifies what stencil operation should occur when the stencil test passes and the depth test passes. |
| `Cmp`       | [Cmp](#Cmp.md)             | Yes       | Cmp specifies the comparison operator to use when comparing stencil data with existing data in the stencil buffer. |

# AABB #
An axis-aligned bounding box definition which contains all points `Min <= P < Max`.
| [Name](#Names.md)   | Type                   | Optional? | Description |
|:--------------------|:-----------------------|:----------|:------------|
| `Min` | Number                 | No       | The lesser point of the box. |
| `Max` | Number                 | No       | The greater point of the box. |

# `BlendOp` #
A list of valid `BlendOp` strings for use in a [BlendState](#BlendState.md).
```
"Zero"
"One"
"SrcColor"
"OneMinusSrcColor"
"OneMinusSrcColor"
"DstColor"
"OneMinusDstColor"
"SrcAlpha"
"OneMinusSrcAlpha"
"DstAlpha"
"OneMinusDstAlpha"
"ConstantColor"
"OneMinusConstantColor"
"ConstantAlpha"
"OneMinusConstantAlpha"

Not applicable for use in BlendState["SrcRGB"]:
"SrcAlphaSaturate"
```

# `BlendEq` #
A list of valid `BlendEq` strings for use in a [BlendState](#BlendState.md).
```
Add represents a blending equation where the src and dst colors are added
to eachother to produce the result.
"Add"

Sub represents a blending equation where the src and dst colors are
subtracted from eachother to produce the result.
"Sub"

ReverseSub represents a blending equation where the src and dst colors
are reverse-subtracted from eachother to produce the result.
"ReverseSub"
```

# `StencilOp` #
A list of valid `StencilOp` operation strings for use in a [StencilState](#StencilState.md).
```
Keep keeps the existing stencil data.
"Keep"

Zero sets the stencil data to zero.
"Zero"

Replace replaces the existing stencil data with the stencil reference
"Replace"

Incr increments the stencil value by one and clamps the result.
"Incr"

IncrWrap increments the stencil value by 1 and wraps the result if necessary.
"IncrWrap"

Decr decrements the stencil value by one and clamps the result.
"Decr"

DecrWrap decrements the stencil value by 1 and wraps the result if necessary.
"DecrWrap"

Invert inverts the stencil data.
"Invert"
```

# `Cmp` #
A list of valid `Cmp` operator strings.
```
"Always"
"Never"
"Less"
"LessOrEqual"
"Greater"
"GreaterOrEqual"
"Equal"
"NotEqual"
```

# `TexFilter` #
A list of valid `TexFilter` filter strings.
```
Samples the nearest pixel.
"Nearest"

Samples the four closest pixels and linearly interpolates them.
"Linear"

Samples point from the closest mipmap. May not be used as a magnification
filter.
"NearestMipmapNearest"

Bilinear filter the pixel from the closest mipmap. May not be used as a
magnification filter.
"LinearMipmapNearest"

Samples the pixel from two closest mipmaps, and linearly blends. May not
be used as a magnification filter.
"NearestMipmapLinear"

(Trilinear filtering) Bilinearly filters the pixel from two mipmaps, and
linearly blends the result. May not be used as a magnification filter.
"LinearMipmapLinear"
```

# `TexWrap` #
A list of valid `TexWrap` wrap mode strings.
```
The extra area of the texture is repeated into infinity.
"Repeat"

The extra area of the texture is represented by stretching the edge
pixels out into infinity.
"Clamp"

The extra area of the texture is represented by the border color
specified on the texture object.
"BorderColor"

The extra area of the texture is represented by itself mirrored into
infinity.
"Mirror"
```

# `TexFormat` #
A list of valid `TexFormat` format strings.
```
RGBA is a standard 32-bit premultiplied alpha image format.
"RGBA"

RGB is a standard 24-bit RGB image format with no alpha component.
"RGB"

DXT1 is a DXT1 texture compression format in RGB form (i.e. fully
opaque) each 4x4 block of pixels take up 64-bits of data, as such when
compared to a standard 24-bit RGB format it provides a 6:1 compression
ratio.
"DXT1"

DXT1RGBA is a DXT1 texture compression format in RGBA form with 1 bit
reserved for alpha (i.e. fully transparent or fully opaque per-pixel
transparency).
"DXT1RGBA"

DXT3 is a RGBA texture compression format with four bits per pixel
reserved for alpha. Each 4x4 block of pixels take up 128-bits of data,
as such when compared to a standard 32-bit RGBA format it provides a 4:1
compression ratio. Color information stored in DXT3 is mostly the same
as DXT1.
"DXT3"

DXT5 is a RGBA format similar to DXT3 except it compresses the alpha
chunk in a similar manner to DXT1's color storage. It provides the same
4:1 compression ratio as DXT3.
"DXT5"
```