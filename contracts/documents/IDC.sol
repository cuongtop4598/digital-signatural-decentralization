// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev Interface for document contract
 */
interface IDC {
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) external returns (bytes32);
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) external returns(bool);
    function saveDoc(string memory userID, bytes memory signature) external returns (bytes32) ;
    function verifyDoc(string memory userID, string memory digest, uint indexDoc) external returns(bool); 
}