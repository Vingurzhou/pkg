package httpz

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Timeout time.Duration
}
type WrapperHttpCli struct {
	httpCli *http.Client
}

func NewHttpCli(config Config) *WrapperHttpCli {
	return &WrapperHttpCli{
		httpCli: &http.Client{Timeout: config.Timeout},
	}
}

func (c *WrapperHttpCli) Do(url, method, payLoadStr string, headerMap map[string]string) ([]byte, error) {
	switch method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		return nil, fmt.Errorf("不支持的方法:%s", method)
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
