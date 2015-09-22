# teleport - like ngrok but reverse

[ngrok](https://ngrok.io) is awesome. For those not in the know, ngrok addresses a use case where you want to expose a local server behind a NAT or firewall to the internet. Say you want to show your work to a colleague not on the local network - just run `ngrok http 3000` (assuming you're running your local server on port 3000) and you're given an HTTP address on the Wild Wild Web that you can share with your colleagues.

Except things sometimes won't work this way - sometimes things in your project only really work when run on localhost. This is the problem [teleport](https://github.com/codequest-eu/teleport) wants to address. It creates a TCP tunnel to a remote endpoint and passes connections from a selected port on the localhost (3000 by default) to the remote endpoint. This is a perfect companion for ngrok.

_Developer A_ wants to show her work:
```bash
$ ngrok tcp 3000  # gets an address like 0.tcp.ngrok.io:42644
```

_Developer B_ wants to look at _A's_ work from the cosy confines of his localhost:
```bash
$ teleport -r 0.tcp.ngrok.io:42644
```

Now for as long as this is running _Developer B_ can access localhost:3000 locally and connect via a magic gateway to _Developer A's_ server. Note that teleport works on TCP level so you will need to expose a TCP port with ngrok.

Of course nothing requires you to use teleport _only_ with ngrok - teleport will work standalone just fine so you can connect to a server running on your VPS as long as you can access it from your local machine.

# Installing
If you have Go installed on your system you can use Go's package manager to install the newest version binary by running `go get -u github.com/codequest-eu/telelport`. Otherwise there are pre-built binaries available from the [Releases](https://github.com/codequest-eu/teleport/releases) page.
