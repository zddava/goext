package goext

import (
	"testing"

	"github.com/zddava/goext/assert"
)

/*
Validate struct by field's tag. the tags currently supported as below:

string: required, ip4addr, ip6addr, ip4net, ip6net, netmask, macaddr, domain, date, time, datetime

int, float: nonzero, positive, negative, gte0, lte0

uint: nonzero, positive

pointer: required and actual type's tag
*/
type ValidateTestSingle struct {
	Str_Required   string  `v:"required"`
	Str_Ip4        string  `v:"ip4addr"`
	Str_Ip6        string  `v:"ip6addr"`
	Str_Ip4net     string  `v:"ip4net"`
	Str_Ip6net     string  `v:"ip6net"`
	Str_Netmask    string  `v:"netmask"`
	Str_Mac        string  `v:"macaddr"`
	Str_Domain     string  `v:"domain"`
	Str_URL        string  `v:"url"`
	Str_Port       string  `v:"port"`
	Str_Date       string  `v:"date"`
	Str_Time       string  `v:"time"`
	Str_DateTime   string  `v:"datetime"`
	Int_Nonzero    int     `v:"nonzero"`
	Int_Positive   int8    `v:"positive"`
	Int_Negative   int16   `v:"negative"`
	Int_Gte0       int32   `v:"gte0"`
	Int_Lte0       int64   `v:"lte0"`
	Int_Port       int     `v:"port"`
	Uint_Nonzero   uint    `v:"nonzero"`
	Uint_Positive  uint32  `v:"positive"`
	Uint_Port      uint64  `v:"port"`
	Float_Nonzero  float32 `v:"nonzero"`
	Float_Positive float32 `v:"positive"`
	Float_Negative float32 `v:"negative"`
	Float_Gte0     float64 `v:"gte0"`
	Float_Lte0     float64 `v:"lte0"`
	Pointer_Ip4    *string `v:"ip4addr"`
}

func TestSuite(t *testing.T) {
	t.Run("simple validate", TestSimpleValidate)
	t.Run("slice validate", TestSliceValidate)
	t.Run("pointer validate", TestPointerValidate)
	t.Run("tag validate", TestPointerCustom)
}

