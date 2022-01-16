package gecko

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// return the MD5 hash of a file
func GetMD5(file *os.File) string {
	hash := md5.New()
	io.Copy(hash, file)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// return the SHA1 hash of a file
func GetSHA1(file *os.File) string {
	hash := sha1.New()
	io.Copy(hash, file)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// return the SHA256 hash of a file
func GetSHA256(file *os.File) string {
	hash := sha256.New()
	io.Copy(hash, file)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// return the base64 encoding of a file content
func GetB64(file *os.File) string {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return base64.StdEncoding.EncodeToString(data)
}

// return the base64 decoding of a file content
func GetD64(file *os.File) string {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	content, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return string(content)
}
