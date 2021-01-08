package timeutils

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func ParseDuration(str string) (time.Duration, error) {
	duration, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return  0, errors.Wrapf(err, "failed parse %s to int", str)
	}
	return time.Duration(duration), nil
}