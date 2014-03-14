# Copyright 2014 The Azul3D Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ar x libxkbcommon-x11.a
ld -r *.o libxkbcommon.a -o ../xkbcommon_amd64.syso
rm -rf *.o


# Copyright 2014 The Azul3D Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Rename static libraries
#
#ar r ../libxkbcommon-x11_amd64.a *.o
#rm -rf *.o

#ar x libxkbcommon.a
#ar r ../libxkbcommon_amd64.a *.o
#rm -rf *.o

