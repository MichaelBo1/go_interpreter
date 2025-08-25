package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MichaelBo1/go_interpreter/lexer"
	"github.com/MichaelBo1/go_interpreter/token"
)

const PROMPT = "-> "

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		if !scanner.Scan() {
			return
		}
		line := scanner.Text()

		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
