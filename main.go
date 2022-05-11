package main

import (
	"NJSGenX/NJSGenX"
	"fmt"
)

func main() {
	a := NJSGenX.NewBlock().
		WithURIRegexMatch("/api/v1/test/\\w+").
		WithBodyReturning("\"127.0.0.1:8082\"")

	b := NJSGenX.NewBlock().
		WithMatchRequestMethod(NJSGenX.MethodPut).
		WithBodyReturning("\"127.0.0.1:8085\"").
		WithElseReturning("\"127.0.0.1:8086\"")

	c := NJSGenX.NewBlock().
		WithMatchRequestMethod(NJSGenX.MethodDelete).
		WithBodyReturning("\"127.0.0.1:8090\"").
		WithElseReturning("\"127.0.0.1:8091\"")

	d := NJSGenX.NewBlock().
		WithMatchQueryParam("test").
		WithBodyReturning("\"127.0.0.1:8090\"").
		WithElseReturning("\"127.0.0.1:8091\"")

	fn, err := NJSGenX.NewFunction("router").
		WithParameters("r").
		WithReturn("\"127.0.0.1:80\"").
		WithBlocks(a, b, c, d).
		WithDebug().
		WriteToFile("test.js")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fn)
}
