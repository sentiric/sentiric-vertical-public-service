module github.com/sentiric/sentiric-vertical-public-service

go 1.24.5

require (
	github.com/joho/godotenv v1.5.1
	github.com/rs/zerolog v1.34.0
	github.com/sentiric/sentiric-contracts v1.9.0
	google.golang.org/grpc v1.75.1
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
// Gerekli diğer indirect bağımlılıklar buraya gelecek (go mod tidy ile)
)
