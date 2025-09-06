generate-mocks:
	mockery --dir pkg/myapi --name API --output mocks --outpkg mocks --with-expecter
	mockery --dir internal/httpx --name Doer --output mocks --outpkg mocks --with-expecter