package main

import (
	"business/cronJobs/jobs/country/config"
	"business/cronJobs/jobs/country/internal/handler"
	"business/cronJobs/jobs/country/internal/svc"
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	var c config.Config
	if err := config.Unmarshal("etc/country.yaml", &c); err != nil {
		log.Fatal(err)
		return
	}
	svcCtx := svc.NewServiceContext(c)

	fmt.Println("crontab job start...")
	handler.StartHandlers(svcCtx)
	fmt.Println("crontab job end...")
}
