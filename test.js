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