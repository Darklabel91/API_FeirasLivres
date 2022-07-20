# API_FeirasLivres

## Dependencies
- [GORM](https://gorm.io)
- [DRIVER](https://github.com/go-gorm/postgres)
- [MUX](https://github.com/gorilla/mux)

## Run
To run the API locally run docker image and set postgres Server.
Next, just run go run main.go and you are good to go.
### Data Base
#### Docker
Run docker compose and get the hostname.
'''
docker-compose up
docker-compose exec postgres sh
hostname -i
'''

#### Postgres
With hostname access [pgAdmin](http://localhost:54321) and create sever, login information can be found in [LoginInfo](https://github.com/Darklabel91/API_FeirasLivres/blob/main/docker-compose.yml).
Create the server
'''
right click on "Servers->Register->Server..."
give it a name and in "Connection" set Hostname/address with the hostname given by docker
username and password can be found on [LoginInfo](https://github.com/Darklabel91/API_FeirasLivres/blob/main/docker-compose.yml)
'''

### API
go run main.go
