### Go 云餐厅
项目虽是 `Go 云餐厅` 但是实现和原版不一样，自己有封装和改进一些代码的写法和模块。例如：
<br/>1. 短信验证码不是存储在数据库中，而是存储在 Redis 中
```go
redisStore.Set(phone, code, time.Minute*5)
return esandSms.SendSms(phone, code)
```
<br/>2. 新增了一个 `BaseController`，其他 `Controller` 都继承它，里面封装了 `BuildResponse` 等方法。
```go
//BuildResponse
func BuildResponse(context *gin.Context, HttpStatus int, MessageCode string, Data interface{}) {
	requestId, _ := context.Get("requestId") // 中间件请求之前设置 requestId，返回中携带 
	context.JSON(HttpStatus, gin.H{
		"data":        Data,
		"message":     enums.ErrorMessage[MessageCode],
		"return_code": MessageCode,
		"response_id": requestId,
	})
}
```
<br/>3. 直接使用 Redis 存储 Session 信息，不使用第三方模块
```go
func SetSession(key string, value map[string]interface{}) {
	redisStore := RedisStoreEngine
	redisStore.SetJson(key, value, time.Minute*60*2)
}

func GetSession(key string) map[string]interface{} {
	redisStore := RedisStoreEngine
	return redisStore.GetJson(key)
}
```
<br/>4. 使用中间件 `LoginRequiredMiddleware` 进行接口的认证校验
```go
func LoginRequiredMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从三个地方获取 phone 的值
		var phone string
		phone = context.DefaultPostForm("phone", "")
		if phone == ""{
			phone = context.DefaultQuery("phone", "")
		}
		if phone == ""{
			phone, _ = context.Cookie("phone")
		}
		// 查询 redis 看当前 phone 是否登录
		userInfo := tool.GetSession(phone)
		if len(userInfo) == 0{
			controller.BuildResponse(context, http.StatusUnauthorized, enums.Unauthorized, "")
			return
		}
		context.Set("userInfo", userInfo)
	}
}
```

#### 后端 
- gin
- 短信服务商: [Esand](https://market.aliyun.com/products/57126001/cmapi00040066.html) 300次/月, 0元
- 数据库: MySQL/SQLite

#### 前端
- [Restaurant-Vue](https://github.com/rexyan/Restaurant-Vue)

#### TODO
- [ ] 数据库支持 SQLite
- [ ] 前后端 Docker 打包

#### 文档
- 接口文档
- 原文档 [HelloGin](https://github.com/rexyan/HelloGin)

#### 预览
