// Package mrlogger is a customized for error handling and login with log levels:
// (4: "DEBUG", 3: "INFO", 2: "WARN", 1: "ERROR")
package mrlogger

import (
	"errors"
	"fmt"
	"log"
	"mrlogger/logtools"
	"os"
	"runtime"
	"strings"
)

// Check maneja errores de forma básica: Recibe elemento "err" y
// cualquier secuencia de strings que se ingresen como argumento
func Check(e error, s ...string) {
	if e != nil {
		log.Panicf("[PANIC] %v: %v", s, e)
		//log.("[FATAL] %v: %v", s, e)
	}
}

// InThisPoint retorna un string que contiene el nombre del archivo de origen, nombre de la función
// y el numero del linea especificada en el call stack
func InThisPoint(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s:%s:%d", chopPath(file), runtime.FuncForPC(function).Name(), line)
}

// chopPath retorna el nombre del archivo de origen después del ultimo slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]
}

func NewLoggingLevel(l string) {
	lv, err := checkLevel(l)
	Check(err, "Logging input:", l, InThisPoint())

	filter := &logtools.LevelFilter{
		Levels:   []logtools.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logtools.LogLevel(lv),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

func Debug(v ...interface{}) {
	log.Println("[DEBUG]", v)
}

func Warn(v ...interface{}) {
	log.Println("[WARN ]", v)
}

func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}

func Info(v ...interface{}) {
	log.Println("[INFO ]", v)
}

func checkLevel(l string) (string, error) {

	ls := []string{
		4: "DEBUG",
		3: "INFO",
		2: "WARN",
		1: "ERROR",
	}

	valid := false
	l = strings.Trim(l, " ")

	for _, v := range ls {
		if strings.EqualFold(l, v) {
			l = v
			valid = true
		}
	}

	if !valid {
		return l, errors.New("Logging level is not: DEBUG, INFO, WARN, ERROR")
	} else {
		return l, nil
	}
}
