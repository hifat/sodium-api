package response

import "github.com/hifat/sodium-api/internal/utils/ernos"

type ErrorResponse struct {
	Error ernos.Ernos `json:"error"`
}

type ErrorMessageResponse struct {
	Message   string `json:"message,omitempty"`
	Code      string `json:"code,omitempty"`
	Attribute any    `json:"attribute,omitempty"`
}

type SuccesResponse struct {
	Item    any    `json:"item,omitempty"`
	Items   []any  `json:"items,omitempty"`
	Total   int    `json:"total,omitempty"`
	Message string `json:"message,omitempty"`
}

func HandleErr(err any) ErrorResponse {
	if _, ok := err.(error); ok {
		return ErrorResponse{
			Error: ernos.Ernos{
				Message: err.(error).Error(),
			},
		}
	}

	return ErrorResponse{
		Error: ernos.Ernos{
			Attribute: err,
		},
	}
}
