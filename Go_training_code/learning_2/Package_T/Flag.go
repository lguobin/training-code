package Package_T

import (
	"flag"
	"fmt"
)

func Flag_Demo() {
	var user, pwd, host string
	var port int

	// &user 就是接收用户命令行中输入的 -u 后面的参数值
	// "u" ,就是 -u 指定参数
	// "" , 默认值
	// "用户名,默认为空" 说明
	flag.StringVar(&user, "user", "", "用户名,默认为空")
	flag.StringVar(&pwd, "p", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)

}
