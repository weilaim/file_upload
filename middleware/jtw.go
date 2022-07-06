package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/util/errmsg"
)

var JWTKEY = []byte(os.Getenv("JWTKEY"))

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

//生产token
func SetToken(username string) (string, int){
	//有效时间
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	res := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := res.SignedString(JWTKEY)
	if err != nil {
		return "", 404
	}

	return token, 200
}

//验证token
func CheckToken(token string) (*MyClaims, int) {
	settoken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, 500
	}

	if key, _ := settoken.Claims.(*MyClaims); settoken.Valid {
		return key, 200
	} else {
		return nil, 404
	}

}

//jwt中间件
// var code int

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCSE
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIT
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return

		}
		checktoken := strings.SplitN(tokenHerder, " ", 2)
		if len(checktoken) != 2 && checktoken[0] != "Bearer" {
			code = errmsg.ERROR_PASSWORD_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, Tcode := CheckToken(checktoken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})

			c.Abort()
			return
		}
		
		//验证token过期时间
		tempNow := time.Now().Unix()
		if tempNow > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusFound, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})

			c.Abort()
			return
		}

		c.Set("username", key.UserName)
		c.Next()
	}
}
