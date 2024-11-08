package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Glorp struct {
	HadError bool
}

func NewGlorp() *Glorp {
	return &Glorp{
		HadError: false,
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
	if ext != "glp" {
		return fmt.Errorf("glorp file (.glp) is required to run, got %s", ext)
	}
	return nil
}

func (g *Glorp) Repl() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		g.Run(line)
		g.HadError = false
	}
	return nil
}

func (g *Glorp) Run(line string) {
	
}

func main() {
	glorp := NewGlorp()
	err := glorp.Start()
	if err != nil {
		fmt.Println("Unable to GLORP: ", err)
		os.Exit(1)
	}
}
