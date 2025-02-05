package markdown_note

import (
	"github.com/dgdts/CloudNoteServer/biz/biz_context"
	"github.com/dgdts/CloudNoteServer/pkg/mongo"
)

const (
	MarkdownNoteCollection = "markdown_notes"
	MarkdownNoteIDField    = "_id"
)

type MarkdownNoteData struct {
	ID      string `bson:"_id"`
	Content string `bson:"content"`
}

func InsertMarkdownNoteData(ctx *biz_context.BizContext, data *MarkdownNoteData) error {
	r := mongo.Inserter(ctx.GlobalCollection(MarkdownNoteCollection)).Insert(ctx, data)
	return r.Error()
}

func GetMarkdownNoteData(ctx *biz_context.BizContext, id string) (*MarkdownNoteData, error) {
	r := mongo.Finder(ctx.GlobalCollection(MarkdownNoteCollection)).FindOne(ctx, MarkdownNoteIDField, id)
	if r.Error() != nil {
		return nil, r.Error()
	}

	var data MarkdownNoteData
	err := r.Read(&data)
	return &data, err
}

func CountMarkdownNoteData(ctx *biz_context.BizContext, id string) (int64, error) {
	r, err := mongo.Finder(ctx.GlobalCollection(MarkdownNoteCollection)).CountDocuments(ctx, MarkdownNoteIDField, id)
	return r, err
}

func UpdateMarkdownNoteData(ctx *biz_context.BizContext, id string, data *MarkdownNoteData) error {
	r := mongo.Updater(ctx.GlobalCollection(MarkdownNoteCollection)).WithEqFilter(MarkdownNoteIDField, id).ReplaceOne(ctx, data)
	return r.Error()
}
