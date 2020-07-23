package upload

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type DirManager struct {
	Root string
	Path string
}

func CreateDir(root, mime string) (*DirManager, error) {
	dm := &DirManager{Root: root}

	dm.CalcPath(mime)
	if err := dm.Create(); err != nil {
		return nil, err
	}

	return dm, nil
}

func (dm *DirManager) Abs() string {
	return filepath.Join(dm.Root, dm.Path)
}

func (dm *DirManager) Create() error {
	return os.MkdirAll(dm.Root+dm.Path, 0755)
}

func (dm *DirManager) CalcPath(mime string) {
	date := time.Now()
	dm.Path = fmt.Sprintf("/%s/%d/%s/%s", mime, date.Year(), yearDay(date), containerName(date))
}

func yearDay(t time.Time) string {
	return strconv.FormatInt(int64(t.YearDay()), 36)
}

func containerName(t time.Time) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
	seconds := t.Hour()*3600 + t.Minute()*60 + t.Second()

	return strconv.FormatInt(int64(seconds*1000+r), 36)
}
