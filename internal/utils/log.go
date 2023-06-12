package utils

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"reflect"
	"runtime"
)

func Log(message string, err error) {
	_, file, line, _ := runtime.Caller(1)

	var fields = logrus.Fields{
		"type":     filepath.Base(file),
		"function": runtime.FuncForPC(reflect.ValueOf(Log).Pointer()).Name(),
		"line":     line,
	}

	if err != nil {
		fields["err"] = err
		logrus.WithFields(fields).Error(message)
	} else {
		logrus.WithFields(fields).Debug(message)
	}

}
