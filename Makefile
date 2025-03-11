run:
	@rm -rf ./out
	@go build -o out/kv cmd/main.go
	@./out/kv $(filter-out $@, $(MAKECMDGOALS))
%:
	@true
build:
	@rm -rf ./out
	@go build -o out/kv cmd/main.go
move:
	@sudo cp ./out/kv /usr/local/bin
reset:
	rm -rf ~/kv
	rm -rf ~/.config/kv
