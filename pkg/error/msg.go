package error

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Invalid request params",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token check failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token check timeout",
	ERROR_AUTH_TOKEN:               "Token generation failed",
	ERROR_AUTH:                     "Wrong token",
	//Add your API error messages
	ERROR_EXAMPLE_EXIST:      "Example exist",
	ERROR_EXAMPLE_NOT_EXIST:  "Example does not exist",
	ERROR_EXIST_EXAMPLE_FAIL: "Example exists but fail",

	ERROR_GET_EXAMPLE_FAIL:    "Get example failed",
	ERROR_COUNT_EXAMPLE_FAIL:  "Count example failed",
	ERROR_ADD_EXAMPLE_FAIL:    "Add example failed",
	ERROR_EDIT_EXAMPLE_FAIL:   "Edit example failed",
	ERROR_DELETE_EXAMPLE_FAIL: "Delete example failed",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
