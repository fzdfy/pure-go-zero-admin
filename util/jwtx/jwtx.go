package jwtx

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString, AccessSecret string) (jwt.MapClaims, error) {
	fn := func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil //校验字符串
	}

	result, err := jwt.Parse(tokenString, fn)
	if err != nil {
		return nil, err //signature is invalid or Token is expired
	}

	//解析存入的jwt信息
	resp := result.Claims.(jwt.MapClaims)

	return resp, nil
}
