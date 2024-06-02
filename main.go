package main

import (
	"flag"
	"fmt"
	"future-trading/marketdataservice/bootstraper"
	"future-trading/marketdataservice/config"
	"future-trading/marketdataservice/server"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "c", "./marketdataservice/config/config.yaml", "configuration file path")
	flag.Parse()
	configuration := config.GetConf(path)

	// new server and start service
	bootstrapper, err := bootstraper.NewBootstrapper(bootstraper.Config{User: configuration.Postgres.User, Host: configuration.Postgres.Host, Password: configuration.Postgres.Password, Port: configuration.Postgres.Port, Name: configuration.Postgres.Name, SSL: configuration.Postgres.SslMode})
	if err != nil {
		//bootstrapper.Logger.Error(context.Background(), "unable to connect to the database: %w", err)
		os.Exit(1)
	}

	r := server.NewRouter(configuration.Mode, bootstrapper)

	err = r.Run(fmt.Sprintf("%s:%s", configuration.Host, configuration.Port))
	if err != nil {
		fmt.Printf("Start server failed, error: %v", err.Error())
		os.Exit(0)
	}
}
