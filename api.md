# RedRock 电商 API 文档

## 通用返回格式
所有接口返回统一格式的 JSON 数据：
```json
{
  "status": 10000,    // 5位状态码，10000 表示成功
  "info": "success",  // 状态描述
  "data": {           // 实际返回数据
    "param1": "...",
    "param2": "..."
  }
}
```

---

## 用户相关接口

### 用户注册
**请求路径**：`POST /user/register`  
**请求头**：`Content-Type: application/json`  
**请求参数**：
| 字段名     | 类型   | 必选 | 说明   |
|-----------|--------|------|--------|
| username  | string | 是   | 用户名 |
| password  | string | 是   | 密码   |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success"
}
```

---

### 用户登录（获取 Token）
**请求路径**：`GET /user/token`  
**请求头**：`Content-Type: application/json`  
**请求参数**：
| 字段名     | 类型   | 必选 | 说明               |
|-----------|--------|------|--------------------|
| username  | string | 是   | 用户名/手机号/邮箱 |
| password  | string | 是   | 密码               |

**返回参数**：
| 字段名          | 类型   | 说明          |
|----------------|--------|---------------|
| refresh_token  | string | 刷新 Token    |
| token          | string | 访问 Token    |

**返回示例**：
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

### 刷新 Token
**请求路径**：`GET /user/token/refresh`  
**请求头**：  
- `Authorization: Bearer [token]`  
- `Content-Type: application/json`  

**请求参数**：
| 字段名          | 位置   | 类型   | 必选 | 说明          |
|----------------|--------|--------|------|---------------|
| refresh_token  | query  | string | 是   | 刷新 Token    |

**返回参数**：  
同登录接口，返回新的 `token` 和 `refresh_token`。

---

### 修改用户密码
**请求路径**：`PUT /user/password`  
**请求头**：  
- `Authorization: Bearer [token]`  
- `Content-Type: application/json`  

**请求参数**：
| 字段名       | 类型   | 必选 | 说明   |
|-------------|--------|------|--------|
| old_password| string | 是   | 旧密码 |
| new_password| string | 是   | 新密码 |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success"
}
```

---

## 商品相关接口

### 获取商品列表
**请求路径**：`GET /product/list`  
**返回参数**：
| 字段名     | 类型               | 说明       |
|-----------|--------------------|------------|
| products  | 复杂数据类型数组   | 商品列表   |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "products": [
      {
        "product_id": "1",
        "name": "傲慢与偏见",
        "price": 9.8,
        "cover": "http://127.0.0.1/picture_url"
      }
    ]
  }
}
```

---

### 加入购物车
**请求路径**：`PUT /product/addCart`  
**请求头**：  
- `Authorization: Bearer [token]`  
- `Content-Type: application/x-www-form-urlencoded`  

**请求参数**：
| 字段名      | 类型   | 必选 | 说明    |
|------------|--------|------|---------|
| product_id | string | 是   | 商品 ID |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success"
}
```

---

## 评论相关接口

### 发表评论
**请求路径**：`POST /comment/{product_id}`  
**请求头**：  
- `Authorization: Bearer [token]`  
- `Content-Type: application/json`  

**请求参数**：
| 字段名     | 类型   | 必选 | 说明     |
|-----------|--------|------|----------|
| content   | string | 是   | 评论内容 |

**返回参数**：
| 字段名     | 类型   | 说明       |
|-----------|--------|------------|
| comment_id| int    | 评论 ID    |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success",
  "data": 123
}
```

---

## 订单相关接口

### 下单
**请求路径**：`POST /operate/order`  
**请求头**：  
- `Authorization: Bearer [token]`  
- `Content-Type: application/json`  

**请求参数**：
| 字段名   | 类型               | 必选 | 说明       |
|---------|--------------------|------|------------|
| orders  | 复杂数据类型数组   | 是   | 订单内容   |
| address | 复杂数据类型       | 是   | 地址信息   |
| total   | float              | 是   | 订单总价   |

**返回参数**：
| 字段名    | 类型   | 说明     |
|----------|--------|----------|
| order_id | string | 订单 ID  |

**返回示例**：
```json
{
  "status": 10000,
  "info": "success",
  "data": {
    "order_id": "233"
  }
}
```