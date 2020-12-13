start-users:
	@echo "start-users"
	cd cmd/user; pwd && ls && \
	go run main.go handlers.go model.go setup.go