package serializer

import (
	"github.com/weilaim/blog-api/model"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// User 用户序列化器
type WxUser struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Code    string `json:"Code"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Token     string `json:"token"`
}


// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}


//BuildWxuser 
func BuildWxuser(user model.User,token string) WxUser {
	return WxUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
		Token: token,
	}

}

//wx 返回序列化器
func BuildWxuserResponse(user model.User, token string, msg string) Response {
	return Response{
		Code: 200,
		Msg:    msg,
		Data:   BuildWxuser(user, token),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Code: 200,
		Msg:    "成功",
		Data:   BuildUser(user),
	}
}

//BuildUsers 序列化用户列表
func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser((item))
		users = append(users, user)
	}
	return users
}
