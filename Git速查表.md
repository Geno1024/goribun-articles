## 配置工具 ##

为所有本地仓库配置用户信息

`$ git config --global user.name "[name]"`
设置在你提交时想要附加的名字

`$ git config --global user.email "[email address]"`
设置你提交时想要附加的email

`$ git config --global color.ui auto`
支持彩色的命令行输出。

## 创建仓库 ##

开始一个新的仓库或者从一个现有的URL获取

`$ git init [project-name]`
使用指定名称创建一个新仓库

`$ git clone [url]`
下载一个项目和它的版本历史

## 做出更改 ##

审查代码和提交
`$ git status`
列出所有新的或者修改的要被提交的文件

`$ git diff`
显示尚未staged的文件差异

`$ git add [file]`
快照文件为版本控制做准备

`$ git diff --staged`
显示staging和最后文件版本的差异

`$ git reset [file]`
Unstage文件，但是保留其内容

`$ git commit -m "[descriptive message]"`
在版本历史中记录文件快照

## 组更改 ##

命名一系列提交，联合完成工作

`$ git branch`
列出当前仓库中的所有分支

`$ git branch [branch-name]`
创建新分支

`$ git checkout [branch-name]`
切换到指定分支并更新目录

`$ git merge [branch]`
将指定的历史分支组合到当前分支

`$ git branch -d [branch-name]`
删除指定分支

## 修改文件名 ##

迁移删除受版本控制的文件

`$ git rm [file]`
从工作目录删除文件并stage删除

`$ git rm --cached [file]`
从版本控制中删除文件但是本地保留文件

`$ git mv [file-original] [file-renamed]`
修改文件名准备提交

## 禁止跟踪 ##

排除临时文件和路径

`*.log`

`build/`

`temp-*`
.gitignore文件禁止那些匹配指定规则的文件和路径的意外多版本

`$ git ls-files --other --ignored --exclude-standard`
列出项目中所有被忽略打文件

## 保存片段 ##

搁置和恢复不完整的变化

`$ git stash`
暂存跟踪文件的所有修改

`$ git stash pop`
把最近的文件恢复

`$ git stash list`
列出所有stash的变更

`$ git stash drop`
丢弃最近的stash的变更

## 审查历史 ##

浏览和检查项目文件的变化

`$ git log`
列出当前分支的版本历史

`$ git log --follow [file]`
列出指定文件的版本历史，包括重命名

`$ git diff [first-branch]...[second-branch]`
显示两个分支之间的内容差异

`$ git show [commit]`
输出指定提交的元数据和内容更改

## 重复提交 ##

擦除错误，复位lis

`$ git reset [commit]`
撤销[commit]后所有的操作，保持本地变化

`$ git reset --hard [commit]`
丢弃所有历史修改返回到指定提交

## 同步修改 ##

注册一个仓库书签，交换版本历史

`$ git fetch [bookmark]`
从仓库书签中下载所有历史

`$ git merge [bookmark]/[branch]`
将书签分支合并到当前本地分支

`$ git push [alias] [branch]`

上传所有本地分支的提交

`$ git pull`
下载书签历史并结合修改






