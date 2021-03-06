package main

import (
	"context"
	"database/sql"

	usecases "admin-skm/src/app/use_cases"

	"admin-skm/src/infra/config"
	postgres "admin-skm/src/infra/persistence/postgress"

	"admin-skm/src/interface/rest"

	ms_log "admin-skm/src/infra/log"

	jawabanUc "admin-skm/src/app/use_cases/jawaban"
	layananUc "admin-skm/src/app/use_cases/layanan"
	loginUc "admin-skm/src/app/use_cases/login"
	pertanyaanUc "admin-skm/src/app/use_cases/pertanyaan"
	userUc "admin-skm/src/app/use_cases/user"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb := postgres.New(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.SqlDB, postgresdb.DB.Name())

	userRepository := postgres.NewUsersRepository(postgresdb.DB)
	pertanyaanRepository := postgres.NewPertanyaanRepository(postgresdb.DB)
	layananRepository := postgres.NewLayananRepository(postgresdb.DB)
	jawabanRepository := postgres.NewJawabanRepository(postgresdb.DB)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{

			UserUseCase:       userUc.NewUserUseCase(userRepository),
			LoginUseCase:      loginUc.NewUserUseCase(userRepository),
			PertanyaanUseCase: pertanyaanUc.NewPertanyaanUseCase(pertanyaanRepository),
			LayananUseCase:    layananUc.NewLayananUseCase(layananRepository),
			JawabanUseCase:    jawabanUc.NewJawabanUseCase(jawabanRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
