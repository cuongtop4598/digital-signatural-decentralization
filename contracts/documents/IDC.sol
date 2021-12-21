// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev ...
 */
interface IDC {
    /**
     */
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,bytes32 publicKey) external returns (bool);

    /**
     */
    function saveDoc(string memory userID, bytes memory signature) external returns (uint256);

    /**
     */
    //function verifyDoc(string memory userID, bytes32 digest, uint256 indexDoc) external returns (bool);

    /**
     */
    function getUserID(bytes32 publicKey) external returns (string memory);
}