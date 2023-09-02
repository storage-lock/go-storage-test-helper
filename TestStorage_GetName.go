package storage_test_helper

import (
	"fmt"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStorage_GetName(t *testing.T, storage storage.Storage) {
	assert.NotNilf(t, storage, "Storage is nil")

	// 名称不能为空
	assert.NotEmptyf(t, storage.GetName(), "Storage name is empty")

	// 名称两侧不能有空格
	assert.Truef(t, storage.GetName() == strings.TrimSpace(storage.GetName()), "Storage name \"%s\" contains whitespace character, it is not allow", storage.GetName())

	t.Log(fmt.Sprintf("Storage %s test GetName done", storage.GetName()))
}
