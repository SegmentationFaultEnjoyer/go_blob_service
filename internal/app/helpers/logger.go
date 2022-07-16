package helpers

import "github.com/sirupsen/logrus"

func ConfigureLogger(LogLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(level)

	return logger, nil
}
