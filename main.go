package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/portal"
)

var (
	apiKey     string = ""
	blinkState lights.State
)

func init() {
	blinkState = lights.State{On: true, Alert: "lselect"}
	flag.StringVar(&apiKey, "key", os.Getenv("HUE_USERNAME"), "hue light api key")
	flag.Parse()
}

func dashboard_failing() {
	cmd := exec.Command("phantomjs", "parse-dev-dashboard.js")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Dav-Dashboard has failing widgets: %q\n", out.String())
}

func main() {
	dashboard_failing()

	return
	// hue lamp test code
	pp, err := portal.GetPortal()

	if err != nil {
		fmt.Println("portal.GetPortal() ERROR: ", err)
		os.Exit(1)
	}
	ll := lights.New(pp[0].InternalIPAddress, apiKey)
	allLights, err := ll.GetAllLights()

	if err != nil {
		fmt.Println("lights.GetAllLights() ERROR: ", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("Lights")
	fmt.Println("------")

	for _, l := range allLights {
		fmt.Printf("ID: %d Name: %s\n", l.ID, l.Name)

	}
	gg := groups.New(pp[0].InternalIPAddress, apiKey)

	allGroups, err := gg.GetAllGroups()

	if err != nil {
		fmt.Println("groups.GetAllGroups() ERROR: ", err)
		os.Exit(1)

	}

	fmt.Println()
	fmt.Println("Groups")
	fmt.Println("------")

	for _, g := range allGroups {
		fmt.Printf("ID: %d Name: %s\n", g.ID, g.Name)
		for _, lll := range g.Lights {
			fmt.Println("\t", lll)

		}
		previousState := g.Action
		_, err := gg.SetGroupState(g.ID, blinkState)
		if err != nil {
			fmt.Println("groups.SetGroupState() ERROR: ", err)
			os.Exit(1)

		}
		time.Sleep(time.Second * time.Duration(10))
		_, err = gg.SetGroupState(g.ID, previousState)
		if err != nil {
			fmt.Println("groups.SetGroupState() ERROR: ", err)
			os.Exit(1)

		}
	}
}
