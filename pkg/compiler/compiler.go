package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ramonhpr/knotc/pkg/generated"
	"github.com/ramonhpr/knotc/pkg/listener"
)

// KnotCompiler is a interface that walt to parser tree and compile it
type KnotCompiler interface {
	Walk(string, generated.KnotListener)
	Compile(string)
}

// BaseKnotCompiler that run the generated parsers and lexers
type BaseKnotCompiler struct {}

// Walk on the Parser tree
func (bc *BaseKnotCompiler) Walk(input string, l generated.KnotListener) {
	lexer := generated.NewKnotLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	
	p := generated.NewKnotParser(stream)
	p.AddErrorListener(new(listener.ExitErrorListener))

	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())
}

// Compile the input
func (bc *BaseKnotCompiler) Compile(input string) {
	bc.Walk(input, &generated.BaseKnotListener{})
}