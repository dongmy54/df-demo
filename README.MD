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
8. 它的session管理默认是文件存储方式，在服务端设置了session后，请求中会返回set-cookies，key: `gfsessionid`
9. 注意在代码里，有v1和v2版本的区别哦
10. 关于分层
   - 只有api才分层（分前后端目录）
   - controller层也不分（controller中通过方法名区分前后端）
11. 测试时一定要注意版本问题 是v1还是v2
12. `ScanList`的用法，它的最大特点是不写死with,动态的组合关联数据
```go
 g.Model("user_scores").
      // ListItemValuesUnique 代表的是从中取值 这里代表从users切片中取每个User字段的Id字段
		Where("user_id", gdb.ListItemValuesUnique(users, "User", "Id")).
      // UserScores 扫描到UserScores字段
      // User代表 与User结构体关联
      // user_id:id 代表表user_scores.user_id 和 users.id关联
		ScanList(&users, "UserScores", "User", "user_id:id")
```
13. 默认非链式安全
```go
// 默认情况下非链式安全
	user := g.Model("user").Safe()
	//m := user.Clone()
	user.Limit(3)
	user.Where("id > ?", 4)

	r, _ := user.All()
	g.Dump(r)
```

### 操作相关
1. 返回一个list如果要用灵活使用`interface{}`注意`[]interface{}`是不能用一个切片的任意类型赋值的哦
2. 对于controller的公用部分放到`api.go`下

### DB
```go
// 冲突会自动更新还是挺方便
g.Model("user").Data(g.Map{
		"id":       1,
		"name":     "啦啦啦啦啦",
		"password": "pw1",
	}).Save()


//返回的是一个map
r, _ := dao.User.Ctx(context.TODO()).Fields("name", "mobile").
		Order("created_at desc").One()
// 通过r["name"]取值
```

1. where支持非常多的方式,map/slice/string/struc还是非常好用的
2. Scan既能扫单个对象 也能扫多个对象（切片）
3. 类似scoped写法,还是比较好，利于逻辑复用
```go
// 一个自定义的Name查询
func NameLike(name string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		return m.Where("name like ?", name+"%")
	}
}

func IdGreate10(m *gdb.Model) *gdb.Model {
	return m.Where("id > ?", 10)
}

dao.User.Ctx(context.TODO()).Handler(NameLike("KKK"), IdGreate10).One()
```