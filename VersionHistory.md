![http://i.imgur.com/fksUHXY.png](http://i.imgur.com/fksUHXY.png)

# IMPORTANT - README #

  * This page is out of date. Please visit http://azul3d.org

# Version History #

  * [Versioning Scheme](VersionHistory#Versioning_Scheme.md)
  * [Version 1 (latest)](VersionHistory#Version_1.md)

## Versioning Scheme ##
Azul3D uses a versioning scheme specific to the project itself. It should be noted that the versioning scheme is different from, that of say www.gopkg.in which some of the Go community uses.

Versions are numeric non-decimal numbers found in the import path URL, like so:

```
import "azul3d.org/v1/pkg"
```

Because the version is at the front of the URL, all Azul3D packages are tied under a guarantee for the specific version (v1 here). Because the versions are non-decimal, each new version is incremented by a single numeric value (v1, v2, v3, v4, but no v1.2).

The code is hosted on Google Code at the moment and the import URL azul3d.org just redirect to the actual git repository (so the code is always available for download no matter what).

New versions are only released when a serious API-incompatible break must be made, for instance if a core package was going to be completely rewritten from the ground up for some reason -- a new version would be released.

But new features, bug fixes, etc are all still added to the latest version even after a version release has been made so the version number is not often incremented. This is important to us because Azul3D develops in a highly backwards compatible way, but we can always just increment the version number and break backwards compatibility without breaking depending code.

It should be noted that versions cannot be mixed safely: a package that depends on v1 should not be used with v2 until it is updated to reflect the new changes, etc. Using tools like [govers](https://launchpad.net/govers) can greatly help detect these such cases and provide an easy way to re-write import paths to update code to new Azul3D versions.

## Version 1 ##
**Below is an outline of the version 1 changes, for a full list see: [Version 1](Version1.md).**

Version 1 includes many refinements and critical fixes:

  * Brand new graphics API (azul3d.org/v1/gfx).
  * Rewrite of the OpenGL 2 renderer (azul3d.org/v1/gfx/gl2).
  * New 3D Audio API using OpenAL (azul3d.org/v1/a3d).
  * Support for windowing toolkits: GLFW3, QML, and Chippy.
  * Support for Windows, i386.
  * Support for Linux, i386 and amd64.