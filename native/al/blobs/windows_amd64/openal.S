.section .rdata,"dr" /* put in COFF section .rdata */
.globl _blobopenal /* no longer need to prepend package name here */
.globl _eblobopenal
_blobopenal:
.incbin "openal"
_eblobopenal:

