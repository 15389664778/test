pragma solidity ^0.4.24;

contract UnderflowSample {
   mapping (address => uint256) public balanceOf;
   // INSECURE
   function transfer(address _to, uint256 _value) {
       /* Add and subtract new balances */
       balanceOf[msg.sender] -= _value;
       balanceOf[_to] += _value;
   }
}