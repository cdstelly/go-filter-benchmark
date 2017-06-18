package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
import "strings"
import "flag"

var (
	myfiles  []metafile
	numToGen int
)

func init() {
	flag.IntVar(&numToGen, "num", 1000000, "Number of files to generate")
	flag.Parse()
}

type metafile struct {
	Name       string
	CreateTime time.Time
	EditTime   time.Time
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randFilename() string {
	extensions := []string{"pdf", "exe", "png", "ini", "txt", "docx", "xls", "bmp", "wav", "mp3"}
	randIndex := rand.Intn(10)
	fn := RandStringBytes(12) + "." + extensions[randIndex]
	return fn
}

func randTime() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenRandomFile() metafile {
	var x metafile
	x.Name = randFilename()
	x.CreateTime = randTime()
	x.EditTime = randTime()
	return x
}

//Does anything in the set satisfy 'f'?
func Any(list []metafile, f func(metafile) bool) bool {
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

//Does everything in the set saisfty 'f'?
func All(list []metafile, f func(metafile) bool) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

//Return the items in te set that satisfy 'f'
func Filter(list []metafile, f func(metafile) bool) []metafile {
	tlist := make([]metafile, 100)
	for _, v := range list {
		if f(v) {
			tlist = append(tlist, v)
		}
	}
	return tlist
}

//Return new slice containing results of applying 'f' to each item in set
func Map(list []metafile, f func(metafile) metafile) []metafile {
	vsm := make([]metafile, len(list))
	for i, v := range list {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	for i := 0; i < numToGen; i++ {
		myfiles = append(myfiles, GenRandomFile())
		//fmt.Println(GenRandomFile())
	}
	fmt.Println("Generated " + strconv.Itoa(numToGen) + " files.")

	//Filter for pdfs
	start := time.Now()
	Filter(myfiles, func(v metafile) bool { return strings.HasSuffix(v.Name, "pdf") })
	elapsed := time.Since(start)
	fmt.Printf("Filter took\t%s, %f files/ms\n", elapsed, float64(numToGen)/float64(elapsed.Nanoseconds())*1000000)

	//Do any match xls?
	start = time.Now()
	Any(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "xls") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
	elapsed = time.Since(start)
	//fmt.Printf("Any took %s, %f files/ms\n", elapsed, float64(numToGen)/float64(elapsed.Nanoseconds())*1000000)

	//Are they all .txt?
	start = time.Now()
	All(myfiles, func(v metafile) bool {
		if strings.HasSuffix(v.Name, "txt") {
			//fmt.Println(v.Name)
			return true
		}
		return false
	})
	elapsed = time.Since(start)
	//fmt.Printf("All took %s, %f files/ms\n", elapsed, float64(numToGen)/float64(elapsed.Nanoseconds())*1000000)

	//Apply .bk to each filename
	start = time.Now()
	Map(myfiles, func(old metafile) metafile {
		new := old
		new.EditTime = time.Now()
		new.Name += ".bk"
		return new
	})
	elapsed = time.Since(start)
	fmt.Printf("Map took \t%s, %f files/ms\n", elapsed, float64(numToGen)/float64(elapsed.Nanoseconds())*1000000)

}
