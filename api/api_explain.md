
# RedRock电商API文档

---

## 用户相关接口

### 1. 用户注册
**请求路径**: `POST /user/register`  
**请求头**: `Content-Type: application/json`  
**请求参数**:
| 名称       | 类型   | 必选 | 说明   |
|------------|--------|------|--------|
| username   | string | 是   | 用户名 |
| password   | string | 是   | 密码   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```
**状态码**:
```
    UserNotExistCode  = 1001	//用户不存在
	PasswordWrongCode    = 1002	//密码错误
	UserExistCode    = 1003		//用户已存在
	DataBaseErrCode  = 1004		//数据库操作错误
	MakeTokenErrCode = 1005		//生成token错误
	TokenErrCode     = 1006		//token错误
	RequestErrCode   = 1007		//请求错误
	ProductNotExistCode = 1008	//商品不存在
	CartNotExistCode    = 1009	//购物车不存在
	NotCommentsCode     = 1010	//没有评论
	SuccessCode      = 10000	//成功

```
---

### 2. 用户登录（获取Token）
**请求路径**: `GET /user/token`  
**请求头**:  
**请求参数**: `application/json`
| 名称       | 类型   | 必选 | 说明               |
|------------|--------|------|--------------------|
| username   | string | 是   | 用户名/手机号/邮箱 |
| password   | string | 是   | 密码               |

**返回参数**:
| 字段名         | 类型   | 说明           |
|----------------|--------|----------------|
| refresh_token  | string | 刷新令牌       |
| token          | string | 访问令牌       |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "refresh_token": "[refresh_token]",
    "token": "[token]"
  }
}
```

---

### 3. 刷新Token
**请求路径**: `GET /user/token/refresh`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/query`
| 名称           | 位置   | 类型   | 必选 | 说明         |
|----------------|--------|--------|------|--------------|
| refresh_token  | query  | string | 是   | 刷新令牌     |

**返回参数**:
| 字段名         | 类型   | 说明           |
|----------------|--------|----------------|
| token          | string | 新访问令牌     |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "token": "[new_token]"
  }
}
```

---

### 4. 修改用户密码
**请求路径**: `PUT /user/password`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称          | 类型   | 必选 | 说明   |
|---------------|--------|------|--------|
| old_password  | string | 是   | 旧密码 |
| new_password  | string | 是   | 新密码 |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

### 5. 获取用户信息
**请求路径**: `GET /user/info/[user_id]`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/query`
| 名称      | 位置   | 类型   | 必选 | 说明   |
|-----------|--------|--------|------|--------|
| user_id   | path   | string | 是   | 用户ID |

**返回参数**:
| 字段名   | 类型   | 说明       |
|----------|--------|------------|
| user     | object | 用户信息   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "user": {
      "id": "123456",
      "nickname": "你好世界",
      "phone": "12345678909",
      "email": "test@gmail.com"
    }
  }
}
```

---

### 6. 修改用户信息
**请求路径**: `PUT /user/info`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称          | 类型   | 必选 | 说明                   |
|---------------|--------|------|------------------------|
| nickname      | string | 否   | 昵称                   |
| avatar        | string | 否   | 头像链接               |
| introduction  | string | 否   | 简介                   |
|telephone      | string | 否   | 电话号码               |
|email          | string | 否   | 邮箱                   |
|qq| string | 否 | QQ号码 |
|gener| string | 否 | 性别 |
|birthday| string | 否 | 生日 |
**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

## 商品相关接口

### 7. 获取商品列表
**请求路径**: `GET /product/list`  
**返回参数**:
| 字段名     | 类型    | 说明       |
|------------|---------|------------|
| products   | array   | 商品列表   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "products": [
      {
        "product_id": "1",
        "name": "傲慢与偏见",
        "price": 9.8
      }
    ]
  }
}
```

---

### 8. 搜索商品
**请求路径**: `GET /product/search`  
**请求参数**: `application/query`
| 名称           | 位置   | 类型   | 必选 | 说明       |
|----------------|--------|--------|------|------------|
| product_name   | query  | string | 是   | 商品名称   |

**返回参数**:
| 字段名     | 类型    | 说明       |
|------------|---------|------------|
| products   | array   | 搜索结果   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "products": [
      {
        "product_id": "1",
        "name": "傲慢与偏见"
      }
    ]
  }
}
```

