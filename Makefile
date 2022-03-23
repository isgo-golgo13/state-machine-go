SERVICE_NAME=app

.PHONY : clean 
.DEFAULT_GOAL : all

all:
	
compile: 
	go build github.com/isgo-golgo13/state-machine-go/svc/app

clean: 
	rm -f ${SERVICE_NAME}

test:
	go test -v