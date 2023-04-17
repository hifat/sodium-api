package ernos

/* ------------------------------ CONSTANT CODE ----------------------------- */

type c struct {
	DUPLICATE_RECORD      string
	RECORD_NOTFOUND       string
	INVALID_CREDENTIALS   string
	UNAUTHORIZED          string
	INTERNAL_SERVER_ERROR string
}

var C = c{
	RECORD_NOTFOUND:       "RECORD_NOTFOUND",
	DUPLICATE_RECORD:      "DUPLICATE_RECORD",
	INVALID_CREDENTIALS:   "INVALID_CREDENTIALS",
	UNAUTHORIZED:          "UNAUTHORIZED",
	INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
}
