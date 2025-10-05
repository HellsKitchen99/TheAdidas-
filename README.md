# TheAdidas-

**TheAdidas-** — это API-сервис на Go, который принимает календарные события, определяет их геолокацию и вычисляет прогноз погоды для каждого события с помощью **OpenWeather** и **2GIS Geocoder**.

---

## Запуск проекта

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/HellsKitchen99/TheAdidas-
   cd TheAdidas-
   ```

2. Настроить `.env`:
   ```env
   APIKEY=<2GIS_API_KEY>
   WEATHERKEY=<OPENWEATHER_KEY>
   ```

3. Настроить `config.json`:
   ```json
   {
     "ip": "ваш IP",
     "port": "ваш Port"
   }
   ```

4. Запустить проект:
   ```bash
   go run main.go
   ```

---

## Структура проекта

```
TheAdidas-/
├── Controllers/             # HTTP-контроллеры (эндпоинты)
│
├── Models/                  # Модели данных (структуры и DTO)
│   ├── config.go
│   ├── requestFromIlya.go
│   ├── ResponseFromGeo.go
│   ├── ResponseFromOpen.go
│   └── ResponseToIlya.go
│
├── Service/                 # Логика приложения и утилиты
│   ├── .env                 # переменные окружения (API ключи и настройки)
│   ├── envWork.go           # работа с .env файлом
│   └── service.go           # основная бизнес-логика (запросы, погода, координаты)
│
├── .gitignore               # исключения для Git
├── config.json              # конфигурационный файл проекта
├── go.mod                   # зависимости Go-модуля
├── go.sum                   # контрольные суммы зависимостей
├── hourly_weather.csv       # CSV с погодными данными
├── main.go                  # точка входа в приложение
└── README.md                # описание проекта
```

---

## Технологии

| Компонент | Используется |
|------------|---------------|
| Язык | Go 1.23.5 |
| Фреймворк | Gin |
| Работа с .env | Godotenv |
| Внешние API | 2GIS Geocoder, OpenWeather |

---

## Эндпоинты

### `POST /adidas`

**Описание:**  
Принимает календарные события с фронтенда, вычисляет прогноз погоды для каждого события с учётом координат и времени.

**Пример запроса:**
```json
{
  "today": [
    {
      "name": "Meeting",
      "location": "Moscow",
      "start_event": "2025-10-05T10:00:00Z",
      "end_event": "2025-10-05T11:00:00Z"
    }
  ],
  "tomorrow": [
    {
      "name": "Gym",
      "location": "Saint Petersburg",
      "start_event": "2025-10-06T08:00:00Z",
      "end_event": "2025-10-06T09:30:00Z"
    }
  ]
}
```

**Пример ответа:**
```json
{
  "today": [
    {
      "name": "Meeting",
      "user_location": {"lat": 55.7558, "lon": 37.6173},
      "event_location": {"lat": 59.9343, "lon": 30.3351},
      "start_time": "2025-10-05T10:00:00Z",
      "end_time": "2025-10-05T11:00:00Z",
      "weather": {
        "temp": 15,
        "feels_like": 14,
        "condition": "Clouds",
        "wind_speed": 5,
        "wind_dir": "NW"
      }
    }
  ],
  "tomorrow": [...]
}
```

## 🔮 Планы развития 

---
