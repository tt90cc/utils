package uniqueid

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/tt90cc/utils"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_HOMESTAY_ORDER SnPrefix = "HSO" // 民宿订单前缀 looklook_order/homestay_order
	SN_PREFIX_THIRD_PAYMENT  SnPrefix = "PMT" // 第三方支付流水记录前缀 third_payment
)

// 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), utils.Krand(8, utils.KC_RAND_KIND_NUM))
}

// 公司编号 日期+纳税人识别码
func GenCompanySN(s string, length int) string {
	if len(s) < length {
		return cast.ToString(GenId())
	}

	num := s[len(s)-length : len(s)]
	return fmt.Sprintf("%s%s", time.Now().Format("20060102"), num)
}
