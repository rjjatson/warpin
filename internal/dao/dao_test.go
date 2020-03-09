package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotfMessageDAO(t *testing.T) {
	nm := New()
	assert.NotNil(t, nm)
}

func TestStoreSuccess(t *testing.T) {
	nm := New()
	err := nm.Store("test")

	assert.Nil(t, err)
	if len(nm.storage) != 1 {
		assert.FailNow(t, "fail storing message")
	}
	assert.Equal(t, "test", nm.storage[0])
}

func TestGetAllSuccess(t *testing.T) {
	nm := New()
	notif := []string{"test-1", "test-2", "test-3"}
	nm.storage = notif
	actual, err := nm.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, notif, actual)
}
