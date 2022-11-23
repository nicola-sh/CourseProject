# Course Project

Hiking microservice description...

## install

- add folders:
```console
cmd/
internal/
pkg/
proto/
```
- add go mod, add module requirements and sums:
```console
'go mod init example.com/m/v2' to initialize a v2 module
```
создает новый модуль, инициализируя файл go.mod, описывающий модуль. 
Вначале он только добавит путь к модулю и версию Go в файл go.mod

```console
go mod tidy
```
обеспечивает соответствие файла go.mod исходному коду модуля

- install protoc and rpc: [grpc.io](https://grpc.io/)
```console
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### docker
load image postgres:latest
create container postgres:latest


### Migration
go db postgres migrate to db/migration. It's create init_shema up and down:
```console
migrate create -ext sql -dir db\migration -seq init_shema
```
then copy from .sql to .up
and
DROP TABLE IF EXISTS nameTable; to .down

start migration with working docker postgresql:
```console
docker run --name hike-postgres -p 5432:5432 -e POSTGRES_USER=root -e PASSWORD=root -d postgres:latest
docker ps
docker ps -a
docker stop name_container
```

доступ к оболочке:
```console
docker exec -it name_container /bin/sh
```
внутри оболочки:
```console
pwd
ls -l
createdb --username=user_name --owner=owner_name db_name [createdb --username=root --owner=root DoHike]
psql DoHike[check that's db working]
dropdb db_name[to delete db]
exit
```
также можно делать все через докер(в т.ч. создать БД):
```console
docker exec -it hike-postgres createdb --username=root --owner=root dohike
docker exec -it hike-postgres psql dohike
makefile migrate
```

###### sqlc

use sqlc crud [sqlc.dev](https://sqlc.dev/)
-install sqlc:
```console
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```
-generate go code from sql

```console
sqlc init
sqlc generate
```

-sqlc configure

[sqlc info](https://github.com/kyleconroy/sqlc/blob/v1.16.0/docs/tutorials/getting-started-postgresql.md)
edit sqlc.yaml
make sqlc

###### pq lib + testify for unit test

go get github.com/lib/pq


###### Unit Tests

*test.go


##### useful

change go env -w GO111MODULE=auto to go env -w GO111MODULE=off

```console
go clean -modcache
```
Эта команда используется для очистки кеша модов, 
который хранится в $GOPATH/pkg/mod. 
Эта команда используется для удаления установленных пакетов. 
Флаг -modcache удаляет весь кеш загрузки модуля, 
включая распакованный исходный код версий зависимостей.




## LeetCode Tasks

[LeetCode Task 1 - 3sum](https://leetcode.com/problems/3sum)

[LeetCode Task 2 - permutations](https://leetcode.com/problems/permutations)
