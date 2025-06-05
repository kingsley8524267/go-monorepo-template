#!/usr/bin/env bash

set -e

APP_NAME="$1"
INPUT_VERSION="$2"

if [ -z "$APP_NAME" ]; then
  echo "Usage: ./build.bash <app-name> [version]"
  exit 1
fi

# 1. 计算路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APP_DIR="$SCRIPT_DIR/../cmd/$APP_NAME"
VERSION_FILE="$SCRIPT_DIR/.$APP_NAME"

# 2. 自动版本号递增
if [ -z "$INPUT_VERSION" ]; then
  if [ -f "$VERSION_FILE" ]; then
    LAST_VERSION=$(cat "$VERSION_FILE")
    # 去掉 v 前缀
    BASE_VERSION="${LAST_VERSION#v}"
    IFS='.' read -r MAJOR MINOR PATCH <<< "$BASE_VERSION"
    PATCH=$((PATCH + 1))
    VERSION="v$MAJOR.$MINOR.$PATCH"
  else
    VERSION="v0.1.0"
  fi
else
  VERSION="$INPUT_VERSION"
fi

# 3. 保存当前版本
echo "$VERSION" > "$VERSION_FILE"

# 4. 自动构建包路径（用于 ldflags）
# 假设 go module 名是当前项目根目录下的 go.mod 中定义的 module
GO_MOD_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
MODULE_NAME=$(grep '^module ' "$GO_MOD_ROOT/go.mod" | awk '{print $2}')

# 获取 app.Version 的完整路径
REL_APP_PKG=$(realpath --relative-to="$GO_MOD_ROOT" "$APP_DIR")
LD_PACKAGE="$MODULE_NAME/${REL_APP_PKG}/app"

# 5. 构建
echo "Building $APP_NAME version $VERSION..."
go build -ldflags "-X ${LD_PACKAGE}.Version=${VERSION}" -o "$SCRIPT_DIR/../bin/$APP_NAME" "$APP_DIR"

echo "✅ Build complete: $APP_NAME ($VERSION)"
