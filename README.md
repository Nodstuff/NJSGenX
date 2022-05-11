# NJSGenX 
### A very simple function generator for NJS

#### Basic Example Output With Debug 
```javascript
const debug = true;

function router(r) {
    if (r.uri.match("/api/v1/test/\w+")) {
        return "127.0.0.1:8082";
    } else {
        debug && r.log(r.uri);
    }

    if (r.method==="PUT") {
        return "127.0.0.1:8085";
    } else {
        debug && r.log(r.method);
        return "127.0.0.1:8086";
    }

    if (r.method==="DELETE") {
        return "127.0.0.1:8090";
    } else {
        debug && r.log(r.method);
        return "127.0.0.1:8091";
    }

    if (decodeURIComponent(r.args.env)==="test") {
        return "127.0.0.1:8090";
    } else {
        debug && r.log(r.args.env);
        return "127.0.0.1:8091";
    }

    if (r.headersIn['key']==="value") {
        return "127.0.0.1:8090";
    } else {
        debug && r.log(r.headersIn['key']);
        return "127.0.0.1:8091";
    }

    return "127.0.0.1:80";
}
export default {router};
```

#### Basic Example Output Without Debug
```javascript
function router(r) {
    if (r.uri.match("/api/v1/test/\w+")) {
        return "127.0.0.1:8082";
    }

    if (r.method==="PUT") {
        return "127.0.0.1:8085";
    } else {
        return "127.0.0.1:8086";
    }

    if (r.method==="DELETE") {
        return "127.0.0.1:8090";
    } else {
        return "127.0.0.1:8091";
    }

    if (decodeURIComponent(r.args.env)==="test") {
        return "127.0.0.1:8090";
    } else {
        return "127.0.0.1:8091";
    }

    if (r.headersIn['key']==="value") {
        return "127.0.0.1:8090";
    } else {
        return "127.0.0.1:8091";
    }

    return "127.0.0.1:80";
}
export default {router};
```