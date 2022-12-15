module github.com/NicholasDotSol/duality

go 1.16

require (
	github.com/cosmos/admin-module v0.0.0
	github.com/cosmos/cosmos-sdk v0.45.12-0.20221116140330-9c145c827001
	github.com/cosmos/ibc-go/v3 v3.4.0
	github.com/cosmos/interchain-security v0.2.2-0.20221213191441-435cf38c38a4
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0 // indirect
	github.com/ignite-hq/cli v0.22.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.34.24
	github.com/tendermint/tm-db v0.6.7
	google.golang.org/genproto v0.0.0-20221114212237-e4508ebdbee1
	google.golang.org/grpc v1.50.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0
	github.com/cosmos/admin-module => github.com/Ethernal-Tech/admin-module v0.0.0-20221102105340-e693f4d379c3
	github.com/cosmos/ibc-go/v3 => github.com/informalsystems/ibc-go/v3 v3.4.1-0.20221202165607-3dc5ba251371
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2

)
