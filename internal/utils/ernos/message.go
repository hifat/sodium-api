package ernos

/* ------------------------------ CONSTANT MESSAGE ----------------------------- */

type m struct {
	DUPLICATE_RECORD      string
	RECORD_NOTFOUND       string
	INVALID_CREDENTIALS   string
	UNAUTHORIZED          string
	INTERNAL_SERVER_ERROR string
}

var M = m{
	RECORD_NOTFOUND:       "record not found",
	DUPLICATE_RECORD:      "duplicate record",
	INVALID_CREDENTIALS:   "invalid credentials",
	UNAUTHORIZED:          "unauthorized",
	INTERNAL_SERVER_ERROR: "internal server error",
}
