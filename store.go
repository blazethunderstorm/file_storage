package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


const defaultRootFolderName="ggnetwork"

func CASPathTransformFunc(key string) PathKey{
	hash:=sha1.sum([]byte(key))
	hashStr:=hex.EncodeToString(hash[:])

	blocksize=5
	sliceLen= len(hashStr) / blocksize
	paths:=make([]string, sliceLen)
	for i=0;i<sliceLen;i++{
		from,to=i*blocksize,i*blocksize+blocksize
		paths[i]=hashStr[from:to]
	}

	return PathKey{
		PathName:strings.Join(paths, "/"),
		FileName:hashStr,
	}
}

type PathTransformFunc func(key string) PathKey

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FirstPathName() string{
	paths:=strings.split(p.PathName,"/")
	if len(paths) ==0 {
		return ""
	}
	return paths[0]
}

func (p PathKey) FullPathName() string {
	
	return fmt.Sprintf("%s/%s", p.PathName, p.FileName)	
}

type StoreOpts struct {
	 
	Root              string
	PathTransformFunc PathTransformFunc
}

var defaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName:key,
		FileName:key,
	}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc==nil{
		opts.PathTransformFunc = defaultPathTransformFunc
	}
	if len(opts.Root)==0{
		opts.Root = defaultRootFolderName
	}

	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(id String , key String) bool{
	pathKey=s.PathTransformFunc(key)
	fullPathNameWithRoot:=fmt.Sprintf("%s/%s", s.Root, pathKey.FullPathName())
	_, err := os.Stat(fullPathNameWithRoot)
	return !errors.Is(err, os.ErrNotExist)
}

func (s *Store) Clear() error{
	return os.RemoveAll(s.Root)
}

func (s *Store) Delete (id String , key String) error{
	pathKey:=s.PathTransformFunc(key)

	defer func(){
		log.Printf("deleted [%s] from disc", pathKey.FileName)
	}()

	firstPathNameWithRoot:= fmt.Sprintf("%s/%s", s.Root, pathKey.FirstPathName())

	return os.RemoveAll(firstPathNameWithRoot)
}

func (s *Store) Write (id String, key String, r io.Reader) (int64 ,error){
    returnn s.WriteStream(id, key, r)
}

func (s *Store) WriteDecrypt(encKey []byte,id String,key String,r io.Reader)(int64 ,error){
	f,err:=s.opeenFileForWriting(id,key)

	if err != nil {
		return 0, err
	}

	n,err:=copyDecrypt(encKey, r, f)

	return int64(n), err
}


func (s *Store) openFileForWriting(id string, key string) (*os.File, error) {
	pathKey := s.PathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s/%s", s.Root, id, pathKey.PathName)
	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return nil, err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s/%s", s.Root, id, pathKey.FullPath())

	return os.Create(fullPathWithRoot)
}

func (s *Store) writeStream(id string, key string, r io.Reader) (int64, error) {
	f, err := s.openFileForWriting(id, key)
	if err != nil {
		return 0, err
	}
	return io.Copy(f, r)
}

func (s *Store) Read(id string, key string) (int64, io.Reader, error) {
	return s.readStream(id, key)
}

func (s *Store) readStream(id string, key string) (int64, io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s/%s", s.Root, id, pathKey.FullPath())

	file, err := os.Open(fullPathWithRoot)
	if err != nil {
		return 0, nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return 0, nil, err
	}

	return fi.Size(), file, nil
}