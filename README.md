[View the documentation](https://pkg.go.dev/github.com/btoll/trivial)

## Test the `/query` Endpoint

To send a single question to the game players, you can run the following query (of course, replace the `URL` with your own) :

```
$ curl -XGET -H "X-TRIVIA-APIKEY: GwCcn6VZntwnI7qp1kly" \
    --data "What year did the Beatles play Budokan?,B,50,1965,1968,1970" \
    127.0.0.1:3000/query
```

An easier way would be to concatenate all of the `csv` files together and then create a variable which will be incremented to pull the questions from the file line-by-line.

```
$ cat *.csv > game.csv
```

Set the variable to increment:

```
$ i=0
```

Push questions through to the game players by issuing a single chained command that first updates the variable and then uses it to pick the respective line from the `game.csv` file:

```
$ i=$((i+1)) && curl -XGET -H "X-TRIVIA-APIKEY: GwCcn6VZntwnI7qp1kly" \
    --data "$(awk 'NR=='$i'' game.csv)" \
    127.0.0.1:3000/query
```

> If using a self-signed TLS certificate, pass the `-k` or `--insecure` switch so `curl` will disable strict certificate checking.
>
> ```
> $ curl --insecure -H "X-TRIVIA-APIKEY: wsj76zSOyxNU93uW8LXhrHms4" \
>     --data "$(awk 'NR=='$i'' questions/roman_history.csv)" \
>     127.0.0.1:3000/query
> ```

## License

[GPLv3](COPYING)

## Author

Benjamin Toll

