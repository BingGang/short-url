package ecode

const (
	OK1                int64 = 1
	OK                 int64 = 20000
	ServerError        int64 = 1 // 服务器错误
	Forbidden          int64 = 10001
	ParamError         int64 = 10002 //参数错误
	SignError          int64 = 10003 //签名错误
	MallUserNoLogin    int64 = 10004 // 暂无用户
	HttpMethodError    int64 = 10005 //请求错误
	AuthError          int64 = 10006 //签名错误
	NotExist           int64 = 10007 //信息不存在
	OperationFail      int64 = 10008 //操作失败
	CacheDataError     int64 = 10009 //数据错误
	UserStatusUnable   int64 = 10010 //用户状态不可用
	WxCodeErr          int64 = 10011 //登录失败
	ProductStock       int64 = 10012 //库存不足
	ProductBuyNumLimit int64 = 10013 //购买限制
	ProductErr         int64 = 10014 //商品错误
	OrderStatusErr     int64 = 10015 //   订单状态错误
	AuthErr            int64 = 30000 //   订单状态错误
	PwdErr             int64 = 30001
	AccountExist       int64 = 30002
	ReturnBucketErr    int64 = 40000

	NotCoupon int64 = 10101
)

var Ecode_intro = map[int64]string{
	OK:                 "成功",
	Forbidden:          "禁止访问",
	ParamError:         "参数错误",
	SignError:          "签名错误",
	MallUserNoLogin:    "账号信息错误",
	HttpMethodError:    "请求错误",
	ServerError:        "服务器错误",
	AuthError:          "鉴权失败",
	NotExist:           "信息不存在",
	OperationFail:      "操作失败",
	CacheDataError:     "数据错误",
	UserStatusUnable:   "用户状态不可用",
	WxCodeErr:          "登录失败",
	ProductStock:       "库存不足",
	ProductBuyNumLimit: "超出限购数量",
	ProductErr:         "商品错误",
	OrderStatusErr:     "订单状态错误",
	AuthErr:            "鉴权失败",
	PwdErr:             "密码错误",
	AccountExist:       "用户不存在",
	ReturnBucketErr:    "还桶数量错误",
	NotCoupon:          "暂无优惠券",
}
