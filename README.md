# Ubiquitous octo

### How to install and run?

Get the git checkout of ***Ubiquitous octo*** module and run in development mode.

```sh
$ git clone https://github.com/ctomatis/ubiquitous-octo-guide.git
$ cd ubiquitous-octo-guide
$ go run main.go
# Or run...
$ make dev
```

Or just clone this repository, build and then run the Dockerize app.

```sh
$ git clone https://github.com/ctomatis/ubiquitous-octo-guide.git
$ cd ubiquitous-octo-guide
$ docker-compose up
# Or run...
$ make dockerize
```

### Testing with CURL
```sh
$ curl --location 'http://localhost/tasks/' \
--header 'Content-Type: application/json' \
--header 'Authorization: vmo'
```