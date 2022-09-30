package kikuero_connect

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AizenResponce struct {
	BaseUrl  string   `json:"base_url"`
	Describe string   `json:"describe"`
	Id       string   `json:"id"`
	Image    string   `json:"image"`
	Tags     []string `json:"tags"`
	Title    string   `json:"title"`
}

func AizenReq(query string, req_url string) (AizenResponce, error) {
	// リクエストの作成
	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		log.Panic(err)
		return AizenResponce{}, err
	}

	// クエリの追加
	q := req.URL.Query()
	q.Add("user_req", query)
	req.URL.RawQuery = q.Encode()

	// クエリの送信、レスポンスの取得
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
		return AizenResponce{}, err
	}
	defer resp.Body.Close()

	// レスポンスを構造体として扱う
	res_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		return AizenResponce{}, err
	}
	var aizen_resp AizenResponce // レスポンスを格納する変数

	if err := json.Unmarshal(res_body, &aizen_resp); err != nil {
		log.Panic(err)
		return AizenResponce{}, err
	}

	return aizen_resp, nil
}
