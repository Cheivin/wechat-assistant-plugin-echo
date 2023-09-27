package function

// Handle a function invocation
func Handle(req Request) Response {
	return Response{Type: Text, Body: req.Message}
}

func Info() Plugin {
	return Plugin{Keyword: "echo", Description: "消息回显"}
}
