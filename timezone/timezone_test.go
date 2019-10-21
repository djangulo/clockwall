package timezone_test

import (
	"testing"

	"github.com/djangulo/clockwall/timezone"
	"github.com/pkg/errors"
)

func TestSystemTimezones(t *testing.T) {
	t.Run("gets a slice of timezones", func(t *testing.T) {
		timezones, err := timezone.SystemTimezones("/usr/share/zoneinfo/")
		if err != nil {
			t.Errorf("received error from Timezones: %v", err)
		}
		if len(timezones.Slice()) == 0 {
			t.Error("wanted len > 0 but got 0")
		}
	})
	t.Run("cannot read non-existent directory", func(t *testing.T) {
		_, err := timezone.SystemTimezones("/tmp/no-existo")
		AssertError(t, errors.Cause(err), timezone.ErrDirNoExist)
	})

}

func TestValidate(t *testing.T) {
	timezones, _ := timezone.SystemTimezones("/usr/share/zoneinfo")
	for _, test := range []struct {
		name, in string
		want     bool
	}{
		{"true", "US/Central", true},
		{"false", "not real", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := timezones.Validate(test.in)
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}


func AssertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Error("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
