package main

import (
	"fmt"

	"hkc.ink/rss/cmd"
)

// @title 研学社
// @version 1.0
// @description 研学社-go-后端接口
// @termsOfService http://swagger.io/terms/

// @contact.name hkc
// @contact.url http://hkc.ink
// @contact.email hkchenjianhui@foxmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api

func main() {
	cmd.Execute()
	fmt.Println()
}
