package ernos

/* ------------------------------ CONSTANT CODE ----------------------------- */

type c struct {
	DUPLICATE_RECORD      string
	RECORD_NOTFOUND       string
	INVALID_CREDENTIALS   string
	UNAUTHORIZED          string
	INTERNAL_SERVER_ERROR string
	BROKEN_TOKEN          string
	NOT_FOUND_BEARER      string
}

var C = c{
	RECORD_NOTFOUND:       "RECORD_NOTFOUND",
	DUPLICATE_RECORD:      "DUPLICATE_RECORD",
	INVALID_CREDENTIALS:   "INVALID_CREDENTIALS",
	UNAUTHORIZED:          "UNAUTHORIZED",
	INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
	BROKEN_TOKEN:          "BROKEN_TOKEN",
	NOT_FOUND_BEARER:      "NOT_FOUND_BEARER",
}
