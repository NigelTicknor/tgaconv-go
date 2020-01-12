/**
* Author: Nigel Ticknor
* Usage: tgaconv "/path/to/dir"
*/

package main

import (
	"github.com/ftrvxmtrx/tga"
	"golang.org/x/image/bmp"
	"os"
	"bufio"
    "fmt"
	"path/filepath"
	"flag"
	"strings"
)

func main() {
	
	flag.Parse()
	
	if len(flag.Args())!=1  {
		fmt.Println("One and only one argument required: Directory. Got",len(flag.Args()))
		return
	}
	
	dirPath := flag.Args()[0];
	
	dirStat, err := os.Stat(dirPath)
    if err != nil {
        fmt.Println(err)
        return
    }
	if(dirStat.Mode().IsDir()!=true){
		fmt.Println("Error: Input is not a directory")
		return
	}
	
	err = filepath.Walk(dirPath,func (path string, info os.FileInfo, err error) error{
		if (strings.Contains(info.Name(),".tga")) {
			convert(path)
		}
		return nil
	})
	
	if err != nil {
		fmt.Println("Error walking through directory:",err)
	}
	
	
	fmt.Println("done")
	
}

func convert(filename string) {
	
	fin, err := os.Open(filename)

	if err != nil {
		fmt.Println("could not open input")
		return
	}

	img, err := tga.Decode(bufio.NewReader(fin))
	
	if err != nil {
		fmt.Println("could not decode image")
		return
	}
	
	fin.Close()
	
	// Later I might make this replacement more robust
	fout, err := os.Create(strings.Replace(filename,".tga",".bmp",1))
	
	if err != nil {
		fmt.Println("could not create new image file")
		return
	}
	
	err = bmp.Encode(fout,img)
	
	if err != nil {
		fmt.Println("could not write new image file")
		return
	}
	fout.Close()
	
	
	fmt.Println("Converted:",filename)
	
	
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("could not delete old image file",err)
		return
	}
	
	
	
}