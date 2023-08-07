package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_UpdateWithVersion(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "s %s initialization error: %#v", s.GetName(), err)

	TestEnsureLockNotExists(t, s, TestLockId)

	err = s.CreateWithVersion(ctx, TestLockId, TestLockVersion, BuildTestLockInformation(t))
	assert.Nilf(t, err, "s %s insert error: %#v", s.GetName(), err)

	newVersion := storage.Version(TestLockVersion + 1)
	err = s.UpdateWithVersion(ctx, TestLockId, TestLockVersion, newVersion, BuildTestLockInformation(t, newVersion))
	assert.Nilf(t, err, "s %s update error: %#v", s.GetName(), err)
}
