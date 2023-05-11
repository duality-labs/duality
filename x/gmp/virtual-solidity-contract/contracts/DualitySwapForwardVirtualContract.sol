pragma solidity ^0.8.0;

contract DualitySwapForwardVirtualContract {

  enum LimitOrderType { 
    GOOD_TIL_CANCELLED,
    FILL_OR_KILL,
    IMMEDIATE_OR_CANCEL,
    JUST_IN_TIME,
    GOOD_TIL_TIME
  }

	function swapAndForward(
    string memory creator,
    string memory receiver,
    string memory tokenIn,
    string memory tokenOut,
    int64 tickIndex,
    uint256 amountIn,
    LimitOrderType orderType,
    // expirationTime is only valid iff orderType == GOOD_TIL_TIME.
    uint expirationTime,
    bool nonRefundable,
	  string memory refundAddress,
    // should be a json-encoded "forward" memo
    // {
    //   receiver: "<bech32 addr>",
    //   port: "<port name>",
    //   channel: "channel-<id>",
    //   timeout: <timeout integer?>,
    //   retries: <num retries>,
    //   next: "<optionally include another memo for execution once the result has been forwarded>"
    // }
	  bytes memory nextArgs
  ) public {}

}