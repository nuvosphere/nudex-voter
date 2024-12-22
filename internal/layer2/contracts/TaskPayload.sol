// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TaskPayload {
    enum AddressCategory {
        BTC,
        EVM,
        SOL,
        SUI
    }

    enum AssetType {
        MAIN,
        ERC20
    }

    event WalletCreationRequest(
        uint8 version,
        uint32 account,
        AddressCategory addressType,
        uint8 index
    );

    event DepositRequest(
        uint8 version,
        string userTssAddress,
        uint64 amount,
        bytes32 chainId,// hex
        string txHash,
        string contractAddress,
        string ticker,
        AssetType assetType,
        uint8 decimal
    );
    
    event WithdrawalRequest(
        uint8 version,
        string userTssAddress,
        uint64 amount,
        bytes32 chainId,// hex
        string txHash,
        string contractAddress,
        bytes32 ticker,
        AssetType assetType,
        uint8 decimal,
        uint64 fee
    );

    event ConsolidationRequest(
        uint8 version,
        string userTssAddress,
        uint64 amount,
        bytes32 chainId,// hex
        string contractAddress,
        bytes32 ticker,
        AssetType assetType,
        uint8 decimal,
        uint64 fee
    );

    event TaskResult(uint8 version, bool success, uint8 errorCode);

    event WalletCreationResult(
        uint8 version,
        bool success,
        uint8 errorCode,
        string walletAddress
    );

    event DepositResult(uint8 version, bool success, uint8 errorCode);

    event WithdrawalResult(uint8 version, bool success, uint8 errorCode);
}
