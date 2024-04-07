
Project Structure

```
├── README.md
├── cmd
│   └── main.go
├── configs
│   ├── default.toml
│   └── lib_test.go
├── internal
│   ├── models
│   │   └── video.go
│   |── boot
│   |    └── boot.go
│   │   └── config.go
│   |── repo
│   |    └── video.go
│   └── controllers
│   |    └── app.go
│   └── router
│       └── router.go
│       └── routes.go
|── go.mod
├── go.sum
|-- Dockerfile
|-- docker-compose.yaml
```

Following general go application structure
Contains main.go in cmd file and it contians start of the program
docker-composer contains both db and application
Dockerfile contains go application build commands
Internal contains code logic
1. Boot contains connections with db & env variables with intialization
2. Models contains db objects
3. controllers contains handler logic
4. repo contains db operations
5. router contains router creation & new routes creation
6. main.go contain simple file which do
    i. Initialization of configs and db connections
    ii. calling external api using go routine to continuosly run in async for every 10 sec to get data from youtube api
    iii. creating routing endpoints


-------

How to test

1. Use docker
2. Run docker compose up --build in cmd
3. Once build is done, look for server starup logs and db connection successful logs as well as         continuous logs for async api and their info storage logs in db.
4. Once app started running, try below apis
    i. For fetch: http://localhost:8080/fetch?page=1&pageSize=100
        here difine page and pagesize
    ii. http://localhost:8080/search?title=Cricbuzz
        here difine title or description with partial seach from the results of fetch api
