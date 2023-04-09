swag init -d ./cmd/app,./internal/app,./internal/transport,./internal/models
http://localhost:8080/swagger/index.html
sudo mongod --dbpath ./data/db

Add JWT_SECRET to .env file