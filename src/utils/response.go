package utils

// import "golang.org/x/text/message"

type SuccessResponseStruct struct {
	Meta Meta        `json:meta`
	Data interface{} `json:data`
}

type ErrorResponseStuct struct {
	Meta Meta        `json:meta`
	Error interface{} `json:error`
}

type Meta struct {
	Message string `json:message`
	Status  string "json:status"
}

func SuccessResponse(message string, data interface{}) SuccessResponseStruct {
	return SuccessResponseStruct{
		Meta: Meta{
			Message: message,
			Status:  "success",
		},
		Data: data,
	}
}


func ErrorResponse(message string, err interface {}) ErrorResponseStuct {
	return ErrorResponseStuct{
		Meta: Meta{
			Message: message,
			Status:  "error",
		},
		Error : err,
	}
}