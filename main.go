package main

import (
	"flag"

	log "unknwon.dev/clog/v2"

	"github.com/vidar-team/srun-login/pkg/srun"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	host := flag.String("host", "https://login.hdu.edu.cn/", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")
	flag.Parse()

	client := srun.NewClient(*host, *username, *password)
	challengeResp, err := client.GetChallenge()
	if err != nil {
		log.Fatal("Failed to get challenge %v", err)
	}
	challenge := challengeResp.Challenge
	log.Trace("Challenge: %q", challenge)

	portalResp, err := client.Portal(challengeResp.Challenge)
	if err != nil {
		log.Fatal("Failed to portal: %v", err)
	}
	log.Trace("%+v", portalResp)
}
