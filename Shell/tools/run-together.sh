#!/bin/bash

# 前后端并发启动脚本
# 自动检测项目结构并使用 concurrently 同时运行前后端

echo "=========================================="
echo "前后端并发启动脚本 v1.0"
echo "=========================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

# 检查是否在项目根目录
if [ ! -f "package.json" ]; then
    echo -e "${RED}错误: 未找到 package.json 文件${NC}"
    echo "请在项目根目录下运行此脚本"
    exit 1
fi

# 前端目录可能的名称（按优先级排序）
FRONTEND_PATTERNS=(
    "frontend"
    "front-end"
    "client"
    "web"
    "app"
    "ui"
    "src/frontend"
)

# 后端目录可能的名称（按优先级排序）
BACKEND_PATTERNS=(
    "backend"
    "back-end"
    "server"
    "api"
    "service"
    "src/backend"
)

# 查找目录的函数
find_directory() {
    local patterns=("$@")
    for pattern in "${patterns[@]}"; do
        # 精确匹配
        if [ -d "$pattern" ] && [ -f "$pattern/package.json" ]; then
            echo "$pattern"
            return 0
        fi
        
        # 模糊匹配（不区分大小写）
        for dir in */; do
            dir=${dir%/}  # 移除末尾的斜杠
            if [[ "${dir,,}" =~ ${pattern,,} ]] && [ -f "$dir/package.json" ]; then
                echo "$dir"
                return 0
            fi
        done
    done
    return 1
}

echo -e "${BLUE}[1/5] 检测项目结构...${NC}"
echo ""

# 查找前端目录
FRONTEND_DIR=$(find_directory "${FRONTEND_PATTERNS[@]}")
if [ -n "$FRONTEND_DIR" ]; then
    echo -e "${GREEN}✓ 找到前端目录: $FRONTEND_DIR${NC}"
else
    echo -e "${RED}✗ 未找到前端目录${NC}"
fi

# 查找后端目录
BACKEND_DIR=$(find_directory "${BACKEND_PATTERNS[@]}")
if [ -n "$BACKEND_DIR" ]; then
    echo -e "${GREEN}✓ 找到后端目录: $BACKEND_DIR${NC}"
else
    echo -e "${RED}✗ 未找到后端目录${NC}"
fi

echo ""

# 如果都没找到，尝试手动输入
if [ -z "$FRONTEND_DIR" ] && [ -z "$BACKEND_DIR" ]; then
    echo -e "${RED}错误: 未能自动检测到前后端目录${NC}"
    echo "当前目录结构："
    ls -d */ 2>/dev/null
    echo ""
    read -p "请输入前端目录名称（或按回车跳过）: " FRONTEND_DIR
    read -p "请输入后端目录名称（或按回车跳过）: " BACKEND_DIR
fi

# 验证至少有一个目录
if [ -z "$FRONTEND_DIR" ] && [ -z "$BACKEND_DIR" ]; then
    echo -e "${RED}错误: 至少需要指定一个目录${NC}"
    exit 1
fi

echo -e "${BLUE}[2/5] 检查 concurrently 依赖...${NC}"
if ! npm list concurrently > /dev/null 2>&1; then
    echo -e "${YELLOW}未安装 concurrently，正在安装...${NC}"
    npm install -D concurrently
    if [ $? -ne 0 ]; then
        echo -e "${RED}concurrently 安装失败${NC}"
        exit 1
    fi
    echo -e "${GREEN}✓ concurrently 安装完成${NC}"
else
    echo -e "${GREEN}✓ concurrently 已安装${NC}"
fi
echo ""

echo -e "${BLUE}[3/5] 检测启动命令...${NC}"

# 检测前端启动命令
FRONTEND_CMD=""
if [ -n "$FRONTEND_DIR" ] && [ -f "$FRONTEND_DIR/package.json" ]; then
    if grep -q '"dev"' "$FRONTEND_DIR/package.json"; then
        FRONTEND_CMD="npm run dev"
    elif grep -q '"start"' "$FRONTEND_DIR/package.json"; then
        FRONTEND_CMD="npm start"
    fi
    echo -e "${GREEN}✓ 前端命令: $FRONTEND_CMD (在 $FRONTEND_DIR 目录)${NC}"
fi

# 检测后端启动命令
BACKEND_CMD=""
if [ -n "$BACKEND_DIR" ] && [ -f "$BACKEND_DIR/package.json" ]; then
    if grep -q '"dev"' "$BACKEND_DIR/package.json"; then
        BACKEND_CMD="npm run dev"
    elif grep -q '"start"' "$BACKEND_DIR/package.json"; then
        BACKEND_CMD="npm start"
    fi
    echo -e "${GREEN}✓ 后端命令: $BACKEND_CMD (在 $BACKEND_DIR 目录)${NC}"
fi
echo ""

echo -e "${BLUE}[4/5] 配置 package.json scripts...${NC}"

