package enum

import (
	"fmt"
	"testing"
)

type ISP Enum

var (
	CMCC = InitEnum[ISP]("cmcc", "中国移动")
	CUCC = InitEnum[ISP]("cucc", "中国联通")
	CTCC = InitEnum[ISP]("ctcc", "中国电信")
)

func TestEmun(t *testing.T) {
	// 根据字符串解析
	isp := ParseEnum[ISP]("cmcc")
	fmt.Println(isp)
}
