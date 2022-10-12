package libs

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func CloudUpload(src string) (string, error) {
	cloud := os.Getenv("CLOUD_NAME")
	keys := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SEC")

	ctx := context.Background()
	imageId := uuid.New()

	cld, err := cloudinary.NewFromParams(cloud, keys, secret)
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{PublicID: imageId.String(), Folder: "goback"})
	if err != nil {
		return "", err
	}

	if err := os.Remove(src); err != nil {
		log.Println(err.Error())
	}

	return resp.SecureURL, nil
}
