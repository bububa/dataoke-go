package requests

import (
	"net/url"
	"strconv"

	"github.com/bububa/dataoke-go/core"
	"github.com/bububa/dataoke-go/util"
)

// GetTbServiceRequest 联盟搜索 API Request
type GetTbServiceRequest struct {
	// PageNo 第几页，默认1
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页条数， 默认20，范围值1~100
	PageSize int `json:"pageSize,omitempty"`
	// Keywords 查询词
	Keywords string `json:"keyWords,omitempty"`
	// Sort 排序指标：销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）,排序方式：排序_des（降序），排序_asc（升序）,示例：升序查询销量：total_sales_asc
	Sort string `json:"sort,omitempty"`
	// Source 是否商城商品，设置为1表示该商品是属于淘宝商城商品，设置为非1或不设置表示不判断这个属性（和overseas字段冲突，若已请求source，请勿再请求overseas）
	Source int `json:"source,omitempty"`
	// Overseas 是否海外商品，设置为1表示该商品是属于海外商品，设置为非1或不设置表示不判断这个属性（和source字段冲突，若已请求overseas，请勿再请求source）
	Overseas int `json:"overseas,omitempty"`
	// EndPrice 折扣价范围上限，单位：元
	EndPrice float64 `json:"endPrice,omitempty"`
	// StartPrice 折扣价范围下限，单位：元
	StartPrice float64 `json:"startPrice,omitempty"`
	// StartTkRate 媒体淘客佣金比率下限，如：1234表示12.34%
	StartTkRate int `json:"startTkRate,omitempty"`
	// EndTkRate 商品筛选-淘客佣金比率上限，如：1234表示12.34%
	EndTkRate int `json:"endTkRate,omitempty"`
	// HasCoupon 是否有优惠券，设置为true表示该商品有优惠券，设置为false或不设置表示不判断这个属性
	HasCoupon bool `json:"hasCoupon,omitempty"`
	// SpecialID 会员运营id
	SpecialID string `json:"specialId,omitempty"`
	// ChannelID 渠道id将会和传入的pid进行验证，验证通过将正常转链，请确认填入的渠道id是正确的channelId对应联盟的relationId
	ChannelID string `json:"channelId,omitempty"`
	// ItemLoc 商品所在地，默认为全部商品，其他值：北京、上海、广州等必须是城市名称，不能带省份(2021/1/15新增字段)
	ItemLoc string `json:"itemLoc,omitempty"`
	// NeedPrepay 商品是否加入消费者保障，1为加入消费者保障，默认全部(2021/1/15新增字段)
	NeedPrepay int `json:"needPrepay,omitempty"`
	// IncludeGoodRate 商品好评率是否高于行业均值，1为高于行业均值，默认全部(2021/1/15新增字段)
	IncludeGoodRate int `json:"includeGoodRate,omitempty"`
}

// Values implement Request interface
func (r GetTbServiceRequest) Values(values url.Values) {
	if r.PageNo < 1 {
		r.PageNo = 1
	}
	values.Set("pageNo", strconv.Itoa(r.PageNo))
	if r.PageSize == 0 {
		r.PageSize = 20
	}
	values.Set("pageSize", strconv.Itoa(r.PageSize))
	values.Set("keyWords", r.Keywords)
	if r.Sort != "" {
		values.Set("sort", r.Sort)
	}
	if r.Source == 1 {
		values.Set("source", "1")
	}
	if r.Overseas == 1 {
		values.Set("overseas", "1")
	}
	if r.EndPrice > 1e-15 {
		values.Set("endPrice", strconv.FormatFloat(r.EndPrice, 'f', 2, 64))
	}
	if r.StartPrice > 1e-15 {
		values.Set("startPrice", strconv.FormatFloat(r.StartPrice, 'f', 2, 64))
	}
	if r.StartTkRate > 0 {
		values.Set("startTkRate", strconv.Itoa(r.StartTkRate))
	}
	if r.EndTkRate > 0 {
		values.Set("endTkRate", strconv.Itoa(r.EndTkRate))
	}
	if r.HasCoupon {
		values.Set("hasCoupon", "true")
	}
	if r.SpecialID != "" {
		values.Set("specialId", r.SpecialID)
	}
	if r.ChannelID != "" {
		values.Set("channelId", r.ChannelID)
	}
	if r.ItemLoc != "" {
		values.Set("itemLoc", r.ItemLoc)
	}
	if r.NeedPrepay == 1 {
		values.Set("needPrepay", "1")
	}
	if r.IncludeGoodRate == 1 {
		values.Set("includeGoodRate", "1")
	}
}

// Url implement Request interface
func (r GetTbServiceRequest) Url() string {
	return "tb-service/get-tb-service"
}

