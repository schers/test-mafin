package upload

import (
	"github.com/schers/test-mafin/db/dto"
	"os"
	"path/filepath"
)

type File struct {
	Id           uint64
	OriginalFile *OriginalFile
	DirManager   *DirManager
	Filename     string
	Size         int64
}

func CreateFile(storage string, ofile *OriginalFile) (*File, error) {
	dm, err := CreateDir(storage, ofile.BaseMime)
	if err != nil {
		return nil, err
	}

	file := &File{
		OriginalFile: ofile,
		DirManager:   dm,
		Filename:     ofile.Filename,
	}

	if err := file.Move(ofile.Filepath); err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) Move(src string) error {
	err := os.Rename(src, f.Filepath())
	if err != nil {
		return err
	}

	file, err := os.Open(f.Filepath())
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	f.Size = fi.Size()

	return nil
}

func (f *File) Filepath() string {
	return filepath.Join(f.DirManager.Abs(), f.Filename)
}

func (f *File) Url() string {
	return filepath.Join(f.DirManager.Path, f.Filename)
}

func (f *File) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":       f.Id,
		"url":      f.Url(),
		"filename": f.Filename,
		"size":     f.Size,
	}
}

func (f *File) SaveInfo() error {
	fileInfo := dto.File{
		Name: f.Filename,
		Size: f.Size,
		Path: f.Url(),
	}

	err := fileInfo.Save()
	if err != nil {
		return err
	}
	f.Id = fileInfo.Id

	return nil
}
