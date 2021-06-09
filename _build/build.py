#!/usr/bin/python
# -*- coding:utf-8 -*-

import os
import platform
import sys
import shutil

project_name = "data-view"
config_name = "config.toml"
build_folder = "_build"

print u"正在打包"

sys_info = platform.system()
print u"系统类型: " + sys_info

code_dir = os.path.abspath(os.path.join(sys.argv[0], "../.."))

print u"项目路径: " + code_dir

is_windows = "" if sys_info != 'Windows' else "SET"

cmd_windows = "cd  %s&&" \
              "%s CGO_ENABLE=0&&" \
              "%s GOOS=windows&&" \
              "%s GOARCH=amd64&&" \
              "go build -ldflags \"-s -w\" -o %s/windows/%s-windows-x86_64.exe" % (
                  code_dir,
                  is_windows,
                  is_windows,
                  is_windows,
                  build_folder,
                  project_name)

cmd_linux = "cd  %s&&" \
            "%s CGO_ENABLE=0&&" \
            "%s GOOS=linux&&" \
            "%s GOARCH=amd64&&" \
            "go build -ldflags \"-s -w\" -o %s/linux/%s-linux-x86_64" % (
                code_dir,
                is_windows,
                is_windows,
                is_windows,
                build_folder,
                project_name)

cmd_darwin = "cd  %s&&" \
             "%s CGO_ENABLE=0&&" \
             "%s GOOS=darwin&&" \
             "%s GOARCH=amd64&&" \
             "go build -ldflags \"-s -w\" -o %s/darwin/%s-darwin-x86_64" % (
                 code_dir,
                 is_windows,
                 is_windows,
                 is_windows,
                 build_folder,
                 project_name)

print u"windows打包"
os.system(cmd_windows)

print u"linux打包"
os.system(cmd_linux)

print u"darwin打包"
os.system(cmd_darwin)

print u"拷贝配置文件"
shutil.copy2(os.path.join(code_dir, config_name),
             os.path.join(code_dir, build_folder, "windows"))
shutil.copy2(os.path.join(code_dir, config_name),
             os.path.join(code_dir, build_folder, "linux"))
shutil.copy2(os.path.join(code_dir, config_name),
             os.path.join(code_dir, build_folder, "darwin"))

print u"拷贝其他文件目录"
shutil.copytree(os.path.join(code_dir, "storage"),
                os.path.join(code_dir, build_folder, "windows", "storage"))
shutil.copytree(os.path.join(code_dir, "storage"),
                os.path.join(code_dir, build_folder, "linux", "storage"))
shutil.copytree(os.path.join(code_dir, "storage"),
                os.path.join(code_dir, build_folder, "darwin", "storage"))
