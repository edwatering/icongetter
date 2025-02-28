# IconGetter

IconGetter是一个多平台应用图标提取工具，支持从Android应用(.apk)、iOS应用(.ipa)和Windows程序(.exe)中提取高质量图标。

![License](https://img.shields.io/github/license/edwatering/icongetter)
![Go Version](https://img.shields.io/badge/Go-1.20+-blue)

## ✨ 功能特点

- 🤖 支持从Android应用包(.apk)中提取图标
- 🍎 支持从iOS应用包(.ipa)中提取图标
- 🪟 支持从Windows可执行文件(.exe)中提取图标
- 🖥️ 双击即可直接对当前目录下文件自动提取，也可命令行窗口使用

## 📥 安装

### 预编译版本

从[Releases](https://github.com/edwatering/icongetter/releases)页面下载最新版本的预编译可执行文件。

### 从源码编译

需要安装Go 1.20或更高版本。

```bash
# 克隆仓库
git clone https://github.com/edwatering/icongetter.git

# 进入项目目录
cd icongetter

# 编译
go build -ldflags "-s -w" -o IconGetter.exe
```

## 🚀 使用方法

### 基本用法

1. 双击运行IconGetter.exe
2. 程序将自动扫描当前目录中的.apk、.ipa和.exe文件
3. 提取的图标将保存在`icons`文件夹中，文件名格式为`{原文件名}_icon.png`

### 命令行参数

IconGetter也支持通过命令行运行，提供更多高级选项：

```bash
# 指定目录扫描
extract --dir=/path/to/apps

# 指定单个文件
extract --file=/path/to/app.apk

# 指定输出目录
extract --output=/path/to/output

```

## 📸 截图

![Screenshot](/images/1.png)

## 🔧 支持的平台

### 运行环境
- Windows 8/10/11

### 可提取图标的文件类型
- Android应用包(.apk)
- iOS应用包(.ipa)
- Windows可执行文件(.exe)

---

**注意**：IconGetter仅供学习和个人使用，请尊重应用开发者的知识产权，不要将提取的图标用于商业用途。