package main

import (
	"log"
	"os"
	"runtime"
)
import "fmt"
import "bytes"
import "io"

type user struct {
	Name string
	Age  int
}

func InitLogger() {
	log.SetPrefix("from-baselog ")
	log.SetFlags(log.Ldate | log.Llongfile)
}

func CustomLogger(w io.Writer) *log.Logger {
	logger := log.New(w, "", log.LstdFlags)
	return logger
}

func useCustomLogger(u user) {
	buf := &bytes.Buffer{}
	logger := CustomLogger(buf)
	logger.Printf("%s login, age %d", u.Name, u.Age)
	fmt.Println(buf.String())
}

func useMultiLogger(u user) {
	w1 := &bytes.Buffer{}
	w2 := os.Stdout
	w3, err := os.OpenFile("base-log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	mutliWriter := io.MultiWriter(w1, w2, w3)
	logger := CustomLogger(mutliWriter)
	logger.Printf("multiwriter: %s login, age %d", u.Name, u.Age)
	fmt.Println(w1.String())
}

func main() {
	defer func() {
		fmt.Println("main exits")
	}()
	InitLogger()
	u := user{
		Name: "roby",
		Age:  40,
	}
	useCustomLogger(u)
	useMultiLogger(u)
	pc, fileName, lineNumber, result := runtime.Caller(0)
	fmt.Printf("PC: %v filename: %s linenumber: %d result: %v\n", pc, fileName, lineNumber, result)
	if wd, err := os.Getwd(); err != nil {
		log.Printf("err get current working directory %v", err)
	} else {
		log.Printf("current program runs in %s", wd)
	}
	if ex, err := os.Executable(); err != nil {
		log.Printf("err get executable %v", err)
	} else {
		log.Printf("current executable is %s", ex)
	}
	log.Printf("%s login, age %d", u.Name, u.Age)
	log.Panicf("System error when %s login", u.Name)
	// Fatalf won't run defer func anymore
	//log.Fatalf("Danger, hacker %s login", u.Name)
}
