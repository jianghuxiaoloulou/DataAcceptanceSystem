# 项目
# ****PACS集成平台****

# 项目描述
* 本平台主要处理PACS数据和其它系统的数据处理
  @import "image/第三方数据处理系统.png"

# 设计依据
* 创建web服务与PACS系统和第三方系统通信
* 第三方系统的业务根据医院实现可配置

# 目录结构
<!-- 
configs：配置文件。
docs：文档集合。
global：全局变量。
internal：内部模块。
model：数据库相关操作。
routers：路由相关逻辑处理。
pkg：项目相关的模块包。
storage：项目生成的临时文件。
scripts：各类构建，安装，分析等操作的脚本。
-->
# 路由
<!-- 在 RESTful API 中 HTTP 方法对应的行为动作分别如下：

GET：读取/检索动作。
POST：新增/新建动作。
PUT：更新动作，用于更新一个完整的资源，要求为幂等。
PATCH：更新动作，用于更新某一个资源的一个组成部分，也就是只需要更新该资源的某一项，就应该使用 PATCH 而不是 PUT，可以不幂等。
DELETE：删除动作。 -->

# 对象管理

# 公共组件
配置管理
数据库连接

Oracle数据库生产环境需要配置客户端程序Tool/instantcloent_21_9

日志写入

# 文件配置文件读取：go get -u github.com/spf13/viper
Viper 是适用于GO 应用程序的完整配置解决方案

# 日志：go get -u gopkg.in/natefinch/lumberjack.v2
它的核心功能是将日志写入滚动文件中，该库支持设置所允许单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数，
如果出现超出设置项的情况，就会对日志文件进行滚动处理。

# 生成接口文档
Swagger 相关的工具集会根据 OpenAPI 规范去生成各式各类的与接口相关联的内容，
常见的流程是编写注解 =》调用生成库-》生成标准描述文件 =》生成/导入到对应的 Swagger 工具
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template

@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	路由，从左到右分别为：路由地址，HTTP 方法

swag init

http://127.0.0.1:8000/swagger/index.html



## 第二次相同项目提交文件到github
# git add README.md
# git commit -m "first commit"
# git push -u origin master

# 代码编译命令
go build -ldflags="-H windowsgui" -o .\DataAcceptanceSystem\DataAcceptanceSystem.exe .\main.go


# 修改记录
* 2023/07/10 增加对接第三方PACS数据接口(上传的申请单数据应该从区域PACS中获取)
* 2023/06/15 重构PACS集成平台设计（修改对接区域PACS，对接多家HIS，增加医院配置，HIS配置，功能模块配置，字典表）
* 2023/06/15 修改数据库配置，拆分门诊/体检/住院数据
* 2023/06/13 修改初始化函数，增加初始化文件包
* 2023/06/08 修改分页逻辑
* 2023/06/07 增加排序，增加申请单字段
* 2023/06/07 增加申请单模糊查询功能
* 2023/06/06 增加急诊请求类型
* 2023/05/19 对接华卓HIS申请单
* 2023/04/28 对接中联HIS接口