// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./IDC.sol";
import "./utils/Context.sol";
import "./StringsUtils.sol";
/**
 * @dev Implementation of the {IDC} interface
 * 
 */

 contract Document is Context, IDC {
                             
    struct User {
        string userID;
        bytes32 infoHash;
        string publicKey;
        uint256 numDoc;
        mapping(uint256 => Doc) docs;
    }
    struct UserList {
        mapping(string => User) _ulistpubKey;
        mapping(string => User) _ulistUID;
    }
    struct Doc {
        string userID;
        bytes signature;
        bytes finDocument;
    }
    UserList _userlist;

    /**
     * @dev Return True if user information is stored successfully, otherwise return False.
     */
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,string memory publicKey) 
    public override returns(bool) {
        bytes32 hashif = hashUserInfo(name,cmnd,dateOB,phone,gmail);
        
        User storage u = _userlist._ulistpubKey[publicKey];
        
        u.userID = userID;
        u.infoHash = hashif;
        u.publicKey = publicKey;

        u = _userlist._ulistUID[userID];

        return true;
    }
    /**
     * @dev Verify user information
    */
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,string memory publicKey)
    public override view returns(bool) {
         bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail);
         if (compareBytes(_userlist._ulistpubKey[publicKey].infoHash, hashInfo)){
             return true;
         } else {
             return false;
         }
    }                            
    /**
     * @dev Save Document to Chain
     * 
     * Return index of Document in the document list of the owner
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
     * @dev Return True if Doc wasn't change else False
     */
    function verifyDoc(string memory userID, bytes32 digest, uint indexDoc) public view override returns(bool) {
        User storage u = _userlist._ulistUID[userID];
        Doc memory d = u.docs[indexDoc];
        return keccak256(abi.encodePacked(recoverSigner(digest,d.signature))) == keccak256(abi.encodePacked(u.publicKey));
    }

    /**
     * @dev Return publicKey 
     * references to ERC1271 Standard Signature Validation Method
     */
    function recoverSigner(bytes32 _hash, bytes memory _signature) private pure returns(address) {
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);
        return ecrecover(_hash, v, r, s);
    }

    /**
     * @dev Return r,s,v in order is ...
     */
    function splitSignature(bytes memory sig) private pure returns(bytes32 r,bytes32 s, uint8 v) {
        require(sig.length == 65, "invalid signature length");

        assembly {
            /*
            First 32 bytes stores the length of the signature
            add(sig, 32) = pointer of sig + 32
            effectively, skips first 32 bytes of signature
            mload(p) loads next 32 bytes starting at the memory address p into memory
            */

            // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
            // second 32 bytes
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        // implicitly return (r, s, v)
    }

    /**
    * @dev Return hash UserInfo 
    */
    function GetHashUserInfo(string calldata userPubkey) public view returns (bytes32) {
        bytes32 hashInfo = _userlist._ulistpubKey[userPubkey].infoHash;
        return hashInfo;
    }  
    /**
     * @dev Return a hash of user information 
     */
    function hashUserInfo(string memory name, string memory cmnd, string memory dOB, string memory phone, string memory gmail) 
    private pure returns (bytes32) {
        return keccak256(abi.encodePacked(name,cmnd,dOB,phone,gmail));
    }

    function compareBytes(bytes32 a, bytes32 b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }
    
}
