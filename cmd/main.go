package main

import (
	"fmt"

	"NJSGenX/NJSGenX"
)

func main() {
	a := NJSGenX.NewRoutingBlock().
		WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
		WithSubRoutes("\"/api/v1/\"",
			NJSGenX.NewRoutingBlock().
				WithURIRegexMatch("test/\\w+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
			NJSGenX.NewRoutingBlock().
				WithURIRegexMatch("test2/\\d+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
			NJSGenX.NewRoutingBlock().
				WithURIRegexMatch("test3/\\d+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
		)

	b := NJSGenX.NewRoutingBlock().
		WithMatchRequestMethod(NJSGenX.MethodGet).
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	c := NJSGenX.NewRoutingBlock().
		WithMatchRequestMethod(NJSGenX.MethodPost).
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	d := NJSGenX.NewRoutingBlock().
		WithMatchQueryParam("thing", "test").
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	e := NJSGenX.NewRoutingBlock().
		WithMatchHeaderValue("key", "value").
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	_, err := NJSGenX.NewFunction("router").
		WithParameters("r").
		WithReturn("\"<UPSTREAM-VALUE-HERE>\"").
		WithDebug().
		WithBlocks(a, b, c, d, e).
		WriteToFile("test-debug.js")

	if err != nil {
		fmt.Println(err)
	}
}
