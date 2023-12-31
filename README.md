# `awesome-trivia`

`awesome-trivia` is a client of the [`trivial`] API server.  It sets up a new trivia game and allows for multiple players to join using a time-sensitive token that is generated when the new game is started.

It is easy to create a new game.  Here are the following steps:

1. Create a new socket server.  This will define the host location of where the game is hosted (which you send to each player) and expose the APIs which the game will use throughout the session.  In addition, it exposes a `Games` map to add the new game to, which allows for multiple games to be hosted on the same socket server.

1. Optionally generate a `TLS` certificate.  The game will expect to find the `cert.pem` and `key.pem` files in the same directory as the `awesome-trivia` binary, so you can "bring your own" or optionally have the [`trivial`] API server generate them for you.

1. Create and register the new game.

1. Start the game.

The `awesome-trivia` game client will take care of most of this for you.  For example, once the binary is placed on the publicly-accessible remote server, a simple command such as the following is all that is needed to start a game:

```bash
$ ./awesome-trivia -generateCert -wss wss://167.114.97.28:3000
---------------------------------------------------------------------------
created new websocket server `wss://167.114.97.28:3000/ws`
generated new TLS certificate for domains `127.0.0.1` and `127.0.0.1`
registered game `default` with key `bZu5SaAQ5d3EEwz1bkEp` on host `https://127.0.0.1:3000`
---------------------------------------------------------------------------
```

- If you already have created your cert, simply omit the `-generateCert` flag.
- The web socket `URL` needs to be the same as the server `IP` address.
- The generated private key (`bZu5SaAQ5d3EEwz1bkEp` in this example) should be distributed to all of the game players.  This is a time-sensitive token that will only allow a player to successfully login up to one hour from the time of the token creation.
- Distribute the `URL` of the game server to all of the players (i.e., `https://167.114.97.28:3000`).  Once there, they can choose a username and enter the private key (`bZu5SaAQ5d3EEwz1bkEp`).  This will allow them entry to the game.

<!--## Testing the `/query` Endpoint-->

## Controlling the Game

The question and the answer(s) are delimited by the pipe (`|`) symbol.  Here is a breakdown of the format:

```
Question?|Number of points given for a correct answer.|The correct answer.|Multiple possible answers, each separated by a pipe (`|`)
```

A question can have more than one correct answer.  Simply separate each correct answer by a comma (`,`).

For instance:

```
Name the Beatles?|50|1,2,3,5|John|Paul|George|Tony|Ringo
```

Currently, game control is facilitated on the command line using `curl`.  Here is an example:

```bash
$ curl -XGET -H "X-TRIVIA-APIKEY: bZu5SaAQ5d3EEwz1bkEp" \
    --data "What year did the Beatles play Budokan?|50|2|1965|1968|1970" \
    127.0.0.1:3000/query
```

> Note that the `X-TRIVIA-APIKEY` header expects the private key that was generated when the server was started.  Refer to the logs above.

This will send the question to the game players.  If you're controlling the game from a remote machine, replace the `loopback` address with the `IP` address or domain of the remote server.

An easier way to send questions to the players would be to concatenate all of the `csv` files together which contain the questions in the format specified above and then create a variable which will be incremented to pull the questions from the file line-by-line:

```bash
$ cat *.csv > game.csv
```

Then, set the variable to increment:

```bash
$ i=0
```

Now, push questions through to the game players by issuing a single chained command that first updates the variable and then uses it to pick the respective line from the `game.csv` file:

```bash
$ i=$((i+1)) && curl -XGET -H "X-TRIVIA-APIKEY: bZu5SaAQ5d3EEwz1bkEp" \
    --data "$(awk 'NR=='$i'' game.csv)" \
    127.0.0.1:3000/query
```

> If using a self-signed `TLS` certificate, pass the `-k` or `--insecure` switch so `curl` will disable strict certificate checking.
>
> ```bash
> $ curl --insecure -H "X-TRIVIA-APIKEY: bZu5SaAQ5d3EEwz1bkEp" \
>     --data "$(awk 'NR=='$i'' questions/roman_history.csv)" \
>     127.0.0.1:3000/query
> ```

## Troubleshooting

If the server that is running the game has an older version of `glibc`, you may encounter the following error when starting the game:

```bash
$ ./awesome-trivia
./awesome-trivia: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.32' not found (required by ./awesome-trivia)
./awesome-trivia: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.34' not found (required by ./awesome-trivia)
```

Here is one way "fix" that error:

```bash
$ CGO_ENABLED=0 go build
```

Or, install a previous version of Go.

## References

- [View the documentation](https://pkg.go.dev/github.com/btoll/trivial)

## License

[GPLv3](COPYING)

## Author

[Benjamin Toll](https://benjamintoll.com)

[`trivial`]: https://github.com/btoll/trivial

