server:
	@go run cmd/main.go

templ:
	@templ generate

mod_init: # or go mod init your_project (when locally)
	@go mod init github.com/yourusername/yourproject