package filesystem

import (
	"fmt"
	"testing"
)

func TestPathIsExist(t *testing.T) {
	expected := false
	if PathIsExist("/Users/xujiajun123") != expected {
		t.Errorf("returned unexpected bool value : got %v want %v", !expected, expected)
	}
}

func TestPathIsExist2(t *testing.T) {
	fmt.Println(PathIsExist(";;;"))
}

func TestCopyDir(t *testing.T)  {
	src:="/Users/xujiajun/go/src/github.com/xujiajun/utils/testdata/test"
	dst:="/Users/xujiajun/go/src/github.com/xujiajun/utils/testdata/test2"
	err:=CopyDir(src,dst)

	fmt.Println("err",err)
}

func TestCopyFile(t *testing.T) {
	src:="/Users/xujiajun/go/src/github.com/xujiajun/utils/testdata/test/xx3.log"
	dst:="/Users/xujiajun/go/src/github.com/xujiajun/utils/testdata/test2/xx3.log"
	err:=CopyFile(src,dst)
	fmt.Println("err",err)
}
