package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "arp-exporter"
	appDescription = "arp table exporter, supports prometheus metrics or json format"
	appMainversion = "0.1"
)

var CLI struct {
	ArpTableFile   string `help:"retrieve arp table from file, default is by command 'arp -an'" short:"f"`
	MacInfo        string `help:"look up and display mac address vendor information" short:"m"`
	ListVendors    bool   `help:"list all vendors in db" short:"l"`
	HostnameConfig string `help:"host name configuration file" short:"c"`
	EnableVendors  bool   `help:"enable displaying mac vendors" short:"e"`
	Server         bool   `help:"run web server" short:"s"`
	Bind           string `help:"where to bind the server to" short:"b" default:":9100"`
	LogFile        string `help:"log file" default:"/dev/stdout"`
	LogLevel       string `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors    bool   `help:"disable output colours, print plain text"`
	LogJSON        bool   `help:"enable json log, instead of text one"`
	VersionFlag    bool   `help:"display version" short:"V"`
}

func parseArgs() {
	homeFolder := getHome()
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"config": returnFirstExistingFile(
				[]string{
					path.Join(getBindir(), appName+".yaml"),
					path.Join(homeFolder, ".conf", appName, "conf.yaml"),
					path.Join(homeFolder, ".config", appName, "conf.yaml"),
				},
			),
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

func getBindir() (s string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	s = filepath.Dir(ex)
	return
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("unable to retrieve user's home folder")
	}
	return usr.HomeDir
}

func isFile(filePath string) bool {
	stat, err := os.Stat(makeAbs(filePath))
	if !os.IsNotExist(err) && !stat.IsDir() {
		return true
	}
	return false
}

func makeAbs(filename string) string {
	filename, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf("can not assemble absolute filename %q\n", err)
		os.Exit(1)
	}
	return filename
}

func returnFirstExistingFile(arr []string) (s string) {
	for _, el := range arr {
		if isFile(el) {
			s = el
			break
		}
	}
	return
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}
