package kingbase

import (
	"fmt"
	"net/url"
	"strings"
)

// DriverName 数据库驱动、连接字符串协议名称
const DriverName = "kingbase"

// BuildUrl 构建 KES 数据库连接字符串
//
//   - 如： kingbase://user:password@host1[,host2]:port/dbname?search_path=public[&propName2=propValue2]…
//   - 可用连接参数请参考 gokb 包文档及： https://github.com/godoes/gorm-kingbase/blob/master/gokb/conn.go#L3277
func BuildUrl(user, password, host string, port int, dbname string, options ...map[string]string) string {
	propQuery := url.Values{}
	var timezone string
	for _, option := range options {
		for key, value := range option {
			if strings.EqualFold(key, "TimeZone") {
				timezone = value
				continue
			}
			propQuery.Add(key, value)
		}
	}

	kingbaseUrl := &url.URL{
		Scheme:   DriverName,
		User:     url.UserPassword(user, password),
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     dbname,
		RawQuery: propQuery.Encode(),
	}
	if timezone != "" {
		kingbaseUrl.RawQuery += fmt.Sprintf("&timezone=%s", timezone)
	}

	return kingbaseUrl.String()
}