---

### 9. 加入购物车
**请求路径**: `PUT /product/addCart`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称         | 类型   | 必选 | 说明   |
|--------------|--------|------|--------|
| product_id   | string | 是   | 商品ID |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

### 10. 获取购物车商品列表
**请求路径**: `GET /product/cart`  
**请求头**: `Authorization: [token]`  
**返回参数**:
| 字段名     | 类型    | 说明         |
|------------|---------|--------------|
| products   | array   | 购物车商品   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "products": [
      {
        "product_id": "1",
        "name": "傲慢与偏见",
        "price": 9.8
      }
    ]
  }
}
```

---

### 11. 获取商品详情
**请求路径**: `GET /product/info/[product_id]`  
**请求参数**:
| 名称         | 位置   | 类型   | 必选 | 说明   |
|--------------|--------|--------|------|--------|
| product_id   | path   | string | 是   | 商品ID |

**返回参数**:
| 字段名          | 类型    | 说明               |
|-----------------|---------|--------------------|
| name            | string  | 商品名称           |
| price           | float   | 价格               |
| is_addedCart    | boolean | 是否在购物车中     |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "name": "傲慢与偏见",
    "price": 9.8,
    "is_addedCart": true
  }
}
```

---

### 12. 获取分类商品列表
**请求路径**: `GET /product/[type]`  
**请求参数**:
| 名称       | 位置   | 类型   | 必选 | 说明     |
|------------|--------|--------|------|----------|
| type       | path   | string | 是   | 商品分类 |

**返回参数**:
| 字段名     | 类型    | 说明       |
|------------|---------|------------|
| products   | array   | 商品列表   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "products": [
      {
        "product_id": "1",
        "name": "傲慢与偏见"
      }
    ]
  }
}
```

---

## 评论相关接口

### 13. 获取商品评论
**请求路径**: `GET /comment/[product_id]`  
**请求参数**:
| 名称         | 位置   | 类型   | 必选 | 说明   |
|--------------|--------|--------|------|--------|
| product_id   | path   | string | 是   | 商品ID |

**返回参数**:
| 字段名     | 类型    | 说明       |
|------------|---------|------------|
| comments   | array   | 评论列表   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "comments": [
      {
        "content": "这本书很好！",
        "nickname": "用户A"
      }
    ]
  }
}
```

---

### 14. 发表评论
**请求路径**: `POST /comment/[product_id]`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称       | 类型   | 必选 | 说明     |
|------------|--------|------|----------|
| content    | string | 是   | 评论内容 |

**返回参数**:
| 字段名       | 类型   | 说明       |
|--------------|--------|------------|
| comment_id   | string | 评论ID     |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": "comment_123"
}
```

---

### 15. 删除评论
**请求路径**: `DELETE /comment/[comment_id]`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称         | 位置   | 类型   | 必选 | 说明   |
|--------------|--------|--------|------|--------|
| comment_id   | path   | string | 是   | 评论ID |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

### 16. 更新评论
**请求路径**: `PUT /comment/[comment_id]`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称       | 类型   | 必选 | 说明     |
|------------|--------|------|----------|
| content    | string | 是   | 新评论内容 |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

### 17. 点赞/点踩
**请求路径**: `PUT /comment/praise`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/json`
| 名称         | 类型   | 必选 | 说明                     |
|--------------|--------|------|--------------------------|
| comment_id   | string | 是   | 评论ID                   |
| model        | int    | 是   | 1-点赞，2-点踩           |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success"
}
```

---

## 操作相关接口

### 18. 下单
**请求路径**: `POST /operate/order`  
**请求头**: `Authorization: [token]`  
**请求参数**: `application/form-data`
| 名称       | 类型   | 必选 | 说明       |
|------------|--------|------|------------|
| user_id    | string | 是   | 用户ID     |
| orders     | array  | 是   | 订单商品   |
| total      | float  | 是   | 订单总价   |

**返回参数**:
| 字段名     | 类型   | 说明     |
|------------|--------|----------|
| order_id   | string | 订单ID   |

**返回示例**:
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "order_id": 1
  }
}
```

---