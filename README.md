# Задание по Docker swarm

# Stack
+ REST API
+ PostgreSQL
+ Nginx
+ Docker, Docker Swarm

# Getting Started with a swarm
1. `git clone https://github.com/evgkhm/task-for-docker`
2. `cd task-for-docker`
3. `make run-swarm`

# Getting Started with a single host
1. `make run`

# API
| Endpoint       | API method |                         Description |
|----------------|:----------:|------------------------------------:|
| /api/time      |    POST    |            Создание временной метки |
| /api/last_time |    GET     | Получение последней временной ветки |


# Example
1. Создание временной метки (post запрос на nginx)
curl -X POST http://localhost:8080/api/time
2. Получение (get запрос на nginx)
curl http://localhost:8080/api/last_time
