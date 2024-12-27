package interfaces

// Cell 定义细胞接口
type Cell interface {
    Entity
    Updateable
    // GetState 获取细胞状态
    GetState() int
    // Clone 克隆细胞
    Clone() (Cell, error)
}

// Grid 定义网格接口
type Grid interface {
    Entity
    Updateable
    // GetCell 获取指定位置的细胞
    GetCell(x, y int) (Cell, error)
    // SetCell 设置指定位置的细胞
    SetCell(x, y int, cell Cell) error
    // GetNeighbors 获取指定位置的邻居细胞
    GetNeighbors(x, y int) ([]Cell, error)
    // GetSize 获取网格大小
    GetSize() (width, height int)
}
