.PHONY: restart
restart:
	docker compose down --remove-orphans && docker compose build && DD_API_KEY=$(DD_API_KEY) docker compose up -d