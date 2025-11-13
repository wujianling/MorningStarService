package schema

import (
	"github.com/wujianling/moringstaradmin/pkg/gormx"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"time"
)

// Defining the `Customer` struct.
type Customer struct {
	gormx.BaseDBStruct
	NickName string     `json:"nickName" gorm:"size:128;comment:名字"`              // 昵称
	UserName string     `json:"userName" gorm:"size:128;uniqueIndex;comment:用户名"` //用户名
	Avatar   string     `json:"avatar" gorm:"size:256;comment:头像"`                // 头像
	Email    string     `json:"email" gorm:"size:128;uniqueIndex;comment:邮箱"`     //邮箱
	Gender   string     `json:"gender" gorm:"size:20;comment:性别"`                 // 性别
	Birthday *time.Time `json:"birthDay" gorm:"comment:生日"`                       // 生日
	Phone    string     `json:"phone" gorm:"size:128;uniqueIndex;comment:电话"`     // 电话
	Password string     `json:"-" gorm:"size:128;comment:密码"`                     // 密码
	Status   string     `json:"status" gorm:"size:128;index;comment:用户状态"`        // 状态  enabled disabled
}

// Defining the query parameters for the `Customer` struct.
type CustomerQueryParam struct {
	util.PaginationParam
	LikeNickName string `form:"nickName" json:"nickName"` //昵称模糊查找
	LikeUserName string `form:"userName" json:"userName"` //用户名模糊查找
	LikeEmail    string `form:"email" json:"email"`       //邮箱模糊查找
	LikePhone    string `form:"phone" json:"phone"`       //电话模糊查找
	Status       string `form:"status" json:"status"`     //状态
}

// Defining the query options for the `Customer` struct.
type CustomerQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Customer` struct.
type CustomerQueryResult struct {
	Data       Customers
	PageResult *util.PaginationResult
}

// Defining the slice of `Customer` struct.
type Customers []*Customer

// Defining the data structure for creating a `Customer` struct.
type CustomerForm struct {
	NickName string     `json:"nickName" ` // 昵称
	UserName string     `json:"userName" ` //用户名
	Avatar   string     `json:"avatar" `   // 头像
	Email    string     `json:"email" `    //邮箱
	Gender   string     `json:"gender" `   // 性别
	Birthday *time.Time `json:"birthDay" ` // 生日
	Phone    string     `json:"phone" `    // 电话
	Password string     `json:"password" ` // 密码
	Status   string     `json:"status" `   // 状态  enabled disabled
}

// A validation function for the `CustomerForm` struct.
func (a *CustomerForm) Validate() error {
	return nil
}

// Convert `CustomerForm` to `Customer` object.
func (a *CustomerForm) FillTo(customer *Customer) error {
	customer.NickName = a.NickName
	customer.UserName = a.UserName
	customer.Avatar = a.Avatar
	customer.Email = a.Email
	customer.Gender = a.Gender
	customer.Birthday = a.Birthday
	customer.Phone = a.Phone
	customer.Password = a.Password
	customer.Status = a.Status
	return nil
}
