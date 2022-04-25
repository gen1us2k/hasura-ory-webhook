package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	os.Setenv("ORY_SDK_URL", "hello")
	h, err := Parse()
	assert.NoError(t, err)
	assert.Equal(t, "hello", h.OrySDKURL)

}
