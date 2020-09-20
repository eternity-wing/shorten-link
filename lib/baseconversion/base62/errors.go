package base62

type UnexpectedCharacterError struct {
}

func (e *UnexpectedCharacterError) Error() string {
	return "Unexpected character!\nAcceptable characters are numerical and alphabetical"
}

type UnexpectedNumberError struct {
}

func (e *UnexpectedNumberError) Error() string {
	return "Unexpected Number!\nAcceptable numbers are between 0 and 61"
}
