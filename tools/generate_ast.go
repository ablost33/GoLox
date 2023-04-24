package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	expressions := map[string]string{
		"Binary":   "Expr left, Token operator, Expr right",
		"Grouping": "Expr expression",
		"Literal":  "value interface{}",
		"Unary ":   "Token operator, Expr right",
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("[Tools] os.Getwd err: %s", err)
		return
	}
	path := filepath.Join(dir, "/golox/ast.go")
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("[Tools] os.Create err: %s", err)
	}
	file.WriteString("package lox\n\n")
	defineAst("Expression", expressions, file)

	if err := file.Close(); err != nil {
		fmt.Printf("[Tools] file.Close() err: %s", err)
	}
}

func defineAst(iface string, types map[string]string, outputFile *os.File) {
	outputFile.WriteString(fmt.Sprintf("// %s representation\n", iface))
	outputFile.WriteString(fmt.Sprintf("type %s interface {\n", iface))
	outputFile.WriteString(fmt.Sprintf("	Accept(v %sVisitor) (interface{}, error)\n", iface))
	outputFile.WriteString("}\n\n")

	outputFile.WriteString(fmt.Sprintf("// %sVisitor defines the visit method of every %s\n", iface, iface))
	outputFile.WriteString(fmt.Sprintf("type %sVisitor interface {\n", iface))
	for name, _ := range types {
		outputFile.WriteString(fmt.Sprintf("	visit%s(e *%s) (interface{}, error)\n", name, name))
	}
	outputFile.WriteString("}\n\n")

	for name, value := range types {
		outputFile.WriteString(fmt.Sprintf("// New%s %s constructor\n", name, iface))
		outputFile.WriteString(fmt.Sprintf("func New%s(%s) *%s {\n", name, value, name))
		outputFile.WriteString(fmt.Sprintf("	return &%s{\n", name))

		for _, line := range strings.Split(value, ", ") {
			prop := strings.Split(line, " ")[0]
			outputFile.WriteString(fmt.Sprintf("		%s: %s,\n", prop, prop))
		}
		outputFile.WriteString("	}\n")
		outputFile.WriteString("}\n\n")

		outputFile.WriteString(fmt.Sprintf("// %s %s implementation\n", name, iface))
		outputFile.WriteString(fmt.Sprintf("type %s struct {\n", name))
		for _, line := range strings.Split(value, ", ") {
			outputFile.WriteString("	" + line + "\n")
		}
		outputFile.WriteString("}\n\n")
		outputFile.WriteString("// Accept method of the visitor pattern it calls the proper visit method\n")
		outputFile.WriteString(fmt.Sprintf("func(e *%s) Accept(v %sVisitor) (interface{}, error) {\n", name, iface))
		outputFile.WriteString(fmt.Sprintf("	return v.visit%s(e)\n", name))
		outputFile.WriteString("}\n\n")
	}
}
