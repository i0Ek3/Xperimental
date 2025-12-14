#!/bin/bash

# 批量重命名脚本：移除文件夹名称后的 -main 或 -master 后缀
# 作用：将 xxx-main / xxx-master 重命名为 xxx
# 使用：在目标目录下执行此脚本

# 遍历当前目录下的所有项
for item in *; do
    # 条件1：是目录；条件2：名称以 -main 或 -master 结尾
    if [[ -d "$item" && ( "$item" == *-main || "$item" == *-master ) ]]; then
        # 核心逻辑：删除最后一个 '-' 及其后面的所有字符（去掉-main/-master）
        new_item="${item%-*}"

        # 检查新名称是否已存在，避免覆盖
        if [[ -e "$new_item" ]]; then
            echo "⚠️  跳过：$item → $new_item（目标名称已存在）"
        else
            # 执行重命名，-v 参数显示重命名过程
            mv -v "$item" "$new_item"
        fi
    fi
done

echo -e "\n✅ 批量重命名操作执行完毕！"
