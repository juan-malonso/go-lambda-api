package log

import (
  "os"
  "github.com/sirupsen/logrus"
)

func InitLogger(name string) *logrus.Logger {
  log := logrus.New()

  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&logrus.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  // log.SetLevel(logrus.InfoLevel)

  return log
}

type Fields = logrus.Fields
