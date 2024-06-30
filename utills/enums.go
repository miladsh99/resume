package utills

import "resume/dto"

const (
	DB dto.ErrorType = iota + 1
	FailedGetDataFromDB
	NotExistEmail
	ExistEmail
	InvalidEmail
	InvalidMethod
	Body
	Unmarshal
	Pass
	Other
	TokenErr
	InvalidToken
	NotReceiveToken
)
