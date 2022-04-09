# golangPath
repo for golang

## data type
word, the smallest of all
this word, create data type

go type:
- basic :int, float
- agregates: array and struct
- reference : map and slices, slices
- interface: interface type

rune . int32
uintptr ... to hold the bit for pointer
% only for integers
overflow ..
silently discarted

unary:
- + ... 0+x
- - ... 0-x
- & and
- | or
- ^ XOR

integer:
- 0123 ... as octal
- 0x231 as hex
- radix format: %d, %o, %x: biasa, octal, hex

rune .. written in '%c' or %q

floating point number:
- float32 and float64
- math package.. maxFloat32, maxFloat64
- %e for exponent, %f for floating

boolean:
- short circuit behaviour

strings:
- immutable sequence of bytes
- sharing underlying byte array
- backstring, no esacape sequence are processed
- unicode code point, in go it is called rune

rune:
- utf32, or utf-8
- most ascii can be form with 8 bit
- 1 until 4 bytes, but ascii only need 1 bytes
- as whole string itu ok, tapi kalau butuh karakter, perlu pakai bantuan
- unicode/utf8
- replacement charakter, ini yg lambang tanda tanya.

