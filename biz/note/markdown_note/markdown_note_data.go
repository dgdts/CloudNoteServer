package markdown_note

import (
	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/pkg/mongo"
)

const MarkdownNoteCollection = "markdown_notes"

type MarkdownNoteData struct {
	ID      string `bson:"_id"`
	Content string `bson:"content"`
}

func InsertMarkdownNoteData(ctx *biz_context.BizContext, data *MarkdownNoteData) error {
	r := mongo.Inserter(ctx.GlobalCollection(MarkdownNoteCollection)).Insert(ctx, data)
	return r.Error()
}
