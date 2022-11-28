package models

import "github.com/assyatier21/admin-deall-technical-test/entity"

const (
	Addres      = "0.0.0.0"
	Port        = "8800"
	AccessToken = "access-token"
)

func SetResponse(Status int, Message string, Data []interface{}) (res entity.Response) {
	res = entity.Response{
		Status:  Status,
		Message: Message,
		Data:    Data,
	}
	return
}