# 备份 package.json
cp package.json package.json.backup
echo -e "${YELLOW}已备份 package.json 到 package.json.backup${NC}"

# 读取现有的 package.json
PACKAGE_JSON=$(cat package.json)

# 构建 concurrently 命令
CONCURRENT_COMMANDS=""
NAME_PREFIX=""

if [ -n "$FRONTEND_DIR" ] && [ -n "$FRONTEND_CMD" ]; then
    CONCURRENT_COMMANDS="\"--prefix $FRONTEND_DIR $FRONTEND_CMD\""
    NAME_PREFIX="\"frontend\""
fi

if [ -n "$BACKEND_DIR" ] && [ -n "$BACKEND_CMD" ]; then
    if [ -n "$CONCURRENT_COMMANDS" ]; then
        CONCURRENT_COMMANDS="$CONCURRENT_COMMANDS \"--prefix $BACKEND_DIR $BACKEND_CMD\""
        NAME_PREFIX="$NAME_PREFIX \"backend\""
    else
        CONCURRENT_COMMANDS="\"--prefix $BACKEND_DIR $BACKEND_CMD\""
        NAME_PREFIX="\"backend\""
    fi
fi

# 使用 Node.js 来修改 package.json（更可靠）
node << EOF
const fs = require('fs');
const packageJson = JSON.parse(fs.readFileSync('package.json', 'utf8'));

if (!packageJson.scripts) {
    packageJson.scripts = {};
}

// 添加并发启动命令
packageJson.scripts['dev'] = "concurrently --names $NAME_PREFIX --prefix-colors \"cyan,magenta\" $CONCURRENT_COMMANDS";
packageJson.scripts['dev:frontend'] = "$FRONTEND_DIR && $FRONTEND_CMD";
packageJson.scripts['dev:backend'] = "$BACKEND_DIR && $BACKEND_CMD";

// 添加安装依赖命令
if ("$FRONTEND_DIR" && "$BACKEND_DIR") {
    packageJson.scripts['install:all'] = "npm install && cd $FRONTEND_DIR && npm install && cd ../$BACKEND_DIR && npm install";
} else if ("$FRONTEND_DIR") {
    packageJson.scripts['install:all'] = "npm install && cd $FRONTEND_DIR && npm install";
} else if ("$BACKEND_DIR") {
    packageJson.scripts['install:all'] = "npm install && cd $BACKEND_DIR && npm install";
}

fs.writeFileSync('package.json', JSON.stringify(packageJson, null, 2));
console.log('✓ package.json 配置完成');
EOF

echo -e "${GREEN}✓ 已添加以下 npm scripts:${NC}"
echo -e "  - ${YELLOW}npm run dev${NC}          : 同时启动前后端"
[ -n "$FRONTEND_DIR" ] && echo -e "  - ${YELLOW}npm run dev:frontend${NC} : 仅启动前端"
[ -n "$BACKEND_DIR" ] && echo -e "  - ${YELLOW}npm run dev:backend${NC}  : 仅启动后端"
echo -e "  - ${YELLOW}npm run install:all${NC}  : 安装所有依赖"
echo ""

echo -e "${BLUE}[5/5] 安装子项目依赖...${NC}"
read -p "是否需要安装前后端依赖？(y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    if [ -n "$FRONTEND_DIR" ] && [ -d "$FRONTEND_DIR" ]; then
        echo -e "${YELLOW}正在安装前端依赖...${NC}"
        (cd "$FRONTEND_DIR" && npm install)
    fi
    
    if [ -n "$BACKEND_DIR" ] && [ -d "$BACKEND_DIR" ]; then
        echo -e "${YELLOW}正在安装后端依赖...${NC}"
        (cd "$BACKEND_DIR" && npm install)
    fi
    echo -e "${GREEN}✓ 依赖安装完成${NC}"
fi
echo ""

echo -e "${GREEN}=========================================="
echo "✓ 配置完成！"
echo "==========================================${NC}\n"

echo -e "${YELLOW}使用说明：${NC}"
echo -e "1. 同时启动前后端: ${GREEN}npm run dev${NC}"
[ -n "$FRONTEND_DIR" ] && echo -e "2. 仅启动前端: ${GREEN}npm run dev:frontend${NC}"
[ -n "$BACKEND_DIR" ] && echo -e "3. 仅启动后端: ${GREEN}npm run dev:backend${NC}"
echo -e "4. 安装所有依赖: ${GREEN}npm run install:all${NC}"
echo ""

echo -e "${YELLOW}项目信息：${NC}"
[ -n "$FRONTEND_DIR" ] && echo -e "- 前端目录: ${BLUE}$FRONTEND_DIR${NC}"
[ -n "$BACKEND_DIR" ] && echo -e "- 后端目录: ${BLUE}$BACKEND_DIR${NC}"
echo ""

# 询问是否立即启动
read -p "是否立即启动开发环境？(y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}正在启动开发环境...${NC}"
    echo ""
    npm run dev
fi