package core

import "fmt"

type PhpErrorType int

const (
	PhpErrorFatal = iota
	PhpThrow
	PhpBreak
	PhpContinue
)

type PhpError struct {
	e error
	l *Loc
	t PhpErrorType
}

func (e *PhpError) Run(ctx Context) (*ZVal, error) {
	return nil, e
}

func (e *PhpError) Loc() *Loc {
	return e.l
}

func (e *PhpError) Error() string {
	if e.l == nil {
		return e.e.Error()
	}
	return fmt.Sprintf("%s in %s on line %d", e.e, e.l.Filename, e.l.Line)
}
