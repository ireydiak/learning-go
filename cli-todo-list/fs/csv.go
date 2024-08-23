package fs

import (
	"os"
)

func OpenOrCreate(fname string) (*os.File, error) {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		f, err := os.Create(fname)
		if err != nil {
			return nil, err
		}
		f.WriteString("id,title,created_at,status\n")
		f.Close()
	}

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
