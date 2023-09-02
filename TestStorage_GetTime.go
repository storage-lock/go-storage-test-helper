package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_GetTime(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "Storage %s init error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	time, err := s.GetTime(ctx2)
	assert.Nilf(t, err, "Storage %s GetTime error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))
	assert.Falsef(t, time.IsZero(), "Storage %s GetTime return zero time", s.GetName())

	t.Log(fmt.Sprintf("Storage %s test GetTime done", s.GetName()))
}
