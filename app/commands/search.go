package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/juju/errors"

	"github.com/codelingo/lingo/app/util"
	"github.com/codelingo/lingo/service"
	grpcclient "github.com/codelingo/lingo/service/grpc"
	"github.com/codelingo/platform/flow/service/client"
	flow "github.com/codelingo/platform/flow/service/flowengine"

	"context"
	"github.com/codegangsta/cli"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	register(&cli.Command{
		Name:        "search",
		Usage:       "Search code following queries in .lingo.",
		Subcommands: cli.Commands{*pullRequestCmd},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  util.OutputFlg.String(),
				Usage: "File to save found results to.",
			},
			cli.StringFlag{
				Name:  util.FormatFlg.String(),
				Value: "json-pretty",
				Usage: "How to format the found results. Possible values are: json, json-pretty.",
			},
			// cli.BoolFlag{
			// 	Name:  util.InteractiveFlg.String(),
			// 	Usage: "Be prompted to confirm each issue.",
			// },
		},
		Description: `
""$ lingo search <filename>" .
`[1:],
		Action: searchAction,
	},
		false,
		homeRq, authRq, configRq, versionRq,
	)
}

func searchAction(ctx *cli.Context) {
	msg, err := searchCMD(ctx)
	if err != nil {

		// Debugging
		util.Logger.Debugw("searchAction", "err_stack", errors.ErrorStack(err))

		util.OSErr(err)
		return
	}

	fmt.Println(msg)
}

func searchCMD(cliCtx *cli.Context) (string, error) {
	defer util.Logger.Sync()
	util.Logger.Debugw("searchCMD called")

	args := cliCtx.Args()
	if len(args) == 0 {
		return "", errors.New("Please specify the filepath to a .lingo file.")
	}

	dotlingo, err := ioutil.ReadFile(args[0])
	if err != nil {
		return "", errors.Trace(err)
	}

	conn, err := service.GrpcConnection(service.LocalClient, service.FlowServer)
	if err != nil {
		return "", errors.Trace(err)
	}

	c := client.NewFlowClient(conn)

	ctx, err := grpcclient.GetGcloudEndpointCtx()
	if err != nil {
		return "", errors.Trace(err)
	}
	ctx, cancel := context.WithCancel(ctx)
	// Cancel the context passed to the platform on exit.
	sigc := make(chan os.Signal, 2)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		cancel()
		os.Exit(1)
	}()

	resultc, errorc, err := c.Search(ctx, &flow.SearchRequest{
		Dotlingo: string(dotlingo),
	})
	if err != nil {
		return "", errors.Trace(err)
	}

	results := []*flow.Result{}

l:
	for {
		select {
		case err, ok := <-errorc:
			if !ok {
				break l
			}

			util.Logger.Debugw(err.Error())
			util.Logger.Sync()

		case result, ok := <-resultc:
			if !ok {
				break l
			}

			results = append(results, result)
		}
	}

	msg, err := OutputResults(results, cliCtx.String("format"), cliCtx.String("output"))
	return msg, errors.Trace(err)
}

func OutputResults(results []*flow.Result, format, outputFile string) (string, error) {
	var data []byte
	var err error
	switch format {
	case "json":
		data, err = json.Marshal(results)
		if err != nil {
			return "", errors.Trace(err)
		}
	case "json-pretty":
		data, err = json.MarshalIndent(results, " ", " ")
		if err != nil {
			return "", errors.Trace(err)
		}
	default:
		return "", errors.Errorf("Unknown format %q", format)
	}

	if outputFile != "" {
		err = ioutil.WriteFile(outputFile, data, 0775)
		if err != nil {
			return "", errors.Annotate(err, "Error writing results to file")
		}
		return fmt.Sprintf("Done! %d results written to %s \n", len(results), outputFile), nil
	}

	return string(data), nil
}
