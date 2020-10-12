package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "param is invalid",

	ERROR_NOT_EXIST_WORD:        "word does not exist",
	ERROR_ADD_WORD_FAIL:         "failed to add a word",
	ERROR_DELETE_WORD_FAIL:      "failed to delete word",
	ERROR_CHECK_EXIST_WORD_FAIL: "failed to check if word exist",
	ERROR_EDIT_WORD_FAIL:        "failed to edit word",
	ERROR_GET_WORDS_FAIL:        "failed to get words",
	ERROR_GET_WORD_FAIL:         "failed to get word",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
