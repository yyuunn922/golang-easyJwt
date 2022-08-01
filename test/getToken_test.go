package test

import (
	"testing"

	easyjwt "github.com/madoleeee/golang-easyJwt"
)

var accessToken string

type _tempStruct struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func TestDefaultGetToken(t *testing.T) {

	temp := easyjwt.New()
	temp.Data = _tempStruct{Id: 3, Name: "123123"}
	temp.SecretKey = "secretKey"

	t.Log(temp.GetAccessToken())
	t.Log(temp.GetRefreshToken())
	t.Log(temp.GetAllToken())
	t.Log(temp.GetAllToken().AccessToken)
	t.Log(temp.GetAllToken().RefreshToken)

	accessToken = temp.GetAccessToken()
}
