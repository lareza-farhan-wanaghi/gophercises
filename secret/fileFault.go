package secret

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type fileFault struct {
	encodingKey string
	filePath    string
}

func (f *fileFault) GetAll() (map[string]string, error) {
	file, err := os.OpenFile(f.filePath, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	pairMap := map[string]string{}
	if fi.Size() <= 0 {
		return pairMap, nil
	}

	reader, err := f.encriptedReader(file)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buffer.Bytes(), &pairMap)
	if err != nil {
		return nil, err
	}

	return pairMap, nil
}

func (f *fileFault) Get(mapKey string) (string, error) {
	pairMap, err := f.GetAll()
	if err != nil {
		return "", err
	}

	return pairMap[mapKey], nil
}

func (f *fileFault) Set(mapKey, mapValue string) error {
	pairMap, err := f.GetAll()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	pairMap[mapKey] = mapValue
	marshaledMap, err := json.Marshal(pairMap)
	if err != nil {
		return err
	}

	writer, err := f.encriptedWriter(file)
	if err != nil {
		return err
	}

	_, err = writer.Write(marshaledMap)
	if err != nil {
		return err
	}

	return nil
}

func (f *fileFault) encriptedWriter(w io.Writer) (*cipher.StreamWriter, error) {
	key, err := hex.DecodeString(f.encodingKey)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	n, err := w.Write(iv)
	if err != nil || n != len(iv) {
		return nil, errors.New("failed to write the initial value")
	}

	stream := cipher.NewOFB(block, iv)
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func (f *fileFault) encriptedReader(r io.Reader) (*cipher.StreamReader, error) {
	key, err := hex.DecodeString(f.encodingKey)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if err != nil || n != len(iv) {
		return nil, errors.New("failed to read the initial value")
	}

	stream := cipher.NewOFB(block, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func (f *fileFault) empty() error {
	file, err := os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func FileFault(encodingKey, filePath string) *fileFault {
	result := &fileFault{
		encodingKey: encodingKey,
		filePath:    filePath,
	}
	return result
}
