.PHONY: default  build pub_local

default: build

BIN = echoServer
NAME = echoServer

PUB_DIR = pub
BUILD_DIR = dist

define build
	@echo "===== $(BIN) ====="
	@mkdir -p git
	@echo "$$GIT_INFO" > git/git.go
	@go run cmd/make/make.go -build -binary $(BIN) -name $(NAME) -build-dir $(BUILD_DIR)
	@tree -Csh -L 3 $(BUILD_DIR)
endef

define pub
	@echo "publish $(1) $(NAME) ..."
	@go run cmd/make/make.go -pub -binary $(BIN) -name $(NAME) -build-dir $(BUILD_DIR) -pub-dir $(PUB_DIR) -env $(1) -version $(2)
	@tree -Csh -L 3 $(PUB_DIR)
endef

# 编译
build:
	$(call build, dev)

# 发布
pub_test: build
	$(call pub, test, test)

pub_preprod: build
	$(call pub, preprod, preprod)

pub_release:
	$(call pub, release, release)

build_image: build
	@docker build -t echoServer:dev .
	@docker push echoServer:dev

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(PUB_DIR)