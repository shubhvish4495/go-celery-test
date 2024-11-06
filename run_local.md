# RUN Local

1. Run docker compose from root folder
```cmd
docker-compose up --build
```

2. Run this command from root folder again
Now run command
```cmd
cd client && go run client.go
```

Would suggest to run this multiple times to see how data is being distributed among workers