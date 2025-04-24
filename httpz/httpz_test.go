package httpz

import (
	"testing"
)

func TestWrapperHttpCli_Do(t *testing.T) {
	cli := NewHttpCli()
	resp, err := cli.Do("http://127.0.0.1:2531/v2/api/tools/getTokenId", "post", "", nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(resp))
}