// TbkItem 淘宝商品信息
type TbkItem struct {
	// Title 商品信息-商品标题
	Title string `json:"title,omitempty"`
	// Volume 商品信息-30天销量
	Volume int64 `json:"volume,omitempty"`
	// Nick 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty"`
	// CouponStartTime 优惠券信息-优惠券开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty"`
	// CouponEndTime 优惠券信息-优惠券结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty"`
	// TkTotalSales 商品信息-淘客30天推广量
	TkTotalSales util.Int64 `json:"tk_total_sales,omitempty"`
	// CouponID 优惠券信息-优惠券id
	CouponID string `json:"coupon_id,omitempty"`
	// PictURL 商品信息-商品主图
	PictURL string `json:"pict_url,omitempty"`
	// SmallImages 商品信息-商品小图列表
	SmallImages struct {
		String []string `json:"string,omitempty"`
	} `json:"small_images,omitempty"`
	// ReservePrice 商品信息-商品一口价格
	ReservePrice util.Float64 `json:"reserve_price,omitempty"`
	// ZkFinalPrice 商品信息-商品折扣价格
	ZkFinalPrice util.Float64 `json:"zk_final_price,omitempty"`
	// UserType 店铺信息-卖家类型，0-表示集市，1-表示天猫
	UserType int `json:"user_type,omitempty"`
	// SellerID 店铺信息-卖家id
	SellerID uint64 `json:"seller_id,omitempty"`
	// CouponTotalCount 优惠券信息-优惠券总量
	CouponTotalCount int64 `json:"coupon_total_count,omitempty"`
	// CouponRemainCount 优惠券信息-优惠券剩余量
	CouponRemainCount int64 `json:"coupon_remain_count,omitempty"`
	// CouponInfo 优惠券信息-优惠券满减信息
	CouponInfo string `json:"coupon_info,omitempty"`
	// ShopTitle 店铺信息-店铺名称
	ShopTitle string `json:"shop_title,omitempty"`
	// ShopDsr 店铺信息-店铺dsr评分
	ShopDsr int64 `json:"shop_dsr,omitempty"`
	// LevelOneCategoryName 商品信息-一级类目名称
	LevelOneCategoryName string `json:"level_one_category_name,omitempty"`
	// LevelOneCategoryID 商品信息-一级类目ID
	LevelOneCategoryID uint64 `json:"level_one_category_id,omitempty"`
	// CaregoryName 商品信息-叶子类目名称
	CategoryName string `json:"category_name,omitempty"`
	// CategoryID 商品信息-叶子类目id
	CategoryID uint64 `json:"category_id,omitempty"`
	// ShortTitle 商品信息-商品短标题
	ShortTitle string `json:"short_title,omitempty"`
	// WhiteImage 商品信息-商品白底图
	WhiteImage string `json:"white_image,omitempty"`
	// CouponStartFee 优惠券信息-优惠券起用门槛
	CouponStartFee util.Float64 `json:"coupon_start_fee,omitempty"`
	// CouponAmount 优惠券信息-优惠券面额
	CouponAmount util.Int64 `json:"coupon_amount,omitempty"`
	// ItemDescription 商品信息-宝贝描述(推荐理由)
	ItemDescription string `json:"item_description,omitempty"`
	// ItemURL uland.taobao.com
	ItemURL string `json:"item_url,omitempty"`
	// URL s.click.taobao
	URL string `json:"url,omitempty"`
	// NumIid 商品信息-宝贝id
	NumIid string `json:"num_iid,omitempty"`
	// ItemID 商品信息-宝贝id
	ItemID string `json:"item_id,omitempty"`
	// CommissionRate 佣金比例
	CommissionRate float64 `json:"commission_rate,omitempty"`
	// YsylJltFace 预估淘礼金
	YsylJltFace util.Float64 `json:"ysyl_jlt_face,omitempty"`
	// PresaleDeposit 预售定金
	PresaleDeposit util.Float64 `json:"presale_deposit,omitempty"`
	// PresaleDiscountFeeText 定金立减
	PresaleDiscountFeeText string `json:"presale_discount_fee_text,omitempty"`
	// Province
	Province string `json:"province,omitempty"`
	// RealPostFee
	RealPostFee util.Float64 `json:"real_post_fee,omitempty"`
	// ActivityID 单单有奖活动id（1/10新增字段）
	ActivityID string `json:"activity_id,omitempty"`
	// CpaRewardAmount 单单有奖奖励金额（1/10新增字段）
	CpaRewardAmount util.Float64 `json:"cpa_reward_amount,omitempty"`
}

// GetTbService 联盟搜索
func GetTbService(clt *core.Client, req *GetTbServiceRequest, ret *[]TbkItem) error {
	return clt.Get(req, ret)
}
