// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

enum State {
    Created,
    Pending,
    Completed,
    Failed
}

struct Operation {
    address managerAddr; // 20 bytes
    State state; // 1 byte
    uint64 taskId; // 8 bytes
    bytes optData;
}


contract Codec {
    constructor(uint256 nonce, Operation[] memory opts) {}
}
