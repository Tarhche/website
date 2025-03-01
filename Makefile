ps:
	docker compose ps

up:
	docker compose up --build -d

down:
	docker compose down --remove-orphans --volumes

sh%:
	docker compose exec -it $* sh

logs%:
	docker compose logs -f $*
