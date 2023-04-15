package response

type ErrorResponse struct {
	Error ErrorMessageResponse `json:"error"`
}

type ErrorMessageResponse struct {
	Message   string `json:"message,omitempty"`
	Code      string `json:"code,omitempty"`
	Attribute any    `json:"attribute,omitempty"`
}

type SuccesResponse struct {
	Item  any `json:"item,omitempty"`
	Items any `json:"items,omitempty"`
	Total any `json:"total,omitempty"`
}

func HandleErr(err any) ErrorResponse {
	if _, ok := err.(error); ok {
		return ErrorResponse{
			Error: ErrorMessageResponse{
				Message: err.(error).Error(),
			},
		}
	}

	return ErrorResponse{
		Error: ErrorMessageResponse{
			Attribute: err,
		},
	}
}
