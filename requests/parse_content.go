package requests

import (
	"net/url"

	"github.com/bububa/dataoke-go/core"
)

// ParseContentRequest 淘系万能解析 API Request
type ParseContentRequest struct {
	// Content 包含淘口令、链接的文本。优先解析淘口令，再按序解析每个链接，直至解出有效信息。如果淘口令失效或者不支持的类型的情况，会按顺序解析链接。如果存在解析失败，请再试一次
	Content string `json:"content,omitempty"`
}

// Values implement Request interface
func (r ParseContentRequest) Values(values url.Values) {
	values.Set("content", r.Content)
}

// Url implement Request interface
func (r ParseContentRequest) Url() string {
	return "tb-service/parse-content"
}

// ParseContentResult 解析结果
type ParseContentResult struct {
	// GoodsID 淘宝商品ID
	GoodsID string `json:"goodsId,omitempty"`
	// OriginURL 链接
	OriginURL string `json:"originUrl,omitempty"`
	// OriginType 链接中的信息类型
	OriginType string `json:"originType,omitempty"`
	// OriginInfo 链接中的信息
	OriginInfo *OriginInfo `json:"originInfo,omitempty"`
	// ItemID 当dataType=goods时，标识商品ID；当dataType=activity时，标识活动会场id；（10/27新增字段）
	ItemID string `json:"itemId,omitempty"`
	// ItemName 当dataType=goods时，标识商品名称；当dataType=activity时，标识活动会场名称；（10/27新增字段
	ItemName string `json:"itemName,omitempty"`
	// MainPic 当dataType=goods时，标识商品主图；当dataType=activity时，标识活动会场主图；（10/27新增字段）
	MainPic string `json:"mainPic,omitempty"`
	// DataType goods标识商品；activity标识活动会场（10/27新增字段）
	DataType string `json:"dataType,omitempty"`
	// CouponSrcScene 优惠券类型，0-全网公开券；1-阿里妈妈券（11/16新增字段）
	CouponSrcScene int `json:"couponSrcScene"`
	// ItemLink 商品链接（11/16新增字段）
	ItemLink string `json:"itemLink,omitempty"`
	// CouponLink 优惠券链接（11/16新增字段）
	CouponLink string `json:"couponLink,omitempty"`
	// BizSceneID 场景ID是联盟为了更好跟踪业务数据设立的一个新的接口参数，在转链时，需要根据场景不同入参不同的场景ID （2022/9/6新增出参）
	BizSceneID int `json:"bizSceneId,omitempty"`
}

// OriginInfo 链接中的信息
type OriginInfo struct {
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// ShopName 店铺名
	ShopName string `json:"shopName,omitempty"`
	// ShopLogo 店铺LOGO图
	ShopLogo string `json:"shopLogo,omitempty"`
	// Image 商品主图
	Image string `json:"image,omitempty"`
	// StartTime 券开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 券结束时间
	EndTime string `json:"endTime,omitempty"`
	// Amount 券金额
	Amount float64 `json:"amount,omitempty"`
	// StartFee 券门槛金额
	StartFee float64 `json:"startFee,omitempty"`
	// Price 商品价格
	Price float64 `json:"price,omitempty"`
	// ActivityID 券ID
	ActitivyID string `json:"activityId,omitempty"`
	// Pid PID
	Pid string `json:"pid,omitempty"`
	// Status 券状态。0:可用; 非0:不可用
	Status int `json:"status,omitempty"`
}

// ParseContent 淘系万能解析
func ParseContent(clt *core.Client, content string, ret *ParseContentResult) error {
	req := ParseContentRequest{
		Content: content,
	}
	return clt.Get(req, ret)
}
