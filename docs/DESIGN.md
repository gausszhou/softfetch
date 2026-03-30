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
│   │   ├── detectors.go     # 具体检测器实现
│   │   └── detectors_test.go
│   ├── display/             # 显示模块
│   │   └── display.go       # 终端输出格式化
│   ├── command/             # 命令执行模块
│   │   └── command.go       # 系统命令封装
│   └── info/                # 版本信息模块
│       └── info.go
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

| 检测器 | 命令 | 类别 |
|--------|------|------|
| GoDetector | go | Language |
| NodeDetector | node | Language |
| PythonDetector | python3/python | Language |
| JavaDetector | java | Language |
| CDetector | gcc/clang/cc | Language |
| CppDetector | g++/clang++/c++ | Language |
| RustDetector | rustc | Language |
| PHPDetector | php | Language |
| DockerDetector | docker | Runtime |
| GitDetector | git | BuildTool |
| RubyDetector | ruby | Language |
| DotNetDetector | dotnet | Runtime |

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

### 4. info 模块

管理应用版本信息。

#### 变量

- Version: 当前版本

### 5. 并行检测

Detect 函数使用 `sync.WaitGroup` 和 channel 实现并行检测：

```go
func Detect(detectors ...Detector) DetectionResult {
    result := DetectionResult{
        Detected: time.Now(),
        OS:       runtime.GOOS,
        Arch:     runtime.GOARCH,
    }

    resultChan := make(chan Tool, len(detectors))
    var wg sync.WaitGroup

    for _, d := range detectors {
        wg.Add(1)
        go func(detector Detector) {
            defer wg.Done()
            resultChan <- detector.Detect()
        }(d)
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    for tool := range resultChan {
        result.Tools = append(result.Tools, tool)
    }

    return result
}
```

**性能优势：**
- 串行检测：O(n × timeout)
- 并行检测：O(timeout)
- 12 个检测器理论提升 10 倍以上

## 执行流程

```
main()
    │
    ├─→ detect.GetCoreDetectors()    # 获取所有检测器（12个）
    │
    ├─→ detect.Detect(...)           # 并行执行检测
    │       │
    │       ├─→ 启动 goroutine 并行调用 Detect()
    │       │       │
    │       │       └─→ 执行系统命令获取版本信息
    │       │           │
    │       │           └─→ parseVersion() 解析版本号
    │       │
    │       └─→ channel 收集结果
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
