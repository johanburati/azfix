package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
    "github.com/johanburati/azfix/util"
)

func main() {
    OSInfo := util.ReadOSRelease("/etc/os-release")
    app := &cli.App{
        Name:  "azfix",
        Version: "v0.0.0-alpha",
        Usage: "azfix is a tool for fixing Linux issues on Azure",
        EnableBashCompletion: true,
        Commands: []*cli.Command{
            {
                Name:    "info",
                Aliases: []string{"i"},
                Usage:   "print system info",
                Action: func(cCtx *cli.Context) error {
                    OSName := OSInfo["NAME"]
                    OSVersion := OSInfo["VERSION"]
                    fmt.Println("Distro:",OSName, OSVersion)
                    return nil
                },
            },
            {
                Name:    "fix",
                Aliases: []string{"f"},
                Usage:   "fix an issue",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("fix: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "check",
                Aliases: []string{"c"},
                Usage:   "check for an issue",
                Subcommands: []*cli.Command{
                    {
                        Name:  "repolist",
                        Usage: "check repolist",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("repolist: ", cCtx.Args().First())
                            RepoUrl:="rhui-1.microsoft.com"
                            ips, err := util.CheckRepoUrl(RepoUrl)
                            if err != nil {
                                log.Fatal("Fail to check Url: ", err)
                            }
                            for _, ip := range ips {
                                fmt.Printf("%s. IN A %s\n", RepoUrl, ip.String())
                            }                       
                            return nil
                        },
                    },
                    {
                        Name:  "wireserver",
                        Usage: "check wireserver",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("wireserver: ", cCtx.Args().First())
                            return nil
                        },
                    },
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}