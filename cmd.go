package main

import "flag"

type Cmd struct {
	bind string
	param string
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.bind, "listen", "0.0.0.0:8089", "listen on ip:port")
	flag.StringVar(&cmd.param, "param", "_pref", "target address query param")
	flag.Parse()
	return cmd
}