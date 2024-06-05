package logFileHandle

import (
		"log"
		"os"
)

func WriteLogFiles(p string) (*os.File, *os.File){
		//open log files
		f1, err1 := os.OpenFile(p+"info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err1 != nil {
				log.Fatal(err1)
		}
		//defer f1.Close()
		f1.Write([]byte("\n\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<INFO LOGGING STARTED>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n"))

		f2, err2 := os.OpenFile(p+"error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err2 != nil {
				log.Fatal(err2)
		}
		//defer f2.Close()
		f2.Write([]byte("\n\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<ERROR LOGGING STARTED>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n"))

		return f1,f2
}
