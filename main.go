package main

import (
	"log"
	"net"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/cifra-city/cifra-sso/service/config"
	"github.com/cifra-city/cifra-sso/service/data/db"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

// AuthServer - структура для реализации gRPC сервиса.
type AuthServer struct {
	ssov1.UnimplementedAuthServer             // Для совместимости с gRPC.
	Queries                       *db.Queries // Объект для работы с базой данных.
}

// NewAuthServer создает новый экземпляр AuthServer.
func NewAuthServer(queries *db.Queries) *AuthServer {
	return &AuthServer{
		Queries: queries,
	}
}

func main() {
	// Загружаем конфигурацию из config.yaml.
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Printf("Database URL: %s", cfg.Database.URL)

	// Настраиваем логгер.
	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC server...")

	// Инициализируем подключение к базе данных.
	conn, err := db.NewDBConnection(cfg.Database.URL)
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
	}
	defer conn.Close() // Закрываем подключение к базе данных при завершении работы.

	// Создаем объект Queries для выполнения запросов.
	queries := db.New(conn)

	// Создаем сервер AuthServer.
	authServer := NewAuthServer(queries)

	// Создаем базовый gRPC сервер.
	grpcServer := grpc.NewServer()

	// Регистрируем AuthServer на базовом gRPC сервере.
	ssov1.RegisterAuthServer(grpcServer, authServer)

	// Создаем сетевой слушатель на порту, указанном в конфигурации.
	listener, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		logger.Fatalf("failed to listen on port %s: %v", cfg.Server.Port, err)
	}

	// Запускаем gRPC сервер.
	logger.Infof("gRPC server is running on port %s", cfg.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve gRPC server: %v", err)
	}
}
