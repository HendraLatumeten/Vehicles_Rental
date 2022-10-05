package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

func FileUpload(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		img := r.FormValue("image")

		uploadedFile, handler, err := r.FormFile("image")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer uploadedFile.Close()

		dt := time.Now()
		filename := dt.Format("15:04:05") + "_" + handler.Filename
		if img != "" {
			filename = fmt.Sprintf("%s%s", img, filepath.Ext(handler.Filename))
		}

		fileExtension := filepath.Ext(filename)

		//validate extension
		if fileExtension != ".jpg" && fileExtension != ".png" && fileExtension != ".jpeg" {
			libs.Respone("extension salah", 400, true).Send(w)
			return
		}

		//fileLocation := filepath.Join(dir, "../bin/vehicles_rental/image", filename)
		fileDestination, err := os.Create("images/" + filename)
		if err != nil {
			libs.Respone(err.Error(), 400, true)
			return
		}

		_, err = io.Copy(fileDestination, uploadedFile)
		if err != nil {
			libs.Respone(err.Error(), 400, true)
			return
		}
		//test
		ctx := context.WithValue(r.Context(), "image", filename)
		next.ServeHTTP(w, r.WithContext(ctx))

	}

}
