package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_Get(t *testing.T, s storage.Storage) {
	assert.NotNilf(t, s, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "Storage %s init error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	TestEnsureLockNotExists(t, s, TestLockId)

	lockInformation := BuildTestLockInformation(t)
	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	err = s.CreateWithVersion(ctx2, TestLockId, TestLockVersion, lockInformation)
	assert.Nilf(t, err, "Storage %s insert error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	ctx3, cancelFunc3 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc3()
	lockInformationJsonStringRs, err := s.Get(ctx3, TestLockId)
	assert.Nilf(t, err, "Storage %s get lock error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))
	assert.Equalf(t, lockInformation.ToJsonString(), lockInformationJsonStringRs, "Storage %s get lock not equals", s.GetName())

	t.Log(fmt.Sprintf("Storage %s test Get done", s.GetName()))

}
