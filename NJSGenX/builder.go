package NJSGenX

import (
	"fmt"
	"os"
	"strings"
)

const (
	space           = "    "
	doubleSpace     = space + "    "
	signatureString = "function %s(%s) {\n"
	nestedString    = "%s%s (%s%s%s) {\n%s%s;\n%s}"
	queryString     = "%s%s (decodeURIComponent(%s)%s%s) {\n%s%s;\n%s}"
	debugLogs       = "debug && r.log(%s);\n"
)

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
		if b.query {
			bldr.WriteString(fmt.Sprintf(queryString, space, b.conditional, b.args.arg1, b.operator, b.args.arg2, doubleSpace, b.body, space))
		} else {
			bldr.WriteString(fmt.Sprintf(nestedString, space, b.conditional, b.args.arg1, b.operator, b.args.arg2, doubleSpace, b.body, space))
		}
		if b.withElse || debug {
			bldr.WriteString(" else { \n")
			if debug {
				bldr.WriteString(fmt.Sprintf(doubleSpace+debugLogs, b.args.arg1))
			}
			if b.elseBody != "" {
				bldr.WriteString(fmt.Sprintf("%s%s;\n", doubleSpace, b.elseBody))
			}
			bldr.WriteString(fmt.Sprintf("%s}\n\n", space))
		} else {
			bldr.WriteString("\n\n")
		}
	}
}
