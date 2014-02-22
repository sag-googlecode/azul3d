ar x libz.a
ar x libbz2.a
ar x libfreetype.a
ar x libpng.a

ld -r *.o -o ../../freetype_linux_amd64.syso

rm -rf *.o
