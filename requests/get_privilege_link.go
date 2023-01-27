package requests

import (
	"net/url"
	"strconv"

	"github.com/bububa/dataoke-go/core"
	"github.com/bububa/dataoke-go/util"
)

// GetPrivilegeLinkRequest 高效转链 API Request
type GetPrivilegeLinkRequest struct {
	// GoodsID 淘宝商品id
	GoodsID string `json:"goodsId,omitempty"`
	// CouponID 商品的优惠券ID，一个商品在联盟可能有多个优惠券，可通过填写该参数的方式选择使用的优惠券，请确认优惠券ID正确，否则无法正常跳转
	CouponID string `json:"couponId,omitempty"`
	// Pid 推广位ID，用户可自由填写当前大淘客账号下已授权淘宝账号的任一pid，若未填写，则默认使用创建应用时绑定的pid
	Pid string `json:"pid,omitempty"`
	// ChannelID 渠道id将会和传入的pid进行验证，验证通过将正常转链，请确认填入的渠道id是正确的 channelId对应联盟的relationId
	ChannelID string `json:"channelId,omitempty"`
	// BizSceneID 场景ID是联盟为了更好跟踪业务数据设立的一个新的接口参数，在转链时，需要根据场景不同入参不同的场景ID （2022/9/6新增入参）
	BizSceneID int `json:"bizSceneId,omitempty"`
	// PromotionType 代理模式必须入参该字段，在实际业务中，【自购省】按钮-入参1，【推广赚】按钮-入参2（2022/9/6新增入参）
	PromotionType int `json:"promotionType,omitempty"`
	// RebateType 付定返红包，0.不使用付定返红包，1.参与付定返红包
	RebateType int `json:"rebateType,omitempty"`
	// SpecialID 会员运营id
	SpecialID string `json:"specialId,omitempty"`
	// ExternalID 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	ExternalID string `json:"externalId,omitempty"`
	// Xid 团长与下游渠道合作的特殊标识，用于统计渠道推广效果 （新增入参）
	Xid string `json:"xid,omitempty"`
	// LeftSymbol 淘口令左边自定义符号,默认￥ （2021/3/9新增入参）
	LeftSymbol string `json:"leftSymbol,omitempty"`
	// RightSymbol 淘口令右边自定义符号,默认￥ （2021/3/9新增入参）
	RightSymbol string `json:"rightSymbol,omitempty"`
	// AuthID 平台的淘宝授权id（获取地址：https://www.dataoke.com/shouquan?type=1&auth_id=1），如果传入了该参数则必须填写对应淘宝联盟授权账号的pid（1/21新增字段）
	AuthID string `json:"authId,omitempty"`
}

// Values implement Request interface
func (r GetPrivilegeLinkRequest) Values(values url.Values) {
	values.Set("goodsId", r.GoodsID)
	if r.CouponID != "" {
		values.Set("couponId", r.CouponID)
	}
	if r.Pid != "" {
		values.Set("pid", r.Pid)
	}
	if r.ChannelID != "" {
		values.Set("channelId", r.ChannelID)
	}
	if r.BizSceneID > 0 {
		values.Set("bizSceneId", strconv.Itoa(r.BizSceneID))
	}
	if r.PromotionType > 0 {
		values.Set("promotionType", strconv.Itoa(r.PromotionType))
	}
	if r.RebateType > 0 {
		values.Set("rebateType", strconv.Itoa(r.RebateType))
	}
	if r.SpecialID != "" {
		values.Set("specialId", r.SpecialID)
	}
	if r.ExternalID != "" {
		values.Set("externalId", r.ExternalID)
	}
	if r.Xid != "" {
		values.Set("xid", r.Xid)
	}
	if r.LeftSymbol != "" {
		values.Set("leftSymbol", r.LeftSymbol)
	}
	if r.RightSymbol != "" {
		values.Set("rightSymbol", r.RightSymbol)
	}
	if r.AuthID != "" {
		values.Set("authId", r.AuthID)
	}
}

// Url implement Request interface
func (r GetPrivilegeLinkRequest) Url() string {
	return "tb-service/get-privilege-link"
}

// PrivilegeLink 转链结果
type PrivilegeLink struct {
	// CouponClickURL 商品优惠券推广链接
	CouponClickURL string `json:"couponClickUrl,omitempty"`
	// CouponEndTime 优惠券结束时间
	CouponEndTime string `json:"couponEndTime,omitempty"`
	// CouponInfo 优惠券面额
	CouponInfo string `json:"couponInfo,omitempty"`
	// CouponStartTime 优惠券开始时间
	CouponStartTime string `json:"couponStartTime,omitempty"`
	// ItemID 商品id
	ItemID string `json:"itemId,omitempty"`
	// CouponTotalCount 优惠券总量
	CouponTotalCount util.Int64 `json:"couponTotalCount,omitempty"`
	// CouponRemainCount 优惠券剩余量
	CouponRemainCount util.Int64 `json:"couponRemainCount,omitempty"`
	// ItemURL 商品淘客链接
	ItemURL string `json:"itemUrl,omitempty"`
	// Tpwd 淘口令
	Tpwd string `json:"tpwd,omitempty"`
	// LongTpwd 针对iOS14版本，增加对应能解析的长口令
	LongTpwd string `json:"longTpwd,omitempty"`
	// MaxCommissionRate 佣金比例
	MaxCommissionRate util.Float64 `json:"maxCommissionRate,omitempty"`
	// ShortURL 短链接
	ShortURL string `json:"shortUrl,omitempty"`
	// MinCommissionRate 当传入请求参数channelId、specialId、externalId时，该字段展示预估最低佣金率(%)(接联盟通知，官方比价订单佣金调整正式生效时间推迟至7月22日)
	MinCommissionRate util.Float64 `json:"minCommissionRate,omitempty"`
	// KuaiZhanURL 商品的快站链接（支持在微信端直接访问商品详情） 仅支持大淘客商品
	KuaiZhanURL string `json:"kuaiZhanUrl,omitempty"`
	// OriginalPrice 商品原价(2020/12/30新增字段)
	OriginalPrice util.Float64 `json:"originalPrice,omitempty"`
	// ActualPrice 券后价(2020/12/30新增字段)
	ActualPrice util.Float64 `json:"actualPrice,omitempty"`
}

// GetPrivilageLink 高效转链
func GetPrivilageLink(clt *core.Client, req *GetPrivilegeLinkRequest, ret *PrivilegeLink) error {
	return clt.Get(req, ret)
}
