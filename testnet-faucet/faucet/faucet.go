package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/ignite/pkg/chaincmd"
	chaincmdrunner "github.com/ignite/cli/ignite/pkg/chaincmd/runner"
)

// Default Config vars
var (
	Port          = "9000"
	AppCli        = "dualityd"
	FaucetAccount = ""
	AmountToSend  = 1000
	DenomsToSend  = []string{"token", "stake"}
	TxTimeoutInt  = 1
	TxTimeout     = time.Duration(TxTimeoutInt) * time.Second
	MaxTxRetry    = 10
	NodeAddress   = "tcp://localhost:26657"
)

type request struct {
	Address string
}

type response struct {
	Message string
}

func newChainCmdRunner() (cr chaincmdrunner.Runner, err error) {
	ccoptions := []chaincmd.Option{
		// chaincmd.WithKeyringPassword(keyringPassword),
		chaincmd.WithKeyringBackend(chaincmd.KeyringBackendTest),
		chaincmd.WithAutoChainIDDetection(),
		chaincmd.WithNodeAddress(NodeAddress),
	}

	cr, err = chaincmdrunner.New(context.Background(), chaincmd.New(AppCli, ccoptions...))

	return
}

func sendToken(cr chaincmdrunner.Runner, address string, amount int, denom string) error {
	amountStr := fmt.Sprintf("%d%s", amount, denom)
	log.Printf("Sending: %s", amountStr)

	tx, err := cr.BankSend(context.Background(), FaucetAccount, address, amountStr)
	if err != nil {
		return err
	}

	// TODO: This is very slow because we can only do 1 tx per block. Ideally need to find a way to batch transactions
	err = cr.WaitTx(context.Background(), tx, TxTimeout, MaxTxRetry)
	if err != nil {
		return err
	}

	return nil
}

func sendAllTokens(address string) error {
	cr, err := newChainCmdRunner()
	if err != nil {
		return err
	}

	for _, denom := range DenomsToSend {
		err := sendToken(cr, address, AmountToSend, denom)
		if err != nil {
			return err
		}
	}

	return nil
}

func faucetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/health" {
		resp := response{Message: "ok"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	address := r.URL.Query().Get("address")

	// Ensure that an address is passed in
	if address == "" {
		http.Error(w, "Key \"address\" not provided in request", http.StatusOK)
		return
	}

	// Ensure that address is valid
	// TODO: this is a very awkward way to validate an adresss since we just convert it back to a string
	accAdress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		http.Error(w, "Address is invalid", http.StatusInternalServerError)
		return
	}

	// Send tokens to address
	err = sendAllTokens(address)
	if err != nil {
		log.Printf("ERROR: %s ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := response{Message: fmt.Sprintf("Tokens sent to: %s", accAdress.String())}
	json.NewEncoder(w).Encode(resp)

	log.Printf("Sent tokens to: %s", address)
}

func startServer() {
	http.HandleFunc("/", faucetHandler)
	log.Printf("Listening on http://localhost:%s", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Port), nil))
}

func main() {
	// Handle flags
	flag.StringVar(&FaucetAccount, "faucet-account", FaucetAccount, "Account to use for faucet")
	flag.StringVar(&Port, "port", Port, "Port to listen on")
	flag.StringVar(&NodeAddress, "node", NodeAddress, "<host>:<port> to tendermint rpc interface for this chain")
	flag.IntVar(&AmountToSend, "token-amount", AmountToSend, "Amount of token to send")
	denomsToSendStr := flag.String("denoms", strings.Join(DenomsToSend, ","), "Denoms to send")
	flag.Parse()

	// Split denoms back to an array
	DenomsToSend = strings.Split(*denomsToSendStr, ",")

	config := map[string]interface{}{
		"Port":          Port,
		"AppCli":        AppCli,
		"FaucetAccount": FaucetAccount,
		"AmountToSend":  AmountToSend,
		"DenomsToSend":  DenomsToSend,
		"TxTimeoutInt":  TxTimeoutInt,
		"MaxTxRetry":    MaxTxRetry,
		"Node":          NodeAddress,
	}
	prettyConfig, _ := json.MarshalIndent(config, "", "  ")
	log.Printf("Starting faucet with config:\n %s", prettyConfig)

	startServer()
}
