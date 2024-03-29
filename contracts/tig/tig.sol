// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Tig is ERC20, Ownable {

    address Dex;

    event Mint(address, uint);
    event CustomTransfer(address, address, uint);

    constructor () ERC20("Tig", "Tig") {}

    modifier onlyDex() {
        require(msg.sender == Dex);
        _;
    }

    function setDex(address dexAddress) public onlyOwner {
        Dex = dexAddress;
    }

    // 유저가 오프체인 화폐를 온체인 토큰으로 변경요청하면 mint해주는 함수
    function mint(address to, uint amount) onlyDex public {
        _mint(to, amount * (10 ** uint256(decimals())));
        emit Mint(to, amount * (10 ** uint256(decimals())));
    }

    // from을 설정할 수 있는 transfer로 함수를 덮어씀
    function transfer(address from, address to, uint amount) onlyDex external virtual returns (bool) {
        address owner = from;
        _transfer(owner, to, amount);
        emit CustomTransfer(from, to, amount);
        return true;
    }
}