# Details

Date : 2026-07-27 10:13:49

Directory d:\\ilyap\\GoProjects\\golang-todoapp

Total : 82 files,  3269 codes, 146 comments, 755 blanks, all 4170 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Makefile](/Makefile) | Makefile | 124 | 10 | 13 | 147 |
| [README.md](/README.md) | Markdown | 1 | 0 | 0 | 1 |
| [cmd/todoapp/main.go](/cmd/todoapp/main.go) | Go | 82 | 18 | 16 | 116 |
| [docker-compose.yaml](/docker-compose.yaml) | YAML | 21 | 0 | 3 | 24 |
| [go.mod](/go.mod) | Go Module File | 23 | 0 | 4 | 27 |
| [go.sum](/go.sum) | Go Checksum File | 54 | 0 | 1 | 55 |
| [internal/core/config/config.go](/internal/core/config/config.go) | Go | 30 | 6 | 8 | 44 |
| [internal/core/domain/nullable.go](/internal/core/domain/nullable.go) | Go | 5 | 7 | 5 | 17 |
| [internal/core/domain/statistics.go](/internal/core/domain/statistics.go) | Go | 21 | 0 | 4 | 25 |
| [internal/core/domain/task.go](/internal/core/domain/task.go) | Go | 162 | 2 | 31 | 195 |
| [internal/core/domain/uninitialized.go](/internal/core/domain/uninitialized.go) | Go | 5 | 0 | 1 | 6 |
| [internal/core/domain/user.go](/internal/core/domain/user.go) | Go | 103 | 3 | 22 | 128 |
| [internal/core/errors/common.go](/internal/core/errors/common.go) | Go | 7 | 0 | 2 | 9 |
| [internal/core/logger/config.go](/internal/core/logger/config.go) | Go | 24 | 0 | 9 | 33 |
| [internal/core/logger/logger.go](/internal/core/logger/logger.go) | Go | 73 | 1 | 20 | 94 |
| [internal/core/repository/postgres/conn/errors.go](/internal/core/repository/postgres/conn/errors.go) | Go | 7 | 0 | 2 | 9 |
| [internal/core/repository/postgres/conn/pgx/adapters.go](/internal/core/repository/postgres/conn/pgx/adapters.go) | Go | 47 | 2 | 12 | 61 |
| [internal/core/repository/postgres/conn/pgx/config.go](/internal/core/repository/postgres/conn/pgx/config.go) | Go | 29 | 1 | 9 | 39 |
| [internal/core/repository/postgres/conn/pgx/pool.go](/internal/core/repository/postgres/conn/pgx/pool.go) | Go | 72 | 0 | 16 | 88 |
| [internal/core/repository/postgres/conn/pool.go](/internal/core/repository/postgres/conn/pool.go) | Go | 24 | 0 | 7 | 31 |
| [internal/core/transport/http/middleware/common.go](/internal/core/transport/http/middleware/common.go) | Go | 77 | 6 | 19 | 102 |
| [internal/core/transport/http/middleware/dummy.go](/internal/core/transport/http/middleware/dummy.go) | Go | 17 | 0 | 7 | 24 |
| [internal/core/transport/http/middleware/middleware.go](/internal/core/transport/http/middleware/middleware.go) | Go | 15 | 1 | 6 | 22 |
| [internal/core/transport/http/request/decode.go](/internal/core/transport/http/request/decode.go) | Go | 38 | 0 | 10 | 48 |
| [internal/core/transport/http/request/path\_values.go](/internal/core/transport/http/request/path_values.go) | Go | 28 | 1 | 5 | 34 |
| [internal/core/transport/http/request/query\_params.go](/internal/core/transport/http/request/query_params.go) | Go | 43 | 0 | 9 | 52 |
| [internal/core/transport/http/response/handler.go](/internal/core/transport/http/response/handler.go) | Go | 86 | 3 | 20 | 109 |
| [internal/core/transport/http/response/writer.go](/internal/core/transport/http/response/writer.go) | Go | 25 | 3 | 8 | 36 |
| [internal/core/transport/http/server/config.go](/internal/core/transport/http/server/config.go) | Go | 25 | 0 | 9 | 34 |
| [internal/core/transport/http/server/route.go](/internal/core/transport/http/server/route.go) | Go | 17 | 2 | 5 | 24 |
| [internal/core/transport/http/server/router.go](/internal/core/transport/http/server/router.go) | Go | 39 | 2 | 10 | 51 |
| [internal/core/transport/http/server/server.go](/internal/core/transport/http/server/server.go) | Go | 72 | 11 | 21 | 104 |
| [internal/core/transport/http/types/nullable.go](/internal/core/transport/http/types/nullable.go) | Go | 27 | 27 | 11 | 65 |
| [internal/features/statistics/repository/postgres/get\_tasks.go](/internal/features/statistics/repository/postgres/get_tasks.go) | Go | 69 | 1 | 18 | 88 |
| [internal/features/statistics/repository/postgres/models.go](/internal/features/statistics/repository/postgres/models.go) | Go | 34 | 0 | 7 | 41 |
| [internal/features/statistics/repository/postgres/repository.go](/internal/features/statistics/repository/postgres/repository.go) | Go | 12 | 0 | 3 | 15 |
| [internal/features/statistics/service/get\_statistics.go](/internal/features/statistics/service/get_statistics.go) | Go | 58 | 3 | 13 | 74 |
| [internal/features/statistics/service/service.go](/internal/features/statistics/service/service.go) | Go | 24 | 0 | 5 | 29 |
| [internal/features/statistics/transport/http/get\_statistics.go](/internal/features/statistics/transport/http/get_statistics.go) | Go | 72 | 1 | 18 | 91 |
| [internal/features/statistics/transport/http/transport.go](/internal/features/statistics/transport/http/transport.go) | Go | 35 | 0 | 7 | 42 |
| [internal/features/tasks/repository/postgres/create\_task.go](/internal/features/tasks/repository/postgres/create_task.go) | Go | 67 | 2 | 11 | 80 |
| [internal/features/tasks/repository/postgres/delete\_task.go](/internal/features/tasks/repository/postgres/delete_task.go) | Go | 29 | 0 | 7 | 36 |
| [internal/features/tasks/repository/postgres/get\_task.go](/internal/features/tasks/repository/postgres/get_task.go) | Go | 45 | 0 | 11 | 56 |
| [internal/features/tasks/repository/postgres/get\_tasks.go](/internal/features/tasks/repository/postgres/get_tasks.go) | Go | 62 | 0 | 14 | 76 |
| [internal/features/tasks/repository/postgres/models.go](/internal/features/tasks/repository/postgres/models.go) | Go | 34 | 0 | 7 | 41 |
| [internal/features/tasks/repository/postgres/patch\_task.go](/internal/features/tasks/repository/postgres/patch_task.go) | Go | 69 | 0 | 13 | 82 |
| [internal/features/tasks/repository/postgres/repository.go](/internal/features/tasks/repository/postgres/repository.go) | Go | 12 | 0 | 3 | 15 |
| [internal/features/tasks/service/create\_task.go](/internal/features/tasks/service/create_task.go) | Go | 25 | 3 | 4 | 32 |
| [internal/features/tasks/service/delete\_task.go](/internal/features/tasks/service/delete_task.go) | Go | 14 | 0 | 3 | 17 |
| [internal/features/tasks/service/get\_task.go](/internal/features/tasks/service/get_task.go) | Go | 16 | 0 | 5 | 21 |
| [internal/features/tasks/service/get\_tasks.go](/internal/features/tasks/service/get_tasks.go) | Go | 31 | 0 | 6 | 37 |
| [internal/features/tasks/service/patch\_task.go](/internal/features/tasks/service/patch_task.go) | Go | 24 | 0 | 6 | 30 |
| [internal/features/tasks/service/service.go](/internal/features/tasks/service/service.go) | Go | 40 | 2 | 10 | 52 |
| [internal/features/tasks/transport/http/create\_task.go](/internal/features/tasks/transport/http/create_task.go) | Go | 42 | 1 | 13 | 56 |
| [internal/features/tasks/transport/http/delete\_task.go](/internal/features/tasks/transport/http/delete_task.go) | Go | 28 | 1 | 8 | 37 |
| [internal/features/tasks/transport/http/dto\_common.go](/internal/features/tasks/transport/http/dto_common.go) | Go | 34 | 0 | 7 | 41 |
| [internal/features/tasks/transport/http/get\_task.go](/internal/features/tasks/transport/http/get_task.go) | Go | 31 | 0 | 10 | 41 |
| [internal/features/tasks/transport/http/get\_tasks.go](/internal/features/tasks/transport/http/get_tasks.go) | Go | 52 | 4 | 15 | 71 |
| [internal/features/tasks/transport/http/patch\_task.go](/internal/features/tasks/transport/http/patch_task.go) | Go | 80 | 2 | 19 | 101 |
| [internal/features/tasks/transport/http/transport.go](/internal/features/tasks/transport/http/transport.go) | Go | 71 | 0 | 11 | 82 |
| [internal/features/users/repository/postgres/create\_user.go](/internal/features/users/repository/postgres/create_user.go) | Go | 36 | 0 | 10 | 46 |
| [internal/features/users/repository/postgres/delete\_user.go](/internal/features/users/repository/postgres/delete_user.go) | Go | 25 | 0 | 6 | 31 |
| [internal/features/users/repository/postgres/get\_user.go](/internal/features/users/repository/postgres/get_user.go) | Go | 46 | 0 | 11 | 57 |
| [internal/features/users/repository/postgres/get\_users.go](/internal/features/users/repository/postgres/get_users.go) | Go | 53 | 0 | 10 | 63 |
| [internal/features/users/repository/postgres/models.go](/internal/features/users/repository/postgres/models.go) | Go | 20 | 0 | 5 | 25 |
| [internal/features/users/repository/postgres/patch\_user.go](/internal/features/users/repository/postgres/patch_user.go) | Go | 62 | 0 | 10 | 72 |
| [internal/features/users/repository/postgres/repository.go](/internal/features/users/repository/postgres/repository.go) | Go | 12 | 0 | 4 | 16 |
| [internal/features/users/service/create\_user.go](/internal/features/users/service/create_user.go) | Go | 19 | 0 | 6 | 25 |
| [internal/features/users/service/delete\_user.go](/internal/features/users/service/delete_user.go) | Go | 14 | 0 | 3 | 17 |
| [internal/features/users/service/get\_user.go](/internal/features/users/service/get_user.go) | Go | 16 | 0 | 4 | 20 |
| [internal/features/users/service/get\_users.go](/internal/features/users/service/get_users.go) | Go | 37 | 0 | 6 | 43 |
| [internal/features/users/service/patch\_user.go](/internal/features/users/service/patch_user.go) | Go | 24 | 3 | 4 | 31 |
| [internal/features/users/service/service.go](/internal/features/users/service/service.go) | Go | 39 | 0 | 10 | 49 |
| [internal/features/users/transport/http/create\_user.go](/internal/features/users/transport/http/create_user.go) | Go | 34 | 4 | 14 | 52 |
| [internal/features/users/transport/http/delete\_user.go](/internal/features/users/transport/http/delete_user.go) | Go | 28 | 2 | 10 | 40 |
| [internal/features/users/transport/http/dto\_common.go](/internal/features/users/transport/http/dto_common.go) | Go | 23 | 2 | 6 | 31 |
| [internal/features/users/transport/http/get\_user.go](/internal/features/users/transport/http/get_user.go) | Go | 31 | 1 | 10 | 42 |
| [internal/features/users/transport/http/get\_users.go](/internal/features/users/transport/http/get_users.go) | Go | 47 | 0 | 14 | 61 |
| [internal/features/users/transport/http/patch\_user.go](/internal/features/users/transport/http/patch_user.go) | Go | 77 | 5 | 18 | 100 |
| [internal/features/users/transport/http/transport.go](/internal/features/users/transport/http/transport.go) | Go | 70 | 3 | 11 | 84 |
| [migrations/000001\_init.down.sql](/migrations/000001_init.down.sql) | MS SQL | 3 | 0 | 0 | 3 |
| [migrations/000001\_init.up.sql](/migrations/000001_init.up.sql) | MS SQL | 20 | 0 | 4 | 24 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)