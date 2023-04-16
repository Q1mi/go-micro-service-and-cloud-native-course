# 短链接项目



## 搭建项目的骨架

1. 建库建表

新建发号器表
```sql
CREATE TABLE `sequence` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stub` varchar(1) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT = '序号表';
```

新建长链接短链接映射表：
```sql
CREATE TABLE `short_url_map` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',
    
    `lurl` varchar(2048) DEFAULT NULL COMMENT '长链接',
    `md5` char(32) DEFAULT NULL COMMENT '长链接MD5',
    `surl` varchar(11) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE(`md5`),
    UNIQUE(`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '长短链映射表';
```

2. 搭建go-zero框架的骨架

2.1 编写`api`文件，使用goctl命令生成代码

```go
/* 短链接项目
* author: q1mi
*/

type ConvertRequest {
	LongUrl string `json:"longUrl"`
}

type ConvertResponse {
	ShortUrl string `json:"shortUrl"`
}

type ShowRequest {
	ShortUrl string `json:"shortUrl"`
}

type ShowResponse {
	LongUrl string `json:"longUrl"`
}

service shortener-api {
	
	@handler ConvertHandler
	post /convert(ConvertRequest) returns(ConvertResponse)
	
    // q1mi.cn/lycsa1
	@handler ShowHandler
	get /:shortUrl(ShowRequest) returns(ShowResponse)
	
}
```

2.2 根据api文件生成go代码

```bash
goctl api go -api shortener.api -dir .
```

3. 根据数据表生成model层代码
```bash
goctl model mysql datasource -url="root:root1234@tcp(127.0.0.1:3306)/db3" -table="short_url_map"  -dir="./model"

goctl model mysql datasource -url="root:root1234@tcp(127.0.0.1:3306)/db3" -table="sequence"  -dir="./model"
```

4. 下载项目依赖
```bash
go mod tidy
```

5. 运行项目
```bash
go run shortener.go
```
看到如下输出表示项目成功启动啦！
```bash
Starting server at 0.0.0.0:8888...
```

6. 修改配置结构体和配置文件
注意：两边一定一定一定要对齐！！！



## 参数校验

1. go-zero使用validator
https://pkg.go.dev/github.com/go-playground/validator/v10

下载依赖：
```bash
go get github.com/go-playground/validator/v10
```

导入依赖：
```bash
import "github.com/go-playground/validator/v10"
```

在api中为结构体添加validate tag，并添加校验规则

