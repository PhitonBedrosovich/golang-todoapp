golang todoapp

Архитеутура:

golang-todoapp/
├── .env.example                  # Пример переменных окружения для todoapp и auth-gateway
├── .gitignore                    # Игнорирование logs/, .env, бинарников, артефактов сборки
├── docker-compose.yml            # Запуск postgres + todoapp + auth-gateway
├── go.mod                        # Основной модуль (go 1.26.4)
├── go.sum                        # Контрольные суммы основного модуля
├── Makefile                      # build/run/test/lint/swagger/tidy/migrate
├── README.md                     # Описание проекта, переменные окружения, запуск
│
├── auth-gateway/
│   ├── Dockerfile                # Сборка образа auth-gateway
│   ├── go.mod                    # Модуль auth-gateway (go 1.22)
│   ├── go.sum                    # Контрольные суммы auth-gateway
│   ├── main.go                   # Точка входа, логика прокси, JWT, cookie, embed login.html
│   └── public/
│       └── login.html            # HTML-страница логина (используется через go:embed)
│
├── cmd/
│   └── todoapp/
│       └── main.go               # Инициализация DI, конфиг, logger, pool, HTTP server, роутинг
|       └── Dockerfile            # Сборка образа
|       └── Caddyfile
│
├── docs/
│   └── docs.go                   # Автогенерация Swagger (swaggo/swag)
|   └── swagger.json
|   └── swagger.yaml
│
├── internal/
│   ├── core/
│   │   ├── config/
│   │   │   └── config.go         # Конфигурация приложения (TIME_ZONE)
│   │   │
│   │   ├── domain/
│   │   │   ├── nullable.go       # Generic тип Nullable[T] для three-state logic
│   │   │   ├── statistics.go     # Доменная модель Statistics
│   │   │   ├── task.go           # Доменная модель Task, TaskPatch, валидация, ApplyPatch
│   │   │   ├── uninitialized.go  # Константы UninitializedID/Version
│   │   │   └── user.go           # Доменная модель User, UserPatch, валидация, ApplyPatch
│   │   │
│   │   ├── errors/
│   │   │   └── common.go         # Стандартные ошибки (ErrNotFound, ErrInvalidArgument, ErrConflict)
│   │   │
│   │   ├── logger/
│   │   │   ├── config.go         # Конфиг логгера (LOGGER_LEVEL, LOGGER_FOLDER)
│   │   │   └── logger.go         # Zap logger, file log, context helpers
│   │   │
│   │   ├── repository/
│   │   │   └── postgres/
│   │   │       └── conn/
│   │   │           ├── errors.go # Ошибки БД (ErrNoRows, ErrViolatesForeignKey, ErrUnknown)
│   │   │           ├── pool.go   # Интерфейсы Pool, Rows, Row, CommandTag
│   │   │           └── pgx/
│   │   │               ├── adapters.go   # Адаптеры pgx под интерфейсы, mapErrors
│   │   │               ├── config.go     # Конфиг подключения к Postgres
│   │   │               └── pool.go       # Реализация пула соединений pgxpool
│   │   │
│   │   └── transport/
│   │       └── http/
│   │           ├── middleware/
│   │           │   ├── common.go     # CORS, RequestID, Logger, Trace, Panic
│   │           │   ├── dummy.go      # Пример middleware
│   │           │   └── middleware.go # Тип Middleware и ChainMiddleware
│   │           │
│   │           ├── request/
│   │           │   ├── decode.go      # DecodeAndValidateRequest (JSON + validator)
│   │           │   ├── path_values.go # Парсинг path параметров
│   │           │   └── query_params.go # Парсинг query параметров (int, date)
│   │           │
│   │           ├── response/
│   │           │   ├── errors.go      # Структура ErrorResponse
│   │           │   ├── handler.go     # HTTPResponseHandler (JSON, HTML, Error, Panic)
│   │           │   └── writer.go      # ResponseWriter wrapper для перехвата status code
│   │           │
│   │           ├── server/
│   │           │   ├── config.go      # Конфиг HTTP сервера (ADDR, SHUTDOWN_TIMEOUT, ORIGINS)
│   │           │   ├── route.go       # Структура Route
│   │           │   ├── router.go      # APIVersionRouter (v1, v2...)
│   │           │   └── server.go      # HTTPServer, graceful shutdown, swagger registration
│   │           │
│   │           └── types/
│   │               └── nullable.go    # HTTP-обертка над domain.Nullable с UnmarshalJSON
│   │
│   └── features/
│       ├── statistics/
│       │   ├── repository/
│       │   │   └── postgres/
│       │   │       ├── get_tasks.go      # Запрос задач для статистики (динамический SQL)
│       │   │       ├── models.go         # TaskModel, маппинг в domain
│       │   │       └── repository.go     # StatisticsRepository struct
│       │   │
│       │   ├── service/
│       │   │   ├── get_statistics.go     # Бизнес-логика расчета статистики
│       │   │   └── service.go            # StatisticsService interface & struct
│       │   │
│       │   └── transport/
│       │       └── http/
│       │           ├── get_statistics.go # Handler GetStatistics, DTO, Swagger docs
│       │           └── transport.go      # Routes registration
│       │
│       ├── tasks/
│       │   ├── repository/
│       │   │   └── postgres/
│       │   │       ├── create_task.go    # INSERT с RETURNING, обработка FK
│       │   │       ├── delete_task.go    # DELETE по ID
│       │   │       ├── get_task.go       # SELECT по ID
│       │   │       ├── get_tasks.go      # Список с пагинацией и фильтрами
│       │   │       ├── models.go         # TaskModel, маппинг
│       │   │       ├── patch_task.go     # UPDATE с optimistic locking (version)
│       │   │       └── repository.go     # TasksRepository struct
│       │   │
│       │   ├── service/
│       │   │   ├── create_task.go        # Валидация + сохранение
│       │   │   ├── delete_task.go        # Удаление
│       │   │   ├── get_task.go           # Получение одной задачи
│       │   │   ├── get_tasks.go          # Валидация пагинации + список
│       │   │   ├── patch_task.go         # Get -> ApplyPatch -> Save (optimistic locking)
│       │   │   └── service.go            # TasksService interface & struct
│       │   │
│       │   └── transport/
│       │       └── http/
│       │           ├── create_task.go    # POST /tasks
│       │           ├── delete_task.go    # DELETE /tasks/{id}
│       │           ├── dto_common.go     # TaskDTOResponse, маппинг
│       │           ├── get_task.go       # GET /tasks/{id}
│       │           ├── get_tasks.go      # GET /tasks
│       │           ├── patch_task.go     # PATCH /tasks/{id}, PatchTaskRequest
│       │           └── transport.go      # Routes registration
│       │
│       ├── users/
│       │   ├── repository/
│       │   │   └── postgres/
│       │   │       ├── create_user.go    # INSERT
│       │   │       ├── delete_user.go    # DELETE
│       │   │       ├── get_user.go       # SELECT by ID
│       │   │       ├── get_users.go      # List с пагинацией
│       │   │       ├── models.go         # UserModel
│       │   │       ├── patch_user.go     # UPDATE с version
│       │   │       └── repository.go     # UsersRepository struct
│       │   │
│       │   ├── service/
│       │   │   ├── create_user.go        # Валидация + сохранение
│       │   │   ├── delete_user.go        # Удаление
│       │   │   ├── get_user.go           # Получение пользователя
│       │   │   ├── get_users.go          # Список с пагинацией
│       │   │   ├── patch_user.go         # Get -> ApplyPatch -> Save
│       │   │   └── service.go            # UsersService interface & struct
│       │   │
│       │   └── transport/
│       │       └── http/
│       │           ├── create_user.go    # POST /users
│       │           ├── delete_user.go    # DELETE /users/{id}
│       │           ├── dto_common.go     # UserDTOResponse, маппинг
│       │           ├── get_user.go       # GET /users/{id}
│       │           ├── get_users.go      # GET /users
│       │           ├── patch_user.go     # PATCH /users/{id}, PatchUserRequest
│       │           └── transport.go      # Routes registration
│       │
│       └── web/
│           ├── repository/
│           │   └── file_system/
│           │       ├── get_file.go       # Чтение файла с диска
│           │       └── repository.go     # WebRepository struct
│           │
│           ├── service/
│           │   ├── get_main_page.go      # Логика получения index.html (PROJECT_ROOT)
│           │   └── service.go            # WebService interface
│           │
│           └── transport/
│               └── http/
│                   ├── get_main_page.go  # Handler отдачи HTML
│                   └── transport.go      # Route "/"
│
├── migrations/
│   └── 00001_init.up.sql              # Создание схемы todoapp, таблиц users/tasks, ограничений
|   └── 00001_init.down.sql            # Удаление схемы todoapp, таблиц users/tasks, ограничений
│
└── public/
    └── index.html                # Главная страница web-приложения (PROJECT_ROOT/public/index.html)