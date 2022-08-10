# CoffeeShop-backend

Backend-сервер для проекта CoffeeShop

На данном этапе сервер отрабатывает два GET запроса, отправляя JSON-файлы с общей информацией о заведении (открыто/закрыто, часы работы, телефон, название и т.д) и меню соответственно.

Также отрабатывает POST запрос на приготовление напитка - принимает название напитка: готовит или не готовит заказ в зависимости от наличия продуктов (информация о них хранится в отдельном JSON-файле).