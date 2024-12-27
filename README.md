# Dao Framework

Dao is a lightweight, flexible framework designed for grid-based simulations. It focuses on simplicity, modularity, and extensibility, allowing developers to create complex simulations with minimal effort.

## Project Structure

```
Dao/
├── go.mod
├── core/
│   ├── interfaces/
│   │   ├── base.go
│   │   └── grid.go
│   ├── concurrent/
│   │   └── worker.go
│   └── util/
│       └── id.go
├── config/
│   └── config.go
└── examples/
    └── README.md
```

## Core Philosophy

The Dao framework is inspired by the philosophy of "Dao" (道), which emphasizes simplicity, naturalness, and balance. The core principles of the framework are:
1. **Simplicity**: Keep the core interfaces and implementations minimal and intuitive.
2. **Modularity**: Allow for easy extension and customization through well-defined interfaces.
3. **Extensibility**: Provide the flexibility to build complex simulations by extending basic components.

## Features

- **Core Interfaces**: Simple and clear interfaces for entities, cells, and grids.
- **Concurrent Processing**: Efficient task management with a built-in worker pool.
- **Unique Identifiers**: Utility for generating unique IDs.
- **Configuration Management**: Easy-to-use configuration system.

## Getting Started

### Installation

To use the Dao framework, you can simply import it into your Go project:

```bash
go get github.com/Corphon/dao
```
## Basic Usage
Here is a basic example demonstrating how to use the Dao framework to create a simple simulation:

``` go
package main

import (
    "github.com/Corphon/dao/config"
    "github.com/Corphon/dao/core/concurrent"
    "github.com/Corphon/dao/core/interfaces"
    "github.com/Corphon/dao/core/util"
    "fmt"
)

// Define a simple Cell implementation
type SimpleCell struct {
    id    string
    state int
}

func NewSimpleCell(state int) *SimpleCell {
    return &SimpleCell{
        id:    util.GenerateID(),
        state: state,
    }
}

func (c *SimpleCell) ID() string {
    return c.id
}

func (c *SimpleCell) Update() error {
    // Update cell state logic
    return nil
}

func (c *SimpleCell) GetState() int {
    return c.state
}

func (c *SimpleCell) Clone() (interfaces.Cell, error) {
    return &SimpleCell{
        id:    util.GenerateID(),
        state: c.state,
    }, nil
}


func main() {
    // Load configuration
    cfg := config.NewDefaultConfig()

    // Create a worker pool
    pool := concurrent.NewWorkerPool(cfg.Workers)
    defer pool.Stop()

    // Example: Create and update cells
    cell := NewSimpleCell(1)
    task := concurrent.Task{
        ID: util.GenerateID(),
        Fn: cell.Update,
    }

    pool.Submit(task)
    pool.Wait()
    
    fmt.Println("Simulation completed.")
}
```

### Open Source and Contributions

This project is open source and welcomes contributions in any form. You can contribute by:

1. Submitting issues and feature requests
2. Submitting code improvements and optimizations
3. Writing and improving documentation

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Thank you for your interest in the Dao System! We hope this project helps with your research and development.
