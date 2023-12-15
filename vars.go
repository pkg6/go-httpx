package httpx

const (
	CharsetUTF8                      = "charset=UTF-8"
	ContentTypeJSON                  = "application/json"
	ContentTypeJSONCharsetUTF8       = ContentTypeJSON + "; " + CharsetUTF8
	ContentTypeJavaScript            = "application/javascript"
	ContentTypeJavaScriptCharsetUTF8 = ContentTypeJavaScript + "; " + CharsetUTF8
	ContentTypeXML                   = "application/xml"
	ContentTypeXMLCharsetUTF8        = ContentTypeXML + "; " + CharsetUTF8
	ContentTypeTextXML               = "text/xml"
	ContentTypeTextXMLCharsetUTF8    = ContentTypeTextXML + "; " + CharsetUTF8
	ContentTypeForm                  = "application/x-www-form-urlencoded"
	ContentTypeProtobuf              = "application/protobuf"
	ContentTypeMsgpack               = "application/msgpack"
	ContentTypeTextHTML              = "text/html"
	ContentTypeTextHTMLCharsetUTF8   = ContentTypeTextHTML + "; " + CharsetUTF8
	ContentTypeTextPlain             = "text/plain"
	ContentTypeTextPlainCharsetUTF8  = ContentTypeTextPlain + "; " + CharsetUTF8
	ContentTypeMultipartForm         = "multipart/form-data"
	ContentTypeOctetStream           = "application/octet-stream"
)

const (
	SchemeHttp  = "http"
	SchemeHttps = "https"
)

// Headers
const (
	HeaderAccept         = "Accept"
	HeaderAcceptEncoding = "Accept-Encoding"

	// HeaderAllow is the name of the "Allow" header field used to list the set of methods
	// advertised as supported by the target resource. Returning an Allow header is mandatory
	// for status 405 (method not found) and useful for the OPTIONS method in responses.
	// See RFC 7231: https://datatracker.ietf.org/doc/html/rfc7231#section-7.4.1

	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderRetryAfter          = "Retry-After"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-Ip"
	HeaderXRequestID          = "X-Request-Id"
	HeaderXCorrelationID      = "X-Correlation-Id"
	HeaderXRequestedWith      = "X-Requested-With"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"
	HeaderCacheControl        = "Cache-Control"
	HeaderConnection          = "Connection"
	HeaderUserAgent           = "User-Agent"

	// Access control

	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

	// Security

	HeaderStrictTransportSecurity         = "Strict-Transport-Security"
	HeaderXContentTypeOptions             = "X-Content-Type-Options"
	HeaderXXSSProtection                  = "X-XSS-Protection"
	HeaderXFrameOptions                   = "X-Frame-Options"
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	HeaderXCSRFToken                      = "X-CSRF-Token"
	HeaderReferrerPolicy                  = "Referrer-Policy"
)
