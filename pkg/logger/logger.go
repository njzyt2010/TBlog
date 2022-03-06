package logger

import (
	"bytes"
	"fmt"
	_ "github.com/rs/zerolog/log"
	"log"
)

//var logPath string

func Debug(err interface{}) {
	var buf bytes.Buffer
	log.New(&buf,"[debug]" , log.Llongfile)

	log.Print(err)
	fmt.Println(&buf)
}

func Info(err interface{}) {
	var buf bytes.Buffer
	log.New(&buf,"[info]",log.Lshortfile)
	log.Print(err)
	fmt.Println(&buf)
}

func Warn(err interface{}) {
	var buf bytes.Buffer
	log.New(&buf,"[warn]",log.Llongfile)
	log.Print(err)
	fmt.Println(&buf)
}

func Panic(err interface{}) {
	var buf bytes.Buffer
	log.New(&buf,"[panic]" , log.Llongfile)
	log.Print(err)
	fmt.Println(&buf)
}