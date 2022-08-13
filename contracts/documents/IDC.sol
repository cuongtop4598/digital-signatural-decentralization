// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev Interface for document contract
 */
interface IDC {
    function storeUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) external returns (bytes32);
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) external returns(bool);
    function createSingleSignerDocument(string memory phone, bytes memory signature) external returns (bytes32 phoneHash) ;
    function createMultipleSignerDocument(address partner, bytes memory signature) external returns (uint256 DocID, uint256 startDate , bool success);
    function signMultipleSignerDocument(uint256 DocID, bytes memory signatureB) external returns (uint256 signDate, bool success);
    function verifyDoc(string memory phone, string memory digest, uint indexDoc) external returns(bool); 
}