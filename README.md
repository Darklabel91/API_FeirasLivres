# API_FeirasLivres
API developed to use data made availebe in [Feiras Livres](http://www.prefeitura.sp.gov.br/cidade/secretarias/upload/chamadas/feiras_livres_1429113213.zip) and apply a simple CRUD.

## Steup
- Run docker and get hostname
```
docker-compose up
docker-compose exec postgres sh
hostname -i
```
- Open [pgAdmin](http://localhost:54321) and create server. Username and Password can be found [here](https://github.com/Darklabel91/API_FeirasLivres/blob/main/docker-compose.yml)
```
right click on "Servers->Register->Server..."
give it a name and in "Connection" set Hostname/address with the hostname given by docker
add login and pass
```
- To run API just make sure that docker image in up and run go main.go
```
docker-compose up
go run main.go
```

## API
This project car create, delete, update a fair and also can read the data filtering by: district, region, name, neighbourhood and ID.
Every method expect Status:200 and JSON content-type as show bellow:

 Req  | Endpoint | Description | Expected Return |
|--------|----------|-----------|---------|
|POST  | /api/fairs | Create a new fair in database | Status:200 - JSON |
|GET | /api/fairs | Read all fairs from database | Status:200 - JSON  | 
|PUT  | /api/fairs/id/{searchParam} | Update a fair by given id | Status:200 - JSON |
|DELETE |/api/fairs/id/{searchParam}  | Delete a fair by given id |Status: 200  |
|GET  | /api/fairs/district/{searchParam} | Read fairs with given distric | Status:200 - JSON  |
|GET  | /api/fairs/region/{searchParam} | Read fairs with given region | Status:200 - JSON  |
|GET  | /api/fairs/name/{searchParam} | Read fairs with given fair name | Status:200 - JSON  |
|GET  | /api/fairs/neighbourhood/{searchParam}  | Read fairs with given neighbourhood | Status:200 - JSON  |
|GET  | /api/fairs/id/{searchParam} | Read fairs with given id | Status:200 - JSON  |

### Details
 Details  | Descripion |
|--------|----------|
|Database | Postgres  |
|content-type | application/json  |
|[Script](https://github.com/Darklabel91/API_FeirasLivres/blob/main/database/migration.go) | Imports sigle .csv on database |
|[Test](https://github.com/Darklabel91/API_FeirasLivres/blob/main/main_test.go) | Test all 9 possible endpoints |
|[Log](https://github.com/Darklabel91/API_FeirasLivres/blob/main/logs.txt)  | Simple .txt log file  |


## Endpoint Examples
- GET - http://localhost:8000/api/fairs/id/810
```
[
    {
        "id": 810,
        "longitude": "-46776674",
        "latitude": "-23674371",
        "set_cen": "355030819000051",
        "area_p": "3550308005232",
        "cod_dist": "19",
        "district": "CAPAO REDONDO",
        "cod_sub_pref": "17",
        "sub_pref": "CAMPO LIMPO",
        "region_Five": "Sul",
        "region_Eight": "Sul 2",
        "name_fair": "JARDIM JANGADEIRO",
        "record": "3100-3",
        "street": "RUA ROSARIO SCAMARTI",
        "neighbourhood": "JD JANGADEIRO",
        "reference": "TV RUA ROSARIO ESCARNADI"
    }
]
```

- POST - http://localhost:8000/api/fairs
```
BODY
    {
        "longitude": "-46776674",
        "latitude": "-23674371",
        "set_cen": "355030819000051",
        "area_p": "3550308005232",
        "cod_dist": "19",
        "district": "CAPAO REDONDO",
        "cod_sub_pref": "17",
        "sub_pref": "CAMPO LIMPO",
        "region_Five": "Sul",
        "region_Eight": "Sul 2",
        "name_fair": "JARDIM JANGADEIRO",
        "record": "3100-3",
        "street": "RUA ROSARIO SCAMARTI",
        "neighbourhood": "JD JANGADEIRO",
        "reference": "TV RUA ROSARIO ESCARNADI"
    }
Response
{
    "id": 888,
    "longitude": "-46776674",
    "latitude": "-23674371",
    "set_cen": "355030819000051",
    "area_p": "3550308005232",
    "cod_dist": "19",
    "district": "CAPAO REDONDO",
    "cod_sub_pref": "17",
    "sub_pref": "CAMPO LIMPO",
    "region_Five": "Sul",
    "region_Eight": "Sul 2",
    "name_fair": "JARDIM JANGADEIRO",
    "record": "3100-3",
    "street": "RUA ROSARIO SCAMARTI",
    "neighbourhood": "JD JANGADEIRO",
    "reference": "TV RUA ROSARIO ESCARNADI"
}
```
