/taco-delivery
   /user-service          // Микросервис для работы с пользователями
      /cmd
         main.go          // Точка входа
      /config             // Конфигурация
      /handlers           // HTTP handlers
         user_handler.go
      /repositories       // Работа с БД (Postgres)
         postgres_repo.go
      /kafka              // Kafka consumers/producers
         kafka_handler.go
      go.mod

   /order-service         // Микросервис для работы с заказами
      /cmd
         main.go
      /config
      /handlers
         order_handler.go
      /repositories       // Работа с БД (MongoDB)
         mongo_repo.go
      /kafka
         kafka_handler.go
      go.mod

   /delivery-service      // Микросервис для работы с доставкой
      /cmd
         main.go
      /config
      /handlers
         delivery_handler.go
      /repositories       // Работа с БД (Postgres)
         postgres_repo.go
      /kafka
         kafka_handler.go
      go.mod

   /common                // Общие пакеты и утилиты
      /models             // Общие структуры данных
      /logger             // Логирование
      /kafka              // Kafka клиент и конфигурация
      /db                 // Базовые настройки подключения к БД
      go.mod

   /gateway               // API Gateway для взаимодействия с веб-клиентами
      /cmd
         main.go
      /routes             // Роутинг REST API
         routes.go
      /services           // Клиенты для микросервисов
         user_client.go
         order_client.go
         delivery_client.go
      go.mod
