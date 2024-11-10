package native

import (
	"fmt"
	"glorp/ast"
	"glorp/environment"
)

type Callable interface {
	Arity() int
	Call(interpreter *Interpreter, args []any) (any, error)
	String() string
}

type GlorpFunction struct {
	Declaration ast.Fun
}

func NewGloxFunction(declaration ast.Fun) Callable {
	return &GlorpFunction{
		Declaration: declaration,
	}
}

// Each function gets its own environment to store local vars
// A new environment is necassary when thinking about recursive funs
// They do not share local vars
func (f *GlorpFunction) Call(interpreter *Interpreter, args []any) (any, error) {
	environment := environment.NewEnvironment(interpreter.Globals)
	for i := 0; i < len(f.Declaration.Params); i++ {
		// Place passed args as accessible in the body locally
		environment.Define(f.Declaration.Params[i].Lexeme, args[i])
	}
	// Call function and discard environ, reverting to prev
	err := interpreter.executeBlock(f.Declaration.Body, environment)
	ret, ok := err.(*ReturnErr)
	if ok {
		return ret.Val, nil
	}
	return nil, err
}

func (f *GlorpFunction) Arity() int {
	return len(f.Declaration.Params)
}

func (f *GlorpFunction) String() string {
	return fmt.Sprintf("<fn %s>", f.Declaration.Name.Lexeme)
}