func TestSimpleValidate(t *testing.T) {
	s := &ValidateTestSingle{}
	assert.HasError(Validate(s), t)

	s.Str_Required = "astring"
	s.Str_Ip4 = "114.114.114.1111"
	assert.HasError(Validate(s), t)

	s.Str_Ip4 = "240e:390:e05:8080:75a1:d7d3:9d6e:b23a"
	assert.HasError(Validate(s), t)

	s.Str_Ip4 = "114.114.114.114"
	s.Str_Ip6 = "114.114.114.114"
	assert.HasError(Validate(s), t)

	s.Str_Ip6 = "240e:390:e05:8080:75a1:d7d3:9d6e:b23a"
	s.Str_Ip4net = "192.168.100.0/33"
	assert.HasError(Validate(s), t)

	s.Str_Ip4net = "192.168.100.0/24"
	s.Str_Ip6net = "192.168.100.0/24"
	assert.HasError(Validate(s), t)

	s.Str_Ip6net = "2001:db8::/32"
	s.Str_Netmask = "255.255.255.1"
	assert.HasError(Validate(s), t)

	s.Str_Netmask = "255.255.255.0"
	s.Str_Mac = "11:22:33:44:55:YY"
	assert.HasError(Validate(s), t)

	s.Str_Mac = "11:22:33:44:55:66"
	s.Str_Domain = "1"
	assert.HasError(Validate(s), t)

	s.Str_Domain = "~!!!.baidu.com"
	assert.HasError(Validate(s), t)

	s.Str_Domain = ".baidu.com.cn"
	assert.HasError(Validate(s), t)

	s.Str_Domain = "baidu.com"
	s.Str_URL = "http://1"
	assert.HasError(Validate(s), t)

	s.Str_URL = "www.163.com"
	assert.HasError(Validate(s), t)

	s.Str_URL = "https://www.163.com/u/r/l/"
	s.Str_Port = "0"
	assert.HasError(Validate(s), t)

	s.Str_Port = "65536"
	assert.HasError(Validate(s), t)

	s.Str_Port = "65535"
	s.Str_Date = "2022-13-02"
	assert.HasError(Validate(s), t)

	s.Str_Date = "2022-02-01"
	s.Str_Time = "12:00:61"
	assert.HasError(Validate(s), t)

	s.Str_Time = "12:00:01"
	s.Str_DateTime = "2022-13-01 12:00:01"
	assert.HasError(Validate(s), t)

	s.Str_DateTime = "2022-01-01 12:00:01"
	s.Int_Nonzero = 0
	assert.HasError(Validate(s), t)

	s.Int_Nonzero = 1
	s.Int_Positive = -1
	assert.HasError(Validate(s), t)

	s.Int_Positive = 1
	s.Int_Negative = 0
	assert.HasError(Validate(s), t)

	s.Int_Negative = -1
	s.Int_Gte0 = -1
	assert.HasError(Validate(s), t)

	s.Int_Gte0 = 0
	s.Int_Lte0 = 1
	assert.HasError(Validate(s), t)

	s.Int_Lte0 = 0
	s.Int_Port = 0
	assert.HasError(Validate(s), t)

	s.Int_Port = 65536
	assert.HasError(Validate(s), t)

	s.Int_Port = 65535
	s.Uint_Nonzero = 0
	assert.HasError(Validate(s), t)

	s.Uint_Nonzero = 1
	s.Uint_Positive = 0
	assert.HasError(Validate(s), t)

	s.Uint_Positive = 1
	s.Uint_Port = 65536
	assert.HasError(Validate(s), t)

	s.Uint_Port = 65535
	s.Float_Nonzero = 0
	assert.HasError(Validate(s), t)

	s.Float_Nonzero = 1
	s.Float_Positive = -1
	assert.HasError(Validate(s), t)

	s.Float_Positive = 1
	s.Float_Negative = 0
	assert.HasError(Validate(s), t)

	s.Float_Negative = -1
	s.Float_Gte0 = -1
	assert.HasError(Validate(s), t)

	s.Float_Gte0 = 0
	s.Float_Lte0 = 1
	assert.HasError(Validate(s), t)

	s.Float_Lte0 = -1
	str := "114.114.111"
	s.Pointer_Ip4 = &str
	assert.HasError(Validate(s), t)

	str = "114.114.114.114"
	assert.NoError(Validate(s), t)
}

type ValidateTestSlice struct {
	StrSlice []string `v:"required,ip4addr"`
}

func TestSliceValidate(t *testing.T) {
	s := ValidateTestSlice{}
	assert.HasError(Validate(s), t)

	s.StrSlice = make([]string, 0)
	assert.HasError(Validate(s), t)

	s.StrSlice = append(s.StrSlice, "114.114.114")
	assert.HasError(Validate(s), t)

	s.StrSlice = []string{}
	s.StrSlice = append(s.StrSlice, "114.114.114.114", "114.114.114.")
	assert.HasError(Validate(s), t)

	s.StrSlice = []string{}
	s.StrSlice = append(s.StrSlice, "114.114.114.114", "114.114.114.115")
	assert.NoError(Validate(s), t)
}

type ValidateTestPointer struct {
	StrPointer *string `v:"required,ip4addr"`
}

func TestPointerValidate(t *testing.T) {
	s := &ValidateTestPointer{}
	assert.HasError(Validate(s), t)

	str := "114.114.111"
	s.StrPointer = &str
	assert.HasError(Validate(s), t)

	str = "114.114.114.114"
	assert.NoError(Validate(s), t)
}

type ValidateCustom struct {
	StrPointer *string `v-custom:"required,ip4addr"`
}

func TestPointerCustom(t *testing.T) {
	s := &ValidateCustom{}
	assert.HasError(ValidateTag(s, "v-custom"), t)

	str := "114.114.112"
	s.StrPointer = &str
	assert.HasError(ValidateTag(s, "v-custom"), t)

	str = "114.114.114.114"
	assert.NoError(ValidateTag(s, "v-custom"), t)
}
