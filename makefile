build: ## Build golang binaries
	
	cd api/achievements && GOOS=linux go build -o ../../bin/achievements
	cd api/comments && GOOS=linux go build -o ../../bin/comments
	cd api/forum && GOOS=linux go build -o ../../bin/forum