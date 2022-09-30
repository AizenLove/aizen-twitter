package kikuero_connect

import (
	"fmt"
	"testing"
)

func TestAizenReq(t *testing.T) {

	aizen_beta_query := "お姉さんを調教したい"
	req_url := "https://recommend.aizen.love/query-beta/"
	test_resp, err := AizenReq(aizen_beta_query, req_url)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(test_resp)

}
