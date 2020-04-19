build: ## Build golang binaries
	
	cd cmd/comments && GOOS=linux go build -o ../../bin/comments
	cd cmd/users && GOOS=linux go build -o ../../bin/users
	cd cmd/achievements && GOOS=linux go build -o ../../bin/achievements
	cd cmd/forum && GOOS=linux go build -o ../../bin/forum
	cd cmd/health && GOOS=linux go build -o ../../bin/health

deploy:
	cd api/comments && GOOS=linux go build -o ../../bin/comments
	cd api/users && GOOS=linux go build -o ../../bin/users
	cd api/achievements && GOOS=linux go build -o ../../bin/achievements
	cd api/forum && GOOS=linux go build -o ../../bin/forum
	cd api/health && GOOS=linux go build -o ../../bin/health
	serverless deploy -s production