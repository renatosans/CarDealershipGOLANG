# CarDealershipGOLANG
Sistema para concessionária de carros usando Golang and React

![screenshot](assets/banner.png)

## Steps to run the project
- docker compose up
- Follow the link http://localhost:8080/api/cars


## Steps to run without docker
- Open the terminal and change the directory to the backend folder:
    > cd golang
- Create .env file with DATABASE_URL
- Run the script to generate prisma client and create the database:
    > go run github.com/steebchen/prisma-client-go db push
- Run the app
    > go run main.go
- Follow the link http://localhost:8080/api/cars
