package utills
//
//import (
//	"encoding/json"
//	"golang.org/x/crypto/bcrypt"
//	"io"
//	"net/http"
//	"project1/dto"
//	"project1/entity"
//)
//
//type Response interface {
//	Register2()
//	Login2()
//	Schedules()
//}
//
//func ReadRequest(ReType RequestType, r *http.Request) (Response, *dto.ErrorHandle) {
//	data, rErr := io.ReadAll(r.Body)
//	if rErr != nil {
//		return nil, &dto.ErrorHandle{Type: Body}
//	}
//
//	switch ReType {
//	case Register:
//		return RegisterRead(data)
//	case Login:
//		return LoginRead(data)
//	case AdminSchedules:
//		return AdminRead(data)
//	default:
//		return nil, &dto.ErrorHandle{Type: Other}
//	}
//
//}
//
//func RegisterRead(data []byte) (*entity.User, *dto.ErrorHandle) {
//
//	var req = dto.RegisterRequest{}
//	mErr := json.Unmarshal(data, &req)
//	if mErr != nil {
//		return nil, &dto.ErrorHandle{Type: Unmarshal}
//	}
//	user := entity.User{
//		Name:  req.Name,
//		Email: req.Email,
//	}
//	pass, pErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	if pErr != nil {
//		return nil, &dto.ErrorHandle{Type: Other}
//	}
//	user.SetPassword(string(pass))
//
//	return &user, nil
//
//}
//
//func LoginRead(data []byte) (*entity.User, *dto.ErrorHandle) {
//
//	var req = dto.LoginRequest{}
//	mErr := json.Unmarshal(data, &req)
//	if mErr != nil {
//		return nil, &dto.ErrorHandle{Type: Unmarshal}
//	}
//	user := entity.User{
//		Email: req.Email,
//	}
//	user.SetPassword(req.Password)
//
//	return &user, nil
//
//}
//
//func AdminRead(data []byte) (*dto.AdminSchedulesRequest, *dto.ErrorHandle) {
//
//	var req = dto.AdminSchedulesRequest{}
//	mErr := json.Unmarshal(data, &req)
//	if mErr != nil {
//		return nil, &dto.ErrorHandle{Type: Unmarshal}
//	}
//	ad := dto.AdminSchedulesRequest{
//		AdminID:     req.AdminID,
//		EventName:   req.EventName,
//		StartTime:   req.StartTime,
//		EndTime:     req.EndTime,
//		Description: req.Description,
//	}
//
//	return &ad, nil
//
//}
