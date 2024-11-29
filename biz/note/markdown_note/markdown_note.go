package markdown_note

import (
	"time"

	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
	"github.com/dgdts/UniversalServer/biz/note/note_handler"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ note_handler.NoteHandler = (*MarkdownNoteHandler)(nil)

type MarkdownNoteHandler struct {
}

func (n *MarkdownNoteHandler) CreateNote(ctx *biz_context.BizContext, req *note.CreateNoteRequest) (*note.CreateNoteResponse, error) {
	markdownNote := req.GetMarkdownNote()
	data := &MarkdownNoteData{
		Content: markdownNote.Content,
		ID:      global_id.GenerateUniqueID(),
	}
	err := InsertMarkdownNoteData(ctx, data)
	if err != nil {
		return nil, err
	}

	resp := &note.CreateNoteResponse{
		Id:        data.ID,
		CreatedAt: timestamppb.New(time.Now()),
	}

	return resp, nil
}

func (n *MarkdownNoteHandler) GetNote(ctx *biz_context.BizContext, req *note.GetNoteRequest) (*note.GetNoteResponse, error) {
	return nil, nil
}

func (n *MarkdownNoteHandler) UpdateNote(ctx *biz_context.BizContext, req *note.UpdateNoteRequest) (*note.UpdateNoteResponse, error) {
	return nil, nil
}

func (n *MarkdownNoteHandler) DeleteNote(ctx *biz_context.BizContext, req *note.DeleteNoteRequest) (*note.DeleteNoteResponse, error) {
	return nil, nil
}
