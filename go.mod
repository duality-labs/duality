module github.com/NicholasDotSol/duality

go 1.16

require (
	github.com/cosmos/admin-module v0.0.0
	github.com/cosmos/cosmos-sdk v0.45.2-0.20220901181011-06d4a64bf808
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/cosmos/interchain-security v0.0.0-20221102103028-d7f8d448be65
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/cosmos/admin-module => github.com/jtieri/admin-module v0.0.0-20221116191954-1d63d5fc9608
	github.com/cosmos/ibc-go/v3 => github.com/jtieri/ibc-go/v3 v3.0.0-beta1.0.20221116191630-01c53c7f66f3
	github.com/cosmos/interchain-security v0.0.0-20221102103028-d7f8d448be65 => github.com/jtieri/interchain-security v0.0.0-20221116194529-59bf07eb134f
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
