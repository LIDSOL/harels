package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"lidsol.org/harels/parser"
)

type hareListener struct {
	*parser.BaseHareListener
	stack []int
}

func (l *hareListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *hareListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty")
	}

	result := l.stack[len(l.stack)-1]

	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *hareListener) isEmpty() bool {
	return len(l.stack) == 0
}

func (l *hareListener) ExitIdentifier(c *parser.IdentifierContext) {
	tok := c.GetStart()
	fmt.Println("Name: ", tok.GetLine(), c.GetText())

	if l.isEmpty() {
		l.push(0)
	} else {
		val := l.pop()
		l.push(val+1)
	}
}

func (l *hareListener) EnterCallExpression(c *parser.CallExpressionContext) {
	tok := c.GetStart()
	fmt.Println("CALL: ", tok.GetLine(), c.GetText())

	if l.isEmpty() {
		l.push(0)
	} else {
		val := l.pop()
		l.push(val+1)
	}
}

func (l *hareListener) ExitType(c *parser.TypeContext) {
	tok := c.GetStart()
	fmt.Println("type: ", tok.GetLine(), c.GetText())

	if l.isEmpty() {
		l.push(0)
	} else {
		val := l.pop()
		l.push(val+1)
	}
}

func hare(input string) int {
	is := antlr.NewInputStream(input)

	lexer := parser.NewHareLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewHareParser(stream)

	var listener hareListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

func main() {
	fmt.Println(hare(`
use fmt;

fn fac(n: int) void = {
	if (n <= 1) {
                return 1;
        };
        return n * fac(n-1);
};

export fn main() void = {
	const greetings = [
		"hello, world",
		"hola mundo",
	];
	for (let greeting .. greetings) {
		fmt::println(greeting)!;
	};
};
`))
}
