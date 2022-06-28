### 1. "用户登录"

1. route definition

- Url: /user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	Token string `json:"token"`
	Uid string `json:"uid"`
}
```

### 2. "获取用户信息"

1. route definition

- Url: /user/info
- Method: GET
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
	Uid string `path:"uid"`
}
```


3. response definition



```golang
type UserInfoResp struct {
	Uid string `path:"uid"`
	Name string `json:"name"`
	Email string `json:"email"`
}
```

