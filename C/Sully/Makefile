# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: malexand <malexand@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2015/08/21 18:23:32 by malexand          #+#    #+#              #
#    Updated: 2017/11/21 11:07:05 by malexand         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

EXEC = Sully

DEBUG = no
CC = clang
OS := $(shell uname -s)

export

ifeq ($(DEBUG), yes)
	CFLAGS = -std=c99 -pedantic -g -ggdb
else
	CFLAGS =  -Wall -Werror -Wextra
endif

# Link lib : "-L FOLDER -lXXX" where XXX = libXXX.a

ifeq ($(OS), Linux)
	LFLAGS =
	INCLUDE = -I./incs
else
	LFLAGS =
	INCLUDE = -I./incs
endif

OUT_DIR = ./
SRC_DIR = ./
INC_DIR = ./

SDIR =		./
SRCS =		Sully.c
SRCC =		$(addprefix $(SDIR),$(SRCS))

ODIR =		./
OBJS =		$(SRCS:.c=.o)
OBCC =		$(addprefix $(ODIR),$(OBJS))

all: $(EXEC)

$(EXEC): $(OBCC)
ifeq ($(OS), Linux)
	@echo -e "\x1B[34m$(EXEC):\x1B[0m"
	@$(CC) $(CFLAGS) -o $@ $(OBCC) $(INCLUDE) $(LFLAGS)
	@echo -e "\x1b[36m  + Compile program:\x1B[0m $@"
else
	@echo "\x1B[34m$(EXEC):\x1B[0m"
	@$(CC) $(CFLAGS) -o $@ $(OBCC) $(INCLUDE) $(LFLAGS)
	@echo "\x1b[36m  + Compile program:\x1B[0m $@"
	@echo "\x1B[31m\c"
	@# @norminette srcs/* incs/* | grep -B 1 "Error" || true
	@echo "\x1B[0m\c"
endif

$(ODIR)%.o: $(SDIR)%.c
	@$(CC) $^ $(CFLAGS) -c -o $@ $(INCLUDE)
ifeq ($(OS), Linux)
	@echo -e "\r\x1B[32m  + Compile:\x1B[0m $(notdir $^)"
else
	@echo "\r\x1B[32m  + Compile:\x1B[0m $(notdir $^)"
endif

clean:
	@rm -rf $(OUT_DIR)*.o
ifeq ($(OS), Linux)
	@echo -e "\x1B[31m  - Remove:\x1B[0m $(EXEC) objs"
else
	@echo "\x1B[31m  - Remove:\x1B[0m $(EXEC) objs"
endif

fclean: clean
	@rm -f $(EXEC)
ifeq ($(OS), Linux)
	@echo -e "\x1B[31m  - Remove:\x1B[0m $(EXEC)"
else
	@echo "\x1B[31m  - Remove:\x1B[0m $(EXEC)"
endif

re: fclean
	@make

run: re
	@./$(EXEC)

norm:
	@echo "\x1B[31m\c"
	@norminette srcs/* incs/* | grep -B 1 "Error" || true
	@echo "\x1B[0m\c"


.PHONY: all clean fclean re run directories cleanlib norm
