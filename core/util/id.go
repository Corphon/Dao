package util

import (
    "sync/atomic"
    "time"
    "fmt"
)

var sequence uint32

// GenerateID 生成唯一标识符
func GenerateID() string {
    // 使用原子操作增加序列号
    seq := atomic.AddUint32(&sequence, 1)
    timestamp := time.Now().UnixNano()
    
    // 组合时间戳和序列号
    return fmt.Sprintf("%d-%d", timestamp, seq)
}
