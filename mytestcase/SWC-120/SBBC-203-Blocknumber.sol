pragma solidity ^0.4.24;

contract KingEth {
    address king;
   	uint king_at;
    function claim() {
        if (msg.sender == king && block.number - king_at >= 30) {
            /* Use block.numer to compare before sending ethereum */
            msg.sender.send(address(this).balance);
        }
    }
    function play() payable {
      if (msg.value >= 0.1 ether) {
          king = msg.sender;
          king_at = block.number;
      }  
    }
}