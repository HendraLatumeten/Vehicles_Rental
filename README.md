
<h1 align="center">
  Vehicle-Rental.App With <a href="https://golang.org/" target="blank">Go Language</a>
</h1>

<p align="center"><img src="https://i.ibb.co/cDq4dnH/login-page-1.png" width="400px" alt="Golang.jpg" /></p>


## Description
Vehicles-Rental.app is a backend system for renting vehicles such as motorbikes, cars and bicycles. The rental can be done based on the nearest city and there are many more advantages.
This application uses a RESTful API and was developed using the GO language with the Gorilla/Mux Framework. The ORM uses GORM and the database uses Postgresql.

## ⚡Features

- CRUD (Vehicles,Users,History)
- Search Query (Vehicles)
- Hashing Password (bcrypt)
- JWT
- Authentication
- Authorization
- Cobra Command
- DB Migration

## 🛠️ Installation Steps
1. Clone the Repository

```bash
git clone https://github.com/HendraLatumeten/Vehicles_Rental.git
```

2. Install Dependencies

```bash
go get -u ./...
```

3. Run the Migration

```bash
#Migration: 
go run main.go migrate --up
# or
#Rollback: 
go run main.go migrate --down
```


4. Add File Env

```sh
  APP_PORT= Your Port
  JWT_KEYS= Your Secret Keys

  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
```


5. Run the App

```bash
go run main.go start
```

## 💻 Built with

- [Golang](https://go.dev/) : programming language
- [gorilla/mux](https://github.com/gorilla/mux): handle http request
- [Postgres](https://www.postgresql.org/): DBMS


## 📌 Connect My Social Media

 - Facebook : [Hendra Latumeten](https://web.facebook.com/hendra.latumeten)
 
 - Instagram : [Hendra Latumeten](https://www.instagram.com/hendralatumeten)
 
 - Linkedin : [Hendra Latumeten](https://www.linkedin.com/in/hendralatumeten/)
 
 
 ## 🔥 supported by
 <p align="center"><img src="https://yt3.ggpht.com/ytc/AKedOLT7YD9x6PiR-CfbBbFC3wz2WatiIZFrI_I0v-6k=s900-c-k-c0x00ffffff-no-rj" width="200px" height="200px" alt="Arkademylogo.svg" /></p>

<p align="center">
    <a href="https://www.fazztrack.com/" target="blank">Our Website</a>
    ·
    <a href="https://www.fazztrack.com/class/backend-golang">Join With Us</a>
    ·
</p>
 

