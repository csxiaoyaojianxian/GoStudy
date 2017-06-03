#Makefile
GCC_PTHREAD         = -lpthread
GCC_FLAGS           = -c
GCC_OUTPUT_FLAGS    = -o
GCC                 = gcc
RM                  = rm
RM_FLAGS            = -f
TARGET              = test
OBJS                = linktable.o menu.o test.o
all:	$(OBJS)
	$(GCC) $(GCC_OUTPUT_FLAGS) $(TARGET) $(OBJS)

.c.o:
	$(GCC) $(GCC_FLAGS) $<

clean:
	$(RM) $(RM_FLAGS) $(OBJS) $(TARGET) *.bak
