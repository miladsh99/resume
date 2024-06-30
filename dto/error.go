package dto

type ErrorType int

type ErrorHandle struct {
	Type    ErrorType
	Message string
}
