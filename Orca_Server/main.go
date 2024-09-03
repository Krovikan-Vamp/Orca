package main

import (
	"Orca_Server/routers"
	"Orca_Server/servers"
	"Orca_Server/setting"
	"Orca_Server/sqlmgmt"
	"Orca_Server/tools/log"
	log2 "log"
	"net/http"
)

func init() {
	setting.Setup()
	log.Setup()
}

func main() {
	// Print logo
	//color.Green("OrcaC2 Server " + define.Version)
	//color.Green(define.Logo)
	// Initialize the database
	sqlmgmt.InitDb()
	// Initialize routes
	routers.Init()
	// Start a timer to send heartbeat signals
	servers.PingTimer()
	log2.Printf("Server started successfully, port: %s", setting.CommonSetting.HttpPort)

	if err := http.ListenAndServe(":"+setting.CommonSetting.HttpPort, nil); err != nil {
		panic(err)
	}
}
