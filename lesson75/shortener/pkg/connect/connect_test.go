package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/unit-test-5/"
		got := Get(url)
		// 断言
		convey.So(got, convey.ShouldEqual, true) // 断言
	})
	convey.Convey("url请求不通的示例", t, func() {
		url := "posts/Go/unit-test-5/"
		got := Get(url)
		// 断言
		convey.ShouldBeFalse(got)
	})
}
