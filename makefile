#!/usr/bin/make -f

install: go.sum
		export GOARCH=amd64 
		export GOOS=darwin
		export GOFLAGS='-buildmode=pie'
		export CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2"
		export CGO_LDFLAGS="-Wl,-z,relro,-z,now -fstack-protector"
		go install $(BUILD_FLAGS) ./cmd/dualityd