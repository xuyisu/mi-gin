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
	groupApi.GET("order/pages", orderInfoAll)
	groupApi.GET("order/:orderNo", orderInfoOne)
	groupApi.POST("order/create", orderInfoCreate)
	groupApi.POST("order/pay", orderInfoPay)
}

//All
func orderInfoAll(c *gin.Context) {
	loginUser := getLoginUser(c)
	orderInfoReq := models.OrderInfo{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	var orderInfoVos []models.OrderInfoVo
	query.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	query.Order = "id desc"
	list, total, err := orderInfoReq.All(query)
	if total > lib.Zero {
		orderInfoVo := models.OrderInfoVo{}
		orderList := *list
		for _, orderRes := range orderList {
			orderInfoVo.OrderInfo = orderRes
			var orderDetailReq models.OrderDetail
			orderQuery := &models.PaginationQuery{}
			orderQuery.Limit = lib.DefaultLimit
			orderQuery.Where = fmt.Sprintf("delete_flag:0,user_id:%s,order_no:%s", strconv.FormatInt(int64(loginUser.Id), 10), orderRes.OrderNo)
			orderDetailList, total, _ := orderDetailReq.All(orderQuery)
			if total > lib.Zero {
				orderInfoVo.Details = *orderDetailList
			}
			orderInfoVos = append(orderInfoVos, orderInfoVo)
		}
	}
	jsonPage(c, orderInfoVos, total, query)
}

//One
func orderInfoOne(c *gin.Context) {
	loginUser := getLoginUser(c)
	var orderInfoReq models.OrderInfo
	orderNo, _ := parseParamOrderNo(c)
	orderInfoReq.OrderNo = orderNo
	orderInfoReq.DeleteFlag = lib.Zero
	orderInfoRes, _ := orderInfoReq.One()
	if orderInfoRes == nil {
		jsonError(c, "当前订单信息不存在!")
		return
	}
	orderInfoVo := models.OrderInfoVo{}
	orderInfoVo.OrderInfo = *orderInfoRes

	var orderDetailReq models.OrderDetail
	orderDetailReq.OrderNo = orderNo
	orderQuery := &models.PaginationQuery{}
	orderQuery.Limit = lib.DefaultLimit
	orderQuery.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	all, total, _ := orderDetailReq.All(orderQuery)
	if total > 0 {
		orderInfoVo.Details = *all
	}
	jsonData(c, orderInfoVo)
}

//Create
func orderInfoCreate(c *gin.Context) {
	loginUser := getLoginUser(c)
	var orderInfoReq models.OrderInfo
	err := c.ShouldBind(&orderInfoReq)
	if handleError(c, err) {
		return
	}
	var addressReq models.UserAddress
	addressReq.AddressId = orderInfoReq.AddressId
	addressReq.DeleteFlag = lib.Zero
	addressRes, _ := addressReq.One()
	if addressRes == nil {
		jsonError(c, "当前地址已不存在，请重新添加地址")
		return
	}
	//根据需要改造
	var orderNo = strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	totalOrderPrice, _ := decimal.NewFromString(".0")
	//查询购物车
	var cartReq models.Cart
	cartQuery := &models.PaginationQuery{}
	cartQuery.Limit = lib.DefaultLimit
	cartQuery.Where = fmt.Sprintf("delete_flag:0,user_id:%s", strconv.FormatInt(int64(loginUser.Id), 10))
	cartList, total, err := cartReq.All(cartQuery)
	if total == 0 {
		jsonError(c, "恭喜您的购物车已经被清空了，再加一车吧")
		return
	}
	carts := *cartList
	now := time.Now()
	for _, cart := range carts {
		//查询商品
		var productReq models.Product
		productReq.ProductId = cart.ProductId
		productReq.Status = lib.One
		productReq.DeleteFlag = lib.Zero
		productRes, _ := productReq.One()
		if productRes != nil && productRes.Stock <= lib.Zero {
			jsonError(c, "商品:"+cart.ProductName+" 已售尽,请选择其它产品")
			return
		}
		//订单明细
		orderDetail := models.OrderDetail{CurrentUnitPrice: productReq.Price}
		//判断活动
		var activityReq models.Activity
		activityRes, _ := activityReq.GetActivityByTime(time.Now())
		if activityRes != nil {
			orderDetail.ActivityId = activityRes.ActivityId
			orderDetail.ActivityName = activityRes.Name
			orderDetail.ActivityMainImage = activityRes.MainImage
		}
		orderDetail.OrderDetailNo = strconv.FormatInt(time.Now().UnixNano(), 10)
		orderDetail.OrderNo = orderNo
		orderDetail.ProductId = productRes.ProductId
		orderDetail.ProductMainImage = productRes.MainImage
		orderDetail.ProductName = productRes.Name
		orderDetail.Quantity = cart.Quantity
		orderDetail.Status = lib.PaymentStatueUnPay
		orderDetail.StatusDesc = lib.PaymentStatueUnPayDesc
		orderDetail.TotalPrice = cart.ProductTotalPrice
		orderDetail.UserId = loginUser.Id
		orderDetail.CreateUser = loginUser.Id
		totalOrderPrice = totalOrderPrice.Add(decimal.NewFromFloat(orderDetail.TotalPrice))
		orderDetail.Create()
		//更新购物车
		cart.DeleteFlag = lib.One
		cart.UpdateTime = &now
		cart.UpdateUser = loginUser.Id
		cart.Update()
		//设置订单状态记录
		statusRecordReq := models.OrderStatusRecord{
			OrderNo:       orderNo,
			OrderDetailNo: orderDetail.OrderDetailNo,
			ProductId:     orderDetail.ProductId,
			ProductName:   orderDetail.ProductName,
			Status:        orderDetail.Status,
			StatusDesc:    orderDetail.StatusDesc,
		}
		statusRecordReq.Create()
	}
	//订单主表
	orderInfoReq = models.OrderInfo{
		OrderNo:         orderNo,
		AddressId:       addressRes.AddressId,
		Area:            addressRes.Area,
		City:            addressRes.City,
		Payment:         totalOrderPrice.InexactFloat64(),
		PaymentType:     lib.PaymentTypeOnline,
		PaymentTypeDesc: lib.PaymentTypeOnlineDesc,
		PostalCode:      addressRes.PostalCode,
		Province:        addressRes.Province,
		ReceiveName:     addressRes.ReceiveName,
		ReceivePhone:    addressRes.ReceivePhone,
		Street:          addressRes.Street,
		Status:          lib.PaymentStatueUnPay,
		StatusDesc:      lib.PaymentStatueUnPayDesc,
		CreateUser:      loginUser.Id,
		UserId:          loginUser.Id,
	}

	err = orderInfoReq.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, orderNo)
}

