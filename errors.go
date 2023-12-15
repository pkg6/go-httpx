package httpx

import "fmt"

// NewErrMsg creates a new CodeMsg.
func NewErrMsg(code int, msg string) error {
	return &CodeMsg{Code: code, Msg: msg}
}

type CodeMsg struct {
	Code int
	Msg  string
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}
