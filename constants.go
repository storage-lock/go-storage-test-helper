package storage_test_helper

import "time"

// 单元测试统一使用的一些常量
const (
	TestDatabaseName = "storage_lock_test"
	TestTableName    = "storage_lock_test"

	TestLockId      = "lock_id_for_test"
	TestLockVersion = 1

	TestOwnerIdA = "owner_id_A"
	TestOwnerIdB = "owner_id_B"
)

var (
	// DefaultContextTimeout 操作需要在五分钟内完成
	DefaultContextTimeout = time.Minute * 5
)
