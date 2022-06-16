branch := $(GIT_BRANCH_FOR_CI)
EMPTY :=

protoPath := proto

ifeq ($(branch),$(EMPTY))
	branch := main
endif

grace:pull test

check:

pull:check
	git checkout $(branch) && git pull origin $(branch)

test:check
	mkdir -p coverage && \
	go test ./... -v -timeout 200s  -cover -coverprofile=./coverage/coverage.out

coverage:test
	go tool cover -func ./coverage/coverage.out && \
	go tool cover -html ./coverage/coverage.out -o ./coverage/index.html

.PHONY: check pull test coverage grace
