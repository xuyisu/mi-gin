package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"go-gin/pkg/lib"
	"strconv"
)

func init() {
	groupApi.GET("product/pages", productAll)
	groupApi.GET("product/:productId", productOne)
	groupApi.POST("product", productCreate)
	groupApi.PATCH("product", productUpdate)
	groupApi.DELETE("product/:id", productDelete)
}

//All
func productAll(c *gin.Context) {
	categoryId := c.Query("categoryId")
	size := c.Query("size")
	categoryIdUint, _ := strconv.ParseUint(categoryId, 10, 64)
	sizeUint, _ := strconv.ParseUint(size, 10, 64)
	mdl := models.Product{}
	mdl.CategoryId = categoryIdUint
	query := &models.PaginationQuery{}

	query.Limit = uint(sizeUint)
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	query.Where = "delete_flag:0"
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPage(c, list, total, query)
}

//One
func productOne(c *gin.Context) {
	var mdl models.Product
	id, err := parseParamProductId(c)
	if handleError(c, err) {
		return
	}
	mdl.ProductId = id
	mdl.DeleteFlag = lib.Zero
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func productCreate(c *gin.Context) {
	var mdl models.Product
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
func productUpdate(c *gin.Context) {
	var mdl models.Product
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
func productDelete(c *gin.Context) {
	var mdl models.Product
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
