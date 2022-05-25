package NJSGenX

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestFunction_Build(t *testing.T) {
	tests := []struct {
		name  string
		given Function
		want  string
	}{
		{
			name: "Test Basic Function",
			given: NewFunction("router").
				WithParameters("r").
				WithReturn("\"10.0.0.1:8080\""),
			want: "function router(r) {\n    return \"10.0.0.1:8080\";\n}\nexport default {router};",
		},
		{
			name: "Test Basic Function with multiple parameters",
			given: NewFunction("router").
				WithParameters("r", "s", "t").
				WithReturn("\"10.0.0.1:8080\""),
			want: "function router(r,s,t) {\n    return \"10.0.0.1:8080\";\n}\nexport default {router};",
		},
		{
			name: "Test Basic Function with Debug",
			given: NewFunction("router").
				WithParameters("r").
				WithReturn("\"10.0.0.1:8080\"").
				WithDebug(),
			want: "const debug = true;\n\nfunction router(r) {\n" +
				"    return \"10.0.0.1:8080\";\n}\n" +
				"export default {router};",
		},
		{
			name: "Test Function with Block",
			given: NewFunction("router").
				WithParameters("r").
				WithReturn("\"10.0.0.1:8080\"").
				WithBlocks(NewRoutingBlock().
					WithMatchRequestMethod(MethodGet).
					WithBodyReturning("\"10.0.0.2:8080\"")),
			want: "function router(r) {\n    if (r.method===\"GET\") {\n" +
				"        return \"10.0.0.2:8080\";\n    }\n\n" +
				"    return \"10.0.0.1:8080\";\n}\nexport default {router};",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := assertions.New(t)
			got := strings.TrimSpace(tc.given.Build())
			a.So(got, assertions.ShouldResemble, strings.TrimSpace(tc.want))
		})
	}
}
