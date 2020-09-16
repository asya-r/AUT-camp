heroku-deploy:
	heroku container:login
	heroku container:push web
	heroku container:release web

heroku-logs:
	heroku logs --tail

local-build:
	docker-compose down
	docker-compose build
	docker-compose up -d

local-restart:
	docker-compose restart app

local-addr:
	docker-compose port app 8080
