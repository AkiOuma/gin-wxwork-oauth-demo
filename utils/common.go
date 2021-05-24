package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"oauth/dto"
)

func JsonHandler(r *http.Response, target dto.Reuslter) (err error) {
	enc := json.NewDecoder(r.Body)
	err = enc.Decode(target)
	if target.GetErrCode() != 0 {
		err = errors.New(target.GetErrmsg())
	}
	return
}
