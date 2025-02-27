// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package tests

// Injectors from wire.go:

func Init() (*Tests, func(), error) {
	logger := NewZapLogger()
	logLogger := NewLogger(logger)
	db, cleanup, err := NewDB(logger, logLogger)
	if err != nil {
		return nil, nil, err
	}
	redisClient, cleanup2, err := NewRDB(logLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tests := New(logger, logLogger, db, redisClient)
	return tests, func() {
		cleanup2()
		cleanup()
	}, nil
}
