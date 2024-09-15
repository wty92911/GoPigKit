package router__test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createFamily(t *testing.T, router *gin.Engine, token string, name string) {
	body := map[string]interface{}{
		"name": name,
	}
	req, _ := http.NewRequest("POST", "/api/v1/family/create", NewJsonReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Log(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestCreateFamily(t *testing.T) {
	router := setupRouter(true)
	token := testWechatLogin(t, router, "test1")
	createFamily(t, router, token, "test1_family")
}

func joinFamily(t *testing.T, router *gin.Engine, token1 string, id string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("PUT", "/api/v1/family/join/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+token1)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Log(w.Body)
	return w
}
func TestJoinFamily(t *testing.T) {
	router := setupRouter(true)
	token := testWechatLogin(t, router, "test1")
	createFamily(t, router, token, "test1_family")
	token2 := testWechatLogin(t, router, "test2")
	w := joinFamily(t, router, token2, "1")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllFamilies(t *testing.T) {
	router := setupRouter(true)
	token1 := testWechatLogin(t, router, "test1")
	createFamily(t, router, token1, "test1_family")
	token2 := testWechatLogin(t, router, "test2")
	createFamily(t, router, token2, "test2_family")
	req, _ := http.NewRequest("GET", "/api/v1/family", nil)
	req.Header.Set("Authorization", "Bearer "+token2)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Log(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
