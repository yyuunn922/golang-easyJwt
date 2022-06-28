package test

import (
	"testing"

	easyjwt "github.com/madoleeee/golang-easyJwt"
)

func TestDefaultGetToken(t *testing.T) {
	temp := easyjwt.New()
	temp.Data = struct {
		Id uint `json:"id"`
	}{
		Id: 1,
	}
	temp.SecretKey = "secretKey"

	t.Log(temp.GetAccessToken())
	t.Log(temp.GetRefreshToken())
	t.Log(temp.GetAllToken())
}
