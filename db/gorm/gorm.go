package gorm

import (
	"fmt"
	"log"
	"time"

	"rest_echo/bootstrap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

// initialize database
func Init() {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")
	log.Printf("this is your POSTGRE TABLRE: %#v", adapter)
	switch adapter {
	case "mysql":
		mysqlConn()
		break
	case "postgre":
		postgresConn()
		break
	default:
		panic("Undefined connection on config.yaml")
	}
}

// setupPostgresConn: setup postgres database connection using the configuration from database.yaml
func postgresConn() {
	var (
		connectionString string
		err              error
	)
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"), bootstrap.App.DBConfig.String("sslmode"))
	if db, err = gorm.Open("postgres", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec(fmt.Sprintf("SET TIMEZONE TO '%s'", bootstrap.App.AppConfig.String("timezone")))
	db.Exec("CREATE EXTENSION postgis")
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(bootstrap.App.DBConfig.Int("maxlifetime")))
	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

// mysqlConn: setup mysql database connection using the configuration from database.yaml
func mysqlConn() {
	var (
		connectionString string
		err              error
	)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"))

	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec(fmt.Sprintf("SET TIMEZONE = '%s'", bootstrap.App.AppConfig.String("timezone")))
	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

/*
 * Gorm: return GORM's postgres database connection instance.
 */
func DBManager() *gorm.DB {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")
	switch adapter {
	case "mysql":
		mysqlConn()
		break
	case "postgre":
		postgresConn()
		break
	default:
		panic("Undefined connection on config.yaml")
	}

	return db
}
