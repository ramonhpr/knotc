# knot-compiler

A compiler to convert knot definition schema in thing (zephyr, arduino) code

# how to generate parser
`antlr -Dlanguage=Go -visitor -o generated -package generated Knot.g4`

