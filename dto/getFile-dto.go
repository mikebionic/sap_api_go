package dto

type GetFile struct {
	File string `json:"File" binding:"required"`
	Guid string `json:"Guid" binding:"required"`
	Size string `json:"Size" binding:"required"`
}
