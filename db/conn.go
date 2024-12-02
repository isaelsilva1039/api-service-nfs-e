package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configurações de conexão PostgreSQL
var (
	host     = mustGetEnv("DB_INSTANCE_CONNECTION_NAME")
	user     = mustGetEnv("DB_USER")
	password = mustGetEnv("DB_PASS")
	dbname   = mustGetEnv("DB_NAME")
)

// ConnectDB cria uma conexão com o banco de dados PostgreSQL usando GORM e retorna um ponteiro para *gorm.DB
func ConnectDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%v user=%s password=%s database=%s sslmode=disable", host, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: "cloudsqlpostgres", DSN: dsn}))
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err == nil {
		fmt.Println("Conectado ao banco de dados PostgreSQL!")
		return db, nil
	}

	return db, nil
}

// Caso necessário, refatorar o código para invocar a função adequada com base em alguma váriavel de ambiente. Ex.:
// func ConnectToLocalDB() (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s", lc.user, lc.pwd, lc.host, lc.port, lc.dbname)
// 	dbPool, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, fmt.Errorf("gorm.Open: %w", err)
// 	}

// 	sqlDB, _ := dbPool.DB()
// 	err = sqlDB.Ping()
// 	if err == nil {
// 		fmt.Println("Conectado ao banco de dados PostgreSQL!")
// 		return dbPool, nil
// 	}

// 	return dbPool, nil
// }

func mustGetEnv(envName string) string {
	envValue := os.Getenv(envName)
	if envName == "" {
		log.Fatalf("não foi possível recuperar env %v", envName)
	}

	return envValue
}
