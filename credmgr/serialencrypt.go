package credmgr

import (
	"os"

	"encoding/gob"

	"github.com/BR3AKR/pwk/cryptor"
)

type SerializeEncryptionWriter struct {
	Data []byte
}

func (s SerializeEncryptionWriter) WriteToFile(filename string, password string) {
	f, _ := os.Create(filename)
	defer f.Close()
	hash, _ := cryptor.CreateHash(password)
	f.Write(cryptor.Encrypt(s.Data, hash))
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
	data, err := cryptor.DecryptFile(filename, hash)
	// TODO Obviously need much better error handling here
	// Honestly, the whole thing needs revisited for better error handling
	// I was just trying to get this thing working first :)
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
