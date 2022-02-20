package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin/models"
	"go-gin/pkg/lib"
	"go-gin/pkg/session"
	"go-gin/pkg/utils"
)

func init() {
	groupApi.GET("user", userAll)
	groupApi.GET("user/:id", userOne)
	groupApi.POST("user", userCreate)
	groupApi.PUT("user", userUpdate)
	groupApi.DELETE("user/:id", userDelete)
	groupApi.POST("user/login", login)
	groupApi.GET("user/getUser", getUser)
	groupApi.POST("user/logout", logout)
}

//All
func userAll(c *gin.Context) {
	mdl := models.User{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	query.Where = "delete_flag:0"
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//One
func userOne(c *gin.Context) {
	var mdl models.User
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func userCreate(c *gin.Context) {
	var mdl models.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

//Update
func userUpdate(c *gin.Context) {
	var mdl models.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Delete
func userDelete(c *gin.Context) {
	var mdl models.User
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//login
func login(c *gin.Context) {
	var loginReq models.LoginReq
	err := c.ShouldBind(&loginReq)
	if handleError(c, err) {
		return
	}

	if loginReq.UserName == "" || loginReq.Password == "" {
		jsonError(c, "用户名密码必填")
		return
	}
	var userReq models.User
	userReq.UserName = loginReq.UserName
	userReq.Password = utils.Md5(loginReq.Password)
	userReq.DeleteFlag = lib.Zero
	userReq.Status = lib.One
	userRes, _ := userReq.One()
	if userRes == nil {
		jsonError(c, "登录失败,请检查用户名密码")
		return
	}
	loginUser := models.LoginUser{
		Id:       userRes.Id,
		UserName: userRes.UserName,
		Email:    userRes.Email,
		Phone:    userRes.Phone,
	}
	//生成令牌
	token := utils.BuildToken()
	userMap := make(map[string]interface{})
	userMap[lib.Authorization] = token
	userMap["userInfo"] = loginUser

	//设置缓存
	marshal, _ := json.Marshal(loginUser)
	session.Set(lib.UserLoginToken+token, marshal, viper.GetInt("redis.session_expire"))

	jsonData(c, userMap)

}

//getUser
func getUser(c *gin.Context) {
	userRes, _ := session.GlobalMap.Load(c.GetHeader(lib.Authorization))
	jsonData(c, userRes)
}

func logout(c *gin.Context) {
	authorization := c.GetHeader(lib.Authorization)
	if authorization != "" {
		session.Delete(lib.UserLoginToken + authorization)
		session.GlobalMap.Delete(authorization)
	}
	jsonSuccess(c)
}
