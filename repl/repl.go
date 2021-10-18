package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_interpreter/lexer"
	"monkey_interpreter/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
