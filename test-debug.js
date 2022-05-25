const debug = true;

function router(r) {
    if (r.uri==="/api/v1/") {
        
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

    if (r.method==="GET") {
        return "<UPSTREAM-VALUE-HERE>";
    } else { 
        debug && r.log(r.method);
    }

    if (r.method==="POST") {
        return "<UPSTREAM-VALUE-HERE>";
    } else { 
        debug && r.log(r.method);
    }

    if (decodeURIComponent(r.args.thing)==="test") {
        return "<UPSTREAM-VALUE-HERE>";
    } else { 
        debug && r.log(r.args.thing);
    }

    if (r.headersIn['key']==="value") {
        return "<UPSTREAM-VALUE-HERE>";
    } else { 
        debug && r.log(r.headersIn['key']);
    }

    return "<UPSTREAM-VALUE-HERE>";
}
export default {router};