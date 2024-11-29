package asset

import (
	"time"
)

// 资源表
type Asset struct {
	ID           string    `bson:"_id"`
	UserID       string    `bson:"user_id"`
	NoteID       string    `bson:"note_id"`
	Type         string    `bson:"type"`      // image/attachment
	Name         string    `bson:"name"`      // 原始文件名
	Size         int64     `bson:"size"`      // 文件大小
	MimeType     string    `bson:"mime_type"` // 文件类型
	Hash         string    `bson:"hash"`      // 文件哈希，用于去重
	OssPath      string    `bson:"oss_path"`  // OSS存储路径
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	IsReferenced bool      `bson:"is_referenced"` // 是否被引用
	References   []string  `bson:"references"`    // 引用此资源的文档ID列表
}

// 在 NoteContent 中添加资源引用记录
type NoteContent struct {
	ID        string    `bson:"_id"`
	NoteID    string    `bson:"note_id"`
	Content   string    `bson:"content"`
	Assets    []string  `bson:"assets"` // 引用的资源ID列表
	Version   int64     `bson:"version"`
	UpdatedAt time.Time `bson:"updated_at"`
}
