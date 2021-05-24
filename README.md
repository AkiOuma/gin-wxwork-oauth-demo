# A Demo for Wechat Work Oauth using Gin

A demo for Oauth in wechat work using github.com/gin-gonic/gin v1.7.2

## Usage
This project using go mod

### Install Dependency
```shell
go mod tidy
```

### Set up your application on wechat work console
Update this part later

### Set the application information in code file

#### webapp/login.html
```html
<script>
  const state = 'you can use a random integer'
  const redirect_uri = 'url of your project';
  const appid = 'cropid of your company';
  const agentid = 'app id of your application set in wechat work console';
  const base = 'https://open.work.weixin.qq.com/wwopen/sso/qrConnect?';

  const link = document.getElementById('login');
  link.href = `${base}appid=${appid}&agentid=${agentid}&redirect_uri=${redirect_uri}&state=${state}`;
</script>
```

#### controller/oauth.go
```go
func GetAccessToken() (result dto.AccessTokenDTO, err error) {
	corpid := "corpid of you company"
	corpsecret := "secret of your app"
	
  // the rest code will be ignore...
	return
}
```

Then you can start the project using `go run .` or debug in vscode