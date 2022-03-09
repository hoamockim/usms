##### Get dependencies
```
   make update
```
##### Run internal
```
 make run internal
```
##### Build profile
```
docker build . -t user-profile:latest  --build-arg SERVICE_TYPE=profile
```

#### Build auth
```
docker build . -t auth:latest  --build-arg SERVICE_TYPE=auth
```

#### Build migration
```
docker build . -t user_profile_migration:latest  --build-arg SERVICE_TYPE=migration --build-arg JOBNAME=init_db
```

### Build external
```
docker build . -t user_profile_external:latest --build-arg SERVICE_TYPE=external
```