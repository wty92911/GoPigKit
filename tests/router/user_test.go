package router__test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const code = "0e3pDbGa1GGnaI0WJAHa1nSKgJ2pDbGr"

func NewJsonReader(body interface{}) io.Reader {
	b, _ := json.Marshal(body)
	return bytes.NewBuffer(b)
}
func testWechatLogin(t *testing.T, router *gin.Engine, openId string) string {
	w := httptest.NewRecorder()

	reqBody := map[string]interface{}{
		"code": code,
		"user_info": map[string]interface{}{
			"open_id":    openId,
			"nickname":   openId,
			"avatar_url": "https://www.baidu.com",
		},
	}
	req, _ := http.NewRequest("POST", "/login", NewJsonReader(reqBody))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Nil(t, err)
	return resp["token"].(string)
}

func TestWechatLogin(t *testing.T) {
	router := setupRouter(true)
	testWechatLogin(t, router, "test_1")
	testWechatLogin(t, router, "test_2")
}
func TestGetUserInfo(t *testing.T) {
	router := setupRouter(true)
	token := testWechatLogin(t, router, "test1")

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
