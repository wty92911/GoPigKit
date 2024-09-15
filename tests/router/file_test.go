package router__test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func uploadFile(t *testing.T, router *gin.Engine, token1 string) *httptest.ResponseRecorder {
	// 打开文件
	file, err := os.Open("./piggy.png")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// 创建一个缓冲区和multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 写入文件
	part, err := writer.CreateFormFile("file", "piggy.png")
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("failed to copy file content: %v", err)
	}

	// 写入路径
	err = writer.WriteField("path", "/test.png")
	if err != nil {
		t.Fatalf("failed to write field: %v", err)
	}

	// 关闭writer
	err = writer.Close()
	if err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", "/api/v1/file", body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token1)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestUploadFileWithoutFamily(t *testing.T) {
	router := setupRouter(true)
	token1 := testWechatLogin(t, router, "test1")
	w := uploadFile(t, router, token1)
	t.Log(w.Body)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestUploadFileWithFamily(t *testing.T) {
	router := setupRouter(true)
	token1 := testWechatLogin(t, router, "test1")
	createFamily(t, router, token1, "test1_family")
	token2 := testWechatLogin(t, router, "test2")
	joinFamily(t, router, token2, "1")
	w := uploadFile(t, router, token2)
	t.Log(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func deleteFile(t *testing.T, router *gin.Engine, token1 string, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/api/v1/file/delete", NewJsonReader(map[string]string{"url": url}))
	req.Header.Set("Authorization", "Bearer "+token1)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Log(w.Body)
	return w
}

func TestDeleteFileWithoutFamily(t *testing.T) {
	router := setupRouter(true)
	token1 := testWechatLogin(t, router, "test1")
	w := deleteFile(t, router, token1, "http://81.70.53.202:9000/test/1/test")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeleteFileWithFamily(t *testing.T) {
	router := setupRouter(true)
	token1 := testWechatLogin(t, router, "test1")
	createFamily(t, router, token1, "test1_family")
	token2 := testWechatLogin(t, router, "test2")
	joinFamily(t, router, token2, "1")
	w := uploadFile(t, router, token2)
	t.Log(w.Body)
	data := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &data)
	url := data["data"].(string)
	w = deleteFile(t, router, token2, url)
	assert.Equal(t, http.StatusOK, w.Code)
}
