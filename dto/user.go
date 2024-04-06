// 存放数据传输模型，定义了与前端交互时的数据结构
package dto

type User struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Avatar   string `json:"Avatar"`
	Bio      string `json:"Bio"`
}
