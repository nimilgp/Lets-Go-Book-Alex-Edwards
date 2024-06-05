package main

import (
		"flag"
		"log"
		"net/http"
		"github.com/nimilgp/snippet-box/logFileHandle"
)

type config struct {
		port string
		staticDir string
		logsDir string
}

type application struct {
		errorLog *log.Logger
		infoLog *log.Logger
}

func main() {
		var cfg config
		flag.StringVar(&cfg.port, "port", ":3333", "HTTP port address")//flag, default val, helper
		flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
		flag.StringVar(&cfg.logsDir, "logs-dir", "./", "Path where to log info and error")
		flag.Parse()//call this before use of the flag variables else will stay at default

		f1, f2 := logFileHandle.WriteLogFiles(cfg.logsDir)

		infoLog := log.New(f1, "INFO\t", log.Ldate|log.Ltime)
		errorLog := log.New(f2, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
		
		//create instance of application
		app := &application{
				errorLog: errorLog,
				infoLog: infoLog,
		}

		infoLog.Println("Server Port given is ", cfg.port)
		infoLog.Println("Static directory given is ", cfg.staticDir)
		infoLog.Println("Logs directory given is ", cfg.logsDir)
		
		//create custom http.server
		srv := &http.Server {
				Addr: cfg.port,
				Handler: app.routes(cfg.staticDir),
				ErrorLog: errorLog,
		}

		//start a new web server at a port, handled by a serveMux
		err := srv.ListenAndServe()
		errorLog.Fatal(err)
}
