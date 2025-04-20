package user

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"my-realworld-go/common"
)

// 去除 "TOKEN " 前缀
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// 请求头的 Authorization 变量
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// 允许从多种来源提取 token
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// 修改当前用户上下文
func UpdateUserModelContext(c *gin.Context, userId uint) {
	var userModel UserModel
	if userId != 0 {
		db := common.GetDB()
		db.First(&userModel, userId)
	}
	c.Set("current_user_id", userId)
	c.Set("current_user_model", userModel)
}

// 身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 修改上下文
		UpdateUserModelContext(c, 0)
		// 从请求头 解析出 token
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		if err != nil {
			// 身份认证失败
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份认证失败，请登录！"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 根据 token 解析出用户Id、修改用户上下文
			currentUserId := uint(claims["id"].(float64))
			UpdateUserModelContext(c, currentUserId)
		} else {
			// 无效 token
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "无效的token！",
			})
			c.Abort()
			return
		}
	}
}
