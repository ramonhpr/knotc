# knot-compiler

A compiler to convert knot definition schema in thing (zephyr, arduino) code

# how to generate parser
`antlr -Dlanguage=Go -o pkg/generated -package generated Knot.g4`

