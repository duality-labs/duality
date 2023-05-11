const DualitySwapForwardVirtualContract = artifacts.require('DualitySwapForwardVirtualContract');
const Web3 = require('web3');
const web3 = new Web3();
const fs = require('fs');

contract('DualitySwapForwardVirtualContract', (accounts) => {
  it('should give me the abi', async () => {
    const instance = await DualitySwapForwardVirtualContract.deployed();

    // find the function in the ABI
    const functionAbi = instance.constructor._json.abi.find(
      (element) => element.name === 'swapAndForward'  // replace with your function's name
    );

    // assuming you have the function parameters
    const params = [
      // string memory creator,
      "alice",
      // string memory receiver,
      "bob",
      // string memory tokenIn,
      "foo",
      // string memory tokenOut,
      "bar",
      // int64 tickIndex,
      0,
      // uint256 amountIn,
      100,
      // LimitOrderType orderType,
      0,
      // uint expirationTime,
      0,
      // bool nonRefundable,
      false,
      // string memory refundAddress,
      "alice",
      // bytes memory nextArgs
      [...Buffer.from('{}')]
    ];  // replace with your function's parameters

    // encode the function call
    const data = web3.eth.abi.encodeFunctionCall(functionAbi, params);

    const path = require('path');

    const currentDirectory = process.cwd();
    console.log('Current working directory:', currentDirectory);

    // Write the buffer to a file
    fs.writeFile('./abi-encoded-args.bin', data, (err) => {
      if (err) {
        const newErr = new Error('Error writing file');
        newErr.cause = err
        throw newErr
      }
    });
  });
});