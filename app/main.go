package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/bxcodec/go-clean-arch/app/docs"
	"github.com/bxcodec/go-clean-arch/bmi"
	grpcServer "github.com/bxcodec/go-clean-arch/internal/grpc"
	mysqlRepo "github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"

	"github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/middleware"
	"github.com/joho/godotenv"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title BMI Calculator API
// @version 1.0
// @description API สำหรับคำนวณค่า BMI
// @host localhost:9090
// @BasePath /
func main() {
	// prepare database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()
	server := grpc.NewServer()
	defer server.GracefulStop()

	// prepare echo

	e := echo.New()
	e.Use(middleware.CORS)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	authorRepo := mysqlRepo.NewAuthorRepository(dbConn)
	articleRepo := mysqlRepo.NewArticleRepository(dbConn)

	// Build service Layer
	svc := article.NewService(articleRepo, authorRepo)
	bmiService := bmi.NewService()

	rest.NewArticleHandler(e, svc)
	rest.NewBmiHandler(e, bmiService)

	bmiGrpc := grpcServer.NewGRPCBmiRoute(server)
	bmiGrpcHandler := grpcServer.NewBmiGrpcHandler(bmiService)
	bmiGrpc.RegisterBmiHandler(bmiGrpcHandler)

	/* serve gprc */
	go func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error))
		}
		startGRPCServer(server)
	}()
	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}

func startGRPCServer(server *grpc.Server) {
	GRPC_PORT := os.Getenv("GRPC_PORT")
	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	/* serve grpc */
	fmt.Println(fmt.Sprintf("Start grpc Server [:%s]", GRPC_PORT))
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
