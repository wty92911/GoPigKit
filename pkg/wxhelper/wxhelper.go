package wxhelper

import (
	"encoding/json"
	"fmt"
	"github.com/wty92911/GoPigKit/configs"
	"io"
	"net/http"
	"net/url"
)

const (
	// Code2Session 微信 code2session 接口
	Code2SessionAPI = "https://api.weixin.qq.com/sns/jscode2session"
)

// Code2SessionResponse 微信 code2session 响应结构体
type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session 请求微信 code2session 接口
func Code2Session(config *configs.AppConfig, jsCode string) (*Code2SessionResponse, error) {
	params := url.Values{}
	params.Add("appid", config.ID)
	params.Add("secret", config.Secret)
	params.Add("js_code", jsCode)
	params.Add("grant_type", "authorization_code")

	resp, err := http.Get(fmt.Sprintf("%s?%s", Code2SessionAPI, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Code2SessionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("wechat code2session error: %s", result.ErrMsg)
	}

	return &result, nil
}
