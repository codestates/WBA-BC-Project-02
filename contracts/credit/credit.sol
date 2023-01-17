// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Credit is ERC20, Ownable {
    address Dex;
    address Multisig;

    mapping (string => uint256) public DexAmount;

    event Mint(address, uint);
    event CustomTransfer(address, address, uint);
    event Swap(address, uint);

    constructor () ERC20("Credit", "pWEMIX") payable{}

    modifier onlyDex() {
        require(msg.sender == Dex);
        _;
    }

    modifier onlyMultisig() {
        require(msg.sender == Multisig);
        _;
    }

    function creditToWemix(address payable _to, uint amount) public onlyMultisig {
        _burn(_to, decimal(amount));
        _to.transfer(decimal(amount));
        emit Swap(_to, amount);
    }

    function decimal(uint256 num) private view returns (uint256) {
        return num * (10 ** uint256(decimals()));
    }

    function setDex(address dex) public onlyOwner {
        Dex = dex;
    }

    // 유저가 판매한 credit을 Dex의 해당하는 pool에 넣어주는 함수
    function addDexAmount(string memory tokenName, uint256 creditAmount) onlyDex external returns (uint256) {
        DexAmount[tokenName] += creditAmount;
        return DexAmount[tokenName];
    }

    function reduceDexAmount(string memory tokenName, uint256 creditAmount) onlyDex external returns (uint256) {
        DexAmount[tokenName] -= creditAmount;
        return DexAmount[tokenName];
    }

    // 현재 해당하는 토큰-크레딧 풀에 얼마만큼의 credit 물량이 있는지 확인하는 함수
    // 이 함수를 통해서 token을 사거나 팔 때의 가격이 정해진다.
    function checkTokenPool(string memory tokenName) external view returns (uint256) {
        return DexAmount[tokenName];
    }

    // from을 설정할 수 있는 transfer로 함수를 덮어씀
    function transfer(address from, address to, uint amount) onlyDex external virtual returns (bool) {
        address owner = from;
        _transfer(owner, to, amount);
        emit CustomTransfer(from, to, amount);
        return true;
    }

    // credit 토큰 초기 세팅 함수
    function mintSelf(uint amount, string memory tokenName) onlyDex external virtual {
        _mint(msg.sender, decimal(amount));
        DexAmount[tokenName] = decimal(amount);
    }
}