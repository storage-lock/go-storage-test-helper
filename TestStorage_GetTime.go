package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_GetTime(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "storage %s init error: %#v", s.GetName(), err)

	time, err := s.GetTime(context.Background())
	assert.Nilf(t, err, "storage %s GetTime error: %#v", s.GetName(), err)
	assert.Falsef(t, time.IsZero(), "storage %s GetTime return zero time", s.GetName())
}
