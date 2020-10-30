// OS
package Package_T

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	filename = "Generate_File/test.txt"
)

func OS_CreateFile() {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
	f.Close()
}

func OS_FileOpen() {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	f.Close()
}

func OS_Write() {
	fmt.Println("按字节写入 -> `Write()`")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	write, err := f.Write([]byte("hello"))
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	fmt.Println("write number = ", write)
	f.Close()

}

func OS_Write_String() {
	fmt.Println("按字符串写 -> `WriteString()`")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	write, err := f.WriteString("AAAAAAAAA")
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	fmt.Println("write number = ", write)
	f.Close()
}

func OS_Seek() {
	f, _ := os.OpenFile(filename, os.O_RDWR, 6)
	off, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off)

	n, _ := f.WriteAt([]byte("111"), off)
	fmt.Println(n)
	f.Close()
}

func OS_Fileinfo() {
	fmt.Println("获取文件描述信息")

	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("stat err: ", err)
		return
	}
	// 获取到的fileInfo内部包含 `文件名Name()`、`大小Size()`、`是否是目录IsDir()` 等操作。
	fmt.Printf("%T\n%v\n", fileInfo, fileInfo) // *os.fileStat
}

func OS_filepath() {
	fmt.Println(filepath.IsAbs(filename)) // false：判断是否是绝对路径
	fmt.Println(filepath.Abs(filename))   // 转换为绝对路径

	// 创建目录
	err := os.Mkdir("./test", os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
		return
	}
	// 创建多级目录
	err = os.MkdirAll("./dd/rr", os.ModePerm)
	if err != nil {
		fmt.Println("mkdirAll err: ", err)
		return
	}

	// 删除目录
	err = os.Remove("./test")
	err = os.RemoveAll("./dd")
}

func Cmd_args() {
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
