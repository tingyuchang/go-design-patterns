package singleton

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type InfoLogger struct {
	once sync.Once
	logger *log.Logger
}

func (l *InfoLogger) Pintf(format string, v ...interface{})  {
	l.lazyInit()
	l.logger.Printf(fmt.Sprintf(format, v...))
}

func (l *InfoLogger) lazyInit()  {
	l.once.Do(func() {
		l.logger = log.New(os.Stdout, "Info: ", log.Lshortfile)
		l.logger.Println("init")
	})
}