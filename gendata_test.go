package main

import (
	"testing"
	"time"
)
import "strings"
import "fmt"
import "strconv"
import "reflect"

var (
	NumToGen int = 10000000
)

func init() {

	for i := 0; i < NumToGen; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	fmt.Println("generated " + strconv.Itoa(len(myfiles)) + " files")
}

func BenchmarkFind(b *testing.B) {
	Filter(myfiles, func(v metafile) bool {
		if strings.Contains(v.Name, "pdf") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkAny(b *testing.B) {
	Any(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "xls") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkAll(b *testing.B) {
	All(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "txt") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkMap(b *testing.B) {
	Map(myfiles, func(old metafile) metafile {
		new := old
		new.EditTime = time.Now()
		new.Name += ".bk"
		return new
	})
}

func TestGenRandomFile(t *testing.T) {
	var x metafile
	x.Name = randFilename()
	x.CreateTime = randTime()
	x.EditTime = randTime()

	var y metafile
	y.Name = randFilename()
	y.CreateTime = randTime()
	y.EditTime = randTime()

	if reflect.DeepEqual(x, y) {
		t.Fail()
	}
}
