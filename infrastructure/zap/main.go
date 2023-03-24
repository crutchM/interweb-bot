package zap

import "go.uber.org/zap"

func New(isDebug bool) (*zap.SugaredLogger, error) {
	zapOpt := []zap.Option{
		zap.AddCallerSkip(2),
	}

	logger, err := zap.NewProduction(zapOpt...)
	if err != nil {
		return nil, err
	}

	if isDebug {
		logger, err = zap.NewDevelopment(zapOpt...)
		if err != nil {
			return nil, err
		}
	}

	return logger.Sugar(), nil
}
