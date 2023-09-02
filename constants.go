package storage_test_helper

import "time"

// 单元测试统一使用的一些常量
const (

	// TestDatabaseName 测试时使用到的数据库的名称
	TestDatabaseName = "storage_lock_test"

	// TestTableName 测试时使用到的表的名称
	TestTableName = "storage_lock_test"

	// TestLockId 测试时使用的锁的名称
	TestLockId = "lock_id_for_test"

	// TestLockVersion 测试时使用到的版本号
	TestLockVersion = 1

	// TestOwnerIdA 竞争锁的其中一个owner
	TestOwnerIdA = "owner_id_A"
	// TestOwnerIdB 竞争锁的另一个owner
	TestOwnerIdB = "owner_id_B"
)

var (
	// DefaultContextTimeout 操作需要在五分钟内完成
	DefaultContextTimeout = time.Minute * 5
)
