package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"go-gin/pkg/lib"
	"go-gin/pkg/utils"
	"strconv"
	"time"
)

func init() {
	groupApi.GET("address/pages", userAddressAll)
	groupApi.GET("address/:addressId", userAddressOne)
	groupApi.POST("address/add", userAddressCreate)
	groupApi.PUT("address/:addressId", userAddressUpdate)
	groupApi.DELETE("address/:addressId", userAddressDelete)
}

//All
func userAddressAll(c *gin.Context) {
	size := c.Query("size")
	sizeUint, _ := strconv.ParseUint(size, 10, 64)
	mdl := models.UserAddress{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	query.Where = fmt.Sprintf("delete_flag:0,create_user:%s", strconv.Itoa(lib.DeFaultUser))
	query.Order = "id desc"
	query.Limit = uint(sizeUint)
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPage(c, list, total, query)
}

//One
func userAddressOne(c *gin.Context) {
	var addressReq models.UserAddress
	id, err := parseParamAddressId(c)
	if handleError(c, err) {
		return
	}
	addressReq.AddressId = id
	addressReq.DeleteFlag = lib.Zero
	data, err := addressReq.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func userAddressCreate(c *gin.Context) {
	var addressReq models.UserAddress
	err := c.ShouldBind(&addressReq)
	if handleError(c, err) {
		return
	}

	//当你分布式的部署你的服务的时候，这个NewWorker的参数记录不同的node配置的值应该不一样
	worker, _ := utils.NewWorker(1)
	addressReq.AddressId = uint64(worker.GetId())
	now := time.Now()
	addressReq.UpdateTime = &now
	addressReq.UpdateUser = lib.DeFaultUser
	addressReq.CreateTime = &now
	addressReq.CreateUser = lib.DeFaultUser

	err = addressReq.Create()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Update
func userAddressUpdate(c *gin.Context) {
	id, err := parseParamAddressId(c)
	if handleError(c, err) {
		return
	}
	var param models.UserAddress
	err = c.ShouldBind(&param)
	if handleError(c, err) {
		return
	}

	var addressReq models.UserAddress
	addressReq.AddressId = id
	addressReq.DeleteFlag = lib.Zero
	addressRes, _ := addressReq.One()
	if addressRes == nil {
		jsonError(c, "不存在地址信息")
		return
	}
	param.Id = addressRes.Id
	err = param.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Delete
func userAddressDelete(c *gin.Context) {
	id, err := parseParamAddressId(c)
	if handleError(c, err) {
		return
	}
	var addressReq models.UserAddress
	addressReq.AddressId = id
	addressReq.DeleteFlag = lib.Zero
	addressRes, _ := addressReq.One()
	if addressRes == nil {
		jsonError(c, "不存在地址信息")
		return
	}
	addressRes.DeleteFlag = lib.One
	err = addressRes.Update()
	jsonSuccess(c)
}
