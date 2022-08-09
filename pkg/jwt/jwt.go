package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"goapi/pkg/app"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"strings"
	"time"
)

var (
	TokenInvalid           error = errors.New("无效Token")
	TokenExpired           error = errors.New("Token已过期")
	TokenExpiredMaxRefresh error = errors.New("Token已过最大刷新时间")
	TokenMalformed         error = errors.New("Token格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

type JWTCustomClaims struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.StandardClaims
}

func NewJwt() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key", "goapi")),
		MaxRefresh: time.Duration(config.GetInt("jwt.max_refresh_time", 120)) * time.Minute,
	}
}

func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return nil, err
	}
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, TokenExpired
			}
		}
		return nil, TokenInvalid
	}
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return "", err
	}

	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}
	claims := token.Claims.(*JWTCustomClaims)
	// 用当前时间减去最大刷新时间，计算得到签名生成时间，如果真实的签名生成时间大于计算出来的签名生成时间，说明还没有超过最大刷新时间
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		claims.StandardClaims.ExpiresAt = jwt.getExpireTime()
		return jwt.createToken(*claims)
	}
	return "", TokenExpiredMaxRefresh
}

func (jwt *JWT) IsuseToken(id, name string) string {
	expireTime := jwt.getExpireTime()
	var claims = JWTCustomClaims{
		UserId:       id,
		UserName:     name,
		ExpireAtTime: expireTime,
		StandardClaims: jwtpkg.StandardClaims{
			Issuer:    config.GetString("app.name"),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
			NotBefore: time.Now().Unix(),
		},
	}
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", TokenInvalid
	}
	s := strings.SplitN(auth, " ", 2)
	if len(s) != 2 && s[0] != "Bearer" {
		return "", TokenMalformed
	}
	return s[1], nil
}

func (jwt *JWT) parseTokenString(token string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(token, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	t := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return t.SignedString(jwt.SignKey)
}

func (jwt *JWT) getExpireTime() int64 {
	now := app.TimenowInTimezone()
	var expireTime int64
	if app.IsDebug() {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return now.Add(expire).Unix()
}
