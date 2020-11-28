package file_and_directory


import (
	"crypto/md5"
	"encoding/hex"
	"io"
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
