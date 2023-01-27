package requests

import (
	"net/url"
	"strconv"

	"github.com/bububa/dataoke-go/core"
	"github.com/bububa/dataoke-go/util"
)

// GetGoodsDetailsRequest 单品详情 API Request
type GetGoodsDetailsRequest struct {
	// ID 大淘客商品id，请求时id或goodsId必填其中一个，若均填写，将优先查找当前单品id
	ID uint64 `json:"id,omitempty"`
	// GoodsID 淘宝商品id，id或goodsId必填其中一个，若均填写，将优先查找当前单品id
	GoodsID string `json:"goodsId,omitempty"`
}

// Values implement Request interface
func (r GetGoodsDetailsRequest) Values(values url.Values) {
	values.Set("id", strconv.FormatUint(r.ID, 10))
	if r.GoodsID != "" {
		values.Set("goodsId", r.GoodsID)
	}
}

// Url implement Request interface
func (r GetGoodsDetailsRequest) Url() string {
	return "goods/get-goods-details"
}

// GoodsDetail 单品详情
type GoodsDetail struct {
	// GoodsID 淘宝商品id
	GoodsID string `json:"goodsId,omitempty"`
	// ID 商品id,若查询结果id=-1，则说明该商品非大淘客平台商品，数据为淘宝直接返回的数据，由于淘宝数据的缓存时间相对较长，会出现商品信息和详情信息不一致的情况
	ID int64 `json:"id,omitempty"`
	// GoodsSign 新商品id，每次转链后，都会变化。同一个淘宝客账号对同一个商品 进行转链，商品ID后半段不变 8/23新增字段，暂时为空
	GoodsSign string `json:"goodsSign,omitempty"`
	// ItemLink 商品淘宝链接
	ItemLink string `json:"itemLink,omitempty"`
	// Title 淘宝标题
	Title string `json:"title,omitempty"`
	// Dtitle 大淘客短标题
	Dtitle string `json:"dtitle,omitempty"`
	// SpecialText 特色文案（商品卖点）
	SpecialText []string `json:"specialText,omitempty"`
	// Desc 推广文案
	Desc string `json:"desc,omitempty"`
	// Cid 商品在大淘客的分类id
	Cid uint64 `json:"cid,omitempty"`
	// SubCid 商品在大淘客的二级分类id，该分类为前端分类，一个商品可能有多个二级分类id
	SubCid []uint64 `json:"subCid,omitempty"`
	// TbCid 商品在淘宝的分类id
	TbCid uint64 `json:"tbcid,omitempty"`
	// MainPic 商品主图链接
	MainPic string `json:"mainPic,omitempty"`
	// MarketingMainPic 营销主图链接
	MarketingMainPic string `json:"marketingMainPic,omitempty"`
	// Video 商品视频（新增字段）
	Video string `json:"video,omitempty"`
	// OriginalPrice 商品原价
	OriginalPrice float64 `json:"originalPrice,omitempty"`
	// ActualPrice 券后价
	ActualPrice float64 `json:"actualPrice,omitempty"`
	// Discounts 折扣力度
	Discounts float64 `json:"discounts,omitempty"`
	// CommissionType 佣金类型，0-通用，1-定向，2-高佣，3-营销计划
	CommissionType int `json:"commissionType,omitempty"`
	// CommissionRate 佣金比例
	CommissionRate float64 `json:"commissionRate,omitempty"`
	// CouponLink 优惠券链接
	CouponLink string `json:"couponLink,omitempty"`
	// CouponTotalNum 券总量
	CouponTotalNum int64 `json:"couponTotalNum,omitempty"`
	// CouponReceiveNum 领券量
	CouponReceiveNum int64 `json:"couponReceiveNum,omitempty"`
	// CouponEndTime 优惠券结束时间
	CouponEndTime string `json:"couponEndTime,omitempty"`
	// CouponStartTime 优惠券开始时间
	CouponStartTime string `json:"couponStartTime,omitempty"`
	// CouponPrice 优惠券金额
	CouponPrice float64 `json:"couponPrice,omitempty"`
	// CouponConditions 优惠券使用条件
	CouponConditions string `json:"couponConditions,omitempty"`
	// MonthSales 30天销量
	MonthSales int64 `json:"monthSales,omitempty"`
	// TwoHoursSales 2小时销量
	TwoHoursSales int64 `json:"twoHoursSales,omitempty"`
	// DailySales 当天销量
	DailySales int64 `json:"dailySales,omitempty"`
	// Brand 是否是品牌商品
	Brand int `json:"brand,omitempty"`
	// BrandID 品牌id
	BrandID uint64 `json:"brandId,omitempty"`
	// BrandName 品牌名称
	BrandName string `json:"brandName,omitempty"`
	// CreateTime 商品上架时间
	CreateTime string `json:"createTime,omitempty"`
	// ActivityType 活动类型，1-无活动，2-淘抢购，3-聚划算
	ActivityType int `json:"activityType,omitempty"`
	// ActivityStartTime 活动开始时间
	ActivityStartTime string `json:"activityStartTime,omitempty"`
	// ActivityEndTime 活动结束时间
	ActivityEndTime string `json:"activityEndTime,omitempty"`
	// ShopType 店铺类型，1-天猫，0-淘宝
	ShopType int `json:"shopType,omitempty"`
	// GoldSellers 是否金牌卖家，1-金牌卖家，0-非金牌卖家
	GoldSellers int `json:"goldSellers,omitempty"`
	// SellerID 淘宝卖家id，也是店铺id。店铺转链可用此字段
	SellerID util.Uint64 `json:"sellerId,omitempty"`
	// ShopName 店铺名称
	ShopName string `json:"shopName,omitempty"`
	// ShopLevel 淘宝店铺等级
	ShopLevel int `json:"shopLevel,omitempty"`
	// DescScore 描述分
	DescScore float64 `json:"descScore,omitempty"`
	// DsrScore 描述相符
	DsrScore float64 `json:"dsrScore,omitempty"`
	// DsrPercent 描述同行比
	DsrPercent float64 `json:"dsrPercent,omitempty"`
	// ShipScore 物流服务
	ShipScore float64 `json:"shipScore,omtempty"`
	// ShipPercent 物流同行比
	ShipPercent float64 `json:"shipPercent,omitempty"`
	// ServiceScore 服务态度
	ServiceScore float64 `json:"serviceScore,omitempty"`
	// ServicePercent 服务同行比
	ServicePercent float64 `json:"servicePercent,omitempty"`
	// HotPush 热推值
	HotPush int64 `json:"hotPush,omitempty"`
	// TeamName 放单人名称
	TeamName string `json:"teamName,omitempty"`
	// DetailPics 商品详情图（需要做适配）
	DetailPics string `json:"detailPics,omitempty"`
	// Imgs 淘宝轮播图
	Imgs string `json:"imgs,omitempty"`
	// Reimgs 相关商品图
	Reimgs string `json:"reimgs,omitempty"`
	// QuanMlink 定金，若无定金，则显示0
	QuanMlink float64 `json:"quanMlink,omitempty"`
	// HzQuanOver 立减，若无立减金额，则显示0
	HzQuanOver float64 `json:"hzQuanOver,omitempty"`
	// Yunfeixian 0.不包运费险 1.包运费险
	Yunfeixian int `json:"yunfeixian,omitempty"`
	// EstimateAmount 预估淘礼金
	EstimateAmount float64 `json:"estimateAmount,omitempty"`
	// ShopLogo 店铺logo
	ShopLogo string `json:"shopLogo,omitempty"`
	// FreeshipRemoteDistrict 偏远地区包邮，0.不包邮，1.包邮
	FreesholdRemoteDistrict int `json:"freeshipRemoteDistrict,omitempty"`
	// IsSubdivision 该商品是否有细分类目：0不是，1是
	IsSubdivision int `json:"isSubdivision,omitempty"`
	// SubdivisionID 该商品所属细分类目id
	SubdivisionID uint64 `json:"subdivisionId,omitempty"`
	// SubdivisionName 该商品所属细分类目名称
	SubdivisionName string `json:"subdivisionName,omitempty"`
	// SubdivisionRank 该商品所属细分类目排名
	SubdivisionRank int `json:"subdivisionRank,omitempty"`
	// DirectCommissionType 定向佣金类型，1非定向，3定向佣金
	DirectCommissionType int `json:"directCommissionType,omitempty"`
	// DirectCommission 定向佣金
	DirectCommission float64 `json:"directCommission,omitempty"`
	// DirectCommissionLink 定向链接
	DirectCommissionLink string `json:"directCommissionLink,omitempty"`
	// Sales24h 24小时销量（8/17新增字段）
	Sales24h int64 `json:"sales24h,omitempty"`
	// Lowest 是否近30天历史最低价，0-否；1-是(8/18新增字段)
	Lowest int `json:"lowest,omitempty"`
	// CouponID 优惠券ID(9/15新增字段)
	CouponID string `json:"couponId,omitempty"`
	// DiscountType 1.购物津贴；2.跨店满减；0.无(10/27新增字段)
	DiscountType int `json:"discountType,omitempty"`
	// DiscountFull 活动满减的满值(10/27新增字段)
	DiscountFull float64 `json:"discountFull,omitempty"`
	// DiscountCut 活动满减的减值(10/27新增字段)
	DiscountCut float64 `json:"discountCut,omitempty"`
	// ActivityInfo 活动信息(10/27新增字段)
	ActivityInfo *ActivityInfo `json:"activityInfo,omitempty"`
	// ActivityID 单单有奖活动id(1/10新增字段)
	ActivityID string `json:"activityId,omitempty"`
	// InspectedGoods 商品是否已经验货，0-否；1-是(1/26新增字段)
	InspectedGoods int `json:"inspectedGoods,omitempty"`
	// BizSceneId
	BizSceneId int `json:"bizSceneId,omitempty"`
}

// ActivityInfo 活动信息(10/27新增字段)
type ActivityInfo struct {
	// ActivityID 单单有奖活动id(1/10新增字段)
	ActivityID uint64 `json:"activityId,omitempty"`
	// ActivityName 活动名称(10/27新增字段)
	ActivityName string `json:"activityName,omitempty"`
}

// GetGoodsDetails 单品详情
func GetGoodsDetails(clt *core.Client, req *GetGoodsDetailsRequest, ret *GoodsDetail) error {
	return clt.Get(req, ret)
}
