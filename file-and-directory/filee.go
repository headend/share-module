package file_and_directory

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

func GetMd5FromFile(path string) (md5string string, err error)  {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return "", err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func GetFileSizeInByte(path string) (fileSize int64, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	// get the size
	size := fi.Size()
	return size, nil
}


type MyFile struct {
	Path	string
}

func (mf *MyFile) Read() (string, error) {
	data, err := ioutil.ReadFile(mf.Path)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// Exists reports whether the named file or directory exists.
func (mf *MyFile) Exists() bool {
	if _, err := os.Stat(mf.Path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (mf *MyFile) GetMd5FromFile(path string) (md5string string, err error)  {
	filePath := ""
	if path == "" {
		filePath = mf.Path
	} else {
		filePath = path
	}
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return "", err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func (mf *MyFile) GetFileSizeInByte(path string) (fileSize int64, err error) {
	filePath := ""
	if path == "" {
		filePath = mf.Path
	} else {
		filePath = path
	}
	fi, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	// get the size
	size := fi.Size()
	return size, nil
}

func (mf *MyFile) WriteString(StringToWrite string) (err error) {
	var returnErr error
	f, err := os.Create(mf.Path)

	if returnErr != nil {
		return returnErr
	}
	defer f.Close()

	_, returnErr = f.WriteString(StringToWrite)

	if returnErr != nil {
		return returnErr
	}
	return returnErr
}

func (mf *MyFile) WriteByte(StringToWriteByte string) (err error) {
	var returnErr error
	f, err := os.Create(mf.Path)

	if returnErr != nil {
		return returnErr
	}
	defer f.Close()

	data := []byte(StringToWriteByte)

	_, returnErr = f.Write(data)

	if returnErr != nil {
		return returnErr
	}
	return returnErr
}