pragma solidity ^0.4.24;

contract ExceptionDisorder {
   mapping (address => uint256) public balanceOf;
   // INSECURE
   function transfer(address _to, uint256 _value) {
       /* Check if sender has balance */
       require(balanceOf[msg.sender] >= _value);
       /* Send ethereum */
       _to.send(_value);
       /* Add and subtract new balances */
       balanceOf[msg.sender] -= _value;
   }
}