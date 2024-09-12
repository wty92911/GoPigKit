package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTrimUrl(t *testing.T) {
	url := "http://127.0.0.1:9000/GoPigKit/1619160061.png"
	endpoint := "http://127.0.0.1:9000"
	bucket := "GoPigKit"
	assert.Equal(t, strings.TrimPrefix(url, fmt.Sprintf("%s/%s/", endpoint, bucket)), "1619160061.png")
}
