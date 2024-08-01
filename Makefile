swagger:
	openapi-generator generate -i ./api/rest/file.yaml \
	-g go-server -o ./internal/restapi