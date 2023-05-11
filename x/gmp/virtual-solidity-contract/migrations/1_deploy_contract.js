const DualitySwapForwardVirtualContract = artifacts.require("DualitySwapForwardVirtualContract");

module.exports = function (deployer) {
  deployer.deploy(DualitySwapForwardVirtualContract);
};