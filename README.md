# GORM KingbaseES Driver

基于 KingbaseES 数据库官方 Go 驱动源码
二次开发整理的开箱即用的 GORM KingbaseES 数据库驱动，
无需单独复制驱动源码到项目中。

获取 KingbaseES 最新官方 Go 驱动请访问
<https://www.kingbase.com.cn/download.html#drive>。

## 最低要求

- go `1.25`：<https://go.dev/dl/>
- gorm `v1.31.2`：<https://github.com/go-gorm/gorm>
- KES `V008R006C009B0014`：<https://www.kingbase.com.cn/download.html#database>

## 快速上手

### 安装

```shell
go get github.com/godoes/gorm-kingbase
```

### 默认用例

```go
import (
    "github.com/godoes/gorm-kingbase"
    "gorm.io/gorm"
)

// https://docs.kingbase.com.cn/cn/KES-V9R1C10/application/client_interface/GO/Gokb/go-2/
dsn := "host=localhost user=SYSTEM password=123456 dbname=TEST port=54321 sslmode=disable"
db, err := gorm.Open(kingbase.Open(dsn), &gorm.Config{})
```

### 配置用例

```go
import (
    "github.com/godoes/gorm-kingbase"
    "gorm.io/gorm"
)

db, err := gorm.Open(kingbase.New(kingbase.Config{
    // data source name, refer https://docs.kingbase.com.cn/cn/KES-V9R1C10/application/client_interface/GO/Gokb/go-2/
    DSN: "host=localhost user=SYSTEM password=123456 dbname=TEST port=54321 sslmode=disable",
    // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
    PreferSimpleProtocol: true,
}), &gorm.Config{})
```

## 参考文档

### KingbaseES 技术文档

- 产品简介：<https://docs.kingbase.com.cn/cn/KES-V9R1C10/introduction>
- 常见问题：<https://docs.kingbase.com.cn/cn/KES-V9R1C10/question/>
- GO 客户端编程接口 <https://docs.kingbase.com.cn/cn/KES-V9R1C10/category/gokb>
- 应用开发指南：<https://docs.kingbase.com.cn/cn/KES-V9R1C10/application/application-develop-guide/>

### GORM 文档

- 文档首页：<https://gorm.io/zh_CN/docs/>
- 声明模型：<https://gorm.io/zh_CN/docs/models.html>
- 连接到数据库：<https://gorm.io/zh_CN/docs/connecting_to_the_database.html>

---
