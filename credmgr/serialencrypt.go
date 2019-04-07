package credmgr

import (
	"os"

	"encoding/gob"

	"github.com/BR3AKR/pwk/cryptor"
)

type SerializeEncryptionWriter struct {
	Data []byte
}

func (s SerializeEncryptionWriter) WriteToFile(filename string, password string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	hash, err := cryptor.CreateHash(password)
	if err != nil {
		return err
	}
	data, err := cryptor.Encrypt(s.Data, hash)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

func (s *SerializeEncryptionWriter) Write(p []byte) (n int, err error) {
	s.Data = append(s.Data, p...)
	return len(p), nil
}

type SerializeEncryptionReader struct {
	Data    []byte
	ReadIdx int
}

func (s *SerializeEncryptionReader) Read(p []byte) (n int, err error) {
	copy(p, s.Data[s.ReadIdx:s.ReadIdx+len(p)])
	s.ReadIdx += len(p)
	return len(p), err
}

func (s *SerializeEncryptionReader) ReadByte() (byte, error) {
	b := s.Data[s.ReadIdx]
	s.ReadIdx++
	return b, nil
}

func DeserializeData(filename, password string) ([]Credential, error) {
	hash, err := cryptor.CreateHash(password)
	if err != nil {
		return nil, err
	}
	data, err := cryptor.DecryptFile(filename, hash)
	if err != nil {
		return nil, err
	}
	sr := new(SerializeEncryptionReader)
	sr.Data = data
	decoder := gob.NewDecoder(sr)
	creds := []Credential{}
	err = decoder.Decode(&creds)
	return creds, err
}

func SerializeData(object interface{}, filename, password string) error {
	sr := new(SerializeEncryptionWriter)
	encoder := gob.NewEncoder(sr)
	encoder.Encode(object)
	sr.WriteToFile(filename, password)
	return nil
}
