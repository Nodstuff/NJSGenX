package NJSGenX

import (
	"fmt"
	"strings"
)

const (
	space           = "    "
	doubleSpace     = space + "    "
	signatureString = "function %s(%s) {\n"
	nestedString    = "%s%s (%s%s%s) {\n%s%s;\n%s}"
	debugLogs       = "debug && r.log(\"%s\");"
)

type function struct {
	name       string
	parameters []string
	blocks     []Block
	returns    string
	debug      bool
	debugLogs  string
}

type Block struct {
	conditional string
	regex       bool
	predicate   string
	args        conditionalArgs
	body        string
	withElse    bool
	elseBody    string
}

type conditionalArgs struct {
	arg1 string
	arg2 string
}

func (f function) Build() string {
	bld := &strings.Builder{}
	funcParams := strings.Join(f.parameters, ",")
	if f.debug {
		bld.WriteString("var debug = true;\n\n")
	}
	bld.WriteString(fmt.Sprintf(signatureString, f.name, funcParams))
	buildNestedBlocks(bld, f.blocks, f.debug)
	bld.WriteString(space + "return " + f.returns + ";\n}")
	bld.WriteString("\nexport default {" + f.name + "};")
	return bld.String()
}

func buildNestedBlocks(bldr *strings.Builder, blks []Block, debug bool) {
	for _, b := range blks {
		bldr.WriteString(fmt.Sprintf(nestedString, space, b.conditional, b.args.arg1, b.predicate, b.args.arg2, doubleSpace, b.body, space))
		if b.withElse {
			bldr.WriteString(" else { ")
			if debug {
				bldr.WriteString(fmt.Sprintf(debugLogs, b.args.arg1))
			}
			bldr.WriteString(fmt.Sprintf(" %s; }\n\n", b.elseBody))
		} else {
			bldr.WriteString("\n\n")
		}
	}
}

func NewFunction(fnName string) function {
	return function{name: fnName, debugLogs: debugLogs}
}

func (f function) WithDebug() function {
	f.debug = true
	return f
}

func (f function) WithParameters(params ...string) function {
	for _, p := range params {
		f.parameters = append(f.parameters, p)
	}
	return f
}

func (f function) WithBlocks(blks ...Block) function {
	for _, b := range blks {
		f.blocks = append(f.blocks, b)
	}
	return f
}

func (f function) WithReturn(rtrnVal string) function {
	f.returns = rtrnVal
	return f
}

func NewBlock() Block {
	return Block{}
}

func (b Block) WithRegexMatch(arg, rgx string) Block {
	b.regex = true
	b.args = conditionalArgs{arg1: arg}
	b.predicate = ".match(\"" + rgx + "\")"
	return b
}

func (b Block) WithConditional(c string) Block {
	b.conditional = c
	return b
}

func (b Block) WithArgs(a1, a2 string) Block {
	b.args = conditionalArgs{arg1: a1, arg2: a2}
	return b
}

func (b Block) WithPredicate(p string) Block {
	b.predicate = p
	return b
}

func (b Block) WithBody(bdy string) Block {
	b.body = bdy
	return b
}

func (b Block) WithBodyReturning(bdy string) Block {
	b.body = "return " + bdy
	return b
}

func (b Block) WithElse(bdy string) Block {
	b.withElse = true
	b.elseBody = bdy
	return b
}

func (b Block) WithElseReturning(bdy string) Block {
	b.withElse = true
	b.elseBody = "return " + bdy
	return b
}
