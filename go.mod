module github.com/NicholasDotSol/duality

go 1.16

require (
	github.com/cosmos/admin-module v0.0.0
	github.com/cosmos/cosmos-sdk v0.45.7-0.20221104161803-456ca5663c5e
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/cosmos/interchain-security v0.0.0-20221104204906-6c0d718d8c59
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ignite-hq/cli v0.22.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.50.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/cosmos/admin-module => github.com/Ethernal-Tech/admin-module v0.0.0-20221102105340-e693f4d379c3
	github.com/cosmos/ibc-go/v3 => github.com/informalsystems/ibc-go/v3 v3.0.0-beta1.0.20220816140824-aba9c2f2b943
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
