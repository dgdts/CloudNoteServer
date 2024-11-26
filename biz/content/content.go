package biz_content

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgdts/UniversalServer/biz/model/content"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"github.com/dgdts/UniversalServer/pkg/minio"
)

func CreateContent(ctx context.Context, req *content.ContentRequest, c *app.RequestContext) (*content.ContentResponse, error) {
	if req.Id == "" {
		contentData := &ContentData{
			ID:        global_id.GenerateUniqueID(),
			Title:     req.Title,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  req.Username,
			IsDeleted: false,
		}

		file, err := c.FormFile("content")
		if err != nil {
			return nil, err
		}
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		objectName := fmt.Sprintf("%s/%s_%s", contentData.Username, file.Filename, contentData.ID)

		url, err := minio.UploadFile(objectName, src, file.Header.Get("Content-Type"), time.Hour*24)
		if err != nil {
			return nil, err
		}
		contentData.Content = url
		err = InsertContent(ctx, contentData)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("content id is not empty")
	}

	return &content.ContentResponse{
		Message: "success",
	}, nil
}
