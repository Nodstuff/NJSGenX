# NJSGenX

### A very simple routing function generator for NJS

#### Example Usage

```go
func main() {
	a := NJSGenX.NewBlock().
		WithURIRegexMatch("/\\/api\\/v1\\/test\\/\\d+/").
		WithSubRoutes("\"/api/v1/\"",
			NJSGenX.NewBlock().
				WithURIRegexMatch("test/\\w+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
			NJSGenX.NewBlock().
				WithURIRegexMatch("test2/\\d+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
			NJSGenX.NewBlock().
				WithURIRegexMatch("test3/\\d+").
				WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\""),
		)

	b := NJSGenX.NewBlock().
		WithMatchRequestMethod(NJSGenX.MethodGet).
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	c := NJSGenX.NewBlock().
		WithMatchRequestMethod(NJSGenX.MethodPost).
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	d := NJSGenX.NewBlock().
		WithMatchQueryParam("thing", "test").
		WithBodyReturning("\"<UPSTREAM-VALUE-HERE>\"")

	e := NJSGenX.NewBlock().
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
```

#### Basic Example Output With Debug

```javascript
const debug = true;

function router(r) {
    if (r.uri === "/api/v1/") {
        if (r.uri.match("test/\w+")) {
            return "<UPSTREAM-VALUE-HERE>";
        }
        if (r.uri.match("test2/\d+")) {
            return "<UPSTREAM-VALUE-HERE>";
        }
        if (r.uri.match("test3/\d+")) {
            return "<UPSTREAM-VALUE-HERE>";
        }
    } else {
        debug && r.log(r.uri);
    }

    if (r.method === "GET") {
        return "<UPSTREAM-VALUE-HERE>";
    } else {
        debug && r.log(r.method);
    }

    if (r.method === "POST") {
        return "<UPSTREAM-VALUE-HERE>";
    } else {
        debug && r.log(r.method);
    }

    if (decodeURIComponent(r.args.thing) === "test") {
        return "<UPSTREAM-VALUE-HERE>";
    } else {
        debug && r.log(r.args.thing);
    }

    if (r.headersIn['key'] === "value") {
        return "<UPSTREAM-VALUE-HERE>";
    } else {
        debug && r.log(r.headersIn['key']);
    }

    return "<UPSTREAM-VALUE-HERE>";
}

export default {router};
```

#### Basic Example Output Without Debug

```javascript
function router(r) {
    if (r.uri === "/api/v1/") {
        if (r.uri.match("test/\w+")) {
            return "<SOME-VALUE-HERE>";
        }
        if (r.uri.match("test2/\d+")) {
            return "<SOME-VALUE-HERE>";
        }
        if (r.uri.match("test3/\d+")) {
            return "<SOME-VALUE-HERE>";
        }
    }

    if (r.method === "GET") {
        return "<SOME-VALUE-HERE>";
    }

    if (r.method === "POST") {
        return "<SOME-VALUE-HERE>";
    }

    if (decodeURIComponent(r.args.thing) === "test") {
        return "<SOME-VALUE-HERE>";
    }

    if (r.headersIn['key'] === "value") {
        return "<SOME-VALUE-HERE>";
    }

    return "<SOME-VALUE-HERE>";
}

export default {router};
```