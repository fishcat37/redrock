# API文档
## 1. 简介
- base URL: `http://localhost:8080`
- 响应格式：JSON
- 成功码：10000
- 失败码：500或400
## 2.API
### 2.1 user
#### 2.1.1 注册
- URL: POST `/user/register`
- 请求头：
- 请求参数：application/json
  |参数名|类型|必填|说明|
  |:---|:---|:---|:---|
  |username|string|是|用户名|
  |password|string|是|密码|
- 返回参数：
- 返回示例：
  **成功**：
  ```json
  {
    "status":10000,
    "info":"success"
  }
  ```
#### 2.1.2 登录
- URL: POST `/user/token`
- 请求头：
- 请求参数：application/json
  |参数名|类型|必填|说明|
  |---|---|---|---|
  |username|string|是|用户名|
  |password|string|是|密码|
- 返回参数：
  |