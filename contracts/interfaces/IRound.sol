// SPDX-License-Identifier: MIT
pragma solidity  ^0.7.0 || ^0.8.0;

import "../libraries/Proof.sol";
import "../libraries/Params.sol";


interface IRound {
    function Enter(address _sender, uint256 _value) external;

    function StartRound() external returns(uint, uint);

    function CreateLot() external returns(address);

    function NewLot(
        address _lotAddr,
        uint256 _timeFirst,
        uint256 _timeSecond,
        uint256 _val,
        Proof.ProofRes memory proof
    ) external returns(uint, uint);

    function BuyLot(
        address _lotAddr,
        Proof.ProofRes memory proof,
        Proof.ProofEnoungPrice memory proofEP 
    ) external returns(uint, uint);


    function GetPlayer(address player) external view returns (uint256);


    function SendLot(
        address _lotAddr,
        Params.InitParams memory initParams
    ) external returns(uint amountOut);


    function ReceiveLot(
        address _lotAddr,
        Params.InitParams calldata _init,
        Proof.ProofRes calldata _proof,
        Params.PlayerParams calldata _params
    ) external returns(bytes memory, uint);



    function GetSnap() external view returns (uint256);

    function GetBalance() external view returns (uint256);

    function GetSnapshot() external view returns(uint256);

    function GetInitSnap() external view returns(uint256);

    function GetExchange() external view returns(address);

    function VerifyProofRes(
        Proof.ProofRes calldata proof
    ) external view returns(bool);

    function VerifyParamsPlayer(
        Params.PlayerParams calldata _params
    ) external view returns(bool);

    function GetParamsSnap() external view returns(uint);

    function GetBalancesSnap() external view returns(uint);

    function GetDeposit() external view returns(uint, uint);

    function Swap1(uint amountIn) external;

    function Swap2(uint amountIn) external;

    function GetPrice1(uint _amountIn) external returns(uint);
    function GetPrice2(uint _amountIn) external returns(uint);
}
