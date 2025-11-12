package schema

// UploadFileRequest upload file request
type UploadFileRequest struct {
	File        []byte         `json:"file"`
	RequestBody UploadFileBody `json:"user"`
}

type UploadFileBody struct {
	User string `json:"user"`
}

// UploadFileResponse upload file response
type UploadFileResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	MineType  string `json:"mine_type"`
	CreateBy  string `json:"create_by"`
	CreateAt  int64  `json:"create_time"`
}
