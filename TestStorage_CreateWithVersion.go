package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_CreateWithVersion(t *testing.T, s storage.Storage) {
	assert.NotNilf(t, s, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "storage %s init error: %#v", s.GetName(), err)

	TestEnsureLockNotExists(t, s, TestLockId)

	err = s.CreateWithVersion(ctx, TestLockId, TestLockVersion, BuildTestLockInformation(t))
	assert.Nilf(t, err, "storage %s insert error: %#v", s.GetName(), err)
}
