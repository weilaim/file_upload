package serializer

import "github.com/weilaim/blog-api/model"

type FileAcc struct {
	ID uint `json:"id"`
	Accname   string ` json:"accname"`
	Fieldid   string `json:"fieldid"`
	CreatedAt int64  `json:"create_at"`
	UpdatedAt int64  `json:"update_at"`
}

func BuildFileAcc(item model.FilesAcc) FileAcc {
	return FileAcc{
		ID: item.ID,
		Accname:item.Accname,
		Fieldid:item.Fieldid,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	  

}

func BuildFileAccs(items []model.FilesAcc) (files []FileAcc){
	for _,item := range items {
		file := BuildFileAcc(item)
		files = append(files, file)
	}

	return files
}

//BuildFileResponse
type FileAccRep struct {
	File interface{} `json:"file"`
}
func BuildFileAccResponse(item interface{}) Response {
	return Response{
		Data: FileRep{
			File: item,
		}, 
	}
}