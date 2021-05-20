# coding=utf-8
import os
import platform

print "正在打包"

sys_info = platform.system()
print "系统类型: " + sys_info

if sys_info == 'Window':
    print "打包Window"
    os.system('cd ..')
    os.system('pwd')
elif sys_info == 'Linux':
    print "Linux打包"
elif sys_info == 'Darwin':
    print "Mac打包"

os.system('cd ..')
os.system('ls')
os.system('pwd')

