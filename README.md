# bibliosphere_gin

Bibliosphere is an advanced library management system developed as a backend service using the Gin Web Framework.

## Initialization Commands

To initialize the database and seed it with test data, run the following commands:

1. Apply database migrations using Goose:

    ```sh
    goose -dir adapters/database/migrations mysql user:password@tcp(127.0.0.1:3306)/bibliosphere?parseTime=true up
    ```

2. Seed the database with test data:

    ```sh
    go run . seed
    ```

Replace `user` and `password` with your MySQL database credentials.
