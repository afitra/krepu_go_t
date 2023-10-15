package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
	_userHandler "krepu_go_t/domains/user/delivery/http"
	_userRepository "krepu_go_t/domains/user/repository"
	_userUseCase "krepu_go_t/domains/user/usecase"
	"krepu_go_t/logger"
	"krepu_go_t/models"
	"net/http"
	"os"
	"time"
)

var ech *echo.Echo

type CustomValidator struct {
	validator *validator.Validate
}

func main() {

	sqlConn, sqlxConn := getDBConn()

	defer sqlConn.Close()

	echoGroup := models.EchoGroup{
		API: ech.Group(""),
	}

	customValidator := validator.New()

	ech.Validator = &CustomValidator{validator: customValidator}

	userRepository := _userRepository.NewPsqlUniqLink(sqlxConn)
	_userUseCase := _userUseCase.NewUserUseCase(userRepository)
	_userHandler.NewUserHandler(echoGroup, _userUseCase)

	ech.GET("/ping", ping)
	err := ech.Start(":" + os.Getenv("APP_PORT"))

	if err != nil {
		logger.Make(nil, nil).Debug(err)
	}
}

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	ech = echo.New()
	ech.Debug = true
	loadEnv()
	logger.Init()
}

func ping(echTx echo.Context) error {

	response := map[string]interface{}{
		"status": "online",
		"month":  "Server Actived!!",
	}

	return echTx.JSON(http.StatusOK, response)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func loadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load()

	if err != nil {
		logger.Make(nil, nil).Fatal("Error loading .env file")
	}
}

func getDBConn() (*sql.DB, *sqlx.DB) {
	dbHost := os.Getenv(`DB_HOST`)
	dbPort := os.Getenv(`DB_PORT`)
	dbUser := os.Getenv(`DB_USER`)
	dbPass := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)

	// connection for sqlx
	sqlxConnection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	sqlxConn, err := sqlx.Connect("postgres", sqlxConnection)

	if err != nil {
		logger.Make(nil, nil).Debug(err)
	}

	// connection for sql/migration
	sqlConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	sqlConn, err := sql.Open(`postgres`, sqlConnection)

	if err != nil {
		logger.Make(nil, nil).Debug(err)
	}

	err = sqlxConn.Ping()

	if err != nil {
		fmt.Println("xoxoxo ", err)
		logger.Make(nil, nil).Debug(err)
		os.Exit(1)
	}

	err = sqlConn.Ping()

	if err != nil {
		logger.Make(nil, nil).Debug(err)
		os.Exit(1)
	}

	return sqlConn, sqlxConn
}
