package interfaces

// Entity 定义基础实体
type Entity interface {
    // ID 返回实体的唯一标识符
    ID() string
}

// Updateable 定义可更新接口
type Updateable interface {
    // Update 更新实体状态
    Update() error
}
