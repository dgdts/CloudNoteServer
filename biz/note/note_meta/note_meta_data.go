package note_meta

import (
	"time"

	"github.com/dgdts/CloudNoteServer/biz/biz_context"
	"github.com/dgdts/CloudNoteServer/pkg/mongo"
)

const (
	NoteMetaIDField        = "_id"
	NoteMetaUserIDField    = "user_id"
	NoteMetaCreatedAtField = "created_at"
	NoteMetaNoteIDField    = "note_id"
)

type NoteMeta struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	Title     string    `bson:"title"`
	Type      string    `bson:"type"`    // markdown/mindmap/table etc
	NoteID    string    `bson:"note_id"` // 指向具体内容的ID
	Version   int64     `bson:"version"`
	IsShare   bool      `bson:"is_share"`
	ShareID   string    `bson:"share_id,omitempty"`
	Tags      []string  `bson:"tags,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

const NoteMetaCollection = "note_meta"

func InsertNoteMeta(ctx *biz_context.BizContext, noteMeta *NoteMeta) error {
	r := mongo.Inserter(ctx.GlobalCollection(NoteMetaCollection)).Insert(ctx, noteMeta)
	return r.Error()
}

func GetNoteMetaByNoteIDAndUserID(ctx *biz_context.BizContext, noteID string, userID string) (*NoteMeta, error) {
	r := mongo.Finder(ctx.GlobalCollection(NoteMetaCollection)).FindOne(ctx, NoteMetaNoteIDField, noteID, NoteMetaUserIDField, userID)
	if r.Error() != nil {
		return nil, r.Error()
	}

	var noteMeta NoteMeta
	err := r.Read(&noteMeta)
	return &noteMeta, err
}

func UpdateNoteMeta(ctx *biz_context.BizContext, noteMeta *NoteMeta) error {
	r := mongo.Updater(ctx.GlobalCollection(NoteMetaCollection)).WithEqFilter(NoteMetaIDField, noteMeta.ID).ReplaceOne(ctx, noteMeta)
	return r.Error()
}

func ListNoteMetas(ctx *biz_context.BizContext, userID string, page, pageSize int64) ([]*NoteMeta, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	filter := []interface{}{
		NoteMetaUserIDField, userID,
	}

	r := mongo.Finder(ctx.GlobalCollection(NoteMetaCollection)).WithSort(NoteMetaCreatedAtField, -1).Find(ctx, page, pageSize, filter...)
	if r.Error() != nil {
		return nil, r.Error()
	}

	notes := make([]*NoteMeta, 0)
	err := r.Read(&notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
