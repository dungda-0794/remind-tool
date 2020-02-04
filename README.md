# remind tool
Remind

### With docker & docker-compose

```sh

## Outsite docker
docker-compose up -d
docker-compose exec remind bash


## Inside docker
dep ensure -update

go run main.go

## Build tool
go build main.go
```

### config and run

- Setting file task.yml file
```yml
job: "Title Job"
time: "Time notify"
message: "Message notify"
```

- Setting crontab
```sh
crontab -e

*/30 * * * * export DISPLAY=:0 && cd ~path/remindTool && ./main >> ~path/remindTool/logs/tool.log 2>&1
```

