# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399# df-demo

### 要点
1. 先定义表结构、生成到数据库中，配置好dao下的数据库连接配置；然后执行`gf gen dao`根据表结构生成对应的model文件（do/entity目录下文件）—— 不可修改
2. api下文件，主要用于定义接口文档（纯对外）
3. model下第一层定义逻辑层用的请求参数、响应参数（纯对内部，做数据处理）
4. 编写logic需要用到的方法接口实现（它是准备工作）
5. 项目下执行`gf gen service`，根据当前logic生成service层（它是logic和controller层的接口抽象）
6. 编写controller层（会用到上面api层定义的请求结构体、数据处理用model层定义的数据结构体，使用service间接调用logic）
7. 添加路由（`cmd.go`中）

### 注意点
1. 它是不管理迁移的，我们是通过先到库中生成好表结构体，然后自动生成表 model相关结构体的
2. 它的controller层不直接与logic层通信，他们通过`service`接口层接耦合
3. 配置数据库有两个地方,要注意区分
   - 1. hack下的，用于dao生成
   - 2. manifest下的，用于服务运行时数据库交互
4. 关于路由定义，我们在`cmd.go`中的路由定义是不全面的，它的方法和路径是在api层定义的；因此真正的路由是`cmd.go`和`api`层加一起来看才是完整的，另外通过/swagger 看也不错的
5. 所有的controller定义都要遵循格式`Handler(ctx context.Context, req *Request) (res *Response, err error)`的
6. 执行完`gf run main.go`后 修改代码其实会自动生效的（热更新）
7. 路径中的参数`user/{id}` 也需要写在api参数中


### 操作相关
1. 返回一个list如果要用灵活使用`interface{}`注意`[]interface{}`是不能用一个切片的任意类型赋值的哦
2. 对于controller的公用部分放到`api.go`下