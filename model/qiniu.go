package model

// QiniuStorage 七牛对象存储
type QiniuStorage struct {
	Model

	Name   string `json:"name"`
	URL    string `json:"url"`
	UID    uint   `json:"uid"`
	UserID uint   `json:"userID"`
}
