package msg

var msgMap = map[int]string {
	SUCCESS: "request sucess",
	ERROR: "request fail",
	INVALID_PARAMS: "wrong request params",
}

func getMsg(code int)  string {
	msg, ok := msgMap[code]

	if ok {
		return msg
	}

	return msgMap[ERROR]
}

