github_user = bohe-in
api_dir = demo-api-client-ts

build:
	go build main.go

run:
	go run main.go

api-gen:
	go run main.go api-doc
	openapi-generator-cli generate -i docs/openAPI.json -g typescript-axios -o ../$(api_dir)/src
	cp docs/openAPI.json ../$(api_dir)/openAPI.json
	rm ../$(api_dir)/src/git_push.sh
	cd ../$(api_dir); npm version -f patch;
	cd ../$(api_dir); git add . ; git commit -m "minor update"; git push;
	cd ../$(api_dir); npm publish;

package-auto:
	rm -rf aws-deployment
	mkdir -p aws-deployment/src/main
	env GOOS=linux GOARCH=amd64 go build main.go
	mv main aws-deployment/main
	cp -R internal/resources aws-deployment/internal/main
	cp Procfile aws-deployment

package-manual:
	rm -f deployment.zip
	env GOOS=linux GOARCH=amd64 go build main.go
	zip -r deployment.zip main internal/resources/ Procfile
	rm main

clean:
	rm -f main
	rm -f deployment.zip
	rm -rf aws-deployment

recreate-db:
	go run main.go recreate-db

update-db:
	go run main.go update-db

test-unit:
	go test ./internal... -v

test-integration:
	go test ./internal/test/test_integration -v -count=1 -p 1
