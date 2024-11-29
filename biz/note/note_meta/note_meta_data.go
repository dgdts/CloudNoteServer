package note_meta

import (
	"time"

	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/pkg/mongo"
)

type NoteMeta struct {
	ID         string    `bson:"_id"`
	UserID     string    `bson:"user_id"`
	Title      string    `bson:"title"`
	Type       string    `bson:"type"`    // markdown/mindmap/table etc
	NoteID     string    `bson:"note_id"` // 指向具体内容的ID
	Version    int64     `bson:"version"`
	IsPublic   bool      `bson:"is_public"`
	ShareToken string    `bson:"share_token,omitempty"`
	Tags       []string  `bson:"tags,omitempty"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

const NoteMetaCollection = "note_meta"

func InsertNoteMeta(ctx *biz_context.BizContext, noteMeta *NoteMeta) error {
	r := mongo.Inserter(ctx.GlobalCollection(NoteMetaCollection)).Insert(ctx, noteMeta)
	return r.Error()
}
