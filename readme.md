# Minab Events Server

- Backend server that contains custom business logic in the form of actions and events for [Minab Events](https://github.com/besufikad17/minab_events) application.

## Technologies used

- [Hasura GraphQL](https://hasura.io/) - GraphQL server
- [http](https://pkg.go.dev/net/http) - Builtin HTTP package
- [golang-jwt](https://github.com/golang-jwt/jwt/v5) - JWT tool
- [godotenv](https://github.com/joho/godotenv) - Package for handling environment 
variables
- [go-qrcode](https://github.com/skip2/go-qrcode) - Package for generating QR code
- [smtp](https://pkg.go.dev/net/smtp) - Builtin package 
- [chapa](https://chapa.co/) - Payment gateway
- [cloudinary](https://cloudinary.com) - Image uploading service

## Setup


1. Clonning the repo
   
   ```bash
    git clone https://github.com/Besufikad17/minab_events_server.git
   ```

2. Installing packages
   
   ```bash
    cd minab_events_server && go install
    ```
3. Connecting database
   
   ```bash
   // creaing .env file
   touch .env
   ```
   ```.env
   // storing environment variables in .env file
   JWT_SECRET=""
   HASURA_URL="http://localhost:8080/v1/graphql"
   SMTP_EMAIL=""
   SMTP_PWD=""
   CHAPA_URL="https://api.chapa.co/v1/transaction/initialize"
   CHAPA_AUTH=""
   CHAPA_RETURN_URL="http://localhost.localdomain:3000/events/verify"
   CHAPA_CALLBACK_URL=""
   CLOUDINARY_URL=""
   CLOUDINARY_UPLOAD_PRESET=""
   ```
4. Running 

    ```bash
    go run server.go
    ```