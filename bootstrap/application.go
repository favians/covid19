package bootstrap

import (
	"fmt"

	"log"
	"strings"
	"time"

	logger "github.com/favians/golang_starter/modules/logger"
	hedwig "github.com/favians/golang_starter/modules/notification"
	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	App *Application
	db  *gorm.DB
	err error
)

type Config viper.Viper

type Application struct {
	Name      string        `json:"name"`
	Version   string        `json:"version"`
	ENV       string        `json:"env"`
	AppConfig Config        `json:"application_config"`
	DBConfig  Config        `json:"database_config"`
	DB        *gorm.DB      `json:"db"`
	Path      string        `json:"path"`
	Log       *logger.Log   `json:"log"`
	Hedwig    hedwig.Hedwig `json:"hedwig"`
}

func init() {
	App = &Application{}
	App.loadAppConfig()
	App.loadDBConfig()
	App.loadENV()

	App.Path = "./"
	if App.ENV == "test" {
		App.Path = "./.."
	}

	App.Name = App.AppConfig.String("app_name")
	App.Version = App.AppConfig.String("app_version")

	App.DB = App.DBInit()
	App.DB.LogMode(false) // set query log = OFF

	App.Log, err = logger.NewLogger(fmt.Sprintf("%s/storage/logs", App.Path), App.Name, "info")
	if err != nil {
		log.Println(err)
	}

	// notificationParams := App.AppConfig.GetStringMapString("notification")
	// App.Hedwig, err = notification.Init(notificationParams["client_id"], notificationParams["client_secret"], notificationParams["email_from"], notificationParams["host"], notificationParams["version"], App.AppConfig.GetSub("notification").GetBool("test_mode"))
	// if err != nil {
	// 	log.Println("email notification error : %s", err)
	// 	App.Log.Logger.WithFields(logrus.Fields{
	// 		"hedwig response": App.Hedwig,
	// 		"error":           err,
	// 	}).Error("email notification error")
	// }

	log.Printf("app path : %s", App.Path)
}

// loadAppConfig: read application config and build viper object
func (app *Application) loadAppConfig() {
	var (
		appConfig *viper.Viper
		err       error
	)
	appConfig = viper.New()
	appConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	appConfig.SetEnvPrefix("APP_")
	appConfig.AutomaticEnv()
	appConfig.SetConfigName("config")
	appConfig.AddConfigPath(".")
	appConfig.SetConfigType("yaml")
	if err = appConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	appConfig.WatchConfig()
	appConfig.OnConfigChange(func(e fsnotify.Event) {
		log.Println("App Config file changed %s:", e.Name)
	})
	app.AppConfig = Config(*appConfig)
}

// loadDBConfig: read application config and build viper object
func (app *Application) loadDBConfig() {
	var (
		dbConfig *viper.Viper
		err      error
	)
	dbConfig = viper.New()
	dbConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	dbConfig.SetEnvPrefix("DB_")
	dbConfig.AutomaticEnv()
	dbConfig.SetConfigName("config")
	dbConfig.AddConfigPath(".")
	dbConfig.SetConfigType("yaml")
	if err = dbConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	dbConfig.WatchConfig()
	dbConfig.OnConfigChange(func(e fsnotify.Event) {
		//	log.Println("App Config file changed %s:", e.Name)
	})
	app.DBConfig = Config(*dbConfig)
}

// loadENV
func (app *Application) loadENV() {

	var APPENV string
	var appConfig viper.Viper
	appConfig = viper.Viper(app.AppConfig)
	APPENV = appConfig.GetString("app_env")

	switch APPENV {
	case "test":
		app.ENV = "test"
		break
	case "dev":
		app.ENV = "dev"
		break
	case "staging":
		app.ENV = "staging"
		break
	case "production":
		app.ENV = "production"
		break
	default:
		app.ENV = "dev"
		break
	}
}

// String: read string value from viper.Viper
func (config *Config) String(key string) string {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetString(fmt.Sprintf("%s.%s", App.ENV, key))
}

// Int: read int value from viper.Viper
func (config *Config) Int(key string) int {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetInt(fmt.Sprintf("%s.%s", App.ENV, key))
}

// Boolean: read boolean value from viper.Viper
func (config *Config) Boolean(key string) bool {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetBool(fmt.Sprintf("%s.%s", App.ENV, key))
}

func (config *Config) GetStringMapString(key string) map[string]string {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetStringMapString(fmt.Sprintf("%s.%s", App.ENV, key))
}

func (config *Config) GetSub(key string) *viper.Viper {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.Sub(fmt.Sprintf("%s.%s", App.ENV, key))
}

func (config *Config) GetTime(key string) time.Time {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetTime(fmt.Sprintf("%s.%s", App.ENV, key))
}

func (app *Application) DBInit() *gorm.DB {
	var adapter string
	adapter = app.DBConfig.String("adapter")
	switch adapter {
	case "mysql":
		mysqlConn()
	case "postgre":
		postgresConn()
	default:
		log.Println("Undefined connection on config.yaml")
	}

	return db
}

// setupPostgresConn: setup postgres database connection using the configuration from database.yaml
func postgresConn() {
	var (
		connectionString string
		err              error
	)
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", App.DBConfig.String("username"), App.DBConfig.String("password"), App.DBConfig.String("host"), App.DBConfig.String("port"), App.DBConfig.String("database"), App.DBConfig.String("sslmode"))
	if db, err = gorm.Open("postgres", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec(fmt.Sprintf("SET TIMEZONE TO '%s'", App.AppConfig.String("timezone")))
	db.Exec("CREATE EXTENSION postgis")
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(App.DBConfig.Int("maxlifetime")))
	db.DB().SetMaxIdleConns(App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(App.DBConfig.Int("open_conns"))
}

// mysqlConn: setup mysql database connection using the configuration from database.yaml
func mysqlConn() {
	var (
		connectionString string
		err              error
	)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DBConfig.String("username"), App.DBConfig.String("password"), App.DBConfig.String("host"), App.DBConfig.String("port"), App.DBConfig.String("database"))

	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec(fmt.Sprintf("SET TIMEZONE = '%s'", App.AppConfig.String("timezone")))
	db.DB().SetMaxIdleConns(App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(App.DBConfig.Int("open_conns"))
}
