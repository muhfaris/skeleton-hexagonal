# Go Hexagonal
## Hexagonal
`The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file. (--Netflix Blog)`

## Definisi Core Concept
Ada 3 core utama dalam skelaton ini yaitu Entities, Repository dan Services.
- Entities adalah domain objek yang tidak perlu tahu dimana dia akan di simpan (e.g Movie, Shooting Location). 
- Repositories adalah list interface kontrak dengan data source yang mengembalikan data dari suatu entitas.
- Services adalah implementasi kompleks bisnis dan validasi ke entitas / domain. 

### Reference Architecture
![Hexagonal Architecture](https://techraaga.files.wordpress.com/2022/01/image-1.png)

## Struktur Folder
```
.
├── cmd                                 # untuk starting command
│   ├── rest.go                         # ada command khusus rest, bisa juga ditambahkan file baru untuk adapter lain (e.g graphql)
│   └── root.go                         # root command, yang di panggil oleh main
├── configs
├── core
│   ├── entities                        # ini entity / domain / modul
│   │   ├── account                     # entity di grouping, ada 2 sub folder
│   │   │   ├── model                   # entity utama
│   │   │       └── README.md
│   │   │   └── service                 # entity untuk response api dari service
│   │   │       └── README.md
│   │   └── userpublic
│   │       ├── model
│   │       |   └── README.md
│   │       └── service
│   ├── repository                      # kontrak service dengan data source
│   │   ├── mysql                       # di grouping berdasarkan tipe data source
│   │   │   ├── account                 # di grouping berdasarkan konteksnya
│   │   │   │   ├── mutation.go
│   │   │   │   ├── mysql.go
│   │   │   │   └── query.go
│   │   │   └── userpublic
│   │   │       ├── mutation.go
│   │   │       ├── mysql.go
│   │   │       └── query.go
│   │   ├── repository.go
│   │   └── sqlite
│   └── services                        # service / usecase, tempat logic validasi data disini
│       └── user_public.service.go
├── docker-compose.yml
├── files
│   └── db
│       └── 01-table-accounts.sql
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   └── app_rest.go
│   ├── config
│   │   └── config.go
│   └── respond                         # library internal untuk response fail / success
│       └── respond.go
├── main.go
├── README.md
└── transport                           # adapter data, bisa graphql, rest api
    ├── http
    │   └── rest                        # menggunakan rest api
    │       ├── handlers
    │       │   ├── healthcheck.handler.go
    │       │   ├── router.go
    │       │   ├── signin.handler.go
    │       │   ├── singup.handler.go
    │       │   └── v1.handler.go
    │       ├── READM.md
    │       └── rest.go                 # initialize rest api
    └── structures                      # object data untuk parse payload
        ├── README.md
        ├── signin.go
        └── signup.go

```

###
Reference:
- https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749
- https://medium.com/wearewaes/ports-and-adapters-as-they-should-be-6aa5da8893b
