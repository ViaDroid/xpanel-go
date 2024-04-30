package init

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func init() {
	web.AddFuncMap("czeq", CustomEqual)
	web.AddFuncMap("czne", CustomNotEqual)
	web.AddFuncMap("nl2br", Nl2br)
	web.AddFuncMap("inc", Inc[int])
	web.AddFuncMap("inc", Inc[int64])
	web.AddFuncMap("inc", Inc[float64])
	web.AddFuncMap("contains", Contains[string])
	// web.AddFuncMap("contains", Contains[int])
}

func CustomNotEqual(in any, in2 any) bool {
	return in != in2
}

func CustomEqual(in any, in2 any) bool {
	return in == in2
}

func Nl2br(text string) string {
	return strings.ReplaceAll(text, "\n", " <br>")
}

func Inc[n num](in n) n {
	return in + 1
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
