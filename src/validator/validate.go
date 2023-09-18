package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"go.vemo/src/models"
)

var ErrField error = errors.New("field has error")

func Validate(body io.ReadCloser, t *models.Task) (fields map[string]string, err error) {
	fields = map[string]string{}
	if err = json.NewDecoder(body).Decode(t); err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &unmarshalTypeError):
			fields[unmarshalTypeError.Field] = fmt.Sprintf("must be of type %s", unmarshalTypeError.Type)
			return
		case strings.HasPrefix(err.Error(), "parsing time"):
			fields["due"] = "invalid value"
			return
		default:
			return
		}
	}
	return t.Check()
}
