build:	
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/validator \
			./
	mv bin/validator-linux-amd64 bin/validator