package lexer

import "io"

type Lexer struct {
	input        io.Reader
	position     int
	readPosition int
	ch           byte
}

func New(input io.Reader) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {

}
