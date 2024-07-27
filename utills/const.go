package utills

type RequestType string

const (
	Login          RequestType = "login"
	Register       RequestType = "register"
	AdminSchedules RequestType = "adminSchedules"
)
