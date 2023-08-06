package storage_test_helper

import (
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStorage_GetName(t *testing.T, storage storage.Storage) {
	assert.NotNilf(t, storage, "storage is nil")

	// 名称不能为空
	assert.NotEmptyf(t, storage.GetName(), "storage name is empty")

	// 名称两侧不能有空格
	assert.Truef(t, storage.GetName() == strings.TrimSpace(storage.GetName()), "Storage name \"%s\" contains whitespace character, it is not allow", storage.GetName())
}
