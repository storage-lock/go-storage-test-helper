package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_Get(t *testing.T, s storage.Storage) {
	assert.NotNilf(t, s, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "storage %s init error: %#v", s.GetName(), err)

	TestEnsureLockNotExists(t, s, TestLockId)

	lockInformation := BuildTestLockInformation(t)
	err = s.CreateWithVersion(ctx, TestLockId, TestLockVersion, lockInformation)
	assert.Nilf(t, err, "storage %s insert error: %#v", s.GetName(), err)

	lockInformationJsonStringRs, err := s.Get(ctx, TestLockId)
	assert.Nilf(t, err, "storage %s get lock error: %#v", s.GetName(), err)
	assert.Equalf(t, lockInformation.ToJsonString(), lockInformationJsonStringRs, "storage %s get lock not equals", s.GetName())

}
