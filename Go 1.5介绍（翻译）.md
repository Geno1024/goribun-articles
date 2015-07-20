Go 1.5终于来了！虽然有些迟，但还是比预想的早了些，赶紧找个时间看下Go 1.5的相关介绍。以下翻译来自官方文档：[http://tip.golang.org/doc/go1.5](http://tip.golang.org/doc/go1.5)，发现翻译错误，欢迎指正。


## Go 1.5简介

Go最新的发布版本Go 1.5是一个重要的发布版本，包括主要架构的更改实施。尽管如此，我们还是期望所有的Go程序可以像以前一样编译运行，因为这次的发布仍然保持Go 1中[兼容性承诺](http://tip.golang.org/doc/go1compat)。

 **重大改进如下：**
 
 - 编译器和运行时完全用Go编写（少量汇编）。C不再参与实现，曾经构建分布所必需的C编译器不见了。
 - 垃圾回收器现在是并发的，暂停时间会显著降低，甚至可能与其他goroutines并行执行。
 - 默认情况下，Go使用GOMAXPROCS设置的可用核心数量来运行，之前默认是1。
 - 为所有的仓库提供[internal packages](http://golang.org/s/go14internal)支持，而不仅仅是Go核心。
 - Go命令现在为“vendoring”外部依赖提供试验性支持。
 - 全新的*go tool trace*命令支持程序执行的细粒度跟踪。
 - 全新的*go doc*命令（与godoc不同）专门用于命令行中使用。

<!--more--> 

 以上和其他变化的实施和工具在下面进行讨论。
 
 该版本还包括一个关于map的很小的语法改动。
 
 最后，版本的发布通常是六个月的间隔时间，这样既可以提供更充足的时间准备重大版本发布，也可以更方便地改变计划。
 
 
## 语法变化 

### Map字面量

由于疏忽，在切片中允许省略元素类型的规则没有应用到map的key。在Go 1.5中已经被改正。下面的例子可以说明这些：
<pre class="prettyprint">
	m := map[Point]string{
	    Point{29.935523, 52.891566}:   "Persepolis",
	    Point{-25.352594, 131.034361}: "Uluru",
	    Point{37.422455, -122.084306}: "Googleplex",
	}
</pre>
可以写成下面这样，没有显式列出Point类型：
<pre class="prettyprint">
	m := map[Point]string{
	    {29.935523, 52.891566}:   "Persepolis",
	    {-25.352594, 131.034361}: "Uluru",
	    {37.422455, -122.084306}: "Googleplex",
	}
</pre>	
## 实现

### 不再使用C

现在，编译器和运行时都用Go和汇编完成，不再使用C。唯一的C代码留在用于测试或者cgo相关的代码树中。C编译器只存在1.4或者更早版本时。它用于构建运行时；保证C代码和goroutines的堆栈管理协同工作是一个自定义编译器的重要组成。从运行时是Go实现开始，C编译器就没有必要使用了。移除C的细节在[其他地方](https://golang.org/s/go13compiler)讨论。

### 编译器和工具

独立的工具已由Go重写，工具的名字也发生来变化。6g、8g等旧名字等已经过时，代替它们的是一个二进制的、更易用的go编译工具，通过 $GOARCH和$GOOS将Go源码编译为适合体系结构和操作系统的二进制程序。同样，现在有一个新的连接器（go tool link）和一个新的汇编器（go tool asm）。连接器已经自动地翻译了旧的C实现。但是汇编器是一个新的原生Go实现，下面会有更详细的讨论。

由于移除了6g、8g等名字，现在编译器和汇编器的的输出是.o后缀而不再是.8、.6等等。

### 垃圾回收

作为Go 1.5设计文档开发大纲的一部分，垃圾回收器已经被重新设计。公共更先进算法的组合，更好调度的回收器和与更多地与用户程序并行运行，使得预期的延迟比以前版本的垃圾回收器的延迟低很多。回收器的“stop the world”时间几乎已经控制在10ms以内，甚至更低。

比如对于网站的响应速度，延迟更低的新的回收器是很重要的，这些系统受益于低延迟。

关于垃圾回收器的更多细节在TODO: GopherCon talk。

### 运行时

在Go 1.5中，旧的goroutines调度已经被改变。调度器的顺序不再通过语言定义，但是由于这个改变，依靠调度顺序的程序可能被损坏。我们已经看到一些程序由此受到影响。如果你的程序隐式的依赖调度顺序，你需要更新它们。

另一个潜在的破坏性的改变是运行时现在通过GOMAXPROCS定义的处理器上可用的核心数设置默认的线程数量来运行程序。以前版本默认时1。不希望运行在多核心上的程序可能会被无意破坏。它们可以通过移除限制或者设置GOMAXPROCS来解决。

### 构建

现在，Go编译器和运行时已经由Go实现，这个Go编译器能够从源码编译发行版。因此，为了编译Go核心，必需有一个可运行的Go发行版。（不在Go核心开发的Go程序员不受影响）。任何Go 1.4以及更高版本将担任该角色。更多详情，请看[设计文档](https://golang.org/s/go15bootstrap)。

## 接口

由于大部分厂商已经放弃32位x86架构，所以在1.5版本中，提供的二进制文件会减少。比如OS X操作系统不在提供386支持，只提供amd64架构支持。同样，对于雪豹（苹果OS X 10.6）的接口支持还在工作，但是不再发布下载和维护。此外，dragonfly/386也不再支持，因为DragonflyBSD本身已经不再支持32位386架构。

一些新的可用接口已经从源码构建完毕，包括darwin/arm和darwin/arm64。新的接口linux/arm64大部分工作已经完成，但是cgo仅支持外部链接。

对于FreeBSD，因为Go 1.5使用了新的系统调用指令，所以Go 1.5需要FreeBSD 8-STABLE+以上的环境。

对于NaCl，因为Go 1.5使用了get_random_bytes系统调用，所以Go 1.5需要pepper-39或更高的SDK版本。

## 工具

### 翻译

作为从代码树上移除C的一部分，编译器和连接器已经由C翻译成Go。这是一个正真的（机器辅助）的翻译，所以新的程序本质上是旧程序的翻译而不是一个带有新bug的新程序。我们确信翻译过程中几乎没有引入新bug，但实际上发现了一些之前的未知bug，不过已经修复了。

无论如何，汇编器是一个新程序，下面有它的描述。

### 重命名

编译器（6g、8g等），汇编器（6a、8a等）和连接器（6l、8l等）一整套程序已经被合并到一个单独的工具，它通过环境变量GOOS和GOARCH配置。旧的名字已经废弃了，通过像*go tool compile*、*go tool asm*和 go tool link等工具机制新的工具，新的工具已经可以使用了。另外，中间文件的后缀.6、.8ye已经废弃了，现在它们都是普通的.o文件。

例如，在Darwin adm64构建和连接程序直接使用这些工具而不是通过go build，像这样：
<pre class="prettyprint">
	$ export GOOS=darwin GOARCH=amd64
	$ go tool compile program.go
	$ go tool link program.o
</pre>	
### 移动

因为 [go/types](http://tip.golang.org/pkg/go/types/)包现在已经移到主仓库（见下文），[vet](http://tip.golang.org/cmd/vet)和[cover](http://tip.golang.org/cmd/cover)工具已经被移入。它们不在保存在扩展的golang.org/x/tools仓库，虽然（过时的）代码还保留并与老版本兼容。

### 编译

根据以上的描述，在Go 1.5中编译器是一个独立的Go程序，从旧的C源码翻译而来，并且代替了6g、8g等等。它的目标是通过环境变量GOOS和GOARCH进行配置。

Go 1.5的编译器和旧编译器是等效的，但是一些内部细节已经改变了，一个显著的变化就是常量的赋值现在是使用 math/big包而不是一个高精度的运算的自定义（没有很好的进行测试的）实现。我们期望这不会影响结果。

仅针对amd64架构，编译器有一个新选项，-dynlink，通过引用在外部共享库中定义的Go符号来帮助实现动态链接。

### 汇编器

像编译器和连接器一样，在Go 1.5中汇编器也是一个单独的程序，它取代了一整套的汇编器（6a、8a等），并且通过环境变量GOOS和GOARCH进行配置。和其他程序不一样的是，汇编器是一个Go编写的新程序。

新的汇编器和旧的汇编器兼容，但是有些改变可能影响到一些汇编器的源文件。关于这些改变的信息可以查看更新的[汇编器指导](http://tip.golang.org/doc/asm)。总之：

首先，用于常量的表达式求值有些不同。现在使用64位无符号运算，运算优先级（+, -, <<等）来自于Go而不是C。我们希望这些变化只影响极少数的程序，不过可能需要手动验证了。

也行更重要的是，在一些机器上SP或PC对于地址寄存器只是一个别名而已。例如在ARM上，R13作堆栈指针，R15作硬件程序计数器，现在不包括符号是非法的。例如，SP和4(SP)是非法的，但是sym+4(SP)是正确的。在这样的机器上，用它真实的R名字指代硬件寄存器。

一个微小的变化是一些旧的汇编器这些标记
<pre class="prettyprint">
    constant=value
</pre>    
去定义一个常量。由于像传统的类C的#define标记总是可用，仍被支持（包括一个简单的C预处理器的汇编器），所以这个功能已被移除。

### 连接器

在Go 1.5中，连接器是一个Go程序，代替了6l、8l等。它的操作系统和指令集通过环境变量GOOS和GOARCH设置。

还有一些其他的变化。更显著的事增加了一个-buildmode来扩展连接方式；现在，它支持构建共享库和允许其他语言调用Go的库。这些在设计文档大纲中有涉及。对于一系列可用的构建模式以及它们使用，运行：
<pre class="prettyprint">
    $ go help buildmode
</pre>    
另一个小变化是连接器不再在Windows可执行文件的头部记录构建时间的时间戳。另外，Windows cgo可执行文件丢失一些DWARF信息，虽然这可以被修复。

### Go命令

[go](http://tip.golang.org/cmd/go)命令的基础操作没有改变，但是有些变化还是值得一提的。

先前版本中，通过go命令引入internal目录至一个不用导入的包。在Go 1.4，核心仓库中引入了一些internal组件进行测试。作为[设计文档](https://golang.org/s/go14internal)中的建议，这个改变已经在所有仓库中可用了。这个规则在设计文档中被说明，但是总结起来就是任何在internal目录或者在internal目录底下的包可能被同级目录导入。已存在的目录组件名为internal的包可能无意地被这种改变破坏，这也是为何它在最后的版本才被通告的原因。

如何处理包的另一个改变是实验性的增加了“vendoring”支持。TODO：go命令本身没有显示这个命令。TODO：在[https://golang.org/s/go15vendor](https://golang.org/s/go15vendor)的出本设计需要更新。

还有一些小改动，详情见[文档](http://tip.golang.org/cmd/go)。

- SWIG支持已经更新，比如.swig和.swigxx现在需要SWIG 3.0.6或更高版本。
- 标准库通配符包名现在排除了命令，一个新的通配符覆盖了该命令。
- 一个新的-toolexec标记用于构建允许替换成一个不同的命令去调用编译器等。这作为一个自定义命令来替换go tool。
- 构建子命令包括一个-buildmode选项和连接器绑定，如上所述。
- -asmflags构建选项已经被添加为了提供汇编器的标记。然而，-ccflags构建选项已经被移除；这个专用于老版本，现在已经删除了C编译器。
- 测试的子命令现在有一个-count标记去注明每一个test和benchmark运行了多少次。通过-test.count标记，[testing](http://tip.golang.org/pkg/testing/)包不会运行。
- generate子命令有了些新特性。-run选项指定了一个正则表达式去选择执行的指令；这是一个提议，但是Go 1.4没有实现。现在执行模式可以获得两个新的环境变量：$GOLINE返回指令的源码的行数，$DOLLAR扩展$符。
- 现在get子命令有一个-insecure标记，如果访问一个不安全的未加密仓库，必需使用该标记。

### vet命令

现在，[go tool vet](http://tip.golang.org/cmd/vet)命令能够更彻底地校验struct标签。

### trace命令

    TODO
    cmd/trace: 新的查看追踪命令 (https://golang.org/cl/3601)

### go doc命令

在前几个版本中，go doc命令已经被删除，因为它不是必须的。可以使用godoc命令代替。在1.5版本中，引入了新的go doc命令，它拥有比godoc更方便的命令行接口。它是为命令行使用而专门设计的，通过调用，提供了更紧凑、专注一个包或者它的元素多文档介绍。它还提供累一个不区分大小写的匹配，并支持显示文档中不可导出的符号。更多详情，请运行go help doc。

### cgo

当解析到cgo的行，调用${SRCDIR}现在会扩展路径至源码目录的。它允许涉及把相对于源码目录作为选项传递给编译器和连接器。当前工作目录发生变化时，没有扩展的路径是无效的。

在Windows中，cgo现在默认使用外部链接。

## 性能

和往常一样，改变时普遍和多样的，以至于对于性能很难做出精确的描述。在该版本中，这些变化设置比通常状况更加广泛，它包含了一个新的垃圾回收器并将运行时转为Go实现。一些程序也许运行的更快，一些可能更慢。平均来说，在Go1基准测试中，Go 1.5比Go 1.4快几个百分点，正如上面提到的，垃圾回收器的暂停时间大幅缩短，几乎一直低于10ms。

Go 1.5的构建会比以往慢2倍。编译器和连接器是C自动翻译成Go的，导致生成的代码不高效，和精心编写的代码相比表现不佳。分析工具和和重构有助于提高代码质量，但是还有很多工作要做。进一步的分析和优化将继续在Go 1.6和以后的版本中完成。更多详情，请看[幻灯片](https://talks.golang.org/2015/gogo.slide)或者相关[视频](https://www.youtube.com/watch?v=cF1zJYkBW4A)。

## 核心库

### flag

flag包的[PrintDefaults](http://tip.golang.org/pkg/flag/#PrintDefaults) 函数，和 [FlagSet](http://tip.golang.org/pkg/flag/#FlagSet)的方法，已被修改以创建更好的使用信息。格式已经变得更人性化，在使用信息中一个词使用\`反引号`括起来，则被视为该flag操作数的名字并在使用信息中显示。例如，调用创建flag：
<pre class="prettyprint">
    cpuFlag = flag.Int("cpu", 1, "run `N` processes in parallel")
 
</pre>
 将展示帮助信息：
<pre class="prettyprint"> 
     -cpu N
    	run N processes in parallel (default 1)
</pre>    	
另外，现在仅当它不是该类型的零值时，默认值会被列出。

### math/big中的Floats

[math/big](http://tip.golang.org/pkg/math/big/)包中有了一个新的基础数据类型，[Float](http://tip.golang.org/pkg/math/big/#Float)，它实现任意精度的浮点数。一个浮点值由一个布尔标记、一个变长尾数和一个32位固定大小的符号指数代表。一个Float（尾数大小）的精度可以被明确指定，或者由创建值的操作决定。一旦创建，一个Float的尾数的大小可以通过[SetPrec](http://tip.golang.org/pkg/math/big/#Float.SetPrec)方法修改。Float支持无穷大和溢出的概念，但是当值等于IEEE 754 NaNs时，将触发发一个pianc。Float操作支持所有IEEE-754舍入模式。当精度是24(53)位时，在float32(float64)范围内的操作得到的结果和IEEE-754算法对应的结果是相同的。

### Go types
[go/types](http://tip.golang.org/pkg/go/types/)包之前一直在golang.org/x仓库维护，直到Go 1.5，它被迁入主仓库。目前，旧的位置已经过时了。还有一个轻微的API变化，将在下面讨论。

随着这次变动， [go/constant](http://tip.golang.org/pkg/go/constant/)包也被移动到主仓库；之前是在golang.org/x/tools/exact。[go/importer](http://tip.golang.org/pkg/go/importer/)包以及上述的一些工具也被移动到主仓库。

### Net

net包中的DNS解析器之前几乎都是用cgo范文系统接口的。Go 1.5中的变化意味着更多的Unix系统的DNS解析将不再需要cgo，这简化了在这些平台上的执行过程。现在，如果系统的网络配置允许，原生的Go解析器就足够了。这个改动重要的影响是每一个DNS解析占用一个goroutine，而不是一个线程，所以多个未完成的DNS请求将消耗更少的系统资源。何时运行解析器应用是在运行期决定的，而不是编译期。netgo构建标记用于强制使用Go解析器，现在已经不是必须的了，虽然它依然有效。

这个改变只适合Unix系统。Windows、Mac OS X和Plan 9系统还和之前一样。

### Reflect

[reflect](http://tip.golang.org/pkg/reflect/)包加入了两个新函数：[ArrayOf](http://tip.golang.org/pkg/reflect/#ArrayOf)和[FuncOf](http://tip.golang.org/pkg/reflect/#FuncOf)。这些函数，类似于[SliceOf](http://tip.golang.org/pkg/reflect/#SliceOf)函数，在运行时创建新类型来描述数组和函数。

### 强化
通过使用go-fuzz工具进行随机测试，在标准库中已经发现了几十个bug。在 archive/tar， archive/zip， compress/flate， encoding/gob， fmt， html/template， image/gif， image/jpeg， image/png， and text/template这些包中的bug已经修复。这些修复强化了针对错误或者恶意输入的处理。


### 微小的库改动

- archive/zip包Writer类型现在加入了SetOffset方法来指定输出流写入文档的位置。
- bufio包的Reader现在有一个Discard方法用于丢弃输入的数据。
- 在bytes包，Buffer类型现在加入了Cap方法来报告缓冲区中分配到的字节数。同样， 在bytes包和strings包， Reader类型现在加入了Size方法，来报告底层切片或者字符串的原始长度。
- bytes包和strings包现在加入了LastIndexByte函数来定位参数值中最右边的字节。
- crypto包加入了新接口- Decrypter，它抽象了用在不对称解密的私钥的行为。
- 在crypto/cipher包，Stream接口的文档中关于当源和目标长度不同时的行为已经被澄清。当目标比源短时，方法会panic。这只是文档的改动，实现没有变化。
- crypto/cipher包，现在，在 AES的Galois/Counter模式，支持不同于96字节的随机长度，其中需要某些协议支持。
- 在crypto/elliptic包，在CurveParams结构体中加入了Name成员变量，并且包中实现的curves已经有了名称。对于一个 curve-dependent的密码系统，这些名称提供了一种更安全的方式来选择一个curve，而不是选择其比特大小。
- crypto/elliptic包中，Unmarshal函数已经验证point实际上是在curve上。（如果不是，函数返回nils）。这个变化是为防止某些攻击。
- crypto/tls包默认是TLS 1.0。如果需要，旧的默认SSLv3通过[Config](http://tip.golang.org/pkg/crypto/tls/#Config)也是可用的。
- crypto/tls包支持在RFC 6962中指定签名证书时间戳（SCTs）。如果在证书结构中列出，那么服务器会提供它们，并且客户端请求它们，如果存在的话，在ConnectionState结构体中暴露它们。crytpo/tls包服务器实现现在将总是在Config结构体调用GetCertificate函数来为没有提供证书的的链接选择一个证书。
- 最后，在crypto/tls包的session ticket keys现在可以轮换（在活动链接期间定期更改）。这通过Config类型新加入的SetSessionTicketKeys方法实现。
- 在crypto/x509包，现在如说明书中定义的那样，只有最左边的标签可以接受通配符。
- 在crypto/x509包，处理未知的关键扩展改变了， 它们经常引起解析错误，但是现在它们只在Verify中被解析并引起错误。Certificate 中，新的成员变量UnhandledCriticalExtensions记录这些扩展。
- database/sql包的DN类型现在加入一个Stats方法，用于检索数据库统计信息。
- debug/dwarf包现在有更广泛补充，为了更好的支持DWARF版本4。查看新类型Class的定义的例子。
- encoding/base64包，现在通过两个编码变量RawStdEncoding和 RawURLEncoding支持unpadded编码。
- encoding/json包，如果一个JSON的值不适合目标变量，或者是一个正在反列集合的部分，那么回返回 UnmarshalTypeError错误。
- flag包有一个新的函数，UnquoteUsage，使用上述的新约定协助使用信息创建。
- 在fmt包，Value类型的值现在打印它拥有的，而不是使用反射。Value的Stringer方法，它产生像<int value>这种东西。
- go/ast包的EMptyStmt类型现在加入一个布尔类型的Implicit成员变量，它记录分号是否被隐式添加或者已经存在于源中。
- 为了对于一些架构向前兼容go/build包储存的GOARCH值，Go可能短期支持。它不只是一个承诺。另外，Package结构体现在加入一个PkgTargetRoot成员变量，存储所安装的架构相关的根目录，如果已知的话。
- go/types（新迁移的）包允许使用新的Qualifier函数类型作为多个函数的参数来控制连接到包级别的前缀。这是一个包的API的变化，但是由于它是新迁入核心的，编码时使用的包必须明确的要求它在新的位置。TODO: 对此应该有个gofix。
－ 在image包，现在Rectangle类型现在实现了Image接口，绘图时遮盖图像。
- 在image包，支持4:1:1和4:1:0的YCbCr子采样和基本的CMYK支持，现在由新的image.CMYK结构图表示，以协助处理一些JPEG图像。
－image/color包加入了基础CMYK支持，通过新的CMYK结构体，CMYKModel颜色模式，和CMYKToRBG函数，对于一些JPEG图片是必须的。
- image/gif包包括了一些概括。一个多帧的GIF文件现在可以有一个不同于所有的单帧的边界的整体边界。另外，现在GIF结构体有一个Disposal成员变量对每帧指定处理方法。
- io包加入了CopyBuffer函数，它像Copy函数，但是它使用调用者提供的缓冲区，允许控制和分配缓冲区的大小。
- log包有了一个新的LUTC标记，以至在UTC时区的时间戳被打印。它还增加累一个SetOutput函数设置对于辨准日志器的输出位置，和一个corresponding方法创建日志器。
- 在Go 1.4，Max没有检测所有可能的NaN位模式。它 在Go 1.5中被修复了，所以数据中使用math.Max包括NaN的程序的行为可能不同，但是现在根据IEEE754正确点定义了NaNs。
- math/big包加入了用于整数的Jacobi函数，为Int类型添加了ModSqrt方法。
- mime包加入了WordDecoder类型用来解码MIME含RFC 204编码的头。还提供了BEncoding和QEncoding作为 RFC 2045和RFC 2047编码方案的实现。
- mime包还加入了ExtensionsByType函数，它返回给定MIME关已知联类型的的MIME扩展。
- 新加入的mime/quotedprintable包，实现了RFC 2045定义的引用可打印编码。
- Go 1.5的net包加入了RFC-6555-compliant拨号，为了在DNS中列出了多个TCP地址的站点。
- 在net包中，返回到错误类型不一致的问题已经解决。现在返回的OpError值有了更多的信息。另外，OpError类型现在包括了Source成员变量，保存本地网络地址。
- net/http包现在已经支持从服务器处理程序设置追踪。详情请看[ResponseWriter](http://tip.golang.org/pkg/net/http/#ResponseWriter)文档。
- 另外，在het/http包，在ServeContent函数中，有忽略零Time值的代码。在Go 1.5，现在还忽略了值等于Unix的Epoch（1970年1月1日00:00:00 UTC）的时间。
- net/http/fcgi包导出了两个新错误，ErrConnClosed和 ErrRequestAborted，用来报告相应的错误情况。
- net/http/cgi包有一个bug，处理不好环境变量REMOTE_ADDR 和 REMOTE_HOST的值。现在已经修复。另外，从Go 1.5开始，设置REMOTE_PORT变量。
- net/mail包加了AddressParser类型，它可以解析Email地址。
-  net/smtp包现在有了对客户端类型的TLSConnectionState访问器，返回客户端的TLS状态。
- os包加入了LookupEnv函数，它类似于Getenv函数，但是可以区分空环境变量或者缺失环境变量。
- os/signal包已经加入了 Ignore和Reset函数。
- runtime, runtime/pprof和net/http/pprof包加入了新函数支持上述的跟踪设施：ReadTrace、StartTrace、StopTrace、StartTrace、StopTrace和Trace。详情见文档。
- runtime/pprof包默认情况下包括所有内存配置的整体统计信息。
- strings包加入了Compare函数，目前提供对称的字节包，但是作为字符串的原生比较，它不是必要。
- sync包的WaitGroup函数现在诊断代码，调用Add而不是从Wait返回。如果检测到这种情况，WaitGroup引发panic。
- 在syscall包，Linux的SysProcAttr结构体现在有GidMappingsEnableSetgroups成员变量，这使Linux 3.19的安全变更时必要的。在所有的Unix系统，结构体加入了新的Foreground和 Pgid 成员变量在执行时提供更多的控制。在Darwin，有新的Syscall9函数支持调用更多的参数。
- testing/quick包将为指针类型产生零值，使用递归的数据结构是可行的。而且现在支持新的数组类型。
- 在text/template和html/template包，整型常量包的整型常量太大，而不能用Go的整型表示，引发一个错误。之前，它们默默地转为浮点数，并丢失精度。
- 在 text/template和html/template包，新的Option类型允许执行过程中定制模板行为。唯一实现的选项允许在索引map时，控制如何处理丢失的key。默认的，他现在可以覆盖，像以前一样：继续一个无效值。
- time包的Time类型有一个新的方法AppendFormat， 用于在打印时间值的时候避免分配。
- unicode包和相关的整个系统已经由7.0版本升级至 Unicode 8.0了。
