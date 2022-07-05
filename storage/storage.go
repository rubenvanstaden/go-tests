package storage

import (
    "log"
	"context"
	"io/ioutil"
	"path/filepath"

	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
)

type Storage struct {
	*blob.Bucket
}

func Open(path string) (*Storage, error) {

	log.Printf("creating bucket in (%s)", path)

	// err := os.MkdirAll(path, 0777)
	// if err != nil {
	//         return nil, err
	// }

	bucket, err := fileblob.OpenBucket(path, nil)
	if err != nil {
        return nil, err
	}

	return &Storage{bucket}, nil
}

func (self *Storage) WriteData(filename, data string) error {
    err := self.WriteAll(context.Background(), filename, []byte(data), nil)
	if err != nil {
		return err
	}
    return nil
}

func (self *Storage) WriteTarball(tarfile string) error {

	log.Printf("pushing file (%s) to cloud bucket", tarfile)

	data, err := ioutil.ReadFile(tarfile)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	ctx := context.Background()
	filename := filepath.Base(tarfile)

	err = self.WriteAll(ctx, filename, data, nil)
	if err != nil {
		return err
	}

	return nil
}
