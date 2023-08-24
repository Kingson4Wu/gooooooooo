+ 官方博文：Go 1.18 工作区模式最佳实践: <https://zhuanlan.zhihu.com/p/495832968>

使用场景1
给上游模块新增feature，然后在你的Module里使用这个新feature

为你的workspace(工作区)创建一个目录。
Clone一份你要修改的上游模块的代码到本地。
本地修改上游模块的代码，增加新的feature。
在workspace目录运行命令go work init [path-to-upstream-mod-dir]。
为了使用上游模块的新feature，修改你自己的Go Module代码。
在workspace目录运行命令 go work use [path-to-your-module] 。
go work use 命令会添加你的Go Module的路径到 go.work 文件里：
go 1.18

use (
./path-to-upstream-mod-dir
./path-to-your-module
)

运行和测试你的Go Module。
发布上游模块的新feature。
发布你自己的Go Module代码。
使用场景2
同一个代码仓库里有多个互相依赖的Go Module

当我们在同一个代码仓库里开发多个互相依赖的Go Module时，我们可以使用go.work，而不是在go.mod里使用replace指令。

为你的workspace(工作区)创建一个目录。
Clone仓库里的代码到你本地。代码存放的位置不一定要放在工作区目录下，因为你可以在go.work里使用use指令来指定Module的相对路径。
在工作区目录运行 go work init [path-to-module-one] [path-to-module-two] 命令。
示例: 你正在开发 example.com/x/tools/groundhog 这个Module，该Module依赖 example.com/x/tools 下的其它Module。
你Clone仓库里的代码到你本地，然后在工作区目录运行命令 go work init tools tools/groundhog 。
go.work 文件里的内容如下所示：
go 1.18

use (
./tools
./tools/groundhog
)

tools路径下其它Module的本地代码修改都会被 tools/groundhog 直接使用到。
使用场景3：切换依赖配置
如果要测试你开发的代码在不同的本地依赖配置下的场景，你有2种选择：

创建多个workspace，每个workspace使用各自的go.work文件，每个go.work里指定一个版本的路径。
创建一个workspace，在go.work里注释掉你不想要的use指令。
对于创建多个workspace的方案：

为每个workspace创建独立的目录。比如你开发的代码依赖了example.com/util这个Go Module，但是想测试example.com/util2个版本的区别，你可以创建2个workspace目录。
在各自的workspace目录运行 go work init 来初始化workspace。
在各自的workspace目录运行 go work use [path-to-dependency]来添加依赖的Go Module特定版本的目录。
在各自的workspace目录运行 go run [path-to-your-module] 来测试go.work里指定的依赖版本。
对于使用同一个workspace的方案，可以直接编辑go.work文件，修改use指令后面的目录地址即可。


