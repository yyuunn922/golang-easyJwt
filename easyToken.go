package easyjwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AlgorihmType string

const (
	HS256 AlgorihmType = "HS256"
	HS384 AlgorihmType = "HS384"
	HS512 AlgorihmType = "HS512"
)

// 이지토큰 데이터
type EasyJwtData struct {
	AlgorihmType     AlgorihmType
	AccessTokenTime  int
	RefreshTokenTime int
	Data             interface{} `json:"data"`
	SecretKey        string
}

// 이지데이터 스트럭쳐를 반환합니다
func New() EasyJwtData {
	return EasyJwtData{
		AlgorihmType:     HS256,
		AccessTokenTime:  10,
		RefreshTokenTime: 20,
	}
}

// 이지데이터 토큰 스트럭쳐
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// 토큰 데이터
type TokenData struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

// accessToken을 반환합니다
func (t EasyJwtData) GetAccessToken() string {
	if t.SecretKey == "" || t.Data == nil {
		panic("클레임 및 보안키를 확인해주세요")
	}

	out := TokenData{
		Data: t.Data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(t.AccessTokenTime)).Unix(),
		},
	}
	atoken := jwt.NewWithClaims(algorihmGet(t.AlgorihmType), &out)
	signedAuthToken, err := atoken.SignedString([]byte(t.SecretKey))
	if err != nil {
		panic("토큰 생성중 에러")
	}
	return signedAuthToken
}

// refreshToken을 반환합니다
func (t EasyJwtData) GetRefreshToken() string {
	if t.SecretKey == "" || t.Data == nil {
		panic("클레임 및 보안키를 확인해주세요")
	}
	out := TokenData{
		Data: t.Data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add((time.Hour * 24) * time.Duration(t.RefreshTokenTime)).Unix(),
		},
	}
	atoken := jwt.NewWithClaims(algorihmGet(t.AlgorihmType), &out)
	signedAuthToken, err := atoken.SignedString([]byte(t.SecretKey))
	if err != nil {
		panic("토큰 생성중 에러")
	}
	return signedAuthToken
}

// accessToken 및 refreshToken을 반환 합니다
func (t EasyJwtData) GetAllToken() Token {
	var out Token

	out.AccessToken = t.GetAccessToken()
	out.RefreshToken = t.GetRefreshToken()

	return out
}

// 알고리즘을 불러옵니다
func algorihmGet(a AlgorihmType) *jwt.SigningMethodHMAC {
	switch a {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	default:
		return jwt.SigningMethodHS256
	}
}
