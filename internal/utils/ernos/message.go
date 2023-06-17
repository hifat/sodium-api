package ernos

/* ------------------------------ CONSTANT MESSAGE ----------------------------- */

type m struct {
	DUPLICATE_RECORD        string
	RECORD_NOTFOUND         string
	INVALID_CREDENTIALS     string
	UNAUTHORIZED            string
	INTERNAL_SERVER_ERROR   string
	NO_AUTHORIZATION_HEADER string
	NOT_FOUND_BEARER        string
	BROKEN_TOKEN            string
	TOO_MANY_REQUEST        string
	MAX_DEVICES_LOGIN       string
}

var M = m{
	RECORD_NOTFOUND:         "record not found",
	DUPLICATE_RECORD:        "duplicate record",
	INVALID_CREDENTIALS:     "invalid credentials",
	UNAUTHORIZED:            "unauthorized",
	INTERNAL_SERVER_ERROR:   "internal server error",
	NO_AUTHORIZATION_HEADER: "no authorization header provided",
	NOT_FOUND_BEARER:        "could not find bearer token in authorization header",
	BROKEN_TOKEN:            "the token is broken",
	TOO_MANY_REQUEST:        "too many request",
	MAX_DEVICES_LOGIN:       "maximum devices limit reached. please log out from one of your other devices before attempting to log in again",
}
