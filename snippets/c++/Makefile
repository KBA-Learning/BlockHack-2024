OS := $(shell uname)

ifeq ($(OS), Darwin)  # macOS
    CC := clang++
    CFLAGS := -std=c++11
    LDFLAGS := -lcurl
else  # Ubuntu or other Linux distributions
    CC := g++
    CFLAGS := -std=c++11
    LDFLAGS := -lcurl
endif

account_balance:
	@$(CC) get_account_balance.cpp $(CFLAGS) $(LDFLAGS)
	@./a.out

block_number:
	@$(CC) get_block_number.cpp $(CFLAGS) $(LDFLAGS)
	@./a.out
