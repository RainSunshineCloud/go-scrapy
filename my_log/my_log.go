package my_log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type loger struct {
	log log.Logger
	file *os.File
	time time.Time
	sync.Mutex
	dir string
}

var my_log  *loger = &loger{
};


func (this *loger) Error() *log.Logger {
	this.log.SetPrefix("[ERROR]");
	return &this.log;
}

func (this *loger) Info() *log.Logger {
	this.log.SetPrefix("[INFO]")
	return &this.log;
}

func init () {
	my_log.log.SetFlags(log.LstdFlags | log.Llongfile)
}

func (this *loger) initLog () {

	if this.file == nil || time.Now().Sub(this.time) >= time.Hour {
		this.Lock()
		defer this.Unlock()

		this.time = time.Now();
		file_path := fmt.Sprintf("%s%s.log",this.dir,this.time.Format("2006010215"))


		if _,err := os.Stat(file_path); err != nil {
			if os.IsNotExist(err) {
				os.Create(file_path)
			}
		}
		file_ptr,err := os.OpenFile(file_path,os.O_APPEND,0744)
		if err != nil {
			panic(err)
		}
		this.file.Close()
		this.log.SetOutput(file_ptr)
	}
}

func Get() *loger{
	my_log.initLog()
	return my_log;
}

func SetDir (dir_path string) {
	my_log.dir = strings.TrimRight(dir_path,"/") + "/";
}