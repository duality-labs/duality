# virtual-solidity-contract

This subdir has a solidity contract with a swapAndForward() function that
represents the signature that Duality accepts from Axelar's GMP message payload.
Solidity contracts should encode payload data the way they would when calling
this contract function.

If this interface changes in the future it can be regenerated using `./run.sh`.
The abi will be printed out, to be copied into
`../swap_forward_memo_transcoder.go`. The test binary will also be printed out,
and should be copied into `../swap_forward_memo_transcoder_test.go` and
`$REPO_ROOT/interchaintest/gmp_swap_forward_test.go`.

To update the input to the test encoding look in
`test/dualitySwapForwardVirtualContract.test.js`.