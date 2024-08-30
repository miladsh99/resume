package utills

import "project1/dto"

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
	ShortPass
	EmptyName
)
