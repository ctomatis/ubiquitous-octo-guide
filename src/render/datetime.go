package render

import (
	"fmt"
	"strings"
	"time"

	"go.vemo/src/settings"
)

type DateTime time.Time

var layout = settings.Get("dt_format")

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(layout, s)
	*dt = DateTime(nt)
	return
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.String()), nil
}

func (dt *DateTime) String() string {
	t := time.Time(*dt)
	return fmt.Sprintf("%q", t.Format(layout))
}
