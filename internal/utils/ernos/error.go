package ernos

type Ernos struct {
	Message string
	Code    string
}

func (e Ernos) Error() string {
	return e.Message
}

func HasAlreadyExists(value string) error {
	msg := M.DUPLICATE_RECORD
	if value != "" {
		msg = value + " has already exists"
	}

	return Ernos{
		Message: msg,
		Code:    C.DUPLICATE_RECORD,
	}
}

func NotFound(value string) error {
	msg := M.RECORD_NOTFOUND
	if value != "" {
		msg = value + " not found"
	}

	return Ernos{
		Message: msg,
		Code:    C.RECORD_NOTFOUND,
	}
}

func Unauthorized(value string, code string) error {
	msg := M.UNAUTHORIZED
	if value != "" {
		msg = value
	}

	return Ernos{
		Message: msg,
		Code:    C.UNAUTHORIZED,
	}
}

func InternalServerError(value string) error {
	msg := M.INTERNAL_SERVER_ERROR
	if value != "" {
		msg = value
	}

	return Ernos{
		Message: msg,
		Code:    C.INTERNAL_SERVER_ERROR,
	}
}

func Other(e Ernos) error {
	return e
}
