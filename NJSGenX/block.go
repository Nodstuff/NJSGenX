package NJSGenX

import "regexp"

type Block struct {
	conditional string
	regex       bool
	operator    string
	query       bool
	args        conditionalArgs
	body        string
	withElse    bool
	elseBody    string
}

type conditionalArgs struct {
	arg1 string
	arg2 string
}

func NewBlock() Block {
	return Block{}
}

func (b Block) WithRegexMatch(arg, rgx string) Block {
	valRgx := regexp.MustCompile(rgx)
	b.regex = true
	b.args = conditionalArgs{arg1: arg}
	b.operator = ".match(\"" + valRgx.String() + "\")"
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

func (b Block) WithQueryParams(a1, a2 string) Block {
	b.query = true
	b.args = conditionalArgs{arg1: a1, arg2: a2}
	return b
}

func (b Block) WithOperator(o string) Block {
	b.operator = o
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
