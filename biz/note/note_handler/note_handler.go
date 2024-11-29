package note_handler

import (
	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
)

type NoteHandler interface {
	CreateNote(ctx *biz_context.BizContext, req *note.CreateNoteRequest) (*note.CreateNoteResponse, error)
	GetNote(ctx *biz_context.BizContext, req *note.GetNoteRequest) (*note.GetNoteResponse, error)
	UpdateNote(ctx *biz_context.BizContext, req *note.UpdateNoteRequest) (*note.UpdateNoteResponse, error)
	DeleteNote(ctx *biz_context.BizContext, req *note.DeleteNoteRequest) (*note.DeleteNoteResponse, error)
}
