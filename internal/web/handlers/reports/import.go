package reports

import (
	"net/http"
	"regexp"
	"sync"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/pkg/amazon/aws"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/pkg/amazon/aws/storage/s3"
)

var (
	contentType = regexp.MustCompile(`Content-Type`).MatchString
)

func (rp *Reports) Import(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(rp.sessionStore, rp.logger, w, r))

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
