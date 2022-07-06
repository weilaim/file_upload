package serializer

import (
	"math/rand"
	"time"

	"github.com/weilaim/blog-api/model"
)

type Loves struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Sex          int      `json:"sex"`
	Content      string   `json:"content"`
	Lov          string   `json:"lov"`
	View         uint64   `json:"view"`
	Professional string   `json:"professional"`
	LoveList     []string `json:"lovelist"`
	CoverHeight  int64    `json:"coverheight"`
	CoverWidth   int64    `json:"coverwidth"`
	CreatedAt    int64    `json:"created_at"`
	Islike       int      `json:"islike"`    //点赞按钮的状态
	Likecount    int      `json:"likecount"` //点赞统计
	Collestu     int      `json:"collestu"`  //收藏按钮的状态
	ColleCount   int      `json:"collecount"`
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//BuilLove 序列化情书
func BuildLo(item model.Loves) Loves {
	rand.Seed(time.Now().UnixNano())
	height := randomInt(400, 500) // [0, max)
	width := randomInt(600, 800)
	return Loves{
		ID:           item.ID,
		Name:         item.Nick,
		Avatar:       item.Avatar,
		LoveList:     BuildLoveImgs(item.Loveimg),
		Sex:          item.Sex,
		Content:      item.Content,
		Lov:          item.Lve,
		Islike:       item.LikeSatus(item.Uid),
		Likecount:    item.LikeView(),
		View:         item.View(),
		Collestu:     item.Collestu(item.Uid),
		ColleCount:   item.ColleCount(),
		Professional: item.Professional,
		CoverHeight:  int64(height),
		CoverWidth:   int64(width),
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

type Lovels struct {
	LoveUrl string
}

func BuildLoveImgs(items []model.Loveimg) (lovels []string) {
	for _, item := range items {
		imglst := item.LoveUrl
		lovels = append(lovels, imglst)
	}
	return lovels
}

//BuildLoves 序列化视频列表
func BuildLoves(items []model.Loves) (loves []Loves) {
	for _, item := range items {
		love := BuildLo(item)
		loves = append(loves, love)
	}
	return loves
}
