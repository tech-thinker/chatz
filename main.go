package main

import (
	"fmt"
	"os"

	"github.com/tech-thinker/chat/agents"
	"github.com/tech-thinker/chat/utils"
	"github.com/urfave/cli/v2"
)

func main() {
    
    var profile string
    var threadId string

    app := cli.NewApp()
    app.Name = "Chat"
    app.Flags = []cli.Flag {
        &cli.StringFlag{
            Name: "profile",
            Aliases: []string{"p"},
            Value: "default",
            Usage: "Profile from .schat.ini",
            Destination: &profile,
        },
        &cli.StringFlag{
            Name: "thread-id",
            Aliases: []string{"t"},
            Value: "",
            Usage: "Thread ID for reply to a message",
            Destination: &threadId,
        },
    }
    app.Action = func(ctx *cli.Context) error {
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
            message = fmt.Sprintf("%s %s",message, a)
        }
        
        env, err := utils.LoadEnv(profile)
        if err!=nil {
            return nil
        }
        var agent agents.Agent
        switch env.Provider {
            case "slack":
                agent = agents.NewSlackAgent(env)
            case "google":
                agent = agents.NewGoogleAgent(env)
            default:
                fmt.Println("No valid provider. Please choose from [slack, google].")
                return nil
        }

        if len(threadId) > 0 {
            res, _ := agent.Reply(threadId, message)
            fmt.Println(res)
            return nil
        }
        res, _ := agent.Post(message)
        fmt.Println(res)

        return nil
    }

    if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}

