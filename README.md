
# Waysbeans (Backend)

Backend for waysbeans using RESTful API.


## Tech Stack

**Server:** Golang (Echo)

**Database:** PostgreSQL

**Others:** JWT, Bycript, Cloudinary
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`SECRET_KEY=<this-is-secret>`

`SERVER_KEY=<midtrans-server-key>`

`EMAIL_SYSTEM=<your-gomail-email@gmail.com>`

`PASSWORD_SYSTEM=<gomail-application-password>`

`DB_HOST=<db-host>`

`DB_NAME=<db-name>`

`DB_PASSWORD=<db-password>`

`DB_PORT=<db-port>`

`DB_USER=<db-user>`

`CLOUD_NAME=<cloudinary-cloud-name>`

`API_KEY=<cloudinary-api-key>`

`API_SECRET=<cloudinary-api-secret>`
## Run Locally

Clone the project

```bash
  git clone https://github.com/VindoKountur/backend-waysbeans-golang
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy ; go mod download
```

Start the server

```bash
  go run main.go
```

