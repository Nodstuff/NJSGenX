package NJSGenX

import (
	"fmt"
	"os"
	"strings"
)

const (
	space           = "    "
	doubleSpace     = space + space
	tripleSpace     = doubleSpace + space
	signatureString = "function %s(%s) {\n"
	debugLogs       = "debug && r.log(%s);\n"
)

type Function struct {
	name       string
	parameters []string
	blocks     []Block
	returns    string
	debug      bool
	debugLogs  string
}

func NewFunction(fnName string) Function {
	return Function{name: fnName, debugLogs: debugLogs}
}

func (f Function) WithDebug() Function {
	f.debug = true
	return f
}

func (f Function) WithParameters(params ...string) Function {
	for _, p := range params {
		f.parameters = append(f.parameters, p)
	}
	return f
}

func (f Function) WithBlocks(blks ...Block) Function {
	f.blocks = append(f.blocks, blks...)
	return f
}

func (f Function) WithReturn(rtrnVal string) Function {
	f.returns = rtrnVal
	return f
}

func (f Function) Build() string {
	bld := &strings.Builder{}
	funcParams := strings.Join(f.parameters, ",")
	if f.debug {
		bld.WriteString("const debug = true;\n\n")
	}
	bld.WriteString(fmt.Sprintf(signatureString, f.name, funcParams))
	buildNestedBlocks(bld, f.blocks, f.debug)
	bld.WriteString(space + "return " + f.returns + ";\n}")
	bld.WriteString("\nexport default {" + f.name + "};")
	return bld.String()
}

func (f Function) WriteToFile(filename string) (string, error) {
	fl, err := os.Create("./" + filename)
	if err != nil {
		return "", err
	}
	defer fl.Close()
	_, err = fl.WriteString(f.Build())
	if err != nil {
		return "", err
	}
	return fl.Name(), nil
}

func buildNestedBlocks(bldr *strings.Builder, blks []Block, debug bool) {
	for _, b := range blks {
		bldr.WriteString(b.Build(debug))
	}
}
