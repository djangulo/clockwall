package timezone

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrDirNoExist directory does not exist
	ErrDirNoExist = errors.New("directory does not exist")
	// ErrDirIsEmpty dir is empty
	ErrDirIsEmpty = errors.New("dir is empty")
	// ErrPermissionDenied permission denied.
	ErrPermissionDenied = errors.New("permission denied")
	// IgnoredDirs sub-directories that are ignored when traversing system files.
	IgnoredDirs = map[string]struct{}{
		"right": struct{}{},
		"posix": struct{}{},
	}
)

// Timezones struct to hold a set of timezones.
type Timezones struct {
	Set map[string]struct{}
}

// Slice returns a slice of strings with the timezones in the set.
func (t *Timezones) Slice() (timezones []string) {
	for tz := range (*t).Set {
		timezones = append(timezones, tz)
	}
	return
}

// Validate validates that tz exists.
func (t *Timezones) Validate(tz string) bool {
	if _, ok := (*t).Set[tz]; !ok {
		return false
	}
	return true
}

// New returns a *Timezones object.
func New() *Timezones {
	return &Timezones{
		Set: map[string]struct{}{},
	}
}

// SystemTimezones traverses dir and collects the names in
func SystemTimezones(dir string) (*Timezones, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		msg := fmt.Sprintf("%q does not exist", dir)
		return nil, errors.Wrap(ErrDirNoExist, msg)
	}

	timezones := New()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Fprintf(os.Stderr, "error walking %v", err)
			return err
		}

		if _, ok := IgnoredDirs[info.Name()]; ok {
			return filepath.SkipDir
		}

		if info.IsDir() && info.Name() == dir {
			return nil
		}

		clean := strings.ReplaceAll(path, filepath.Clean(dir)+string(filepath.Separator), "")
		if filepath.Separator == '\\' {
			clean = strings.ReplaceAll(clean, "\\", "/")
		}
		timezones.Set[clean] = struct{}{}
		return nil
	})
	if err != nil {
		msg := fmt.Sprintf("failed to walk %q", dir)
		return nil, errors.Wrap(err, msg)
	}

	return timezones, nil
}
