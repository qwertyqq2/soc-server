// SPDX-License-Identifier: MIT
pragma solidity  ^0.7.0 || ^0.8.0;

import "./Round.sol";
import "./interfaces/IRound.sol";

import "./libraries/Proof.sol";
import "./libraries/JumpSnap.sol";

contract Group {
    event CreateRoundEvent(address _roundAddress);

    event EnterRoundEvent(address _sender);

    event StartRoundEvent(uint _bsnap, uint _psnap);

    event CreatedLotEvent(address _lotAddr);

    event NewLotEvent(
         address _lotAddr,
         uint256 _timeFirst,
         uint256 _timeSecond,
         uint256 _price,
         uint256 _val,
         uint256 _lotSnap,
         uint256 _bsnap
     );


    event BuyLotEvent(
        address _lotAddr,
        address _sender,
        uint256 _price,
        uint256 _lotSnap,
        uint256 _bsnap
    );

    event SendLotEvent(
        address _lotAddr,
        uint256 _receiveTokens
    );

    event UpdatePlayerParams(
        address _owner,
        uint256 _nwin, 
        uint256 _n,
        uint256 _spos, 
        uint256 _sneg
    );


    event ReceiveLotEvent(
        address _lotAddr,
        address _owner,
        uint _balance,
        uint _psnap,
        uint _bsnap
    );



    event Price1(uint _amountOut);
    event Price2(uint _amountOut);


    address owner;
    address roundAddr;
    address exchangeAddress;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "You`re not the owner");
        _;
    }

    function CreateRound(uint256 _deposit) public onlyOwner{
        Round round = new Round(_deposit);
        roundAddr = address(round);
        emit CreateRoundEvent(address(round));
    }

    function Enter() public payable {
        IRound round = IRound(roundAddr);
        round.Enter(msg.sender, msg.value);
        payable(roundAddr).transfer(msg.value);
        emit EnterRoundEvent(msg.sender);
    }

    function StartRound() public {
        IRound round = IRound(roundAddr);
        uint bsnap;
        uint psnap;
        (bsnap, psnap) = round.StartRound();
        emit StartRoundEvent(bsnap, psnap);
    }

    function CreateLot() external{
        IRound round = IRound(roundAddr);
        address lotAddr = round.CreateLot();
        emit CreatedLotEvent(lotAddr);
    }

    modifier onlyPlayer() {
        IRound round = IRound(roundAddr);
        require(round.GetPlayer(msg.sender) > 0, "You're not a player");
        _;
    }

    function NewLot(
        address _lotAddr,
        uint256 _timeFirst,
        uint256 _timeSecond,
        uint256 _price,
        uint256 _val,
        uint _Hres,
        uint _balance
    ) external {
        Proof.ProofRes memory proof = Proof.NewProof(msg.sender, _price, _Hres, _balance);
        IRound round = IRound(roundAddr);
        uint lotSnap;
        uint balancesSnap;
        (lotSnap, balancesSnap) = round.NewLot(_lotAddr, _timeFirst, _timeSecond, _val, proof);
         emit NewLotEvent(
             _lotAddr,
             _timeFirst,
             _timeSecond,
             _price,
             _val,
             lotSnap,
             balancesSnap
     );
    }

    function BuyLot(
        address _lotAddr,
        uint256 _price,
        uint _Hres,
        uint _Hd,
        uint256 _balance,
        uint256 _prevBalance,
        address _prevOwner,
        uint256 _prevPrice,
        uint256 _prevSnap
        ) external {
        Proof.ProofRes memory proofRes = Proof.NewProof(msg.sender, _price, _Hres, _balance);
        proofRes.prevOwner = _prevOwner;
        proofRes.prevBalance = _prevBalance;
        proofRes.Hd = _Hd;
        Proof.ProofEnoungPrice memory proofEP = Proof.NewProofEnoughPrice(_prevOwner, _prevPrice, _prevSnap);
        IRound round = IRound(roundAddr);
        uint lotSnap;
        uint balancesSnap;
        (lotSnap, balancesSnap) = round.BuyLot(_lotAddr, proofRes, proofEP);
        emit BuyLotEvent(
            _lotAddr,
            msg.sender,
            _price,
            lotSnap,
            balancesSnap
        );
    }


    function SendLot(
        address _lotAddr,
        bytes memory initParamsData
    ) external{
        Params.InitParams memory initParams = Params.DecodeInitParams(initParamsData);
        IRound round = IRound(roundAddr);
        uint amountOut = round.SendLot(_lotAddr, initParams);
        emit SendLotEvent(_lotAddr, amountOut);
    }

    function ReceiveLot(
        address _lotAddr,
        address _owner,
        bytes memory initParamsData,
        bytes memory proofResData,
        bytes memory playerParamsData
    ) external{
        Params.InitParams memory initParams = Params.DecodeInitParams(initParamsData);
        Proof.ProofRes memory proof = Proof.DecodeProofRes(proofResData);
        proof.owner = _owner;
        Params.PlayerParams memory params = Params.DecodePlayerParams(playerParamsData);
        IRound round = IRound(roundAddr);
        bytes memory newParamsData;
        uint newBalance;
        (newParamsData, newBalance) = round.ReceiveLot(_lotAddr, initParams, proof, params);
        Params.PlayerParams memory NewParams = Params.DecodePlayerParamsInTuple(newParamsData);
        emit UpdatePlayerParams(
            NewParams.owner, 
            NewParams.nwin,
            NewParams.n,
            NewParams.spos,
            NewParams.sneg);
        uint bsnap = round.GetBalancesSnap();
        uint psnap = round.GetParamsSnap();
        emit ReceiveLotEvent(
            _lotAddr,
            _owner,
            newBalance,
            psnap,
            bsnap
        );
    }


    function GetSnap() public view returns (uint256) {
        IRound round = IRound(roundAddr);
        return round.GetSnap();
    }

    function GetRound() public view returns(address){
        return roundAddr;
    }

    function GetSnapRound() public view returns(uint256){
        IRound round = IRound(roundAddr);
        return round.GetSnapshot();
    }

    function GetInitSnapRound() public view returns(uint256){
        IRound round = IRound(roundAddr);
        return round.GetInitSnap();
    }

    function VerifyProofRes(
        address _addr,
        uint _H,
        uint _balance
    ) public view returns(bool){
        Proof.ProofRes memory proofRes = Proof.NewProof(_addr, 0, _H, _balance);
        IRound round = IRound(roundAddr);
        return round.VerifyProofRes(proofRes);
    }

    function VerifyParamsPlayer(
        bytes memory playerParamsData
    ) external view returns(bool){
        Params.PlayerParams memory params = Params.DecodePlayerParams(playerParamsData);
        IRound round = IRound(roundAddr);
        return round.VerifyParamsPlayer(params);
    }

    function GetParamsSnapshot() external view returns(uint){
        IRound round = IRound(roundAddr);
        return round.GetParamsSnap();
    }

    function GetDepositRound() external view returns(uint, uint){
        IRound round = IRound(roundAddr);
        return round.GetDeposit();
    }

    function SwapEthToDai(uint amountIn) external{
        IRound round = IRound(roundAddr);
        round.Swap1(amountIn);
    }

    function SwapDaiToEth(uint amountIn) external{
        IRound round = IRound(roundAddr);
        round.Swap2(amountIn);
    }

    function CurrentPrice1(uint _amountIn) external{
        IRound round = IRound(roundAddr);
        uint price =  round.GetPrice1(_amountIn);
        emit Price1(price);
    }

    function CurrentPrice2(uint _amountIn) external{
        IRound round = IRound(roundAddr);
        uint price =  round.GetPrice2(_amountIn);
        emit Price2(price);
    }
}  
