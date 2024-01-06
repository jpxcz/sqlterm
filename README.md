# SQLTerm

Connect in a easy way to the database you need without having to memorize the credentials.

## Instalation

### Credentials storage

```
$ mkdir ~/.config/sqlterm
$ cd ~/.config/sqlterm
$ touch databases.json
```

#### Credentials File Format

```json
{
  "databases": [
    {
      "shortname": "<short name to be displayed>",
      "username": "<database username>",
      "hostname": "<database hostname or ip",
      "password": "<database password>",
      "port": "<database port>"
    }
    ...
  ]
}

```

## Quick Start
```sh
$ git clone https://github.com/jpxcz/sqlterm
# in the future we will need to get the dependencies `go get .`
$ cd sqlterm
$ go build -o sqlterm ./cmd/main.go 
$ ./sqlterm
```

## Limitations 
Currently only supported MySQL.

## TODOS:
- [ ] Add BubbleTea for a better UX interface 
- [ ] Allow to send a sql files to the connection
- - [ ] Allow to send sql files to multiple databases
- [ ] Command to do dump of one or multiple databases
- [ ] Allow support to multiple databases types
