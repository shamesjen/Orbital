namespace go api

struct Request {
    1: string message
}

struct Response {
    1: string message
}

service test {
    Response reply(1: Request req)
}