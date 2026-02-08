package main

import (
	"fmt"
	"learn-go/internal/database"
	"learn-go/internal/handler"
	"learn-go/internal/repository"
	"learn-go/internal/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"port"`
	DBConn string `mapstructure:"db_conn"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	})
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)
	http.HandleFunc("/api/report/hari-ini", transactionHandler.HandleReportToday)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
