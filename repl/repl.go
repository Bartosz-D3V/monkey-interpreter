package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_interpreter/evaluator"
	"monkey_interpreter/lexer"
	"monkey_interpreter/object"
	"monkey_interpreter/parser"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		code := scanner.Text()
		l := lexer.New(code)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()

		if len(p.Error()) > 0 {
			printParseError(out, p.Error())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParseError(out io.Writer, errors []string) {
	for _, err := range errors {
		_, _ = io.WriteString(out, fmt.Sprintf("%s \n", err))
	}
}
