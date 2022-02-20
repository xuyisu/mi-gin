package controller

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"go-gin/pkg/lib"
	"strconv"
	"time"
)

func jsonError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, gin.H{"code": lib.Error, "msg": msg})
}
func jsonData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"code": lib.Success, "data": data})
}
func jsonPagination(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
	c.JSON(200, gin.H{"code": lib.Success, "data": list, "total": total, "offset": query.Offset, "limit": query.Limit})
}

func jsonPage(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
	page := Page{PageNo: query.Offset, PageSize: query.Limit, TotalCount: total, Records: list}
	jsonData(c, page)
}

func jsonSuccess(c *gin.Context) {
	c.JSON(200, gin.H{"code": lib.Success, "msg": "操作成功"})
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func parseParamID(c *gin.Context) (uint64, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return parseId, nil
}

func parseParamOrderNo(c *gin.Context) (string, error) {
	orderNo := c.Param("orderNo")
	return orderNo, nil
}

func parseParamProductId(c *gin.Context) (uint64, error) {
	id := c.Param("productId")
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return parseId, nil
}

func parseParamAddressId(c *gin.Context) (uint64, error) {
	id := c.Param("addressId")
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return parseId, nil
}

func enableCorsMiddleware() {
	//TODO:: customize your own CORS
	//https://github.com/gin-contrib/cors
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //https://foo.com
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, //enable cookie
		AllowOriginFunc: func(origin string) bool {
			return true
			//return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, //cache options result decrease request lag
	}))
}

//分页返回数据
type Page struct {
	PageNo     uint        `json:"current"`
	PageSize   uint        `json:"size"`
	TotalCount uint        `json:"total"`
	Records    interface{} `json:"records"`
}
