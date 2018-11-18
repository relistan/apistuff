package apistuff

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetCookieByName(t *testing.T) {
	Convey("GetCookieByName()", t, func() {
		cookies := []*http.Cookie{
			&http.Cookie{
				Name:  "key1",
				Value: "value1",
			},
			&http.Cookie{
				Name:  "key2",
				Value: "value2",
			},
		}

		So(GetCookieByName(cookies, "key2"), ShouldEqual, "value2")
		So(GetCookieByName(cookies, "key1"), ShouldEqual, "value1")
		So(GetCookieByName(cookies, "asdf"), ShouldEqual, "")
	})
}
