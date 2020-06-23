package listener

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ExitErrorListener a listener that exits with status 1
type ExitErrorListener struct {
	*antlr.ConsoleErrorListener
}

// SyntaxError exit if the error is syntax
func (c *ExitErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	os.Exit(1)
}
