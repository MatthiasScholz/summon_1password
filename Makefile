connect-server:
	docker compose up -d
	docker compose ps

summon-1password:
	earthly +build


item := 1password/DemoPassword/password
vault := iraokmv7alkt7wppkl3dmnjuyi
host := http://localhost:8080
token := <DEMO_ACCESS_TOKEN>
summon-1password-local:
	OP_VAULT=$(vault) OP_CONNECT_HOST=$(host) OP_CONNECT_TOKEN=$(token) go run cmd/main.go $(item)

macosx:
	go build -o summon-1password cmd/main.go

logs:
	docker compose logs

logs-follow:
	docker compose logs --follow

summon-test: macosx
	OP_VAULT=$(vault) OP_CONNECT_HOST=$(host) OP_CONNECT_TOKEN=$(token) summon --provider ./summon-1password -f secrets.yml env | grep DEMO
