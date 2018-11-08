package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"git.atonline.com/tristantech/gophp/core/tokenizer"
)

type Global struct {
	context.Context
	p *Process

	globalFuncs map[ZString]Callable

	out io.Writer
}

func NewGlobal(ctx context.Context, p *Process) *Global {
	res := &Global{
		Context: ctx,
		p:       p,
		out:     os.Stdout,

		globalFuncs: make(map[ZString]Callable),
	}

	// import global funcs from ext
	for _, e := range globalExtMap {
		for k, v := range e.Functions {
			res.globalFuncs[ZString(k)] = v
		}
	}
	return res
}

func (g *Global) RunFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	// tokenize
	t := tokenizer.NewLexer(f, fn)

	ctx := NewContext(g)
	// compile
	c := compile(ctx, t)

	_, err = c.Run(ctx)
	return err
}

func (g *Global) Write(v []byte) (int, error) {
	return g.out.Write(v)
}

func (g *Global) GetVariable(name ZString) (*ZVal, error) {
	// TODO
	return nil, nil
}

func (g *Global) SetVariable(name ZString, v *ZVal) error {
	// TODO
	return nil
}

func (g *Global) RegisterFunction(name ZString, f Callable) error {
	name = name.ToLower()
	if _, exists := g.globalFuncs[name]; exists {
		return errors.New("duplicate function name in declaration")
	}
	g.globalFuncs[name] = f
	return nil
}

func (g *Global) GetFunction(name ZString) (Callable, error) {
	if f, ok := g.globalFuncs[name.ToLower()]; ok {
		return f, nil
	}
	return nil, fmt.Errorf("Call to undefined function %s", name)
}