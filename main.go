package main

import (
	"fmt"
	"os"

	"github.com/tech-thinker/chatz/models"
	"github.com/tech-thinker/chatz/providers"
	"github.com/tech-thinker/chatz/utils"
	"github.com/urfave/cli/v2"
)

var (
	AppVersion = "v0.0.1"
	CommitHash = "unknown"
	BuildDate  = "unknown"
)

func main() {

	var version bool
	var profile string
	var threadId string
	var output bool
	var fromEnv bool
	var subject string
	var priority int

	app := cli.NewApp()
	app.Name = "chatz"
	app.Description = "chatz is a versatile messaging app designed to send notifications to Google Chat, Slack, Discord, Telegram and Redis."
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Print output to stdout",
			Destination: &output,
		},
		&cli.StringFlag{
			Name:        "profile",
			Aliases:     []string{"p"},
			Value:       "default",
			Usage:       "Profile from .chatz.ini",
			Destination: &profile,
		},
		&cli.StringFlag{
			Name:        "thread-id",
			Aliases:     []string{"t"},
			Value:       "",
			Usage:       "Thread ID for reply to a message",
			Destination: &threadId,
		},
		&cli.BoolFlag{
			Name:        "version",
			Aliases:     []string{"v"},
			Usage:       "Print the version number",
			Destination: &version,
		},
		&cli.BoolFlag{
			Name:        "from-env",
			Aliases:     []string{"e"},
			Usage:       "To use config from environment variables",
			Destination: &fromEnv,
		},
		&cli.StringFlag{
			Name:        "subject",
			Aliases:     []string{"s"},
			Usage:       "Subject for provider which supports subject or title",
			Destination: &subject,
		},
		&cli.IntFlag{
			Name:        "priority",
			Aliases:     []string{"pr"},
			Usage:       "Priority for gotify notification",
			Destination: &priority,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if version {
			fmt.Println("chatz version: ", AppVersion)
			fmt.Println("Commit Hash: ", CommitHash)
			fmt.Println("Build Date: ", BuildDate)
			return nil
		}

		var message string
		if ctx.Args().Len() == 0 {
			fmt.Println("Please provide a message.")
			return nil
		}
		for i, a := range ctx.Args().Slice() {
			if i == 0 {
				message = a
				continue
			}
			message = fmt.Sprintf("%s %s", message, a)
		}

		env, err := utils.LoadEnv(profile, fromEnv)
		if err != nil {
			return nil
		}

		provider, err := providers.NewProvider(env)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		option := models.Option{}

		if ctx.IsSet("subject") {
			option.Title = &subject
			option.Subject = &subject
		}

		if ctx.IsSet("priority") {
			option.Priority = &priority
		}

		if len(threadId) > 0 {
			res, _ := provider.Reply(threadId, message, option)
			if output {
				fmt.Println(res)
			}
			return nil
		}

		res, err := provider.Post(message, option)
		if output {
			fmt.Println(res)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
