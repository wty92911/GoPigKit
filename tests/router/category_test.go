package router__test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createCategory(t *testing.T, router *gin.Engine, token1 string, body map[string]interface{}) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/api/v1/category", NewJsonReader(body))
	req.Header.Set("Authorization", "Bearer "+token1)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Log(w.Body)
	return w
}
func TestCreateCategory(t *testing.T) {
	router := setupRouter(true)
	token := testWechatLogin(t, router, "test1")
	createFamily(t, router, token, "test1_family")

	category := map[string]interface{}{
		"top_name":  "test1_top",
		"mid_name":  "test1_mid",
		"name":      "test1_category",
		"image_url": "https://www.baidu.com",
	}
	w := createCategory(t, router, token, category)
	assert.Equal(t, http.StatusOK, w.Code)
}
