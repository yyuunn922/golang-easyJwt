토큰생성 
```golang

type _tempStruct struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}


temp := easyjwt.New()
temp.Data = _tempStruct{Id: 3, Name: "user_name"}
temp.SecretKey = "secretKey"

t.Log(temp.GetAccessToken()) // accessToken을 리턴합니다
t.Log(temp.GetRefreshToken()) // refreshToken을 리턴합니다
t.Log(temp.GetAllToken()) // accessToken 및 refreshToken 리턴합니다
t.Log(temp.GetAllToken().AccessToken)
t.Log(temp.GetAllToken().RefreshToken)

log....
getToken_test.go:22: eyJhbGciOiJIUzI1NiIsInR5cCI6Ik....
getToken_test.go:23: eyJhbGciOiJIUzI1NiIsInR5cCI6Ik....
getToken_test.go:24: {eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.......M}
getToken_test.go:25: eyJhbGciOiJIUzI1NiIsInR5cCI6Ik....
getToken_test.go:26: eyJhbGciOiJIUzI1NiIsInR5cCI6Ik....
```
토큰검증
```golang
var result _tempstruct
temp := easyjwt.New()
temp.SecretKey = "secretKey"

tokenData, err := temp.Checker(token, &result)

if err != nil {
		panic(err.Error())
}

t.log(tokenData)
t.log(result)



log....
tokenCheck_test.go:25: &{eyJhbGciOiJIUzI1NiIsInR5cCI6Ik.... 0x14000136198 map[alg:HS256 typ:JWT] 0x1400016e000 5HWxYmOlgC2zYa....AInHdqtb0PhLA true}
tokenCheck_test.go:26: {3 user_name}

```
