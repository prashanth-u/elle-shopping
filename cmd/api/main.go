package main

import (
    "log"
    "shopping/internal/server"
	"shopping/internal/db"
	"os"
	"shopping/internal/config"
	"encoding/json"
    "fmt"
)

func LoadConfigFromFile(filename string, config interface{}) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return err
    }
    return nil
}

func main() {
    var serverConfig config.ServerConfig
    errServer := LoadConfigFromFile("server_config.json", &serverConfig)
    if errServer != nil {
        log.Fatalf("Failed to load server configuration: %s", errServer)
    }
    var mailConfig config.MailConfig
    errMail := LoadConfigFromFile("mail.json", &mailConfig)
    if errMail != nil {
        log.Fatalf("Failed to load mail configuration: %s", errMail)
    }
    var dbConfig config.DatabaseConfig
	errConfigDb := LoadConfigFromFile("db.json", &dbConfig)
    if errConfigDb != nil {
        log.Fatalf("Failed to load configuration: %s", errConfigDb)
    }
    databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	db, errDb := db.Connect(databaseConnectionString)
    if errDb != nil {
        log.Fatal(errDb)
    }
	defer db.Close()
	listenAddr := fmt.Sprintf(":%d", serverConfig.ServerPort)
	server := server.NewServer(listenAddr, db)
	log.Println("Starting server on port " + listenAddr + "...")
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}