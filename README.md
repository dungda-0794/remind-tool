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

- setting cron
```sh
crontab -e

*/30 * * * * cd /remindTool && ./main >> /remindTool/logs/tool.log 2>&1
```

