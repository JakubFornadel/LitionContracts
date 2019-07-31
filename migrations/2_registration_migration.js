var Registry = artifacts.require("LitionRegistry");
var TestToken = artifacts.require("TestToken");

module.exports = function(deployer) {
   deployer.deploy(TestToken).then(function() {
      return deployer.deploy(Registry, TestToken.address);
   });

}
