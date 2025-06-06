# Go Monorepo 模板 - 目录结构说明

这个仓库是一个用于管理多个 Go 应用的 monorepo 模板。下面是每个主要文件夹的作用：

## 目录结构说明

```json lines
├───apps                # 存放每个独立应用的核心业务逻辑。│
├───cmd                 # 存放每个可执行应用的入口文件 `main.go`。
├───configs             # 存放每个应用独立的配置文件（如 YAML 格式）。
├───deploy              # 存放所有应用的部署相关文件（如 Dockerfile 和 Kubernetes 配置）。
├───internal            # 存放模块内部使用，不对外暴露的包。
│   ├───config          # 负责配置的加载和公共配置结构定义。
│   │       config.go   # 通用的配置加载逻辑（如 Viper 设置）。
│   ├───logger          # 集中式日志工具。
│   └───signal          # 操作系统信号处理，用于应用优雅关闭。│
├───scripts             # 存放开发和运维辅助脚本。
│       build.ps1       # Windows 平台构建应用的 PowerShell 脚本。
│       build.sh        # Linux/macOS 平台构建应用的 Bash 脚本。│
└───tools               # 存放开发工具，如代码生成器。
    └───generate        # 用于生成新应用骨架的代码生成工具。

```

---

## 快速开始

要开始使用这个模板，你可以利用内置的代码生成工具来快速创建一个新的应用。

1.  **构建生成工具：**

    ```bash
    go build -o generate.exe .\tools\generate\
    ```

    （在 Linux/macOS 上，你可以使用 `go build -o generate ./tools/generate/`）

2.  **创建新应用：**

    ```bash
    .\generate.exe new <你的应用名称>
    ```

    将 `<你的应用名称>` 替换为你希望创建的应用名称（例如 `my-new-service`）。这个命令将自动在 `apps`、`cmd`、`configs`、`deploy` 和 `internal/config` 目录下创建相应的骨架文件。