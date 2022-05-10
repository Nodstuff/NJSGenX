package NJSGenX

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
	for _, b := range blks {
		f.blocks = append(f.blocks, b)
	}
	return f
}

func (f Function) WithReturn(rtrnVal string) Function {
	f.returns = rtrnVal
	return f
}
