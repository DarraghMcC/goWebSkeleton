package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort_default(t *testing.T) {
	_ = os.Setenv("PORT", "")

	assert.Equal(t, getPort(), ":8090")
}

func TestGetPort_overrideValue(t *testing.T) {
	_ = os.Setenv("PORT", "10000")

	assert.Equal(t, getPort(), ":10000")
}

func TestIsTimeStampDisabled_default(t *testing.T) {
	_ = os.Setenv("DISABLE_LOG_TIMESTAMP", "")

	assert.False(t, isTimeStampDisabled())
}

func TestIsTimeStampDisabled_override(t *testing.T) {
	_ = os.Setenv("DISABLE_LOG_TIMESTAMP", "true")

	assert.True(t, isTimeStampDisabled())
}

func TestIsTimeStampDisabled_invalidValue(t *testing.T) {
	_ = os.Setenv("DISABLE_LOG_TIMESTAMP", "bleh")

	assert.False(t, isTimeStampDisabled())
}
