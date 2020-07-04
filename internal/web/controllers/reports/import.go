package reports

import (
	"net/http"
	"regexp"
	"sync"

	"gitlab.com/getlytica/lytica-app/internal/web/lib/amazon/aws"
	"gitlab.com/getlytica/lytica-app/internal/web/lib/amazon/aws/storage/s3"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
)

var (
	contentType = regexp.MustCompile(`Content-Type`).MatchString
)

func (rp *Reports) Import(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(rp.sessionStore, rp.logger, w, r)
	user := session.Values["User"].(models.User)

	maxSize := int64(40960000)

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		if contentType(err.Error()) {
			rp.logger.Error().Err(err).Msg("issues with the form content-type")
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sess, err := aws.NewSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup

	files := r.MultipartForm.File["file"]
	for i := 0; i < len(files); i++ {
		file, err := files[i].Open()

		if err == nil {
			wg.Add(1)

			go func() {
				err = s3.Upload(user.UserId, sess, file, files[i], &wg)
				if err != nil {
					rp.logger.Error().Err(err).Msg("failed to upload")
				}
			}()

			wg.Wait()
		}
	}

	w.WriteHeader(http.StatusOK)
}
