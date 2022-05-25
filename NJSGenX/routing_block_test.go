package NJSGenX

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestBlockBuilder(t *testing.T) {
	tests := []struct {
		name  string
		given RoutingBlock
		want  string
		debug bool
	}{
		{
			"Test Regex match with body return and no else or debug",
			NewRoutingBlock().
				WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
				WithBodyReturning("\"10.0.0.1:8080\""),

			"if (r.uri.match(\"/\\/api\\/v1\\/test\\/\\d+/\")) {\n" +
				"        return \"10.0.0.1:8080\";\n    }",
			false,
		},
		{
			"Test Regex match with body return and else but no debug",
			NewRoutingBlock().
				WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.2:8080\""),
			"if (r.uri.match(\"/\\/api\\/v1\\/test\\/\\d+/\")) {\n" +
				"        return \"10.0.0.1:8080\";\n    }" +
				" else { \n        return \"10.0.0.2:8080\"\n    }",
			false,
		},
		{
			"Test Regex match with body return, else and debug",
			NewRoutingBlock().
				WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.2:8080\""),
			"if (r.uri.match(\"/\\/api\\/v1\\/test\\/\\d+/\")) {\n " +
				"       return \"10.0.0.1:8080\";\n    } " +
				"else { \n        debug && r.log(r.uri);\n" +
				"        return \"10.0.0.2:8080\"\n    }",
			true,
		},
		{
			"Test match GET request method with body return, no else or debug",
			NewRoutingBlock().
				WithMatchRequestMethod(MethodGet).
				WithBodyReturning("\"10.0.0.1:8080\""),
			"if (r.method===\"GET\") {\n        return \"10.0.0.1:8080\";\n    }",
			false,
		},
		{
			"Test match POST request method with body return, no else or debug",
			NewRoutingBlock().
				WithMatchRequestMethod(MethodPost).
				WithBodyReturning("\"10.0.0.1:8080\""),
			"if (r.method===\"POST\") {\n        return \"10.0.0.1:8080\";\n    }",
			false,
		},
		{
			"Test match OPTIONS request method with body return, else but no debug",
			NewRoutingBlock().
				WithMatchRequestMethod(MethodOptions).
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (r.method===\"OPTIONS\") {\n" +
				"        return \"10.0.0.1:8080\";\n" +
				"    } else { \n        return \"10.0.0.5:8080\"\n    }",
			false,
		},
		{
			"Test match OPTIONS request method with body return, else and debug",
			NewRoutingBlock().
				WithMatchRequestMethod(MethodOptions).
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (r.method===\"OPTIONS\") {\n" +
				"        return \"10.0.0.1:8080\";\n" +
				"    } else { \n        debug && r.log(r.method);\n" +
				"        return \"10.0.0.5:8080\"\n    }",
			true,
		},
		{
			"Test match header value with body return, no else or debug",
			NewRoutingBlock().
				WithMatchHeaderValue("key", "value").
				WithBodyReturning("\"10.0.0.1:8080\""),
			"if (r.headersIn['key']===\"value\") {\n        return \"10.0.0.1:8080\";\n    }",
			false,
		},
		{
			"Test match header value with body return, else but no debug",
			NewRoutingBlock().
				WithMatchHeaderValue("key", "value").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (r.headersIn['key']===\"value\") {\n" +
				"        return \"10.0.0.1:8080\";\n" +
				"    } else { \n        return \"10.0.0.5:8080\"\n    }",
			false,
		},
		{
			"Test match header value with body return, else and debug",
			NewRoutingBlock().
				WithMatchHeaderValue("key", "value").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (r.headersIn['key']===\"value\") {\n" +
				"        return \"10.0.0.1:8080\";\n    } " +
				"else { \n        debug && r.log(r.headersIn['key']);\n" +
				"        return \"10.0.0.5:8080\"\n    }",
			true,
		},
		{
			"Test match query param with body return, no else or debug",
			NewRoutingBlock().
				WithMatchQueryParam("param", "value").
				WithBodyReturning("\"10.0.0.1:8080\""),
			"if (decodeURIComponent(r.args.param)===\"value\") {\n" +
				"        return \"10.0.0.1:8080\";\n    }",
			false,
		},
		{
			"Test match query param with body return, else but no debug",
			NewRoutingBlock().
				WithMatchQueryParam("param", "value").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (decodeURIComponent(r.args.param)===\"value\") {\n" +
				"        return \"10.0.0.1:8080\";\n    } else" +
				" { \n        return \"10.0.0.5:8080\"\n    }",
			false,
		},
		{
			"Test match query param with body return, else and debug",
			NewRoutingBlock().
				WithMatchQueryParam("param", "value").
				WithBodyReturning("\"10.0.0.1:8080\"").
				WithElseReturning("\"10.0.0.5:8080\""),
			"if (decodeURIComponent(r.args.param)===\"value\") {\n" +
				"        return \"10.0.0.1:8080\";\n    } else" +
				" { \n        debug && r.log(r.args.param);\n" +
				"        return \"10.0.0.5:8080\"\n    }",
			true,
		},
		{
			"Test Regex match with sub-routes with body return",
			NewRoutingBlock().
				WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
				WithSubRoutes("\"/api/v1/\"",
					NewRoutingBlock().
						WithURIRegexMatch("test/\\w+").
						WithBodyReturning("\"10.0.0.1\""),
					NewRoutingBlock().
						WithURIRegexMatch("test2/\\d+").
						WithBodyReturning("\"10.0.0.2\""),
					NewRoutingBlock().
						WithURIRegexMatch("test3/\\d+").
						WithBodyReturning("\"10.0.0.3\""),
				),
			"if (r.uri===\"/api/v1/\") {\n" +
				"        \n        if (r.uri.match(\"test/\\w+\")) {\n" +
				"            return \"10.0.0.1\";\n        }\n" +
				"        if (r.uri.match(\"test2/\\d+\")) {\n" +
				"            return \"10.0.0.2\";\n        }\n" +
				"        if (r.uri.match(\"test3/\\d+\")) {\n" +
				"            return \"10.0.0.3\";\n        }\n" +
				"    } else { \n        debug && r.log(r.uri);\n    }",
			true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := assertions.New(t)
			got := strings.TrimSpace(tc.given.Build(tc.debug))
			a.So(got, assertions.ShouldResemble, strings.TrimSpace(tc.want))
		})
	}
}
