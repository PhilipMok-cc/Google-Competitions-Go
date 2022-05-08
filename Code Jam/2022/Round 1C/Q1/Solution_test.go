package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	var tests = []struct {
		file string
	}{
		{"Test.txt"},
		{"Test2.txt"},
		// {"Test3.txt"},
		// {"ts3_input.txt"},
	}
	for tc, tt := range tests {
		fmt.Println("---------------------------------")
		fmt.Printf("Test %d, Input file:%v\n", tc+1, tt.file)
		fmt.Println("---------------------------------")

		Test_Main_perfile := func(t *testing.T) {
			tmpfile, _ := os.Open(tt.file)
			out := os.Stdout
			in := os.Stdin
			defer func() { os.Stdout, os.Stdin = out, in }() // Restore original Stdout,Stdin
			os.Stdin = tmpfile
			outfile, err := os.Create(tt.file + ".out")
			if err != nil {
				panic(err)
			}
			r, w, _ := os.Pipe()
			mw := io.MultiWriter(out, outfile)
			tr := io.TeeReader(r, mw)
			os.Stdout = w
			defer outfile.Close()
			main()
			// Close the pipe at the end of test
			w.Close()
			// Flush the pipe to trigger tee reader
			ioutil.ReadAll(tr)
			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}
		}
		t.Run(tt.file, Test_Main_perfile)
	}
}

func Test_PlayGround(t *testing.T) {
	// var i uint
	// var a int = 1e9

	fmt.Printf("Testing i:%v\n", checkValid("HASH"))
	fmt.Printf("Testing i:%v\n", checkValid("CATTAX"))

}
