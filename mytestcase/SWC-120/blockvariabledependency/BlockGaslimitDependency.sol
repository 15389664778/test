pragma solidity ^0.4.24;
contract  BlockGaslimitDependencySample {
    address king;
    uint256 limit;
    constructor() public{
        king=msg.sender;
    }
    function claim( uint256 gasUsed,uint256 gasPrice) public{
         limit=mul(gasUsed,gasPrice);
        if (msg.sender == king && limit<=block.gaslimit) {
            msg.sender.transfer(address(this).balance);
        }
    }
     function mul(uint256 a, uint256 b) internal pure returns (uint256) {
      if (a == 0) {
        return 0;
      }
      uint256 c = a * b;
      require(c / a == b);
      return c;
    }
   
}