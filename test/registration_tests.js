const BigNumber = web3.BigNumber;

var Registry = artifacts.require("LitionRegistry");
var DummyChainValidator = artifacts.require("DummyChainValidator");
var TestToken = artifacts.require("TestToken");

async function assertRevert(promise) {
    try {
        await promise;
        assert.fail('Expected revert not received');
    } catch (error) {
        const revertFound = error.message.search('revert') >= 0;
        assert(revertFound, `Expected "revert", got ${error} instead`);
    }
};


contract("RegistryTests", function(accounts) {
   it("Shall do something", async function() {
      let dv = await DummyChainValidator.new();
      let tt = await TestToken.new(); 
      let r = await Registry.new(tt.address);
      await tt.mint(accounts[0], "1000000000000000000000");
      await tt.mint(accounts[1], "1000000000000000000000");
      await tt.approve(r.address, "1000000000000000000000");
     
      
      await r.register_chain("great chain", dv.address, "200000000000000000000", "127.0.0.1:9999");





   });

});
