package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"
	"github.com/wwwwshwww/spot-sandbox/graph"
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/cache"
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/google_maps"
	"github.com/wwwwshwww/spot-sandbox/internal/config"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const defaultPort = "8080"
const dbName = "yeah"

func init() {
	config.Configure()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB:   getMySQL(),
		GMC:  getGoogleMapsClient(),
		DuCC: getDurationCacheClient(),
		DiCC: getDistanceCacheClient(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getMySQL() *gorm.DB {
	cfg := mysql.Config{
		DBName:               dbName,
		User:                 config.MySQL.User,
		Passwd:               config.MySQL.Passwd,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.MySQL.Host, config.MySQL.Port),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := gorm.Open(gormysql.Open(cfg.FormatDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect DB")
	}
	return db
}

func getGoogleMapsClient() *google_maps.GoogleMapsClient {
	gmc, err := google_maps.NewGoogleMapsClient(config.GoogleMapsAPIKey)
	if err != nil {
		panic("failed to get GoogleMapsClient")
	}
	return gmc
}

func getDurationCacheClient() *cache.DurationCacheClient {
	c, err := cache.NewDurationCacheClient(config.Redis.Host, config.Redis.Port)
	if err != nil {
		panic("failed to get DurationCacheClient")
	}
	return c
}

func getDistanceCacheClient() *cache.DistanceCacheClient {
	c, err := cache.NewDistanceCacheClient(config.Redis.Host, config.Redis.Port)
	if err != nil {
		panic("failed to get DistanceCacheClient")
	}
	return c
}
