package storage_test_helper

import (
	"context"
	"fmt"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_CreateWithVersion(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "Storage %s init error: %s", s.GetName(), errText(err))

	TestEnsureLockNotExists(t, s, TestLockId)

	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	err = s.CreateWithVersion(ctx2, TestLockId, TestLockVersion, BuildTestLockInformation(t))
	assert.Nilf(t, err, "Storage %s insert error: %s", s.GetName(), errText(err))

	t.Log(fmt.Sprintf("Storage %s test CreateWithVersion done", s.GetName()))
}
