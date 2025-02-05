package binary

import (
	"fmt"
	"time"

	"github.com/dgdts/CloudNoteServer/biz/biz_context"
	"github.com/dgdts/CloudNoteServer/biz/model/binary"
	"github.com/dgdts/CloudNoteServer/pkg/minio"
)

func generateFilename(userID string, filename string) string {
	return fmt.Sprintf("%s/%s/%s", userID, time.Now().Format("20240101"), filename)
}

// GetUploadToken get upload token
// do not consider the file size and capacity of the user for now
// need a completely system to deal with the oss usage control for the user
// TODO: add the oss usage control system
func GetUploadToken(ctx *biz_context.BizContext, req *binary.GetUploadTokenRequest) (*binary.GetUploadTokenResponse, error) {
	filename := generateFilename(ctx.UserID, req.Filename)

	uploadURL, err := minio.GetUploadPresignedURL(filename, time.Hour*24)
	if err != nil {
		return nil, err
	}
	return &binary.GetUploadTokenResponse{
		UploadUrl: uploadURL.String(),
		ExpiresIn: 24 * 60 * 60,
	}, nil
}
