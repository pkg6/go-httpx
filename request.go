package httpx

import (
	"github.com/pkg6/go-httpx/regexpx"
	"net/http"
	"strings"
)

var (
	RegRequestIsScript = "curl|wget|collectd|python|urllib|java|jakarta|httpclient|phpcrawl|libwww|perl|go-http|okhttp|lua-resty|winhttp|awesomium"
)

func IsGet(req *http.Request) bool {
	return Method(req) == http.MethodGet
}
func IsHead(req *http.Request) bool {
	return Method(req) == http.MethodHead
}
func IsPost(req *http.Request) bool {
	return Method(req) == http.MethodPost
}
func IsPut(req *http.Request) bool {
	return Method(req) == http.MethodPut
}
func IsPatch(req *http.Request) bool {
	return Method(req) == http.MethodPatch
}
func IsDelete(req *http.Request) bool {
	return Method(req) == http.MethodDelete
}
func IsConnect(req *http.Request) bool {
	return Method(req) == http.MethodConnect
}
func IsOptions(req *http.Request) bool {
	return Method(req) == http.MethodOptions
}
func IsTrace(req *http.Request) bool {
	return Method(req) == http.MethodTrace
}
func Method(req *http.Request) string {
	return req.Method
}

func IsAjax(req *http.Request) bool {
	return req.Header.Get(HeaderXRequestedWith) == "XMLHttpRequest"
}

func IsScript(req *http.Request) bool {
	s := strings.ToLower(req.Header.Get(HeaderUserAgent))
	return regexpx.IsMatchString(RegRequestIsScript, s)
}
func IsSSL(req *http.Request) bool {
	return strings.EqualFold(req.URL.Scheme, SchemeHttps) || req.TLS != nil
}

func Form(r *http.Request, postMaxMemory int64) (form map[string][]string, found bool) {
	if form := r.Form; len(form) > 0 {
		return form, true
	}
	if form := r.PostForm; len(form) > 0 {
		return form, true
	}
	if m := r.MultipartForm; m != nil {
		if len(m.Value) > 0 {
			return m.Value, true
		}
	}
	err := r.ParseMultipartForm(postMaxMemory)
	if err != nil && err != http.ErrNotMultipart {
		return nil, false
	}
	if form := r.Form; len(form) > 0 {
		return form, true
	}
	if form := r.PostForm; len(form) > 0 {
		return form, true
	}
	if m := r.MultipartForm; m != nil {
		if len(m.Value) > 0 {
			return m.Value, true
		}
	}
	return nil, false
}
