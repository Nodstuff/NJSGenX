package main

import (
	"NJSGenX/NJSGenX"
	"fmt"
)

func main() {
	var blocks []NJSGenX.Block

	for i := 0; i < 3; i++ {
		blocks = append(blocks, NJSGenX.NewBlock().
			WithConditional("if").
			WithRegexMatch("r.uri", fmt.Sprintf("/api/v%d/test/\\w+", i+1)).
			WithBodyReturning("true").
			WithElseReturning("false"))
	}

	fn := NJSGenX.NewFunction("router").
		WithParameters("r").
		WithReturn("false").
		WithBlocks(blocks...).
		WithDebug().
		Build()

	fmt.Println(fn)
}
