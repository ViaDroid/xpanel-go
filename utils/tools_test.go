package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoBytes(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AutoBytes(1023), "1023B")
	assert.Equal(AutoBytes(1024), "1KB")
	assert.Equal(AutoBytes(2048), "2KB")
	assert.Equal(AutoBytes(2048000), "2MB")
	assert.Equal(AutoBytes(204800000), "195MB")
	assert.Equal(AutoBytes(2048000000), "2GB")
	assert.Equal(AutoBytes(2048000000000), "2TB")
	assert.Equal(AutoBytes(2048000000000000), "2PB")
}
func TestAutoBytesR(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AutoBytesR("1000KB"), 1024000, "Should be the same")
	assert.Equal(AutoBytesR("1024TB"), 1125899906842624, "Should be the same")
}

func TestIsEmailLegal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IsEmailLegal("abcd@gmail.com")["ret"], 1)
}
