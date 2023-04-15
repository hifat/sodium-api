package ernos

type Ernos struct {
	Message string
	Code    string
}

func (e Ernos) Error() string {
	return e.Message
}

func HasAlreadyExists(value string) error {
	msg := "duplicate record"
	if value != "" {
		msg = value + " has already exists"
	}

	return Ernos{
		Message: msg,
		Code:    C.DUPLICATE_RECORD,
	}
}

func Other(e Ernos) error {
	return e
}
