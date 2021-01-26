# GitHub OAuth2 登录 Demo

使用 GitHub 进行快捷登录/第三方登录的示例代码。

## 使用方式

1. 下载项目代码
```
git clone https://github.com/gettg/oauth-demo.git
cd oauth-demo
```
2. 替换信息
首先需要申请 OAuth APP ，地址: [Register a new OAuth application](https://github.com/settings/applications/new)

替换代码中的 TODO 信息
- /public/index.html --> client_id
- main.go --> clientSecret , clientID

3. 运行项目
```
go build 
./oauth-demo
./oauth-demo.exe
或者
go run main.go
```
4. 根据实际情况进行补充修改
 

参考文章:
> https://docs.github.com/en/developers/apps/authorizing-oauth-apps
> https://www.ruanyifeng.com/blog/2019/04/github-oauth.html