SOURCES:=$(wildcard *.c)
OBJS:=$(patsubst %.c,%.o,$(SOURCES))
ELF:=main
CC:=gcc
CFLAGS:=-g
$(ELF):$(OBJS)
	gcc $^ -o $@
.PHONY:clean
clean:
	rm -rf $(OBJS) $(ELF)
