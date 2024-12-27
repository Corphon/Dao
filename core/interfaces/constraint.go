package interfaces

// Constraint 约束接口
type Constraint interface {
    // ID 获取约束ID
    ID() string
    
    // Evaluate 评估约束条件
    Evaluate(state int) float64
    
    // Apply 应用约束
    Apply(cell Cell) error
}

// ConstraintSystem 约束系统接口
type ConstraintSystem interface {
    // ID 获取系统ID
    ID() string
    
    // AddConstraint 添加约束
    AddConstraint(constraint Constraint) error
    
    // ApplyConstraints 应用所有约束
    ApplyConstraints(cell Cell) error
}
