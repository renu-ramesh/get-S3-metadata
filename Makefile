generate:
	go run schema/schema.go

lb:
	echo "Building binary...."
	go build -o eveara-meta-data
	
api:
	echo "Running api server...."
	go run main.go --run-api

zip:
	rm -rf eveara-meta-data.zip
	zip -r eveara-meta-data.zip eveara-meta-data schema/json
