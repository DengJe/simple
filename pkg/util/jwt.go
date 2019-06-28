package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

const (
	ACCESS  = "ACCESS"
	REFRESH = "REFRESH"
)

var (
	JwtInstance *Jwt
)

type Jwt struct {
	Key                    []byte
	AccessTokenExpireTime  time.Duration
	RefreshTokenExpireTime time.Duration
}

func init() {
	JwtInstance = &Jwt{
		Key:                    []byte("\x88W\xf09\x91\x07\x98\x89\x87\x96\xa0A\xc68\xf9\xecJJU\x17\xc5V\xbe\x8b\xef\xd7\xd8\xd3\xe6\x95*4"),
		AccessTokenExpireTime:  time.Hour * time.Duration(2),
		RefreshTokenExpireTime: time.Hour * 24 * 7,
	}
}

func GenerateTokens(identify string) (string, string, error) {
	var (
		accessToken, refreshToken string
		err                       error
	)
	accessToken, err = GenerateAccessToken(identify)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = GenerateRefreshToken(identify)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *Jwt) GenerateAccessToken(identify string) (string, error) {
	return j.generateToken(j.AccessTokenExpireTime, ACCESS, identify)
}


func GenerateAccessToken(identify string) (string, error) {
	return JwtInstance.generateToken(JwtInstance.AccessTokenExpireTime, ACCESS, identify)
}

func GenerateRefreshToken(identify string) (string, error) {
	return JwtInstance.generateToken(JwtInstance.AccessTokenExpireTime, REFRESH, identify)
}

func (j *Jwt) generateToken(timeOut time.Duration, tokenType, identify string) (string, error) {
	key := []byte(JwtInstance.Key)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expire := time.Now().Add(timeOut)
	// exp express token expire time
	claims["exp"] = expire.Unix()
	claims["type"] = tokenType
	// identi
	// get the exp timefy express username
	claims["identify"] = identify
	//exp := int64(claims["exp"].(float64))
	return token.SignedString(key)
}

func VerifyAccessTokenInHeader(authHeader string) (jwt.MapClaims, error) {
	var tokenStr string

	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("请填写正确的Bearer字段")
	}
	tokenStr = parts[1]
	return JwtInstance.verifyToken(tokenStr, ACCESS)
}

func VerifyRefreshToken(authHeader string) (jwt.MapClaims, error) {
	return JwtInstance.verifyToken(authHeader, REFRESH)
}

func (j *Jwt) verifyToken(tokenStr, tokenType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		//if jwt.SigningMethodHS256 != t.Method {
		//	return nil, errors.New("token解析错误")
		//}
		claims := t.Claims.(jwt.MapClaims)
		innerTokenType := claims["type"].(string)
		if innerTokenType != tokenType {
			return nil, errors.New("令牌类型错误")
		}
		return j.Key, nil
	})
	if token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("令牌损坏")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.New("令牌过期")
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
