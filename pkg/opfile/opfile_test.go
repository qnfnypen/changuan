package opfile

import (
	"os"
	"testing"
)

func TestCreateFileWithTimeStamp(t *testing.T) {
	fp := "log/test/test.txt"
	f, err := CreateFileWithTimeStamp(fp)
	if err != nil {
		t.Fatal(err)
	}
	expFp := f.Name()
	f.Close()

	_, err = os.Stat(expFp)
	if os.IsNotExist(err) {
		t.Errorf("文件创建失败,error: %v", err)
	}

	t.Log(expFp)
	os.RemoveAll("log")
}
