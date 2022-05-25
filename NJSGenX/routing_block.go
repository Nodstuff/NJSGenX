package NJSGenX

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type RequestMethod string

const (
	MethodGet     RequestMethod = http.MethodGet
	MethodHead    RequestMethod = http.MethodHead
	MethodPost    RequestMethod = http.MethodPost
	MethodPut     RequestMethod = http.MethodPut
	MethodPatch   RequestMethod = http.MethodPatch
	MethodDelete  RequestMethod = http.MethodDelete
	MethodConnect RequestMethod = http.MethodConnect
	MethodOptions RequestMethod = http.MethodOptions
	MethodTrace   RequestMethod = http.MethodTrace
)

type RoutingBlock struct {
	conditional string
	regex       bool
	operator    string
	query       bool
	args        conditionalArgs
	body        string
	hasElse     bool
	elseBody    string
	subRoutes   []RoutingBlock
}

type conditionalArgs struct {
	arg1 string
	arg2 string
}

func NewRoutingBlock() RoutingBlock {
	return RoutingBlock{}
}

func (b RoutingBlock) WithSubRoutes(baseRoute string, subRoutes ...RoutingBlock) RoutingBlock {
	b.conditional = "if"
	b.args = conditionalArgs{arg1: "r.uri", arg2: baseRoute}
	b.operator = "==="
	b.subRoutes = subRoutes
	return b
}

func (b RoutingBlock) WithURIRegexMatch(rgx string) RoutingBlock {
	valRgx := regexp.MustCompile(rgx)
	b.regex = true
	b.conditional = "if"
	b.args = conditionalArgs{arg1: "r.uri"}
	b.operator = ".match(\"" + valRgx.String() + "\")"
	return b
}

func (b RoutingBlock) WithMatchRequestMethod(method RequestMethod) RoutingBlock {
	return b.withConditional("if").
		withOperator("===").
		withArgs("r.method",
			fmt.Sprintf("\"%s\"", method))
}

func (b RoutingBlock) WithMatchQueryParam(param, value string) RoutingBlock {
	return b.withConditional("if").
		withQueryParams(
			fmt.Sprintf("r.args.%s", param),
			fmt.Sprintf("\"%s\"", value)).
		withOperator("===")
}

func (b RoutingBlock) WithMatchHeaderValue(key, value string) RoutingBlock {
	return b.withConditional("if").
		withArgs(
			fmt.Sprintf("r.headersIn['%s']",
				strings.ToLower(key)),
			fmt.Sprintf("\"%s\"",
				strings.ToLower(value))).
		withOperator("===")
}

func (b RoutingBlock) WithBodyReturning(bdy string) RoutingBlock {
	b.body = "return " + bdy + ";"
	return b
}

func (b RoutingBlock) WithElseReturning(bdy string) RoutingBlock {
	b.hasElse = true
	b.elseBody = "return " + bdy
	return b
}

func (b RoutingBlock) withConditional(c string) RoutingBlock {
	b.conditional = c
	return b
}

func (b RoutingBlock) withArgs(a1, a2 string) RoutingBlock {
	b.args = conditionalArgs{arg1: a1, arg2: a2}
	return b
}

func (b RoutingBlock) withQueryParams(a1, a2 string) RoutingBlock {
	b.query = true
	b.args = conditionalArgs{arg1: a1, arg2: a2}
	return b
}

func (b RoutingBlock) withOperator(o string) RoutingBlock {
	b.operator = o
	return b
}

func (b RoutingBlock) withBody(bdy string) RoutingBlock {
	b.body = bdy
	return b
}

func (b RoutingBlock) withElse(bdy string) RoutingBlock {
	b.hasElse = true
	b.elseBody = bdy
	return b
}

func (b RoutingBlock) Build(debug bool) string {
	bldr := strings.Builder{}
	if b.query {
		bldr.WriteString(fmt.Sprintf(queryString, space, b.conditional, b.args.arg1, b.operator, b.args.arg2, doubleSpace, b.body, space))
	} else {
		bldr.WriteString(fmt.Sprintf(nestedString, space, b.conditional, b.args.arg1, b.operator, b.args.arg2, doubleSpace, b.body))
		if b.subRoutes != nil && len(b.subRoutes) > 0 {
			for _, r := range b.subRoutes {
				bldr.WriteString(fmt.Sprintf(nestedString, doubleSpace, r.conditional, r.args.arg1, r.operator, r.args.arg2, tripleSpace, r.body))
				bldr.WriteString(fmt.Sprintf("%s}\n", doubleSpace))
			}
		}
		bldr.WriteString(fmt.Sprintf("%s}", space))
	}
	if b.hasElse || debug {
		bldr.WriteString(" else { \n")
		if debug {
			bldr.WriteString(fmt.Sprintf(doubleSpace+debugLogs, b.args.arg1))
		}
		if b.elseBody != "" {
			bldr.WriteString(fmt.Sprintf("%s%s\n", doubleSpace, b.elseBody))
		}
		bldr.WriteString(fmt.Sprintf("%s}\n\n", space))
	} else {
		bldr.WriteString("\n\n")
	}

	return bldr.String()
}
