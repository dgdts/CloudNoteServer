package markdown_note

import (
	"encoding/json"
	"time"

	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
	"github.com/dgdts/UniversalServer/biz/note/model"
	"github.com/dgdts/UniversalServer/biz/note/types"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ types.NoteHandler = (*MarkdownNoteHandler)(nil)

type MarkdownNoteHandler struct {
}

func validateAndParseNote(req *model.Node) (*MarkdownNoteData, error) {
	var markdownNote MarkdownNoteData
	err := json.Unmarshal(req.Note, &markdownNote)
	if err != nil {
		return nil, err
	}

	return &markdownNote, nil
}

func (n *MarkdownNoteHandler) CreateNote(ctx *biz_context.BizContext, req *model.Node) (*note.CreateNoteResponse, error) {
	markdownNote, err := validateAndParseNote(req)
	if err != nil {
		return nil, err
	}

	markdownNote.ID = global_id.GenerateUniqueID()

	err = InsertMarkdownNoteData(ctx, markdownNote)
	if err != nil {
		return nil, err
	}

	resp := &note.CreateNoteResponse{
		Id:        markdownNote.ID,
		CreatedAt: timestamppb.New(time.Now()),
	}

	return resp, nil
}

func (n *MarkdownNoteHandler) GetNote(ctx *biz_context.BizContext, req *note.GetNoteRequest) (*model.Node, error) {
	markdownNote, err := GetMarkdownNoteData(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	noteContent, err := json.Marshal(markdownNote.Content)
	if err != nil {
		return nil, err
	}

	ret := &model.Node{
		Type: req.Type,
		Note: noteContent,
	}

	return ret, nil
}

func (n *MarkdownNoteHandler) UpdateNote(ctx *biz_context.BizContext, req *model.UpdateNode) (*note.UpdateNoteResponse, error) {
	markdownNote, err := validateAndParseNote(&req.Node)
	if err != nil {
		return nil, err
	}

	markdownNote.ID = req.ID
	err = UpdateMarkdownNoteData(ctx, req.ID, markdownNote)
	if err != nil {
		return nil, err
	}

	return &note.UpdateNoteResponse{}, nil
}

func (n *MarkdownNoteHandler) DeleteNote(ctx *biz_context.BizContext, req *note.DeleteNoteRequest) (*note.DeleteNoteResponse, error) {
	return nil, nil
}
