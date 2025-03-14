# Fund Transfer in Go
An account transaction repository in Go!

Instructions to install
1. Install Go -> https://go.dev/doc/install
2. Install postgres (based on OS) -> https://www.postgresql.org/download/macosx/
3. Install database management tool -> For example TablePlus
4. Install postman

After installing Go, we need to get following dependencies
1. go get github.com/lib/pq       -> (postgres Driver for Go)
2. go get github.com/jinzhu/gorm  -> (ORM library for Go)
3. go get github.com/gorilla/mux  -> (for routing requests)

In db.go file, enter the database credentials according to your local database setup, i.e. line 18 to line 22

Clone the repo and once we are inside the repo, hit the following command

go run cmd/api/main.go

Postman collection for testing the endpoints is uploaded in the repo!
