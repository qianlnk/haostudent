/************************************************
File:		haosudent.go
Version:		0.0.1
Description:	main
Author:		Qianno.Xie
Email:		qianno.xie@appcoachs.com
History:		2015.12.06 created by Qianno.Xie
************************************************/
package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/golang/glog"
	"haostudent/router"
	"haostudent/config"
	"os"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			glog.V(1).Infof("ERROR: recover err:%+v", err)
		}
	}()

	defer glog.Flush()
	ticker := time.NewTicker(1000 * 1000 * 100)
	go func() {
		for _ = range ticker.C {
			glog.Flush()
		}
	}()
	defer ticker.Stop()
	flag.Parse()
	conf := config.GetConfig()
	fmt.Println(conf)
	app := cli.NewApp()
	app.Name = "haostudent server"
	app.Author = "Qianno.Xie"
	app.Version = "0.0.1"
	app.Usage = "haostudent [port] default 1314"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		port := ":9000"
		m := router.NewRouter()
		fmt.Printf("c is %+v, m is %+v\n", c.Args().First(), m)
		if c.Args().First() != "" {
			port = ":" + c.Args().First()
		}
		m.RunOnAddr(port)
	}
	app.Run(os.Args)

}
