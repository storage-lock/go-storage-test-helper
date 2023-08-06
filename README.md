# Storage Test Helper

# 一、这是什么

用于辅助测试[Storage](https://github.com/storage-lock/go-storage)的实现是否OK的测试工具，辅助提高开发效率。

# 二、安装依赖

```bash
go get -u github.com/storage-lock/go-storage-test-helper
```

# 三、如何使用

## 3.1 添加依赖

在你自己的Storage实现的差不多了的时候，在此项目中执行：

```bash
go get -u github.com/storage-lock/go-storage-test-helper
```

把此测试库添加到你自己的Storage的依赖中。

## 3.2 创建单元测试

为你的Storage创建一个单元测试，比如下面是一个MySQL Storage的单元测试：

```go
package mysql_storage

import (
	"context"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewMySQLStorage(t *testing.T) {
	envName := "STORAGE_LOCK_MYSQL_DSN"
	dsn := os.Getenv(envName)
	assert.NotEmpty(t, dsn)
	connectionGetter := NewMySQLConnectionManagerFromDSN(dsn)
	s, err := NewMySQLStorage(context.Background(), &MySQLStorageOptions{
		ConnectionManager: connectionGetter,
		TableName:         "storage_lock_test",
	})
	assert.Nil(t, err)
    // 重点在与这一句，把 *testing.T 和 storage.Storage 传进去 
	storage_test_helper.TestStorage(t, s)
}
```

保证此单元测试通过，你可以在CI中执行单元测试以保证你每次修改之后Storage都能够正常工作。
