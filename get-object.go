package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// GetObjectHandler downloads an object to the client
func GetObjectHandler(s3 S3Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		objectName := vars["objectName"]

		object, err := s3.GetObject(vars["bucketName"], objectName)
		if err != nil {
			msg := "error getting object"
			handleHTTPError(w, msg, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectName))
		w.Header().Set("Content-Type", "application/octet-stream")

		_, err = io.Copy(w, object)
		if err != nil {
			msg := "error copying object"
			handleHTTPError(w, msg, err, http.StatusInternalServerError)
			return
		}
	})
}