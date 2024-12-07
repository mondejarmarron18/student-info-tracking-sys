package utils

type ErrorMessage struct {
	NotFound            string
	InternalServerError string
	InvalidRequest      string
	BadRequest          string
}

func NewErrorMessage() ErrorMessage {
	return ErrorMessage{
		NotFound:            "We couldn't find what you're looking for. Please check the information and try again.",
		InternalServerError: "Something went wrong on our end. Please try again later.",
		InvalidRequest:      "The information you provided is incorrect. Please review and try again.",
		BadRequest:          "There was an issue with your request. Please check and try again.",
	}
}
