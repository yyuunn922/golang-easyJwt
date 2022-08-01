package test

import (
	"testing"

	easyjwt "github.com/madoleeee/golang-easyJwt"
)

type ttempStruct struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func TestChecker(t *testing.T) {
	temp := easyjwt.New()
	temp.SecretKey = "secretKey"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjozLCJuYW1lIjoiMTIzMTIzIn0sImV4cCI6MTY1ODE5ODA4NH0.5HWxYmOlgC2zYatcDiIIu3qmRq5xhgAInHdqtb0PhLA"

	var aa ttempStruct

	ba, err := temp.Checker(token, &aa)
	if err != nil {
		panic(err.Error())
	}
	t.Log(ba)
	t.Log(aa)
}
