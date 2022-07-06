package serializer

import "github.com/weilaim/blog-api/model"

// openid 序列化器

type Wxopen struct {
	Wxuid string `json:"wxuid"`
	Id    int64    `json:"id"`
	ErrCode   int    `json:"errcode"`
	ErrMSG    string `json:"errmsg"`
	Nick      string `form:"nick" json:"nick"`
	AvatarUrl string `form:"avatarurl" json:"avatarurl"`
	Sex       int32  `form:"sex" json:"sex"`
}

// 序列化格式
func BuildOpen(wxopen model.Wxuser) Wxopen {

	return Wxopen{
		Wxuid:   wxopen.Openid,
		Id:        int64(wxopen.Id),
		ErrCode: wxopen.ErrCode,
		// ErrMSG:    wxopen.ErrMSG,
		AvatarUrl: wxopen.AvatarUrl,
		Nick:      wxopen.Nick,
		Sex:       wxopen.Sex,
	}
}

func BuildOpenResponse(wxopen model.Wxuser, token string) Response {
	return Response{
		Code: 200,
		Msg:    "成功",
		Data:   BuildOpen(wxopen),
	}
}
