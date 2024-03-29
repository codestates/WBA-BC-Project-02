// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "Draco.sol";
import "Tig.sol";
import "Credit.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract DEX is Ownable {
    address _owner;
    address Multisig;
    Draco draco;
    Tig tig;
    Credit credit;
    event ratio(string, uint256 , uint256);

    // 토큰들을 먼저 발급해놓고, contructor에 넣어서 그 때 setting함수를 돌리자
    constructor (
            address dracoAddress,
            address tigAddress,
            address payable creditAddress,
            address multisigAddress
        ) payable {
        // draco 객체 생성
        draco = Draco(dracoAddress);
        // tig 객체 생성
        tig = Tig(tigAddress);
        // credit 객체 생성
        credit = Credit(creditAddress);
        // multisig 등록
        Multisig = multisigAddress;
    }

    modifier onlyMultisig() {
        require(msg.sender == Multisig);
        _;
    }

    // decimal 숫자를 추가해주는 함수
    function decimal(uint256 num) private view returns (uint256) {
        return num * (10 ** uint256(draco.decimals()));
    }

    // draco 초기세팅 함수
    function dracoInitialSetting() public onlyOwner {
        draco.mint(address(this), 10000);
    }

    // tig 초기세팅 함수
    function tigInitialSetting() public onlyOwner {
        tig.mint(address(this), 10000);
    }

    // credit 초기세팅 함수
    function creditInitialSetting(uint dracoPoolAmount, uint tigPoolAmount) public onlyOwner {
        credit.mintSelf(dracoPoolAmount, "Draco");
        credit.mintSelf(tigPoolAmount, "Tig");
    }

    // user가 draco를 팔고 credit을 구매하는 함수
    function buyCreditByDraco(address user, uint256 tokenAmount) public onlyMultisig {
        // draco 하나의 가격은 pool에 있는 credit의 양 / DEX가 가지고 있는 draco의 양이다.
        uint256 tokenPrice =  credit.checkTokenPool("Draco") * 10000 / draco.balanceOf(address(this));

        // 일단 user의 tokenAmount가 요청한 만큼 있는지 확인한다.
        require(draco.balanceOf(user) >= decimal(tokenAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 credit을 pool에 가지고 있는지 확인한다.
        require(credit.checkTokenPool("Draco") >= decimal(tokenAmount * tokenPrice) / 10000);

        // 확인이 됐으면, user의 draco를 DEX에게로 옮긴다.
        draco.transfer(user, address(this), decimal(tokenAmount));

        // 이 때 토큰이 옮겨진 값으로 다시 tokenPrice 업데이트
        tokenPrice =  credit.checkTokenPool("Draco") * 10000 / draco.balanceOf(address(this));

        // 이후 DEX의 credit을 user에게 옮긴다. 이 때 0.9%만큼의 수수료를 credit으로 제하고 보내게 된다.
        credit.transfer(address(this), user, decimal(tokenPrice * tokenAmount) / 1000 * 991 / 10000);
        
        // Dex의 pool의 남은 credit 개수를 user에게 보낸 양을 제한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.reduceDexAmount("Draco", decimal(tokenPrice * tokenAmount) / 1000 * 991 / 10000);

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Draco 보유량, Draco pool의 Credit 보유량을 emit한다.
        emit ratio("Draco", draco.balanceOf(address(this)), leftPoolAmount);
    }

    // user가 credit을 팔고 draco를 구매하는 함수
    function buyDracoByCredit(address user, uint256 creditAmount) public onlyMultisig {
        // draco 하나의 가격은 pool에 있는 credit의 양 / DEX가 가지고 있는 draco의 양이다.
        uint256 tokenPrice = credit.checkTokenPool("Draco") * 10000 / draco.balanceOf(address(this));

        // 일단 user의 creditAmount가 요청한 만큼 있는지 확인한다.
        require(credit.balanceOf(user) >= decimal(creditAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 draco를 가지고 있는지 확인한다.
        require(draco.balanceOf(address(this)) >= decimal(creditAmount) / tokenPrice * 10000);

        // 확인이 됐으면, user의 credit을 DEX에게로 옮긴다.
        credit.transfer(user, address(this), decimal(creditAmount));

        // 유저가 credit으로 draco를 구매할 때는 credit 수수료가 따로 부과되지 않는다.
        // Dex의 pool의 남은 credit 개수를 user에게서 받은 양을 더한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.addDexAmount("Draco", decimal(creditAmount));

        // credit이 추가된 양으로 다시 tokenPrice 업데이트
        tokenPrice = credit.checkTokenPool("Draco") * 10000 / draco.balanceOf(address(this));

        // 이후 DEX의 token을 user에게 옮긴다.
        draco.transfer(address(this), user, decimal(creditAmount) / tokenPrice * 10000);

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Draco 보유량, Draco pool의 Credit 보유량을 emit한다.
        emit ratio("Draco", draco.balanceOf(address(this)), leftPoolAmount);
    }


    // user가 Tig를 팔고 credit을 구매하는 함수
    function buyCreditByTig(address user, uint256 tokenAmount) public onlyMultisig {
        // draco 하나의 가격은 pool에 있는 credit의 양 / DEX가 가지고 있는 Tig의 양이다.
        uint256 tokenPrice =  credit.checkTokenPool("Tig") * 10000 / tig.balanceOf(address(this));

        // 일단 user의 tokenAmount가 요청한 만큼 있는지 확인한다.
        require(tig.balanceOf(user) >= decimal(tokenAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 credit을 pool에 가지고 있는지 확인한다.
        require(credit.checkTokenPool("Tig") >= decimal(tokenAmount * tokenPrice) / 10000);

        // 확인이 됐으면, user의 tig를 DEX에게로 옮긴다.
        tig.transfer(user, address(this), decimal(tokenAmount));

        // 이 때 토큰이 옮겨진 값으로 다시 tokenPrice 업데이트
        tokenPrice =  credit.checkTokenPool("Tig") * 10000 / tig.balanceOf(address(this));

        // 이후 DEX의 credit을 user에게 옮긴다. 이 때 0.9%만큼의 수수료를 credit으로 제하고 보내게 된다.
        credit.transfer(address(this), user, decimal(tokenPrice * tokenAmount) / 1000 * 991 / 10000);
        
        // Dex의 pool의 남은 credit 개수를 user에게 보낸 양을 제한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.reduceDexAmount("Tig", decimal(tokenPrice * tokenAmount) / 1000 * 991 / 10000);

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Tig 보유량, Tig pool의 Credit 보유량을 emit한다.
        emit ratio("Tig", tig.balanceOf(address(this)), leftPoolAmount);
    }


    // user가 credit을 팔고 tig를 구매하는 함수
    function buyTigByCredit(address user, uint256 creditAmount) public onlyMultisig {
        // tig 하나의 가격은 pool에 있는 credit의 양 / DEX가 가지고 있는 tig의 양이다.
        uint256 tokenPrice = credit.checkTokenPool("Tig") * 10000 / tig.balanceOf(address(this));

        // 일단 user의 creditAmount가 요청한 만큼 있는지 확인한다.
        require(credit.balanceOf(user) >= decimal(creditAmount));
        // 그리고 DEX가 user에게 지급할 수 있는만큼의 tig를 가지고 있는지 확인한다.
        require(tig.balanceOf(address(this)) >= decimal(creditAmount) / tokenPrice * 10000);

        // 확인이 됐으면, user의 credit을 DEX에게로 옮긴다.
        credit.transfer(user, address(this), decimal(creditAmount));

        // 유저가 credit으로 tig를 구매할 때는 credit 수수료가 따로 부과되지 않는다.
        // Dex의 pool의 남은 credit 개수를 user에게서 받은 양을 더한 값으로 업데이트해준다.
        uint256 leftPoolAmount = credit.addDexAmount("Tig", decimal(creditAmount));

        // credit이 추가된 양으로 다시 tokenPrice 업데이트
        tokenPrice = credit.checkTokenPool("Tig") * 10000 / tig.balanceOf(address(this));

        // 이후 DEX의 token을 user에게 옮긴다.
        tig.transfer(address(this), user, decimal(creditAmount) / tokenPrice * 10000);

        // Daemon에서 현재 DEX 상태를 업데이트할 수 있게 Dex의 Tig 보유량, Tig pool의 Credit 보유량을 emit한다.
        emit ratio("Tig", tig.balanceOf(address(this)), leftPoolAmount);
    }

    // user에게 draco를 민팅해주는 함수
    function mintDraco(address user, uint256 amount) public onlyMultisig {
        draco.mint(user, amount);
    }

    // user에게 tig를 민팅해주는 함수
    function mintTig(address user, uint256 amount) public onlyMultisig {
        tig.mint(user, amount);
    }

    function creditToWemix(address user, uint256 amount) public onlyMultisig {
        credit.creditToWemix(payable(user), amount);
    }

    function setMultisig(address multisigAddress) public onlyOwner {
        Multisig = multisigAddress;
    }
}