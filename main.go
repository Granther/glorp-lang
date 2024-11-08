package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"glorp/scanner"
)

type Glorp struct {
	HadError bool
	Scanner  *scanner.Scanner
}

func NewGlorp() *Glorp {
	return &Glorp{
		HadError: false,
		Scanner: scanner.NewScanner(),
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
	tokens, err := g.Scanner.Scan(source)
	if err != nil {
		return err
	}

	for _, tok := range tokens {
		fmt.Println(tok.Type)
	}
	// Scan our source file
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
