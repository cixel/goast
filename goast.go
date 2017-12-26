package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	args := os.Args[1:]
	fileName := args[0]

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("problem reading source file %s:\n%s", fileName, err)
		return
	}

	fmt.Println(fileName)
	ast.Print(fileSet, file)
}
