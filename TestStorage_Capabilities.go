package storage_test_helper

import (
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStorage_Capabilities 测试 Storage 实现的能力声明
func TestStorage_Capabilities(t *testing.T, s storage.Storage) {

	capabilities := s.Capabilities()

	// Capabilities 不应为空
	assert.NotEmpty(t, capabilities, "Storage Capabilities() should not return empty slice")

	// 检查是否声明了必要的能力
	hasCAS := false
	hasReliableTime := false
	for _, c := range capabilities {
		if c == storage.CapabilityCAS {
			hasCAS = true
		}
		if c == storage.CapabilityReliableTime {
			hasReliableTime = true
		}
	}

	// 如果缺少必要能力，给出警告（但不强制失败，因为 FileSystemStorage 等实现确实不支持 CAS）
	if !hasCAS {
		t.Logf("WARNING: Storage %s does not declare CapabilityCAS, distributed lock mutual exclusion cannot be guaranteed", s.GetName())
	}
	if !hasReliableTime {
		t.Logf("WARNING: Storage %s does not declare CapabilityReliableTime, clock skew may break lock correctness", s.GetName())
	}

	t.Logf("Storage %s capabilities: %v", s.GetName(), capabilities)
}
