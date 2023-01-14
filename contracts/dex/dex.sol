// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "Draco.sol";
import "Credit.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract DEX is Ownable {
    address _owner;
    Draco draco;
    Credit credit;
    event ratio(string, uint256 , uint256);

    constructor () payable {}

    // decimal 숫자를 추가해주는 함수
    function decimal(uint256 num) private view returns (uint256) {
        return num * (10 ** uint256(draco.decimals()));
    }

    // Draco 컨트랙트를 설정해주는 함수
    function setDraco(address tokenAddress) public onlyOwner {
        require(tokenAddress != address(0x0));
        draco = Draco(tokenAddress);
        draco.mint(address(this), 3800);
    }

    // Credit 컨트랙트를 설정해주는 함수
    function setCredit(address creditAddress) public onlyOwner {
        require(creditAddress != address(0x0));
        credit = Credit(creditAddress);
        credit.mintSelf(10000, "Draco");
    }


    // user가 draco를 팔고 credit을 구매하는 함수
    function buyCreditByDraco(address user, uint256 tokenAmount) public onlyOwner {
        // draco 하나의 가격은 DEX가 가지고 있는 draco의 양 / pool에 있는 credit의 양 이다.
        uint256 tokenPrice = draco.balanceOf(address(this)) / credit.checkTokenPool("Draco");

        // 일단 user의 tokenAmount가 요청한 만큼 있는지 확인한다.
        require(draco.balanceOf(user) >= decimal(tokenAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 credit을 pool에 가지고 있는지 확인한다.
        require(credit.checkTokenPool("Draco") >= decimal(tokenAmount * tokenPrice));

        // 확인이 됐으면, user의 draco를 DEX에게로 옮긴다.
        draco.transfer(user, address(this), decimal(tokenAmount));
        
        // 이후 DEX의 credit을 user에게 옮긴다. 이 때 0.9%만큼의 수수료를 credit으로 제하고 보내게 된다.
        credit.transfer(address(this), user, decimal(tokenPrice * tokenAmount) / 1000 * 991);
        // Dex의 pool의 남은 credit 개수를 user에게 보낸 양을 제한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.reduceDexAmount("Draco", decimal(tokenPrice * tokenAmount) / 1000 * 991);

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Draco 보유량, Draco pool의 Credit 보유량을 emit한다.
        emit ratio("Draco", draco.balanceOf(address(this)), leftPoolAmount);
    }

    // user가 credit을 팔고 draco를 구매하는 함수
    function buyDracoByCredit(address user, uint256 creditAmount) public onlyOwner {
        // draco 하나의 가격은 DEX가 가지고 있는 draco의 양 / pool에 있는 credit의 양 이다.
        uint256 tokenPrice = draco.balanceOf(address(this)) / credit.checkTokenPool("Draco");

        // 일단 user의 creditAmount가 요청한 만큼 있는지 확인한다.
        require(credit.balanceOf(user) >= decimal(creditAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 draco를 가지고 있는지 확인한다.
        require(draco.balanceOf(address(this)) >= decimal(creditAmount) / tokenPrice);

        // 확인이 됐으면, user의 credit을 DEX에게로 옮긴다.
        credit.transfer(user, address(this), decimal(creditAmount));
        // 이후 DEX의 token을 user에게 옮긴다.
        draco.transfer(address(this), user, decimal(creditAmount) / tokenPrice);

        // 유저가 credit으로 draco를 구매할 때는 credit 수수료가 따로 부과되지 않는다.
        // Dex의 pool의 남은 credit 개수를 user에게서 받은 양을 더한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.addDexAmount("Draco", decimal(creditAmount));

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Draco 보유량, Draco pool의 Credit 보유량을 emit한다.
        emit ratio("Draco", draco.balanceOf(address(this)), leftPoolAmount);
    }
}