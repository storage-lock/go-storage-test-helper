package examples

import (
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"testing"
)

func TestFooStorage(t *testing.T) {
	// 在要测试的Storage实现的仓库中创建好Storage，然后传递给方法测试
	storage_test_helper.TestStorage(t, &FooStorage{})
}
