package models

import "admin/entity"

const (
	Addres      = "127.0.0.1"
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
