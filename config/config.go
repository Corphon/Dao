package config

import (
    "encoding/json"
    "os"
)

// Config 定义配置结构
type Config struct {
    Workers int `json:"workers"`
    Grid    struct {
        Width  int `json:"width"`
        Height int `json:"height"`
    } `json:"grid"`
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    config := &Config{}
    if err := json.Unmarshal(data, config); err != nil {
        return nil, err
    }

    return config, nil
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
    return &Config{
        Workers: 4,
        Grid: struct {
            Width  int `json:"width"`
            Height int `json:"height"`
        }{
            Width:  100,
            Height: 100,
        },
    }
}
