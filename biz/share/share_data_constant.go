package share

const (
	ShareNoteCollection        = "share_notes"
	ShareNoteCommentCollection = "share_note_comments"
)

const (
	IDField          = "_id"
	ShareNoteIDField = "share_note_id"
)

type ShareNoteStatus int

const (
	ShareNoteStatusDefault ShareNoteStatus = iota
	ShareNoteStatusExpired
	ShareNoteStatusCancel
)

type ShareNoteShareType int

const (
	ShareTypeDefault ShareNoteShareType = iota
	ShareTypeCanView
	ShareTypeCanComment
	ShareTypeCanEdit
)

var ShareNoteShareTypeMap = map[ShareNoteShareType]string{
	ShareTypeDefault:    "default",
	ShareTypeCanView:    "can_view",
	ShareTypeCanComment: "can_comment",
	ShareTypeCanEdit:    "can_edit",
}

var ShareNoteShareTypeReverseMap = map[string]ShareNoteShareType{
	"default":     ShareTypeDefault,
	"can_view":    ShareTypeCanView,
	"can_comment": ShareTypeCanComment,
	"can_edit":    ShareTypeCanEdit,
}

var ShareNoteStatusMap = map[ShareNoteStatus]string{
	ShareNoteStatusDefault: "default",
	ShareNoteStatusExpired: "expired",
	ShareNoteStatusCancel:  "cancel",
}

var ShareNoteStatusReverseMap = map[string]ShareNoteStatus{
	"default": ShareNoteStatusDefault,
	"expired": ShareNoteStatusExpired,
	"cancel":  ShareNoteStatusCancel,
}
