package common

import (
	"fmt"
	"os"
)

func InitFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//0777: 拥有者+同组+所有人 可读可写可执行
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return fmt.Errorf("静态文件夹%s创建失败：%s", path, err.Error())
		}
	}
	return nil
}
