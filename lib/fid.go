package fid

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// FileIdentifer records Sha1, Sha256, Md5
type FileIdentifer struct {
	Path   string
	Sha1   string
	Sha256 string
	Md5    string
}

// NewFileIdentifer from path
func NewFileIdentifer(p string) (fid *FileIdentifer, err error) {

	sha1er := sha1.New()
	sha256er := sha256.New()
	md5er := md5.New()

	w := io.MultiWriter(sha1er, sha256er, md5er)

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		return nil, err
	}

	abs, err := filepath.Abs(p)
	if err != nil {
		return nil, err
	}

	return &FileIdentifer{
		Path:   abs,
		Sha1:   hex.EncodeToString(sha1er.Sum(nil)),
		Sha256: hex.EncodeToString(sha256er.Sum(nil)),
		Md5:    hex.EncodeToString(md5er.Sum(nil)),
	}, nil
}
