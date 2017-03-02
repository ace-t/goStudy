package main
//
//import (
//	"net/url"
//	"net/http"
//	"io/ioutil"
//	"encoding/json"
//)


//func StockInfoRequest(query string) (*StockResult, error) {
//	// build url
//	baseURL := "http://finance.daum.net/item/company.daum"
//	u, err := url.Parse(baseURL)
//	if err != nil {
//		return nil, err
//	}
//	q := u.Query()
//	q.Set("q", query)
//	u.RawQuery = q.Encode()
//
//	// request
//	resp, err := http.Get(u.String())
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//
//	// parse json
//	var result StockResult
//	if err := json.Unmarshal(body, &result); err != nil {
//		return nil, err
//	}
//	return &result, nil
//}