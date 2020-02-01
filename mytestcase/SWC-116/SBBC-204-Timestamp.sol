pragma solidity ^0.4.24;

contract KingEth {
    address king;
   	uint king_at;
    function claim() {
        if (msg.sender == king && now - king_at >= 21600) {
            /* Use now to compare before sending ethereum */
            msg.sender.send(address(this).balance);
        }
    }
    function play() payable {
      if (msg.value >= 0.1 ether) {
          king = msg.sender;
          king_at = now;
      }  
    }
}