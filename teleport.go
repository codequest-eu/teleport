package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codegangsta/cli"
	"github.com/marcinwyszynski/tcproxy"
	"github.com/mgutz/ansi"
)

func handleConn(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error connecting to remote host %s: %v\n",
			remoteAddr,
			err,
		)
		os.Exit(3)
	}
	tcproxy.TCProxy(remote.(*net.TCPConn), local.(*net.TCPConn))
}

func serve(ctx *cli.Context) {
	fmt.Fprintln(
		os.Stderr,
		"Made with",
		ansi.Color("<3", "red"),
		"by codequest.com",
	)
	remoteAddr, localPort := ctx.String("remote"), ctx.Int("port")
	fmt.Fprintf(
		os.Stderr,
		"Creating a tunnel between %q and localhost:%d\n",
		remoteAddr,
		localPort,
	)
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("127.0.0.1:%d", localPort),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting local server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	for {
		local, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Error accepting incoming connection: %v\n",
				err,
			)
			os.Exit(2)
		}
		go handleConn(local, remoteAddr)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "teleport"
	app.Usage = "Like ngrok, but reverse - tunnel a remote endpoint to your localhost"
	app.Version = "0.0.1alpha"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "remote, r",
			Usage: "remote endpoint - eg. 0.tcp.ngrok.io:42644",
			Value: "",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "local port to expose",
			Value: 3000,
		},
	}
	app.Action = serve
	app.Run(os.Args)
}