//Update
func orderInfoPay(c *gin.Context) {
	loginUser := getLoginUser(c)
	var pay models.PayReq
	err := c.ShouldBind(&pay)
	if handleError(c, err) {
		return
	}
	var orderInfoReq models.OrderInfo
	orderInfoReq.OrderNo = pay.OrderNo
	orderInfoReq.DeleteFlag = lib.Zero
	orderInfoRes, _ := orderInfoReq.One()
	if orderInfoRes != nil {
		if orderInfoRes.UserId != loginUser.Id {
			jsonError(c, "您无权查询他人订单")
			return
		}
		if orderInfoRes.Status != lib.PaymentStatueUnPay {
			jsonError(c, "您没有待支付的订单")
			return
		}
		orderInfoRes.Status = lib.PaymentStatuePay
		orderInfoRes.StatusDesc = lib.PaymentStatuePayDesc
		date := time.Now()
		orderInfoRes.PaymentTime = &date
		err = orderInfoRes.Update()
		if handleError(c, err) {
			return
		}
		//查询订单明细
		var orderDetailReq models.OrderDetail
		orderQuery := &models.PaginationQuery{}
		orderQuery.Limit = lib.DefaultLimit
		orderQuery.Where = fmt.Sprintf("delete_flag:0,user_id:%s,order_no:%s", strconv.FormatInt(int64(loginUser.Id), 10), orderInfoRes.OrderNo)
		all, total, _ := orderDetailReq.All(orderQuery)
		if total > 0 {
			orderDetailList := *all
			for _, detail := range orderDetailList {
				detail.Status = lib.PaymentStatuePay
				detail.StatusDesc = lib.PaymentStatuePayDesc
				detail.UpdateUser = loginUser.Id
				err := detail.Update()
				if handleError(c, err) {
					return
				}
				//设置订单状态记录
				statusRecord := models.OrderStatusRecord{
					OrderNo:       detail.OrderNo,
					OrderDetailNo: detail.OrderDetailNo,
					ProductId:     detail.ProductId,
					ProductName:   detail.ProductName,
					Status:        detail.Status,
					StatusDesc:    detail.StatusDesc,
				}
				err = statusRecord.Create()
				if handleError(c, err) {
					return
				}
			}
		}
	}
	jsonSuccess(c)
}
