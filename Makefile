REVISION := $(shell git rev-parse --short HEAD)

.PHONY: deploy
deploy:
	goapp deploy -version ${REVISION} backend/app/app_dev.yaml

.PHONY: deploy-production
deploy-production:
	goapp deploy -version ${REVISION} backend/app/app.yaml

.PHONY: test
test:
	goapp test ./backend/app -v
