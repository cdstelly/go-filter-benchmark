package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	b.ResetTimer()

	Filter(myfiles, func(v metafile) bool {
		if strings.Contains(v.Name, "pdf") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	b.ResetTimer()

	Any(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "xls") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	b.ResetTimer()

	All(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "txt") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	b.ResetTimer()

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
