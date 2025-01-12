# Конфигурация
* адрес и порт запуска сервиса: переменная окружения ОС `RUN_ADDRESS` или флаг `-a`;  
* адрес подключения к базе данных: переменная окружения ОС `DATABASE_URI` или флаг `-d`;  
* адрес системы расчёта начислений: переменная окружения ОС `ACCRUAL_SYSTEM_ADDRESS` или флаг `-r`;  
* секрет для jwt токена флаг `-s`.


# Локальная разработка

Запустить БД:

```bash
docker compose up -d
```

Создание файла миргации:

```bash
make create-migration n=migration_name
```
Файл появится в  `./internal/db/migrations/`.

Выполнить миграции:

```bash
make migrations-up
```

Запустить линтер:

```bash
make golangci-lint-run
```

