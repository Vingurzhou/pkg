package httpz

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type WrapperHttpCli struct {
	httpCli *http.Client
}

func NewHttpCli() *WrapperHttpCli {
	return &WrapperHttpCli{
		httpCli: &http.Client{},
	}
}

func (c *WrapperHttpCli) Do(url, method, payLoadStr string, headerMap map[string]string) ([]byte, error) {
	switch method {
	case "POST":
	case "GET":
	default:
		return nil, errors.New("不支持的方法")
	}
	req, err := http.NewRequest(method, url, strings.NewReader(payLoadStr))
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	res, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
