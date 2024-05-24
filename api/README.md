# E-commerce Api

## Requirements
1. Golang v1.22.1
2. Docker Desktop or Docker engine

## Installation and running
1. Clone repository
    ```bash
    git clone https://github.com/helloWRLDs/go-boilerplate.git
    cd yourrepository
    ```
2. Install dependencies
    ```bash
    go mod tidy
    ```
3. create .env file just like in example [.env.example](./.env.example) in the root and put your data to it
3. Start the database from docker
    ```bash
    docker-compose up -d
    ```
4. Up the migrations to init tables, mock data and admin user.
    ```bash
    goose -dir migrations postgres "postgresql://admin:admin@localhost:5432/mydb?sslmode=disable" up
    ```
5. Run the application
    ```bash
    go run ./cmd/api
    ```

## Api endpoint examples:
Authentication:
- `POST` /api/v1/auth/login: User login
- `POST` /api/v1/auth/register: User registration
- `POST` /api/v1/auth/1/verify/432544: Verify email code by userId
- `POST` /api/v1/auth/1/resend: Resend verification code

Users:
- `GET` /api/v1/users: Get all users
- `GET` /api/v1/users/{id}: Get user by ID
- `PUT` /api/v1/users/{id}: Update user by ID
- `DELETE` /api/v1/users/{id}: Delete user by ID

Admin:
- `POST` /api/v1/admin/promote/1: Promote user to admin
- `POST` /api/v1/admin/demote/1: Demote user from admin
- `POST` /api/v1/admin/notify: Send email notification to all the users

Products:
- `GET` /api/v1/products: Get all products
- `GET` /api/v1/products/{id}: Get product by ID
- `POST` /api/v1/products: Create a new product
- `PUT` /api/v1/products/{id}: Update product by ID
- `DELETE` /api/v1/products/{id}: Delete product by ID

## Middleware
- `SecureHeaders`: Add security headers to each response
- `Cors`: Enables Cross-Origin Resource Sharing policy
- `LogRequest`: Logs each request to terminal and api.log file in the root
- `Authenticate`: Checks and decodes jwt token, then puts decoded data into request header
- `AuthenticateAdmin`: Processes data from previous middleware and ensures that it's admin user.
- `AuthenticateSelfOrAdmin`: Processes decoded data and compares id's to access services only to make changes to personal data or give access to admins. 
- `AuthenticateVerified`: Processes decoded data and ensures that user has verified email.

## Testing
- Unit Testing
    ```bash
    go test ./internal/repository/user/ -v
    ```

- Intergration Testing
    ```bash
    go test ./internal/delivery/http/user/ -v
    ```

- End-to-end Testing
    ```bash
    go test ./cmd/api -v
    ```