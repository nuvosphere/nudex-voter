// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TaskPayload {
    enum Chain {
        BITCOIN,
        ETHEREUM,
        SOLANA,
        SUI
    }

    enum AssetType {
        MAIN,
        ERC20
    }

    event WalletCreationRequest(
        uint8 version,
        uint8 taskType,
        uint32 account,
        Chain chain,
        uint8 index
    );

    event DepositRequest(
        uint8 version,
        uint8 taskType,
        string targetAddress,
        uint64 amount,
        Chain chain,
        uint32 chainId,
        uint64 blockHeight,
        string txHash,
        string contractAddress,
        string ticker,
        AssetType assetType,
        uint8 decimal
    );

    // 事件：提现请求
    event WithdrawalRequest(
        uint8 version,
        uint8 taskType,
        string targetAddress,
        uint64 amount,
        Chain chain,
        uint32 chainId,
        uint64 blockHeight,
        string txHash,
        string contractAddress,
        string ticker,
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
