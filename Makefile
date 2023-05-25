TEST?=test/add.easm		# you can use the enviornment variable TEST to specify what binary the "test" target will run  
BIN=${TEST:%.easm=%.bin}

build/vm:
	mkdir -p build
	(cd vm && go build -o ../$@)

build/as: 	
	(cd as && go build -o ../$@)

test: build/vm ${BIN}
	./build/vm ${BIN}

clean:
	rm -rf build test/*.bin

%.bin: %.easm build/as
	./build/as $< -o $@ 

.PHONY: test
