push:
	@echo "Committing code to current branch..."
	git add .
	@powershell -Command " \
		$$message = Read-Host 'Please enter commit message'; \
		if ($$message -eq '') { echo 'Commit message cannot be empty'; exit 1 } \
		git commit -m $$message; \
		git push; \
	"

run:
	go run main.go