# DOCKER / GOLANG CHALLENGE

This project uses a Docker_compose file to do the following.

1. Start a Postgress container.
2. Do a multi Stage compile down to a scratch go container.
3. Starts a pgadmin container with connection to the Postgress server 
4. Demonstrate unit test goLang container


Once the Docker_compose has fully deployed, the golang application will run a todo application with REST API for create, delete, list.  All of these API functions are to be unit tested.

NOTE:  the goLang app will pre populte the DB with 4 records.

TODO:
1. add unit test
2. Create a Travis-CI pipline for test automation on source code push to remote GIT repo
3. Fix the init pg stuff.

# DELIVERABLE
ROOT
    docker-compose.yaml 
    README.md

init
    // for pre populating a PG db.  how ever I am not using it for now.

***

# ROUTINE

1. clear out Images...

```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker image prune -f
```
TO shut down the Docker containers..



2. Start up the Docker Container.

from the root directory...
..* Docker_compose up

To shut down the Docker containers.
..* Docker_compose down



3. PGadmin

http://localhost:5555/

[Duck Duck Go](https://duckduckgo.com)
[PGADMIN](http://localhost:5555/)
### To log into the PGADMIN site:
      PGADMIN_DEFAULT_EMAIL: pgadmin4@pgadmin.org
      PGADMIN_DEFAULT_PASSWORD: admin

### To log into the Server
      POSTGRES_USER: postgres
      POSTGRES_DBNAME: db_amex01
      POSTGRES_PASS: postgres

The pgadmin Db is to_dos.
The sql statement is..
select * from to_dos
< insert instructions to do a SQL statement >

To add the Posgres server, use hostname postgres, port 5432.



## GOLANG ENDPOINTS
In additon to PGADMIN, there is also a GoLang REST endpoint

[TODO FULL LIST](http://localhost:8000/api01/listalltodo)
http://localhost:8000/api01/listalltodo

http://localhost:8000/api01/gettodo/100

http://localhost:8000/

http://localhost:8000/api01/initialize


http://localhost:8000/api01/createtodo/100/another%20todoooddododo

http://localhost:8000/api01/deletetodo/100

http://localhost:8000/api01/closetodo/100

http://localhost:8000/api01/Updatetodo/100/all%20done



# 









 




 login to pgadmin 
 http://localhost:5555/login?next=%2F
 
 login as user pgadmin4@pgadmin.org, pwd: admin. 
 
 
 


 docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker image prune -f


