pragma solidity ^0.4.24;
contract BlockCoinbaseDependencySample {
    address owner;
    address blockaddress;
    function claim() public {
        if (msg.sender == owner && blockaddress==block.coinbase ) {
            msg.sender.transfer(address(this).balance);
        }
    }
    function buy() payable public {
      if (msg.value >= 0.1 ether) {
          owner = msg.sender;
          blockaddress=block.coinbase;
      }  
    }
}