package main

import (
	"bufio"
	"fmt"
	"glorp/interpreter"
	"glorp/parser"
	"glorp/scanner"
	"os"
	"path/filepath"
)

type Glorp struct {
	HadError    bool
	Scanner     *scanner.Scanner
	Parser		*parser.Parser
	Interpreter *interpreter.Interpreter
}

func NewGlorp() *Glorp {
	return &Glorp{
		HadError: false,
		Scanner:  scanner.NewScanner(),
		Parser:    parser.NewParser(),
	}
}

func (g *Glorp) Start() error {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glorp [file.glp]")
		return nil
	} else if len(args) == 2 {
		return g.Runfile(args[1])
	} else {
		return g.Repl()
	}
}

func (g *Glorp) Runfile(file string) error {
	ext := filepath.Ext(file)
	if ext != ".glp" {
		return fmt.Errorf("glorp file (.glp) is required to run, got %s", ext)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return g.Run(string(data))
}

func (g *Glorp) Repl() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		g.Run(line)
		g.HadError = false
	}
	return nil
}

func (g *Glorp) Run(source string) error {

	tokens, err := g.Scanner.ScanTokens(source)
	if err != nil {
		return err
	}

	// Parse tokens to expressions
	statements := g.Parser.Parse(tokens)

	if g.HadError {
		fmt.Println("Error encountered in Run")
		return nil
	}

	if g.Interpreter.HadRuntimeError {
		fmt.Println("Runtime Error encountered in Run")
		return nil
	}

	// fmt.Println(ast.NewAstPrinter().Print(expr))

	g.Interpreter.Interpret(statements)
	return nil
}

func main() {
	glorp := NewGlorp()
	err := glorp.Start()
	if err != nil {
		fmt.Println("Unable to GLORP: ", err)
		os.Exit(1)
	}
}