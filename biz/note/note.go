package note

import (
	"fmt"

	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
	"github.com/dgdts/UniversalServer/biz/note/markdown_note"
	"github.com/dgdts/UniversalServer/biz/note/note_handler"
	"github.com/dgdts/UniversalServer/biz/note/note_meta"
	"github.com/dgdts/UniversalServer/pkg/global_id"
)

func validateCreateNoteRequest(req *note.CreateNoteRequest) error {
	switch req.Type {
	case NoteTypeMarkdown:
		if req.GetMarkdownNote() == nil {
			return fmt.Errorf("markdown note is required")
		}
	default:
		return fmt.Errorf("invalid note type, got %s", req.Type)
	}
	return nil
}

func CreateNote(ctx *biz_context.BizContext, req *note.CreateNoteRequest) (*note.CreateNoteResponse, error) {
	err := validateCreateNoteRequest(req)
	if err != nil {
		return nil, err
	}

	handler := NewNoteHandler(req.Type)
	if handler == nil {
		return nil, fmt.Errorf("invalid note type, got %s", req.Type)
	}

	resp, err := handler.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	noteMeta := &note_meta.NoteMeta{
		ID:        global_id.GenerateUniqueID(),
		UserID:    ctx.UserID,
		Title:     req.Title,
		Type:      req.Type,
		NoteID:    resp.Id,
		IsPublic:  false,
		Tags:      req.Tags,
		CreatedAt: resp.CreatedAt.AsTime(),
		UpdatedAt: resp.CreatedAt.AsTime(),
	}
	err = note_meta.InsertNoteMeta(ctx, noteMeta)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewNoteHandler(noteType string) note_handler.NoteHandler {
	switch noteType {
	case NoteTypeMarkdown:
		return &markdown_note.MarkdownNoteHandler{}
	default:
		return nil
	}
}
