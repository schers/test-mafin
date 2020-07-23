package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type OriginalFile struct {
	BaseMime string
	Filepath string
	Filename string
	Size     int64
}

func Upload(req *http.Request, storage string) (*OriginalFile, error) {
	meta, err := ParseMeta(req)
	if err != nil {
		return nil, err
	}

	up := &Uploader{Root: storage, Meta: meta, Body: req.Body}

	file, err := up.SaveFile()
	if err != nil {
		return nil, err
	}

	return file, nil
}

type Uploader struct {
	Root string
	Meta *Meta
	Body io.Reader
}

func (up *Uploader) SaveFile() (*OriginalFile, error) {

	tempFile, err := up.TempFile()
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	if err = up.Write(tempFile, up.Body); err != nil {
		return nil, err
	}

	fi, err := tempFile.Stat()
	if err != nil {
		return nil, err
	}

	ofile := &OriginalFile{Filename: up.Meta.Filename, Filepath: tempFile.Name(), Size: fi.Size()}

	ofile.BaseMime, err = IdentifyMime(ofile.Filepath)
	if err != nil {
		return nil, err
	}

	return ofile, nil
}

func (up *Uploader) TempFile() (*os.File, error) {
	return ioutil.TempFile(os.TempDir(), up.Root)
}

func (up *Uploader) Write(temp_file *os.File, body io.Reader) error {
	_, err := io.Copy(temp_file, body)
	if err != nil {
		return err
	}
	return nil
}

func IdentifyMime(file string) (string, error) {
	out, err := exec.Command("file", "--mime-type", file).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("identify: err: %s; detail: %s", err, string(out))
	}

	mime := strings.Split(strings.Split(string(out), ": ")[1], "/")[0]

	return mime, nil
}
