package httpx

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg6/go-httpx/mapx"
	"net/http"
	"strings"
)

const defaultMemory = 32 << 20

func Bind(req *http.Request, obj any) error {
	if req.Method == http.MethodGet {
		return BindQuery(req, obj)
	}
	contentType := req.Header.Get(HeaderContentType)
	if strings.HasPrefix(contentType, ContentTypeJSON) {
		return BindJSON(req, obj)
	}
	if strings.HasPrefix(contentType, ContentTypeXML) {
		return BindXML(req, obj)
	}
	return BindForm(req, obj)
}

func BindJSON(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return json.NewDecoder(req.Body).Decode(obj)
}

func BindXML(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return xml.NewDecoder(req.Body).Decode(obj)
}

func BindHeaders(req *http.Request, obj any) error {
	return mapx.Mapping(req.Header, "header", obj)
}

func BindQuery(req *http.Request, obj any) error {
	return mapx.Mapping(req.URL.Query(), "query", obj)
}

func BindForm(req *http.Request, obj any) error {
	if strings.HasPrefix(req.Header.Get(HeaderContentType), ContentTypeMultipartForm) {
		err := req.ParseMultipartForm(defaultMemory)
		if err != nil {
			return err
		}
	} else {
		if err := req.ParseForm(); err != nil {
			return err
		}
	}
	return mapx.Mapping(req.Form, "form", obj)
}
