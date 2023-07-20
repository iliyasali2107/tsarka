server:
	go run ./cmd/main.go


mock_storage:
	mockgen -destination=internal/mocks/mock_repository.go --build_flags=--mod=mod -package=mocks tsarka/internal/repository/counter_repository CounterRepository


mock_service:
	mockgen -destination=internal/mocks/mock_service.go --build_flags=--mod=mod -package=mocks tsarka/internal/service SelfService

test:
	go test -v -cover ./...

