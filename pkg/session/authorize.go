package session

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-gin/models"
	"go-gin/pkg/lib"
	"go-gin/pkg/utils"
	"net/http"
	"strings"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		logrus.Info("请求地址:" + path)

		hasSuffix := strings.HasPrefix(path, "/api/cart/") ||
			strings.HasPrefix(path, "/api/order/") ||
			strings.HasPrefix(path, "/api/address/") ||
			strings.HasPrefix(path, "/api/user/getUser")

		if hasSuffix {
			authorization := c.GetHeader(lib.Authorization)
			if authorization == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": lib.NotLogin, "msg": "请重新登录"})
				return
			} else {
				userRedis := Get(lib.UserLoginToken + authorization)
				if userRedis == "" {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": lib.NotLogin, "msg": "请重新登录"})
					return
				} else {
					sessionExpire := viper.GetInt("redis.session_expire")
					Set(lib.UserLoginToken+authorization, userRedis, sessionExpire)
					loginUser := models.LoginUser{}
					utils.JsonToObject(userRedis, &loginUser)
					GlobalMap.Store(authorization, loginUser)
					c.Next()
				}
			}

		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Next()
		}
	}
}
