package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
	"github.com/Haybu/saw/actions"

	"time"
)



func main() {
	app := cli.NewApp()
	app.Name = "saw"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Haytham Mohamed",
			Email: "haytham_mohamed@homedepot.com",
		},
	}
	app.Copyright = "(c) 2017 The Home Depot"
	app.HelpName = "saw"
	app.Usage = "saw CLI to manage clusters and application deployments in THD Cloud"
	app.UsageText = "saw - The Home Depot Cloud cluster and application management CLI"
	app.ArgsUsage = "[-n CLUSTER -p PROJECT -z ZONE -d nodes]"

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "cluster",
			Aliases:     []string{"cl"},
			Category:    "cluster",
			Usage:       "saw cluster -n [cluster-name] -p [project-name] -z [zone-name] -d [number-of-nodes]",
			UsageText:   "-n [cluster-name] -p [project-name] -z [zone-name] -d [number-of-nodes]",
			Description: "You can create, describe and delete a cluster",
			ArgsUsage:   "[-n NAME -p PROJECT -z ZONE -d NUMBER]",

			Subcommands: cli.Commands{
				cli.Command{
					Name:   "create",
					Action: actions.ClusterCreateAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
						cli.Int64Flag{Name: "nodes, d", Value: 3},
					},
				},
				cli.Command{
					Name:   "nodes",
					Action: actions.ClusterNodesAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
					},
				},
				cli.Command{
					Name:   "delete",
					Action: actions.ClusterDeleteAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
					},
				},
				cli.Command{
					Name:   "info",
					Action: actions.ClusterInfoAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
					},
				},
			},

			SkipFlagParsing: false,
			HideHelp:        false,
			Hidden:          false,
			HelpName:        "help",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "create delete info nodes\n")
			},
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "Before cluster command\n")
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "\nAfter cluster command\n")
				return nil
			},
			Action: func(c *cli.Context) error {
				//flags := c.Command.VisibleFlags()
				fmt.Fprintf(c.App.Writer, "in cluster action\n")
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "error occured\n")
				return err
			},
		},   // end cluster command
		{
			Name:        "application",
			Aliases:     []string{"app"},
			Category:    "application",
			Usage:       "saw app -n [cluster-name] -p [project-name] -z [zone-name] -s [namespace] -a [appname] -v [version] -i [image] -r [replicas]",
			UsageText:   "-n [cluster-name] -p [project-name] -z [zone-name] -s [namespace] -a [appname] -v [version] -i [image] -r [replicas]",
			Description: "You can build, deploy, expose and ignite an application",
			ArgsUsage:   "n [cluster-name] -p [project-name] -z [zone-name] -s [namespace] -a [appname] -v [version] -i [image] -r [replicas]",

			Subcommands: cli.Commands{
				cli.Command{
					Name:   "build",
					Action: actions.BuildApplicationAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "image, n"},
					},
				},
				cli.Command{
					Name:   "deploy",
					Action: actions.DeployApplicationAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
						cli.StringFlag{Name: "appname, a"},
						cli.Int64Flag{Name: "replicas, r"},
						cli.StringFlag{Name: "image, i"},
						cli.StringFlag{Name: "version, v"},
						cli.StringFlag{Name: "namespace, s", Value: "default"},
					},
				},
				cli.Command{
					Name:   "pods",
					Action: actions.ListPodsAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name, n", Value: "my-cluster"},
						cli.StringFlag{Name: "project, p", Value: "my-project"},
						cli.StringFlag{Name: "zone, z", Value: "us-central1-b"},
						cli.StringFlag{Name: "namespace, s", Value: "default"},
					},
				},
			},

			SkipFlagParsing: false,
			HideHelp:        false,
			Hidden:          false,
			HelpName:        "help",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "build deploy pods scale ignite\n")
			},
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "Before application command\n")
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "\nAfter application command\n")
				return nil
			},
			Action: func(c *cli.Context) error {
				//flags := c.Command.VisibleFlags()
				fmt.Fprintf(c.App.Writer, "in application action\n")
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "error occured\n")
				return err
			},
		}, // end app command
	}

	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false
	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "cluster\napp\n")
	}
	app.Before = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "Before saw cli starts\n")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "saw cli is all completed!\n")
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "command %q not found.\n", command)
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
		return nil
	}

	app.Run(os.Args)
}
