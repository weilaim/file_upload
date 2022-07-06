package serializer

import "github.com/weilaim/blog-api/model"

type File struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Title     string `json:"title" `
	FileUrl   string ` json:"file_url"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"create_at"`
	UpdatedAt int64  `json:"update_at"`
}

func BuildFile(item model.Files) File {
	return File{
		ID:        item.ID,
		UserId:    item.UserId,
		Title:     item.Title,
		FileUrl:   item.FileUrl,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}

func BuildFiles(items []model.Files) (files []File){
	for _,item := range items {
		file := BuildFile(item)
		files = append(files, file)
	}

	return files
}

//BuildFileResponse
type FileRep struct {
	File interface{} `json:"file"`
}
func BuildFileResponse(item interface{}) Response {
	return Response{
		Data: FileRep{
			File: item,
		}, 
	}
}