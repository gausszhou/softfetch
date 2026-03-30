# SoftFetch 设计文档

## 项目概述

SoftFetch 是一个用 Go 编写的命令行工具，用于检测系统上安装的开发工具和编程语言，并展示其版本和安装路径。

## 项目结构

```
softfetch/
├── main.go                  # 主入口程序
├── main_test.go             # 主程序测试
├── internal/
│   ├── detect/              # 检测模块
│   │   ├── types.go         # 类型定义和核心接口
│   │   ├── detectors.go    # 具体检测器实现
│   │   └── detectors_test.go
│   ├── display/             # 显示模块
│   │   └── display.go       # 终端输出格式化
│   ├── command/             # 命令执行模块
│   │   └── command.go       # 系统命令封装
│   └── version/             # 版本信息模块
│       ├── version.go
│       └── version_test.go
└── docs/
    └── DESIGN.md            # 本文档
```

## 核心模块

### 1. detect 模块

负责检测系统中安装的开发工具。

#### 类型定义

- **Tool**: 表示一个检测到的工具
  - Name: 工具名称
  - Version: 版本号
  - Path: 安装路径
  - Detected: 是否检测到
  - Symbol: 展示符号
  - Category: 工具类别

- **DetectionResult**: 检测结果
  - Tools: 工具列表
  - OS: 操作系统
  - Arch: 系统架构
  - Detected: 检测时间

- **Category**: 工具类别常量
  - Language: 编程语言
  - PackageMgr: 包管理器
  - BuildTool: 构建工具
  - Compiler: 编译器
  - Runtime: 运行时
  - Other: 其他

#### Detector 接口

所有检测器实现以下接口：

```go
type Detector interface {
    Detect() Tool
    Name() string
    Category() Category
}
```

#### 内置检测器

- GoDetector: 检测 Go
- NodeDetector: 检测 Node.js
- PythonDetector: 检测 Python
- JavaDetector: 检测 Java
- CDetector: 检测 C (gcc/clang/cc)
- CppDetector: 检测 C++ (g++/clang++/c++)
- RustDetector: 检测 Rust
- PHPDetector: 检测 PHP

### 2. display 模块

负责在终端中以美观的格式展示检测结果。

#### 功能

- PrintResult: 以带颜色和边框的表格形式打印结果
- PrintSimple: 以简单列表形式打印结果

#### 颜色方案

- 青色 (Cyan): 边框
- 蓝色 (Blue): 标题
- 黄色 (Yellow): 工具名称
- 绿色 (Green): 已安装版本
- 红色 (Red): 未安装
- 灰色 (Gray): 路径和系统信息

### 3. command 模块

封装系统命令执行功能。

#### 功能

- Execute: 执行命令（默认 5 秒超时）
- ExecuteWithTimeout: 执行命令（自定义超时）
- LookPath: 查找命令路径
- Getenv/GetenvOrDefault: 环境变量操作

### 4. version 模块

管理应用版本信息。

#### 变量

- Version: 当前版本
- BuildDate: 构建日期
- GitCommit: Git 提交哈希

#### 函数

- GetVersion: 获取版本字符串
- GetBuildInfo: 获取完整的构建信息

## 执行流程

```
main()
    │
    ├─→ detect.GetCoreDetectors()    # 获取所有检测器
    │
    ├─→ detect.Detect(...)           # 执行检测
    │       │
    │       └─→ 遍历检测器，调用 Detect() 方法
    │           │
    │           └─→ 执行系统命令获取版本信息
    │               │
    │               └─→ parseVersion() 解析版本号
    │
    ├─→ display.PrintResult()        # 打印检测结果
    │
    └─→ 检查命令行参数
            │
            └─→ --version / -v: 打印版本信息
```

## 扩展开发

### 添加新的检测器

1. 在 `internal/detect/detectors.go` 中创建新的检测器结构体
2. 嵌入 `baseDetector` 并实现 `Detect()` 方法
3. 在 `GetCoreDetectors()` 中添加新检测器

示例：

```go
type RubyDetector struct {
    baseDetector
}

func NewRubyDetector() *RubyDetector {
    return &RubyDetector{
        baseDetector: baseDetector{
            name:     "Ruby",
            category: CategoryLanguage,
        },
    }
}

func (d *RubyDetector) Detect() Tool {
    // 实现检测逻辑
}
```

### 添加新的显示模式

在 `internal/display/display.go` 中添加新的打印函数。

## 命令行参数

- 无参数: 运行检测并显示结果
- `--version` 或 `-v`: 显示版本信息

## 依赖

- Go 标准库（无外部依赖）
