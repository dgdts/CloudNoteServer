package note

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgdts/UniversalServer/biz/biz_config"
	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
	"github.com/dgdts/UniversalServer/biz/note/markdown_note"
	"github.com/dgdts/UniversalServer/biz/note/model"
	"github.com/dgdts/UniversalServer/biz/note/note_meta"
	"github.com/dgdts/UniversalServer/biz/note/types"
	"github.com/dgdts/UniversalServer/biz/share"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewNoteHandler(noteType string) types.NoteHandler {
	switch noteType {
	case types.NoteTypeMarkdown:
		return &markdown_note.MarkdownNoteHandler{}
	default:
		return nil
	}
}

func CreateNote(ctx *biz_context.BizContext, req *model.Note) (*note.CreateNoteResponse, error) {
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
		IsShare:   false,
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

func ListNotes(ctx *biz_context.BizContext, req *note.ListNotesRequest) (*note.ListNotesResponse, error) {
	metas, err := note_meta.ListNoteMetas(ctx, ctx.UserID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	noteMetas := make([]*note.NoteMeta, 0)

	for _, meta := range metas {
		noteMetas = append(noteMetas, &note.NoteMeta{
			UserId:    meta.UserID,
			Title:     meta.Title,
			Type:      meta.Type,
			NoteId:    meta.NoteID,
			Version:   meta.Version,
			IsShare:   meta.IsShare,
			ShareId:   meta.ShareID,
			Tags:      meta.Tags,
			CreatedAt: timestamppb.New(meta.CreatedAt),
			UpdatedAt: timestamppb.New(meta.UpdatedAt),
		})
	}

	return &note.ListNotesResponse{
		Notes:    noteMetas,
		Total:    int64(len(metas)),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func GetNote(ctx *biz_context.BizContext, req *note.GetNoteRequest) (*model.Note, error) {
	handler := NewNoteHandler(req.Type)
	if handler == nil {
		return nil, fmt.Errorf("invalid note type, got %s", req.Type)
	}

	resp, err := handler.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func validateUpdateNoteRequest(ctx *biz_context.BizContext, req *model.UpdateNote) (int64, error) {
	if req.ID == "" {
		return 0, fmt.Errorf("note id is required")
	}

	meta, err := note_meta.GetNoteMetaByNoteIDAndUserID(ctx, req.ID, ctx.UserID)
	if err != nil {
		return 0, err
	}
	if meta.Version != req.Version {
		return 0, fmt.Errorf("version mismatch, got %d", req.Version)
	}

	meta.Version++
	meta.UpdatedAt = time.Now()
	meta.Tags = req.Tags
	meta.Title = req.Title

	err = note_meta.UpdateNoteMeta(ctx, meta)
	if err != nil {
		return 0, err
	}

	return meta.Version, nil
}

func UpdateNote(ctx *biz_context.BizContext, req *model.UpdateNote) (*note.UpdateNoteResponse, error) {
	version, err := validateUpdateNoteRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	handler := NewNoteHandler(req.Type)
	if handler == nil {
		return nil, fmt.Errorf("invalid note type, got %s", req.Type)
	}
	resp, err := handler.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.Version = version

	return resp, nil
}

func CreateOrUpdateShareNote(ctx *biz_context.BizContext, req *note.ShareNoteRequest) (*note.ShareNoteResponse, error) {
	noteMeta, err := note_meta.GetNoteMetaByNoteIDAndUserID(ctx, req.NoteId, ctx.UserID)
	if err != nil {
		return nil, err
	}

	if noteMeta.ShareID != "" {
		return UpdateShareNote(ctx, noteMeta, req)
	}
	return CreateShareNote(ctx, noteMeta, req)
}

func CreateShareNote(ctx *biz_context.BizContext, noteMeta *note_meta.NoteMeta, req *note.ShareNoteRequest) (*note.ShareNoteResponse, error) {
	if noteMeta.IsShare {
		return nil, errors.New("note already shared")
	}

	var shareType share.ShareNoteShareType
	var okShareType bool
	if shareType, okShareType = share.ShareNoteShareTypeReverseMap[req.ShareType]; !okShareType {
		return nil, fmt.Errorf("invalid share type, got %s", req.ShareType)
	}

	shareNote := &share.ShareNote{
		ID:        global_id.GenerateUniqueID(),
		NoteID:    req.NoteId,
		UserID:    ctx.UserID,
		ShareType: shareType,
		Status:    share.ShareNoteStatusDefault,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	shareNote.ShareURL = fmt.Sprintf("https://%s/share/%s", biz_config.GetBizConfigInstance().ShareDomain, shareNote.ID)

	err := share.InsertShareNote(ctx, shareNote)
	if err != nil {
		return nil, err
	}

	// update note meta
	noteMeta.IsShare = true
	noteMeta.ShareID = shareNote.ID
	err = note_meta.UpdateNoteMeta(ctx, noteMeta)
	if err != nil {
		return nil, err
	}

	return &note.ShareNoteResponse{
		ShareUrl: shareNote.ShareURL,
	}, nil

}

func UpdateShareNote(ctx *biz_context.BizContext, noteMeta *note_meta.NoteMeta, req *note.ShareNoteRequest) (*note.ShareNoteResponse, error) {
	var shareType share.ShareNoteShareType
	var okShareType bool
	if shareType, okShareType = share.ShareNoteShareTypeReverseMap[req.ShareType]; !okShareType {
		return nil, fmt.Errorf("invalid share type, got %s", req.ShareType)
	}

	shareNote, err := share.GetShareNote(ctx, noteMeta.ShareID)
	if err != nil {
		return nil, err
	}

	shareNote.ShareType = shareType
	// update status if provided
	if req.Status != "" {
		var status share.ShareNoteStatus
		var okStatus bool
		if status, okStatus = share.ShareNoteStatusReverseMap[req.Status]; !okStatus {
			return nil, fmt.Errorf("invalid share status, got %s", req.Status)
		}
		shareNote.Status = status
	}

	shareNote.UpdatedAt = time.Now()

	err = share.UpdateShareNote(ctx, shareNote)
	if err != nil {
		return nil, err
	}

	noteMeta.IsShare = shareNote.Status == share.ShareNoteStatusDefault
	err = note_meta.UpdateNoteMeta(ctx, noteMeta)
	if err != nil {
		return nil, err
	}

	return &note.ShareNoteResponse{
		ShareUrl: shareNote.ShareURL,
	}, nil
}
