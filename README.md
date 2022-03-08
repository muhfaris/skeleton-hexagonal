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

### How to Switch Database
#### Change to MongoDB
Reference branch ada di [refactor/switch-to-mongo](https://github.com/muhfaris/skeleton-hexagonal/tree/refactor/switch-to-mongo).

- repository 
    disini kita akan membuat folder baru `mongodb` dan membuat file baru untuk implementasi interface repository
```
│   ├── repository
│   │   ├── mongodb
│   │   │   ├── account
│   │   │   │   ├── mongodb.go
│   │   │   │   ├── mutation.go
│   │   │   │   └── query.go
│   │   │   └── userpublic
│   │   │       ├── mongodb.go
│   │   │       ├── mutation.go
│   │   │       └── query.go
│   │   ├── mysql
│   │   │   ├── account
│   │   │   │   ├── mutation.go
│   │   │   │   ├── mysql.go
│   │   │   │   └── query.go
│   │   │   └── userpublic
│   │   │       ├── mutation.go
│   │   │       ├── mysql.go
│   │   │       └── query.go
│   │   └── repository.go
```

- services
    kita tambahkan sedikit kondisi, dimana ada kode yang baca error dari data source ketika `signup`   
```diff
@@ -40,19 +41,19 @@ func (service *userPublicService) Login(ctx context.Context, params *structures.
	}

	if err := account.ComparePassword(params.Password); err != nil {
-		return svcmodel.AccountResponse{}, result.Error
+		return svcmodel.AccountResponse{}, err
	}

	return *account.Response(), nil
}

func (service *userPublicService) SignUp(ctx context.Context, params *structures.SignUpRead) error {
	result := <-service.accountRepo.FindByEmail(ctx, params.Email)
-	if result.Error != nil && result.Error != pgx.ErrNoRows {
+	if result.Error != nil && result.Error != pgx.ErrNoRows && result.Error != mongo.ErrNoDocuments {
		return result.Error
	}

-	if result.Error == pgx.ErrNoRows {
+	if result.Error == pgx.ErrNoRows || result.Error == mongo.ErrNoDocuments {
		account := model.CreateAccount(params)
		if err := account.GenerateHashPassword(); err != nil {
			return err
```

- internal/config
    buat fungsi untuk melakukan initialize mongodb

```diff 
type ConfigApp struct {
-	Port int
-	DB   *pgx.Conn
+	Port   int
+	DB     *pgx.Conn
+	Client *mongo.Client
}

func CreateConfigApp() *ConfigApp {
-	db, err := initDatabase()
+	db, err := initDatabaseMySQL()
+	if err != nil {
+		panic(err)
+	}
+
+	client, err := initDatabaseMongoDB()
	if err != nil {
		panic(err)
	}

	return &ConfigApp{
-		DB: db,
+		DB:     db,
+		Client: client,
	}
}

-func initDatabase() (*pgx.Conn, error) {
+func initDatabaseMySQL() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseURL := fmt.Sprintf("postgres://admin123:admin123@localhost:5432/skelaton_db")
	db, err := pgx.Connect(context.Background(), databaseURL)
@@ -33,3 +42,23 @@ func initDatabase() (*pgx.Conn, error) {

	return db, nil
}

+func initDatabaseMongoDB() (*mongo.Client, error) {
+	credential := options.Credential{
+		Username: "admin123",
+		Password: "admin123",
+	}
+
+	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
+	client, err := mongo.NewClient(clientOptions)
+	if err != nil {
+		return nil, err
+	}
+
+	err = client.Connect(context.Background())
+	if err != nil {
+		return nil, err
+	}
+
+	return client, nil
+}
// transport/http/rest/rest.go
@@ -42,7 +42,7 @@ func NewRest(port int) *Rest {
		_ = r.Shutdown()
	}()

-	servicesApp := app.NewServiceApp(cApp.DB)
+	servicesApp := app.NewServiceApp(cApp.Client)

	rest := &Rest{
		port:     port,
```

###
Reference:
- https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749
- https://medium.com/wearewaes/ports-and-adapters-as-they-should-be-6aa5da8893b
