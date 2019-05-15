package cookie_test

import (
	"net/http"
	"testing"

	"github.com/anihex/server-utils/cookie"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemory(t *testing.T) {
	Convey("Using a memory cookie should work perfectly.", t, func() {
		tmpFunc := cookie.NewMemory(map[string]interface{}{
			"license":     "12345-12345-12345-12345-12345",
			"license_id":  1,
			"client_type": "user",
		})

		Convey("The CookieFunc should return no errors, but and an instance of MemoryCookie", func() {
			tmpCookie, err := tmpFunc(&FakeResponse{}, &http.Request{}, "demo")
			finalCookie, _ := tmpCookie.(*cookie.MemoryCookie)
			So(finalCookie, ShouldHaveSameTypeAs, &cookie.MemoryCookie{})
			So(err, ShouldBeNil)
		})
	})
}

type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

func New(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (r *FakeResponse) Assert(status int, body string) {
	if r.status != status {
		r.t.Errorf("expected status %+v to equal %+v", r.status, status)
	}
	if string(r.body) != body {
		r.t.Errorf("expected body %+v to equal %+v", string(r.body), body)
	}
}
