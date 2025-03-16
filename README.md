# Проект на golang 



## Name
tender-app

## Description
Этот проект представляет собой бэкенд-приложение, которое позволит бизнесу создать тендер на оказание каких-либо услуг. А пользователи/другие бизнесы будут предлагать свои выгодные условия для получения данного тендера.

## Инструкция по запуску
Запуск с помощью Docker Compose: для сборки и запуска контейнеров в корневой директории проекта выполните команду:

"docker-compose up --build"

Запускается с порта :8080

Для деплоя на Kubernetes(использовала данные из файлов задания) выполнить команду "export KUBECONFIG=./cnrprod1725728503-team-77203.kube.config".

## Стек
Проект написан на языке Golang + Postgres + использованы такие golang библиотеки, как: fiber и sqlx.

Базы данных в файлах migration\0001_init.sql - заполнила по 5 значений таблицы employee, organization и organization_responsible
                     migration\0002_fill.sql - заполнила по 5 значений таблицы tenders и bids

## Функционал

Реализованы следующие эндпоинты:

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET   | `/ping` | Проверка доступности сервера |
| GET   | `/tenders` | Получение списка тендеров |
| POST  | `/tenders/new` | Создание нового тендера |
| GET   | `/tenders/my` | Получение тендеров пользователя |
| GET   | `/tenders/{tenderId}/status` | Получение текущего статуса тендера |
| PUT   | `/tenders/{tenderId}/status` | Изменение статуса тендера |
| PATCH | `/tenders/{tenderId}/edit` | Редактирование тендера |
| PUT   | `/tenders/{tenderId}/rollback/{version}` | Откат версии тендера |
| POST  | `/bids/new` | Создание нового предложения |
| GET   | `/bids/my` | Получение списка ваших предложений |
| GET   | `/bids/{tenderId}/list` | Получение списка предложений для тендера |
| GET   | `/bids/{bidId}/status` | Получение текущего статуса предложения |
| PUT   | `/bids/{bidId}/status` | Изменение статуса предложения |
| PATCH | `/bids/{bidId}/edit` | Редактирование параметров предложения |
| PUT   | `/bids/{bidId}/submit_decision` | Отправка решения по предложению |
| PUT   | `/bids/{bidId}/feedback` | Отправка отзыва по предложению |
| PUT   | `/bids/{bidId}/rollback/{version}` | Откат версии предложения |
| GET   | `/bids/{tenderId}/reviews` | Просмотр отзывов на прошлые предложения |

## Тестирование
Тестирование провела качественное, вручную в приложении для тестирования Postman, примеры запросов для тестирования:

http://localhost:8080/api/tenders/new?organizationId=1&creatorId=1 -Создание нового тендера
{
    "name": "Тендер на строительство",
    "description": "Строительство нового офисного здания",
    "serviceType": "Construction"
}
http://localhost:8080/api/tenders/my?username=asmith - Получение тендеров пользователя
http://localhost:8080/api/tenders/4/rollback/2 - Откат версии тендера
http://localhost:8080/api/bids/my?userId=2 - Получение списка ваших предложений
http://localhost:8080/api/bids/13/status -Изменение статуса предложения 
{
    "status": "PUBLISHED"
}
http://localhost:8080/api/bids/2/feedback- Отправка отзыва по предложению
{
    "comment": "Отличное предложение!",
    "rating": 4,
    "organization_id": 1,
    "reviewer": "jdoe"
}


