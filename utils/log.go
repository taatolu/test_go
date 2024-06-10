package utils

import (
	"log"
	"os"
	"io"
)

func LoggingSettings(cfgLogfile string){
	Logfile, err := os.OpenFile(cfgLogfile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0766)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogfile := io.MultiWriter(Logfile,os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogfile)
}