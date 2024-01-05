package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/rolacher/go-smartme/smartmeapi"
	"github.com/tkanos/gonfig"
)

type Globals struct {
	Debug      bool        `help:"Enable debug mode." short:"d"`
	Configfile string      `type:"path" help:"Specify configuration file." short:"c" default:"~/.smartmeapi-config.json"`
	Pretty     bool        `short:"p" help:"Pretty format for json" default:"false"`
	Version    VersionFlag `name:"version" help:"Print version information and quit"`
}

type CLI struct {
	Globals

	Devices DevCmd `cmd help:"Get parameters of all devices or a selected device."`
	Values  ValCmd `cmd help:"Get values from a device, also from the past and multiple."`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

// --------------------------------------

var globals = Globals{
	Version: VersionFlag("0.0.1"),
}

type Configuration struct {
	Host     string `env:"SMARTMEAPI_HOST"`     // smart-me IP, https://smart-me.com
	Username string `env:"SMARTMEAPI_USER"`     // smart-me username
	Password string `env:"SMARTMEAPI_PASSWORD"` // smart-me password
}

var configuration = Configuration{}
var cli CLI

func main() {
	cli = CLI{
		Globals: globals,
	}

	ctx := kong.Parse(&cli,
		kong.Name("smecli"),
		kong.Description("A command line tool for reading the smart-me API"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: false,
			Summary: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		},
	)
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}

func initializeApi(configfile string) error {
	err := gonfig.GetConf(configfile, &configuration)
	if err != nil {
		log.Fatal("GetConfig: " + err.Error())
		return err
	}

	var (
		buf      bytes.Buffer
		logger   = log.New(&buf, "", log.LstdFlags)
		logLevel = 1 // 0: quiet / 1: normal / 2: verbose (show also all trigger outs)
	)
	logger.SetOutput(os.Stdout)

	smartmeapi.ConfigureApi(configuration.Host, configuration.Username, configuration.Password, logger, logLevel)
	return nil
}

func unmarshall(v interface{}) ([]byte, error) {
	if cli.Globals.Pretty {
		return json2.MarshalIndent(v, "", "  ")
	} else {
		return json2.Marshal(v)
	}
}
