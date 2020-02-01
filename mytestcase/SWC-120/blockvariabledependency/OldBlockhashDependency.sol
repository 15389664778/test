pragma solidity ^0.4.24;
contract  OldBlockhashDependencySample{
    uint64 randomNumber;
function random(uint64 upper) public  {
  uint64 _seed;
  _seed = uint64(keccak256(keccak256(block.blockhash(block.number), _seed), now));
   randomNumber=_seed % upper;
}
}