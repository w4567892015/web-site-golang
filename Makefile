MAKEFLAGS += --silent
ENV=${shell cat .env}

APP_DIR=apps
LIB_DIR=libs

NAME=$$name
OPTS=$$opts

help:
	echo "[Golang Application Tasks]"
	echo "- make new.app name=<NAME>"
	echo "- make new.lib name=<NAME>"
	echo "- make start.dev name=<NAME>"
	echo "- make release name=<NAME>"
	echo "- make test name=<NAME>"
	echo "- make clean"

new.app:
	mkdir -p "${APP_DIR}/${NAME}"
	cd "${APP_DIR}/${NAME}" && go mod init ${APP_DIR}/${name}
	go work use ${APP_DIR}/${name}

new.lib:
	mkdir -p "${LIB_DIR}/${NAME}"
	cd "${LIB_DIR}/${NAME}" && go mod init ${LIB_DIR}/${name}
	go work use ${LIB_DIR}/${name}

start.dev:
	go run ./${APP_DIR}/${name}
