#!/bin/bash

# Tailwind CSS 配置修复脚本
# 适用于 React 项目的 Tailwind CSS 快速配置

echo "=================================="
echo "Tailwind CSS 配置修复脚本 v1.0"
echo "=================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 检查是否在项目根目录
if [ ! -f "package.json" ]; then
    echo -e "${RED}错误: 未找到 package.json 文件${NC}"
    echo "请在项目根目录下运行此脚本"
    exit 1
fi

echo -e "${YELLOW}[步骤 1/5] 安装 Tailwind CSS 依赖...${NC}"
npm install -D tailwindcss@3 postcss autoprefixer
if [ $? -ne 0 ]; then
    echo -e "${RED}依赖安装失败，请检查网络连接${NC}"
    exit 1
fi
echo -e "${GREEN}✓ 依赖安装完成${NC}\n"

echo -e "${YELLOW}[步骤 2/5] 初始化 Tailwind 配置文件...${NC}"
npx tailwindcss init
if [ $? -ne 0 ]; then
    echo -e "${RED}Tailwind 初始化失败${NC}"
    exit 1
fi
echo -e "${GREEN}✓ 配置文件创建完成${NC}\n"

echo -e "${YELLOW}[步骤 3/5] 修改 tailwind.config.js...${NC}"
cat > tailwind.config.js << 'EOF'
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
EOF
echo -e "${GREEN}✓ tailwind.config.js 配置完成${NC}\n"

echo -e "${YELLOW}[步骤 4/5] 配置 src/index.css...${NC}"
if [ ! -f "src/index.css" ]; then
    echo -e "${YELLOW}  未找到 src/index.css，正在创建...${NC}"
    touch src/index.css
fi

# 检查是否已经包含 Tailwind 指令
if grep -q "@tailwind base" src/index.css; then
    echo -e "${YELLOW}  src/index.css 已包含 Tailwind 指令，跳过修改${NC}"
else
    # 备份原文件
    cp src/index.css src/index.css.backup
    echo -e "${YELLOW}  已备份原文件到 src/index.css.backup${NC}"
    
    # 在文件开头添加 Tailwind 指令
    echo "@tailwind base;
@tailwind components;
@tailwind utilities;

$(cat src/index.css)" > src/index.css
    echo -e "${GREEN}✓ src/index.css 配置完成${NC}"
fi
echo ""

echo -e "${YELLOW}[步骤 5/5] 配置 frontend/index.html...${NC}"
INDEX_HTML="frontend/index.html"
if [ ! -f "$INDEX_HTML" ]; then
    INDEX_HTML="index.html"
fi

if [ -f "$INDEX_HTML" ]; then
    # 检查是否已经包含 Tailwind CDN
    if grep -q "cdn.tailwindcss.com" "$INDEX_HTML"; then
        echo -e "${YELLOW}  index.html 已包含 Tailwind CDN，跳过修改${NC}"
    else
        # 备份原文件
        cp "$INDEX_HTML" "${INDEX_HTML}.backup"
        echo -e "${YELLOW}  已备份原文件到 ${INDEX_HTML}.backup${NC}"
        
        # 在 </head> 前添加 CDN 脚本
        sed -i.tmp 's|</head>|  <script src="https://cdn.tailwindcss.com"></script>\n</head>|' "$INDEX_HTML"
        rm -f "${INDEX_HTML}.tmp"
        echo -e "${GREEN}✓ index.html 配置完成${NC}"
    fi
else
    echo -e "${RED}  警告: 未找到 index.html 文件，请手动添加 CDN 链接${NC}"
fi
echo ""

echo -e "${GREEN}=================================="
echo "✓ 配置完成！"
echo "==================================${NC}\n"

echo -e "${YELLOW}后续步骤：${NC}"
echo "1. 运行 ${GREEN}npm start${NC} 启动开发服务器"
echo "2. 如果样式仍未生效，请尝试清除缓存后重启"
echo ""
echo -e "${YELLOW}备注：${NC}"
echo "- 生产环境建议移除 CDN 方式，仅使用 npm 安装的版本"
echo "- 原文件已备份为 .backup 后缀"
echo ""

# 询问是否立即启动开发服务器
read -p "是否立即启动开发服务器？(y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}正在启动开发服务器...${NC}"
    npm start
fi