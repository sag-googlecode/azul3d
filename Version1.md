![http://i.imgur.com/fksUHXY.png](http://i.imgur.com/fksUHXY.png)

# IMPORTANT - README #

  * This page is out of date. Please visit http://azul3d.org

# Azul3D Version 1 #

  * [New graphics API](Version1#New_graphics_API.md)
  * [3D Audio API](Version1#3D_Audio_API.md)
  * [New Supported Platforms](Version1#New_Supported_Platforms.md)
  * [Removed Packages](Version1#Removed_Packages.md)
  * [Moved Packages](Version1#Moved_Packages.md)

## New graphics API ##
Azul3D v1 will include among other large changes a new graphics API (azul3d.org/v1/gfx) essentially _rewriting the core of the engine_. The new graphics API will:
  * Be a small wrapper on top of existing OpenGL, OpenGL ES, WebGL, etc API's.
  * Give developers control over draw order, blend modes, and other important features.
  * Serve as a important base for package developers who would like to extend the graphics API.

In compliment to the new graphics API is the new OpenGL 2.0+ renderer (azul3d.org/v1/gfx/gl2). It works with existing OpenGL windowing libraries, like Chippy, GLFW 3, QML, etc.

## 3D Audio API ##
A new 3D audio API is created (which replaces azul3d.org/v0/audio/aio) and is based on OpenAL.

## New Supported Platforms ##
With version 1 the following platforms become officially supported:
  * windows/386
  * linux/amd64
  * linux/386

## Removed Packages ##
Version one is all about creating a stable, easy to use game engine in Go. Part of this is ensuring we only keep around packages that are extremely important to users (not every package _has_ to be in Azul3D). The following packages are marked for removal:

| Package import path | Reason for removal |
|:--------------------|:-------------------|
| azul3d.org/v1/scene | Incompatible with new graphics API, may be rewritten in a later release. |
| azul3d.org/v1/scene/text | Incompatible with new graphics API, may be rewritten. |
| azul3d.org/v1/scene/sprite | Incompatible with new graphics API, may be rewritten. |
| azul3d.org/v1/tmx/tmxmesh | Incompatible with new graphics API, may be rewritten. |
| azul3d.org/v1/audio/aio | Replaced by new 3D Audio API. |
| azul3d.org/v1/embed | Other Go alternatives are available, doesn't belong here. |
| azul3d.org/v1/event | Package-managed event APIs are generally better, it's toxic. |
| azul3d.org/v1/scene/camera | Replaced by the new graphics API. |
| azul3d.org/v1/scene/color | Replaced by the new graphics API. |
| azul3d.org/v1/scene/geom | Replaced by the new graphics API. |
| azul3d.org/v1/scene/renderer | Replaced by the new graphics API. |
| azul3d.org/v1/scene/shader | Replaced by the new graphics API. |
| azul3d.org/v1/scene/texture | Replaced by the new graphics API. |
| azul3d.org/v1/scene/transparency | Replaced by the new graphics API. |
| azul3d.org/v1/scene/bucket | Mostly replaced by the new graphics API. |
| azul3d.org/v1/engine | Was just a thin setup wrapper around the scene packages. |
| azul3d.org/v1/scene/geom/procedural | A community-ran package would be better (also a bad name). |

## Moved Packages ##
The following packages are moved to a new location:
| Package import path | New import path | Reasoning |
|:--------------------|:----------------|:----------|
| azul3d.org/v1/chippy/keyboard | azul3d.org/v1/keyboard | Not specific to Chippy. |
| azul3d.org/v1/chippy/mouse | azul3d.org/v1/mouse | Not specific to Chippy. |
| azul3d.org/v1/math | azul3d.org/v1/gmath | Avoid annoying collision with Go's "math" package. |
| azul3d.org/v1/native/cp | gopkg.in/slimsag/cp.v1 | Not used by most users. |
| azul3d.org/v1/file2go | gopkg.in/slimsag/file2go.v1 | Not used by most users. |
| azul3d.org/v1/dstarlite | gopkg.in/slimsag/dstarlite.v1 | Not used by most users. |
| azul3d.org/v1/binpack | gopkg.in/slimsag/binpack.v1 | Not used by most users. May not be useful in future text package. |