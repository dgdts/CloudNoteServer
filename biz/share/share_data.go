package share

import (
	"time"

	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/pkg/mongo"
)

type ShareNote struct {
	ID        string             `bson:"_id"`
	NoteID    string             `bson:"note_id"`
	UserID    string             `bson:"user_id"`
	NoteType  string             `bson:"note_type"`
	ShareType ShareNoteShareType `bson:"share_type"`
	ShareURL  string             `bson:"share_url"`
	ViewCount int                `bson:"view_count"`
	Status    ShareNoteStatus    `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type ShareNoteComment struct {
	ID          string    `bson:"_id"`
	ShareNoteID string    `bson:"share_note_id"`
	Alias       string    `bson:"alias"`
	Content     string    `bson:"content"`
	IP          string    `bson:"ip"`
	CreatedAt   time.Time `bson:"created_at"`
}

func InsertShareNote(ctx *biz_context.BizContext, shareNote *ShareNote) error {
	r := mongo.Inserter(ctx.GlobalCollection(ShareNoteCollection)).Insert(ctx, shareNote)
	return r.Error()
}

func UpdateShareNote(ctx *biz_context.BizContext, shareNote *ShareNote) error {
	r := mongo.Updater(ctx.GlobalCollection(ShareNoteCollection)).WithEqFilter(IDField, shareNote.ID).ReplaceOne(ctx, shareNote)
	return r.Error()
}

func GetShareNote(ctx *biz_context.BizContext, shareID string) (*ShareNote, error) {
	r := mongo.Finder(ctx.GlobalCollection(ShareNoteCollection)).FindOne(ctx, IDField, shareID)
	if r.Error() != nil {
		return nil, r.Error()
	}

	var shareNote ShareNote
	err := r.Read(&shareNote)
	return &shareNote, err
}

func InsertShareNoteComment(ctx *biz_context.BizContext, comment *ShareNoteComment) error {
	r := mongo.Inserter(ctx.GlobalCollection(ShareNoteCommentCollection)).Insert(ctx, comment)
	return r.Error()
}

func GetShareNoteCommentsWithShareNoteID(ctx *biz_context.BizContext, shareNoteID string, page int, pageSize int) ([]*ShareNoteComment, error) {
	r := mongo.Finder(ctx.GlobalCollection(ShareNoteCommentCollection)).Find(ctx, int64(page), int64(pageSize), ShareNoteIDField, shareNoteID)
	if r.Error() != nil {
		return nil, r.Error()
	}

	var comments []*ShareNoteComment
	err := r.Read(&comments)
	return comments, err
}
