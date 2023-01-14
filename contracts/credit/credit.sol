// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Credit is ERC20, Ownable {

    mapping (string => uint256) DexAmount;

    event Mint(address, uint);
    event CustomTransfer(address, address, uint);

    constructor () ERC20("Credit", "pWEMIX") {}

    // 유저가 판매한 credit을 Dex의 해당하는 pool에 넣어주는 함수
    function addDexAmount(string memory tokenName, uint256 creditAmount) external returns (uint256) {
        DexAmount[tokenName] += creditAmount;
        return DexAmount[tokenName];
    }

    function reduceDexAmount(string memory tokenName, uint256 creditAmount) external returns (uint256) {
        DexAmount[tokenName] -= creditAmount;
        return DexAmount[tokenName];
    }

    // 현재 해당하는 토큰-크레딧 풀에 얼마만큼의 credit 물량이 있는지 확인하는 함수
    // 이 함수를 통해서 token을 사거나 팔 때의 가격이 정해진다.
    function checkTokenPool(string memory tokenName) external view returns (uint256) {
        return DexAmount[tokenName];
    }

    // from을 설정할 수 있는 transfer로 함수를 덮어씀
    function transfer(address from, address to, uint amount) external virtual returns (bool) {
        address owner = from;
        _transfer(owner, to, amount);
        emit CustomTransfer(from, to, amount);
        return true;
    }
}