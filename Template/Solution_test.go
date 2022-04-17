package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func Test_Main(t *testing.T) {
	var tests = []struct {
		file string
	}{
		{"Test.txt"},
		{"Test2.txt"},
		// {"Test3.txt"},
		// {"ts2_input.txt"},
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
	// a := []int{1, 2, 3}
	// for i = 0; i < 30; i++ {
	// 	fmt.Printf("Testing i:%v v:%v \n", i, 1<<i)
	// }
	// fmt.Printf("Testing i:%v v:%v \n", "a", arrayToString(a))

}

func Test_Interactive(t *testing.T) {
	var tests = []struct {
		file string
	}{
		{"Test_Client.py"},
	}
	for tc, tt := range tests {
		fmt.Println("---------------------------------")
		fmt.Printf("Test %d, Test Client:%v\n", tc+1, tt.file)
		fmt.Println("---------------------------------")
		// Define external program
		client := exec.Command("python3", tt.file)
		// Define log file
		logfile, err := os.Create(tt.file + ".log")
		if err != nil {
			panic(err)
		}
		defer logfile.Close()
		out := os.Stdout
		in := os.Stdin
		defer func() { os.Stdout, os.Stdin = out, in }() // Restore original Stdout,Stdin

		// Create pipe connect os.Stdout to client.Stdin
		gr, gw, _ := os.Pipe()

		// Connect os.Stdout to writer side of pipe
		os.Stdout = gw

		// Create MultiWriter to write to logfile and os.Stdout at the same time
		gmw := io.MultiWriter(out, logfile)

		// Create a tee reader read from reader side of the pipe and flow to the MultiWriter
		// Repleace the cmd.Stdin with TeeReader

		client.Stdin = io.TeeReader(gr, gmw)

		// Create a pipe to connect client.Stdout to os.Stdin
		cr, cw, _ := os.Pipe()

		// Create MultWriter to client stdout
		cmw := io.MultiWriter(cw, logfile, out)
		client.Stdout = cmw

		// Connect os stdin to another end of the pipe
		os.Stdin = cr

		// Start Client
		client.Start()
		// Start main
		main()
		// os.Stdout = out
		// Check Testing program error
		if err := client.Process.Release(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					log.Printf("Exit Status: %d", status.ExitStatus())
					t.Errorf("Tester return error\n")
				}
			} else {
				log.Fatalf("cmd.Wait: %v", err)
				t.Errorf("Tester return error\n")
			}
		}
	}
}
