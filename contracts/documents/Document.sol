// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./IDC.sol";
import "./extensions/IDCMetadata.sol";
import "../utils/Context.sol";

/**
 * @dev Implementation of the {IDC} interface
 * 
 */
contract Document is Context, IDC, IDCMetadata {

    struct User {
        string userID;
        bytes32 infoHash;
        bytes32 publicKey;
        uint256 numDoc;
        mapping(uint256 => Doc) docs;
    }
    struct UserList {
        mapping(bytes32 => User) _ulistpubKey;
        mapping(string => User) _ulistUID;
    }
    struct Doc {
        string userID;
        bytes signature;
        bytes finDocument;
    }
    UserList _userlist;
    /**
     */
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,bytes32 publicKey) 
    public override returns(bool) {
        bytes32 hashif = hashUserInfo(userID,name,cmnd,dateOB,phone,gmail);
        
        User storage u = _userlist._ulistpubKey[publicKey];
        
        u.userID = userID;
        u.infoHash = hashif;
        u.publicKey = publicKey;

        u = _userlist._ulistUID[userID];

        return true;
    }

    /**
     */
    function getUserID(bytes32 publicKey) public view override returns (string memory) {
        return _userlist._ulistpubKey[publicKey].userID;
    }
    
    /**
     * @dev Return 
     */
    function getPublicKey(string memory userID) private view returns (bytes32 ) {
        return _userlist._ulistUID[userID].publicKey;
    }

    /**
     */
    function saveDoc(string memory userID,bytes memory signature) public override returns (uint256) {
        uint256 numDoc = _userlist._ulistUID[userID].numDoc;
        numDoc++;
        Doc storage doc = _userlist._ulistUID[userID].docs[numDoc];
        doc.userID = userID;
        doc.signature = signature;
        doc.finDocument = signature;
        return numDoc;
    }

    /**
     */
    // function verifyDoc(string memory userID, bytes32 digest, uint indexDoc) public override returns(bool) {
    //     User storage u = _userlist._ulistUID[userID];
    //     Doc memory d = u.docs[indexDoc];
    //     return recoverSigner(digest,d.signature) == u.publicKey;
    // }

    // /**
    //  * @dev Return publicKey 
    //  * references to ERC1271 Standard Signature Validation Method
    //  */
    // function recoverSigner(bytes32 _hash, bytes memory _signature) private returns(address) {
    //     (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);
    //     return ecrecover(_hash, v, r, s);
    // }
    // /**
    //  */
    // function splitSignature(bytes memory sig) private pure returns(bytes32 r,bytes32 s, uint8 v) {
    //     require(sig.length == 65, "invalid signature length");

    //     assembly {
    //         /*
    //         First 32 bytes stores the length of the signature
    //         add(sig, 32) = pointer of sig + 32
    //         effectively, skips first 32 bytes of signature
    //         mload(p) loads next 32 bytes starting at the memory address p into memory
    //         */

    //         // first 32 bytes, after the length prefix
    //         r := mload(add(sig, 32))
    //         // second 32 bytes
    //         s := mload(add(sig, 64))
    //         // final byte (first byte of the next 32 bytes)
    //         v := byte(0, mload(add(sig, 96)))
    //     }

    //     // implicitly return (r, s, v)
    // }
    /**
     */
    function hashUserInfo(string memory uID, string memory name, string memory cmnd, string memory dOB, string memory phone, string memory gmail) 
    private pure returns (bytes32) {
        return keccak256(abi.encodePacked(uID,name,cmnd,dOB,phone,gmail));
    }
}