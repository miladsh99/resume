package utills

const (
	DB ErrorType = iota
	NotExist
	Exist
	Invalid
	InvalidMethod
	Body
	Unmarshal
	Pass
	Other
)
