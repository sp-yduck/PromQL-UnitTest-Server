# Promtool Server

Simple Web Application/API to run unit test for PromQL(alerting rules & recording rules) by using prometheus/promtool.

## Execute promtool

Endpoint looks : `/promtool/<command>/<subcommand>/<flag>`

You can give any arguments & flags to promtool as url params.

### Endpoint example

If you would like exec `promtool --version`, endpoint is `/promtool/--version`

If you would like exec `promtool check --help`, endpoint is `/promtool/check/--help`

## Docker build & run

Do not forget to publish 1323 port.

```bash
# example
docker build -t promtool-server .
docker run -p 1323:1323 -it promtool-server
```

## License

MIT License. see [LICENSE](./LICENSE)

## Author

Teppei Sudo (https://github.com/sp-yduck)