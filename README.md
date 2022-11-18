# Course Project

Hiking microservice description...

## install

    1. add folders:
        cmd/
        internal/
        pkg/
        proto/

    2. add go mod, add module requirements and sums:
        cd hike
        go mod init example.com/m/v2[создает новый модуль, инициализируя файл go.mod, описывающий модуль. Вначале он только добавит путь к модулю и версию Go в файл go.mod]
        go mod tidy[обеспечивает соответствие файла go.mod исходному коду модуля]

### useful

        go clean -modcache[Эта команда используется для очистки кеша модов, который хранится в $GOPATH/pkg/mod. Эта команда используется для удаления установленных пакетов. Флаг -modcache удаляет весь кеш загрузки модуля, включая распакованный исходный код версий зависимостей.]

#### LeetCode Tasks

[LeetCode Task 1 - 3sum](https://leetcode.com/problems/3sum)

[LeetCode Task 2 - permutations](https://leetcode.com/problems/permutations)
