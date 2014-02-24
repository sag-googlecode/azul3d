# Copyright 2014 The Azul3D Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ar x libz.a
ar x libbz2.a
ar x libfreetype.a
ar x libpng.a

ld -r *.o /usr/lib/x86_64-linux-gnu/libm.a -o ../../freetype_linux_amd64.syso

rm -rf *.o
