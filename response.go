package httpx

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

var (
	// BusinessCodeOK represents the business code for success.
	BusinessCodeOK = 0
	// BusinessMsgOK represents the business message for success.
	BusinessMsgOK = "ok"
	// BusinessCodeError represents the business code for error.
	BusinessCodeError = -1
)

const (
	xmlVersion  = "1.0"
	xmlEncoding = "UTF-8"
)

type BaseResponseErr interface {
	GetCode() int
	error
}

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	// Code represents the business code, not the http status code.
	Code int `json:"code" xml:"code"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgSuccess.
	Msg string `json:"msg" xml:"msg"`
	// Data represents the business data.
	Data T `json:"data,omitempty" xml:"data,omitempty"`
}

type baseXmlResponse[T any] struct {
	XMLName  xml.Name `xml:"xml"`
	Version  string   `xml:"version,attr"`
	Encoding string   `xml:"encoding,attr"`
	BaseResponse[T]
}

func wrapXmlBaseResponse(v any) baseXmlResponse[any] {
	base := wrapBaseResponse(v)
	return baseXmlResponse[any]{
		Version:      xmlVersion,
		Encoding:     xmlEncoding,
		BaseResponse: base,
	}
}

func wrapBaseResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]
	switch data := v.(type) {
	case *CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case BaseResponseErr:
		resp.Code = data.GetCode()
		resp.Msg = data.Error()
	case error:
		resp.Code = BusinessCodeError
		resp.Msg = data.Error()
	default:
		resp.Code = BusinessCodeOK
		resp.Msg = BusinessMsgOK
		resp.Data = v
	}
	return resp
}

// OKResponse writes HTTP 200 OK into w.
func OKResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// OKJSONBaseResponse writes v into w with http.StatusOK.
func OKJSONBaseResponse(w http.ResponseWriter, v any) error {
	return JSONBaseResponse(w, http.StatusOK, v)
}

func JSONBaseResponse(w http.ResponseWriter, code int, v any) error {
	return JSONResponseWithCode(w, code, wrapBaseResponse(v))
}

// JSONResponseWithCode writes v as json string into w with code.
func JSONResponseWithCode(w http.ResponseWriter, code int, v any) error {
	return doWriteJson(w, code, v)
}

func doWriteJson(w http.ResponseWriter, code int, v any) error {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal json failed, error: %w", err)
	}
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(code)
	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if !errors.Is(err, http.ErrHandlerTimeout) {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
	return nil
}

// OKXMLBaseResponse writes v into w with http.StatusOK.
func OKXMLBaseResponse(w http.ResponseWriter, v any) error {
	return XMLBaseResponse(w, http.StatusOK, v)
}

func XMLBaseResponse(w http.ResponseWriter, code int, v any) error {
	return XMLResponseWithCode(w, code, wrapXmlBaseResponse(v))
}

// XMLResponseWithCode writes v as xml string into w with code.
func XMLResponseWithCode(w http.ResponseWriter, code int, v any) error {
	return doWriteXml(w, code, v)
}

func doWriteXml(w http.ResponseWriter, code int, v any) error {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal xml failed, error: %w", err)
	}
	w.Header().Set(HeaderContentType, ContentTypeXML)
	w.WriteHeader(code)
	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
	return nil
}
