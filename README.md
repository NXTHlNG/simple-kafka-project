# Лабораторная работа 3 - Kafka
### б1-ИФСТ-31 Жаворонков Артём

## Описание
**Выполнены все указанные требования.**
### API Service
Предоставляет HTTP API для доступа извне Docker-сети. 

>Работает на порту 8000

HTTP API содержит следующие конечные точки:

- Добавления новой порции данных. Порция данных отправляется в Kafka.
- Поиск по добавленным порциям данных. Поиск производится путем обращения к конечной точке HTTP API Data Service.
- Получение отчетов на основе добавленных данных. Получение отчетов производится путем обращения к конечной точке HTTP API Data Service.

### Data Service
Получает из Kafka порции данных для записи, записывает их в БД.

Предоставляет HTTP API со следующими конечными точками:

- Поиск по добавленным порциям данных. Поиск производится путем выборки из БД.
- Получение отчетов на основе добавленных данных. Получение отчетов производится путем выборки из БД с использованием агрегации и т.п.

## Для запуска:
> docker-compose up --build -d

## Порты
- API Service - 8000

## Конечные точки
### POST /api/video
Добавить видео

### GET /api/videos
Получить все видео

### GET /api/videos/top10
Получить Топ-10 видео по просмотрам

### GET /api/videos/tags_rate
Получить количество использований тэгов

### POST /api/author
Добавить автора

### GET /api/authors
Получить всех авторов

### GET /api/authors/top10
Получить Топ-10 авторов по просмотрам на видеороликах

## Схемы

### Author

```
{
  id: ObjectId,
  name: String,
  email: String,
  avatar: String
}
```

### Video

```
{
  id: ObjectId,
  title: String,
  description: String,
  views: Number,
  likes: Number,
  dislikes: Number,
  tags: [String],
  author_id: ObjectId,
  duration: Number,
  comments: [
    {
      author_id: ObjectId,
      content: String,
      replies: [
        {
          author_id: ObjectId,
          content: String,
        }
      ]
    }
  ]
}
```