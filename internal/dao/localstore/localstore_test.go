package localstore

import (
	"testing"
	"time"
	"warpin/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotfMessageDAO(t *testing.T) {
	nm := New()
	assert.NotNil(t, nm)
}

func TestStoreSuccess(t *testing.T) {
	nm := New()
	timeNow := time.Now().UTC()
	err := nm.Store("test", timeNow)

	assert.Nil(t, err)
	if len(nm.storage) != 1 {
		assert.FailNow(t, "fail storing message")
	}
	assert.Equal(t, model.NotificationStore{Message: "test", Time: timeNow}, nm.storage[0])
}

func TestGetAllSuccess(t *testing.T) {
	nm := New()
	notif := []model.NotificationStore{
		{Message: "test1",
			Time: time.Now().UTC()},
		{Message: "test2",
			Time: time.Now().UTC().Add(time.Second)},
	}
	nm.storage = notif
	actual, err := nm.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, notif, actual)
}
