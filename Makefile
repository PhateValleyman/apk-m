# Detect current platform
UNAME_S := $(shell uname -s 2>/dev/null)
UNAME_M := $(shell uname -m 2>/dev/null)

# Colors
RED    := \033[0;31m
GREEN  := \033[0;32m
YELLOW := \033[1;33m
BLUE   := \033[0;34m
RESET  := \033[0m

# Default target
.DEFAULT_GOAL := help

# Detect target
ifeq ($(UNAME_S),Linux)

    # ZyXEL NSA320 / FFP ARMv5
    ifeq ($(UNAME_M),armv5tel)
        PLATFORM := zyxel
        GOOS := linux
        GOARCH := arm
        GOARM := 5

    # Android Termux ARM64
    else ifeq ($(shell getprop ro.product.cpu.abi 2>/dev/null),arm64-v8a)
        PLATFORM := redmi
        GOOS := android
        GOARCH := arm64

    # Generic ARM64 Linux
    else ifeq ($(UNAME_M),aarch64)
        PLATFORM := linux-arm64
        GOOS := linux
        GOARCH := arm64

    # Generic x86_64 Linux
    else ifeq ($(UNAME_M),x86_64)
        PLATFORM := linux-amd64
        GOOS := linux
        GOARCH := amd64

    else
        PLATFORM := unknown
    endif

else
    PLATFORM := unknown
endif


# Binary name
APP := $(notdir $(CURDIR))


.PHONY: all build help clean info


all: build


build:
	@echo "$(GREEN)Building $(APP) for $(PLATFORM)$(RESET)"
	@echo "$(BLUE)GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM)$(RESET)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -trimpath -mod=vendor -x


info:
	@echo "$(YELLOW)Platform detection$(RESET)"
	@echo "System : $(UNAME_S)"
	@echo "Arch   : $(UNAME_M)"
	@echo "Target : $(PLATFORM)"
	@echo "GOOS   : $(GOOS)"
	@echo "GOARCH : $(GOARCH)"
	@echo "GOARM  : $(GOARM)"


clean:
	@echo "$(RED)Cleaning$(RESET)"
	rm -f $(APP)


help:
	@echo ""
	@echo "$(GREEN)Usage:$(RESET)"
	@echo "  make $(YELLOW)<target>$(RESET)"
	@echo ""
	@echo "$(GREEN)Targets:$(RESET)"
	@echo "  $(BLUE)build$(RESET)   Build application using detected platform"
	@echo "  $(BLUE)info$(RESET)    Show detected platform information"
	@echo "  $(BLUE)clean$(RESET)   Remove compiled binary"
	@echo "  $(BLUE)help$(RESET)    Show this help"
	@echo ""
	@echo "$(GREEN)Examples:$(RESET)"
	@echo "  $(YELLOW)make$(RESET)"
	@echo "      Build automatically for current device"
	@echo ""
	@echo "  $(YELLOW)make build$(RESET)"
	@echo "      Run Go build with vendor modules"
	@echo ""
	@echo "  $(YELLOW)make info$(RESET)"
	@echo "      Display detected platform"
	@echo ""
