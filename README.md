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
## Configuration
The configuration system allows you to easily manage and customize the settings for your simulation. Here is an example of how to load and use a configuration file:
``` go
package main

import (
    "github.com/Corphon/dao/config"
    "fmt"
)

func main() {
    // Load configuration from a file
    cfg, err := config.LoadConfig("config.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    fmt.Printf("Loaded config: %+v\n", cfg)
}
```
Example config.json:
``` json
{
    "workers": 4,
    "grid": {
        "width": 100,
        "height": 100
    }
}
```

### Extending the Framework
Dao is designed to be easily extensible. You can create your own implementations of the core interfaces to customize the behavior of your simulation.
---
## Extending Cell
To create a custom cell, simply implement the Cell interface:
``` go
package main

import (
    "github.com/Corphon/dao/core/interfaces"
    "github.com/Corphon/dao/core/util"
)

type CustomCell struct {
    id    string
    state int
}

func NewCustomCell(state int) *CustomCell {
    return &CustomCell{
        id:    util.GenerateID(),
        state: state,
    }
}

func (c *CustomCell) ID() string {
    return c.id
}

func (c *CustomCell) Update() error {
    // Custom update logic
    return nil
}

func (c *CustomCell) GetState() int {
    return c.state
}

func (c *CustomCell) Clone() (interfaces.Cell, error) {
    return &CustomCell{
        id:    util.GenerateID(),
        state: c.state,
    }, nil
}
```
## Extending Grid
Similarly, you can extend the Grid interface to create a custom grid:
``` go
package main

import (
    "github.com/Corphon/dao/core/interfaces"
    "github.com/Corphon/dao/core/util"
)

type CustomGrid struct {
    id    string
    cells [][]interfaces.Cell
    width int
    height int
}

func NewCustomGrid(width, height int) *CustomGrid {
    cells := make([][]interfaces.Cell, height)
    for i := range cells {
        cells[i] = make([]interfaces.Cell, width)
    }
    return &CustomGrid{
        id:    util.GenerateID(),
        cells: cells,
        width: width,
        height: height,
    }
}

func (g *CustomGrid) ID() string {
    return g.id
}

func (g *CustomGrid) Update() error {
    // Custom update logic
    return nil
}

func (g *CustomGrid) GetCell(x, y int) (interfaces.Cell, error) {
    if x < 0 || x >= g.width || y < 0 || y >= g.height {
        return nil, fmt.Errorf("coordinates out of bounds")
    }
    return g.cells[y][x], nil
}

func (g *CustomGrid) SetCell(x, y int, cell interfaces.Cell) error {
    if x < 0 || x >= g.width || y < 0 || y >= g.height {
        return fmt.Errorf("coordinates out of bounds")
    }
    g.cells[y][x] = cell
    return nil
}

func (g *CustomGrid) GetNeighbors(x, y int) ([]interfaces.Cell, error) {
    // Custom logic to get neighbors
    return nil, nil
}

func (g *CustomGrid) GetSize() (width, height int) {
    return g.width, g.height
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
