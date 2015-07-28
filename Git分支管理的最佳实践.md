*原文地址：*[http://nvie.com/posts/a-successful-git-branching-model/](http://nvie.com/posts/a-successful-git-branching-model/)

**注：**该文章由作者Vincent Driessen写于2010年1月

在这篇文章中，我将介绍一个开发的模式；早在一年前，我就将它引入我的所有项目（工作和私人），它已经被证明是非常成功的。我早就打算写一篇关于它的文章，但是直到现在才找到时间去完整的介绍它。我不会讨论任何项目细节，只是讨论分支策略和发布管理。
![git-model－01.png][1]

它重点围绕Git作为我们所有源码的版本控制工具。

# 为什么是Git？

在网上，有很关于Git与其他源代码版本控制工具的优劣的透彻分析。在这里存在激烈的争论。作为一个开发者，与其他源代码版本控制相比，我更喜欢今天讨论的Git。Git真正改变了开发者关于合并与分支的思维模式。我也经历过经典的CVS/Subversion世界，合并与分支总是被认为有点吓人（当心合并冲突，它们会咬你！）并且在一段时间里，你只能每隔一段时间做一些事。

但是，使用Git后这些动作会变的廉价和简单，它们被认为是你的日常工作流的核心部分之一。例如，在有关CVS/Subversion的书中，分支与合并时在后面的章节中首先讨论的（针对高级用户），而在每一本Git书中，它已经被包含在第三章了（基础知识）。

由于分支与合的简单性和重复性，所以这不再是值得担心的事情了。比起其他的事，版本控制工具应该更注重对分支/合并的支持。

关于工具已经谈得够多了，让我们回到开发模型的话题上来。我将提出的模型基本上是一组流程，每个团队成员必须遵守，以此来管理软件开发过程。

# 分散但是集中

我们使用的仓库能很好的和分支系统工作，这是一个“真实的”仓库。请注意，这个仓库唯一的中央仓库（因为Git时一个DVCS，在技术层面，在技术层面上，作为中央仓库是没有问题的）。我们称它为origin，因为所有Git用户都熟悉origin这个名称。
![centr-decentr－02.png][2]

每个开发者都可以从origin拉（pull）和推（push）资源。不过，除了集中的推拉关系，每个开发者也可以从其他开发者那拉变更，以此形成一个子团队。例如，在进展中的工作推到origin之前，两个或者更多的开发者需要在开发同一个大的新功能，这也许是很有用的。在上图中，有Alice和Bob， Alice和David，Clair和David的子团队。

从技术上讲，无非是Alice定义了一个远程分支，命名为bob，指向Bob的仓库，反之亦然。

# 主要分支

在核心，开发模型很大程度上受现存模型的影响。中央仓库无期限的持有两个主要分支：

- 主分支（master）
- 开发分支（develop）

origin上的主分支每个Git用户应该都熟悉。平行于主分支，另一个分支被称为开发分支。我们认为origin/master是主要分支，是因为分支的HEAD指向的源代码总是反映了一个生产就绪的状态。我们认为origin/develop是主要分支，因为分支的HEAD指向的代码总是反映下一版本中带有最新的开发变化。有些人称它为“集成分支”。每晚的自动构建都源于它。

![main-branches03.png][3]

当开发分支上的源代码达到稳定，并且准备发布时，所有的改变都应该合并回主分支，然后打上一个版本号标签。细节操作将在后面讨论。

因此，每次更改合并回主分支时，将会产生一个新的可以发布的产品状态。在主分支我们往往非常严格，每当主分支上有提交时，我们可以使用Git钩子脚本自动地构建，并推送软件到我们的生产服务器。

# 支持分支

接着主分支和开发分支，我们开发模型采用了多种支持分支来帮助开发团队并行开发，缓解功能跟踪，准备生产版本，协助快速修复线上产品问题。不同于主分支，这些分支总是有一个有限的生命周期，因为它们最终会被删除。我们可以使用的分支类型：

－ 功能分支（Feature branches）
－ 发布分支（Release branches）
－ 修复分支（Hotfix branches）

每个分支都有特定的目的，并且对于哪些分支可能是它们起源的分支，哪些分支必须是它合并的目标，有着严格的规则。我们将在一分钟内过一遍。

从技术角度来看，这些分支并不是特殊的。分支类型是由我们如何使用它们来划分的。它们当然是普通的原始的Git分支。

## 功能分支

如果分支来自开发分支，那么必须合并回开发分支。分支名约定：除了master、develop、release-\*和 hotfix-\*的任何名字。

![fb-04.png][4]

功能分支（有时叫做主题分支）用于为即将到来或者较远的版本开发新功能。当启动一个功能分支开发，目标版本对于这个功能分支的合并可能是未知的。功能分支的本质是它的存在周期和功能开发一样长，但是最终会合并回开发分支（为即将发布的版本添加新功能）或者被丢弃（如果是个令人失望的实验）。

功能分支通常存在于开发者的仓库，不在origin中。

### 创建一个功能分支

当要开发一个新功能时，从开发分支分出一个功能分支。

    $ git checkout -b myfeature develop
    Switched to a new branch "myfeature"


### 将完成的功能分支合并回开发分支

完成功能可能被合并到开发分支，一定会添加它们至即将发布版本：

    $ git checkout develop
    Switched to branch 'develop'
    $ git merge --no-ff myfeature
    Updating ea1b82a..05e9557
    (Summary of changes)
    $ git branch -d myfeature
    Deleted branch myfeature (was 05e9557).
    $ git push origin develop

--no-ff标签会使合并时总会创建一个新的提交对象，即使合并可以使用快速合并。这避免了丢失功能分支的历史信息和团队共同添加到功能分支所有提交。对比：

![merge-without-ff－05.png][5]

在第二种情况下，你不可能从Git提交对象历史中看到实现的功能，你不得不手动地读取所有的日志消息。并且还原所有功能（即一组提交）是非常头疼大；如果使用了--no-ff标记那会变的很容易。

的确，它会创建更大的提交对象，不过受益会更大！

不幸的是，我还没找到让Git合并的默认使用--no-ff的方法，不过它应该可以。

## 发布分支

如果分支来自开发分支，那么必须合并回开发分支和主分支，分支命名约定：release-*。

发布分支支持新产品版本的预发布。它们允许最后一分钟的细节检查。另外，它们允许小的bug修复，为发布（版本号，构建日期等）准备元数据。通过在发布分支做的这些工作，开发分支被清理以接受下一个大版本的功能。

从开发分支分出一个新的发布分支的时刻是在开发分支反映了新版本的理想状态。在这个时间点上，至少针对构建发行版本的所有功能必须合并到开发分支。针对未来发行版本的所有功能未必一定等到分出发布分支之后。

将发布的版本被分配一个版本号正是一个发布分支的开始，不能更早了。直到那时，开发分支反映下一版本的变化，但是目前还不清楚“下一版本”会变成0.3还是1.0，直到发布分支开始。这一决定是在发布分支的开始做出的，并通过了打版本号的项目规则。

### 创建发布分支

发布分支创建自开发分支。例如，称1.1.5版本是当前产品发布版本，为将有个大版本到来。开发分支的状态是准备“下一版本”，并且我们已经觉得将是1.2版本（而不是1.1.6或者2.0）。所有我们分一个分支，给发布分支一个反映新版本号的名字。

    $ git checkout -b release-1.2 develop
    Switched to a new branch "release-1.2"
    $ ./bump-version.sh 1.2
    Files modified successfully, version bumped to 1.2.
    $ git commit -a -m "Bumped version number to 1.2"
    [release-1.2 74d9424] Bumped version number to 1.2
    1 files changed, 1 insertions(+), 1 deletions(-)
    
创建一个新分支并且切换过去后，我们打上版本号，这里的bump-version.sh时一个虚构的shell脚本，改变在工作副本的文件来反映新的版本。这当然可以是一个手动的修改，主要是文件的修改。然后，打上的版本号被提交。

 这个新分支可能存在一段时间，直到发布版确定推出。在此期间，bug修复会应用在这个分支上（而不是开发分支）。在这里添加大的新功能是被禁止的。它们必须被合并到开发分支，等待下次的大版本。
 
### 完成发布分支

当发布分支状态时准备成为一个真正的版本需要进行一些操作。首先，发布分支合并到主分支（请记住，因为主分支的每次提交都是一个新版本）。接下来，主分支的提交必须打上标签以便将来参考历史版本。最后，发布分支的改动需要合并回开发分支，以至于将来的版本也包含来这些bug修复。

在Git中的前两个步骤：

    $ git checkout master
    Switched to branch 'master'
    $ git merge --no-ff release-1.2
    Merge made by recursive.
    (Summary of changes)
    $ git tag -a 1.2
    
现在发布完成，并且打上了标记以便将来参考。

>**编辑：**你也许想使用-s或者-u<key>标记来加密签名你的标签。

保留发布版本所做的修改，虽然我们需要把它们合并回开发分支，在Git中：

    $ git checkout develop
    Switched to branch 'develop'
    $ git merge --no-ff release-1.2
    Merge made by recursive.
    (Summary of changes)  
 
这一步可能导致合并冲突（可能因为我们已经改变了版本号）。如果是的话，修复并提交。

现在，我们真正的完成了，发布分支可以被删除了，因为我们不需要它了。

    $ git branch -d release-1.2
    Deleted branch release-1.2 (was ff452fe).
    
## 修复分支

如果分支来自主分支，那么必须合并回开发分支和主分支，分支命名约定：hotfix-*。

![mhotfix-branches-07.png.png][6]

修复分支很像发布分支，因为它们也是为新产品发布做准备，虽然是个意外。它们产生的必要性是处理生产版本中不受欢迎的状态。当生产版本中一个关键的bug需要立即修复时，修复分支可以起自主分支中对应的记录产品版本的标签。

本质上团队成员的工作（在开发分支）可以继续，而另一个人准备一个快速的生产修复。

### 创建修复分支

修复分支创建自主分支。例如，1.2版本是当前生运行的产版本，一个严重的bug导致了故障。但是在开发分支的修改是不稳定的，我们可以开一个修复分支并开始修复问题：

    $ git checkout -b hotfix-1.2.1 master
    Switched to a new branch "hotfix-1.2.1"
    $ ./bump-version.sh 1.2.1
    Files modified successfully, version bumped to 1.2.1.
    $ git commit -a -m "Bumped version number to 1.2.1"
    [hotfix-1.2.1 41e61bb] Bumped version number to 1.2.1
    1 files changed, 1 insertions(+), 1 deletions(-)
    
开分支后不要忘了打版本号。

然后，修复bug并提交一个或多个bug修复提交。

    $ git commit -m "Fixed severe production problem"
    [hotfix-1.2.1 abbe5d6] Fixed severe production problem
    5 files changed, 32 insertions(+), 17 deletions(-)
    
### 完成修复分支      

当修复完成，修复分支不仅需要合并回主分支，还要合并回开发分支，为了保证bug修复也包含在下一版本。这和完成发布分支非常相似。

首先，更新主分支并打上发布标签。

    $ git checkout master
    Switched to branch 'master'
    $ git merge --no-ff hotfix-1.2.1
    Merge made by recursive.
    (Summary of changes)
    $ git tag -a 1.2.1
**编辑：**你可能想使用-s或者-u<key>标记去加密签名的标签。

接下来，包含bug修复至开发分支：

    $ git checkout develop
    Switched to branch 'develop'
    $ git merge --no-ff hotfix-1.2.1
    Merge made by recursive.
    (Summary of changes)
    
一个例外的规则是，**当一个发布分支存在，修复分支的修改需要被合并到发布分支，而不是**开发分支。合并回发布分支也意味着在发布分支完成时，bug修复会被合并到开发分支。（如果开发分支上的工作需要这个bug修复，并且不能等待发布分支完成，你也可以安全地合并bug修复到现在的开发分支。）

最后，删除临时分支：

    $ git branch -d hotfix-1.2.1
    Deleted branch hotfix-1.2.1 (was abbe5d6).

# 总结

虽然这个分支模型没有什么真正令人震撼的，不过文章开始的大图在外面的项目中被证明是非常有用的。它形成了一个优雅的心智模型，很容易理解并且允许团队成员对分支和发布过程有一个共同的认识。







[1]: images/git-model/git-model-01.png
[2]: images/git-model/centr-decentr－02.png
[3]: images/git-model/main-branches03.png
[4]: images/git-model/fb-04.png
[5]: images/git-model/merge-without-ff－05.png
[6]: images/git-model/mhotfix-branches-07.png