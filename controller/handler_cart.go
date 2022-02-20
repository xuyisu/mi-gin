package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go-gin/models"
	"go-gin/pkg/lib"
	"strconv"
	"time"
)

func init() {
	groupApi.GET("cart/list", cartAll)
	groupApi.GET("cart/:id", cartOne)
	groupApi.PUT("cart/:productId", cartUpdate)
	groupApi.DELETE("cart/:productId", cartDelete)
	groupApi.GET("cart/sum", cartCount)
	groupApi.POST("cart/add", cartAdd)
	groupApi.PUT("cart/selectAll", selectAll)
	groupApi.PUT("cart/unSelectAll", unSelectAll)
}

//All
func cartAll(c *gin.Context) {
	loginUser := getLoginUser(c)
	mdl := models.Cart{}
	query := &models.PaginationQuery{}
	query.Limit = lib.DefaultLimit
	query.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, _, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	totalQuantity := lib.Zero
	totalPrice, _ := decimal.NewFromString(".0")
	selectAll := true
	cartList := *list
	for _, v := range cartList {
		//计算价格
		if v.Selected < lib.One {
			selectAll = false
		}
		totalQuantity += v.Quantity
		totalPrice = totalPrice.Add(decimal.NewFromFloat(v.ProductTotalPrice))
	}
	resp := models.CartResp{CartTotalPrice: totalPrice.String(), CartTotalQuantity: totalQuantity, SelectedAll: selectAll, CartProductList: list}

	jsonData(c, resp)
}

//One
func cartOne(c *gin.Context) {
	var mdl models.Cart
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

//Update
func cartUpdate(c *gin.Context) {

	id, err := parseParamProductId(c)
	if handleError(c, err) {
		return
	}
	//查询商品是否已下架
	var product models.Product
	product.ProductId = id
	product.DeleteFlag = 0
	product.Status = 1
	productRes, err := product.One()
	if err != nil {
		jsonError(c, "当前商品已下架或删除")
		return
	}

	//查询当前购物车的产品信息
	var cartProduct models.Cart
	cartProduct.ProductId = id
	cartProduct.DeleteFlag = 0
	cartProductRes, err := cartProduct.One()

	cartReq := models.CartReq{}
	err = c.ShouldBind(&cartReq)
	if handleError(c, err) {
		return
	}
	if cartReq.Type != nil {
		if int8(cartReq.Type.(float64)) == lib.One {
			cartProductRes.Quantity = cartProductRes.Quantity + lib.One
		} else {
			if cartProductRes.Quantity <= lib.One {
				jsonError(c, "不能再减了,要减没了")
				return
			}
			cartProductRes.Quantity = cartProductRes.Quantity - lib.One
		}
	}
	if cartReq.Selected != nil {
		cartProductRes.Selected = int8(cartReq.Selected.(float64))
		if cartProductRes.Selected == 0 {
			cartProductRes.Selected = -1
		}
	}
	totalQuantity := decimal.NewFromInt(int64(cartProductRes.Quantity))
	productPrice := decimal.NewFromFloat(productRes.Price)
	cartProductRes.ProductTotalPrice = totalQuantity.Mul(productPrice).InexactFloat64()

	err = cartProductRes.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Delete
func cartDelete(c *gin.Context) {
	var mdl models.Cart
	id, err := parseParamProductId(c)
	if handleError(c, err) {
		return
	}

	mdl.ProductId = id
	cartRes, err := mdl.One()
	if err != nil {
		jsonError(c, "该商品不在购物车")
		return
	}
	cartRes.DeleteFlag = lib.One
	err = cartRes.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//查询购物车数量
func cartCount(c *gin.Context) {
	var mdl models.Cart
	count := mdl.GetCartCount(getLoginUser(c).Id)
	jsonData(c, count)
}

// 购物车添加
func cartAdd(c *gin.Context) {
	loginUser := getLoginUser(c)
	var v models.Cart
	err := c.ShouldBind(&v)
	if handleError(c, err) {
		return
	}

	//查询商品是否已下架
	var product models.Product
	product.ProductId = v.ProductId
	product.DeleteFlag = 0
	product.Status = 1
	productRes, err := product.One()
	if err != nil {
		jsonError(c, "当前商品已下架或删除")
		return
	}

	//查询当前购物车的产品信息
	var cartProduct models.Cart
	cartProduct.ProductId = v.ProductId
	cartProductRes, err := cartProduct.One()
	if err != nil {
		v.ProductName = productRes.Name
		v.ProductSubtitle = productRes.SubTitle
		v.ProductUnitPrice = productRes.Price
		v.ProductMainImage = productRes.MainImage
		v.Quantity = lib.One
		totalPrice := decimal.NewFromInt(int64(v.Quantity))
		productPrice := decimal.NewFromFloat(productRes.Price)
		v.ProductTotalPrice = totalPrice.Mul(productPrice).InexactFloat64()

		//查询活动
		var activity models.Activity
		activityRes, err := activity.GetActivityByTime(time.Now())
		if err == nil {
			v.ActivityName = activityRes.Name
			v.ActivityId = activityRes.ActivityId
		}
		v.Selected = lib.One
		v.CreateUser = loginUser.Id
		v.UserId = loginUser.Id
		v.UpdateUser = loginUser.Id
		now := time.Now()
		v.CreateTime = &now
		v.UpdateTime = &now
		err = v.Create()
		if handleError(c, err) {
			return
		}
	} else {
		now := time.Now()
		cartProductRes.UpdateTime = &now
		cartProductRes.Quantity = cartProductRes.Quantity + lib.One
		totalPrice := decimal.NewFromInt(int64(cartProductRes.Quantity))
		productPrice := decimal.NewFromFloat(productRes.Price)
		cartProductRes.ProductTotalPrice = totalPrice.Mul(productPrice).InexactFloat64()
		err := cartProductRes.Update()
		if handleError(c, err) {
			return
		}
	}
	jsonData(c, v.GetCartCount(loginUser.Id))

}

func unSelectAll(c *gin.Context) {
	loginUser := getLoginUser(c)
	mdl := models.Cart{}
	mdl.UserId = loginUser.Id
	mdl.DeleteFlag = lib.Zero
	query := &models.PaginationQuery{}
	query.Limit = lib.DefaultLimit
	query.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	query.Order = "id desc"
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	cartList, total, _ := mdl.All(query)
	if total > 0 {
		carts := *cartList
		for _, cart := range carts {
			cart.Selected = lib.FOne
			cart.UpdateUser = loginUser.Id
			now := time.Now()
			cart.UpdateTime = &now
			cart.Update()
		}

	}
	jsonSuccess(c)
}

func selectAll(c *gin.Context) {
	loginUser := getLoginUser(c)
	mdl := models.Cart{}
	mdl.UserId = loginUser.Id
	mdl.DeleteFlag = lib.Zero
	query := &models.PaginationQuery{}
	query.Limit = lib.DefaultLimit
	query.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	cartList, total, _ := mdl.All(query)
	if total > 0 {
		carts := *cartList
		for _, cart := range carts {
			cart.Selected = lib.One
			cart.UpdateUser = loginUser.Id
			now := time.Now()
			cart.UpdateTime = &now
			cart.Update()
		}

	}
	jsonSuccess(c)
}
