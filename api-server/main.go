package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/viking0312/golang-mongodb-http-server/internal/db"
	"github.com/viking0312/golang-mongodb-http-server/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SERVER_PORT_CONFIG_NAME = "server.port"
)

var dbClient *mongo.Client

func init() {
	loadAppConfig()
}

func main() {

	log.Println("staring server on port", viper.Get(SERVER_PORT_CONFIG_NAME), "...")

    var err error
	dbClient, err = db.GetDbClient()

    if err != nil {
        util.PanicWithError("getting DB client", err)
    }

	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			util.PanicWithError("db disconnect", err)
		}
	}()

	httpHandlers()

	err = http.ListenAndServe(fmt.Sprintf("localhost:%d", viper.Get(SERVER_PORT_CONFIG_NAME)), nil)
	if err != nil {
		util.PanicWithError("server startup", err)
	}
}

func httpHandlers() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there!!")
	})

	http.HandleFunc("/record/movie/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("http request received", "path:", r.URL.Path, "method:", r.Method)
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Bad ID value", http.StatusBadRequest)
			return
		}

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				util.PanicWithError("reading body", err)
			}

			var movie db.Movies
			json.Unmarshal(body, &movie)

			updatedCount, err := db.UpdateMovie(dbClient, objId, movie)

			if err != nil {
				util.PanicWithError("updating DB record", err)
			}

			fmt.Fprintf(w, "%v", updatedCount)
		} else if r.Method == "DELETE" {
			deletedCount, err := db.DeleteMovie(dbClient, objId)
			if err != nil {
				util.PanicWithError("deleting record from DB", err)
			}

			fmt.Fprintf(w, "%+v", deletedCount)
		} else {
			res, err := db.GetMovie(dbClient, objId)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					http.NotFound(w, r)
					return
				}
				util.PanicWithError(fmt.Sprintf("movie query for id %v", id), err)
			}

			fmt.Fprintf(w, "%+v", res)
		}
	})

	http.HandleFunc("/record/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				util.PanicWithError("convertig body to bson", err)
			}

			var movie db.Movies
			json.Unmarshal(body, &movie)

			res, err := db.CreateMovie(dbClient, movie)
			if err != nil {
				util.PanicWithError("creating movie record", err)
			}

			fmt.Fprintf(w, "%v", res)
		}
	})
}

func loadAppConfig() {
	//this sets a common env var prefix
	viper.SetEnvPrefix("GMWS")
	viper.BindEnv("env")

	//this will get an env var with the name GMWS_ENV as prefix was set earlier
	envName := viper.Get("env")
	if envName == nil || envName == "" {
		fmt.Println("GMWS_ENV environment variable missing")
		envName = "dev"
	}

	log.Println("app profile set to", envName)

	//read the config file
	viper.SetConfigName(fmt.Sprintf("application-%s", envName))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		util.PanicWithError("reading app config", err)
	}
}
