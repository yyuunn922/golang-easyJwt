package easyjwt

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
)

func (t EasyJwtData) Checker(accessToken string, claimData interface{}) (*jwt.Token, error) {

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(t.SecretKey), nil
	}

	claims := TokenData{}
	tok, err := jwt.ParseWithClaims(accessToken, &claims, key)
	if err != nil {
		return nil, err
	}

	outByte, err := json.Marshal(tok.Claims)
	if err != nil {
		return nil, err
	}

	var aa TokenData
	json.Unmarshal(outByte, &aa)

	bb, err := json.Marshal(aa.Data)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bb, claimData)

	return tok, nil

}
