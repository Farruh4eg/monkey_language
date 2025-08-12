// Package repl is READ EVAL PRINT LOOP for monkey language
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/farruh4eg/monkey_interpreter/pkg/lexer"
	"github.com/farruh4eg/monkey_interpreter/pkg/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
