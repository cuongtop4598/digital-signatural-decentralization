// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./IDC.sol";
import "./utils/Context.sol";
import "./StringsUtils.sol";

contract Document is Context, IDC {
                             
    struct User {
        string userID;
        bytes32 infoHash;
        bytes32 phoneHash;
        string publicKey;
        uint documentSize;
        mapping(uint => Doc) doc;
    }

    User[] users;

    struct Doc {
        bytes signature;
    }

    event IndexDocument(string publickey, uint256 numdoc, bytes signature);
    event StoreUserStatus(bytes32 infoHash, string publickey);
    event Status(string);
    /**
     * @dev Return True if user information is stored successfully, otherwise return False.
     */
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,string memory publicKey) 
    public override returns(bytes32) {
        bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail,publicKey);
        uint256 idx = users.length;
        users.push();
        User storage u = users[idx];
        u.userID = userID;
        u.infoHash = hashInfo;
        u.publicKey = publicKey;
        u.phoneHash = hashPhoneNumber(phone);
        return u.phoneHash;
    }
    /**
     * @dev Verify user information
    */
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,string memory publicKey)
    public override view returns(bool) {
         bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail,publicKey);
         uint i = 0;
         for (i = 0;  i <= users.length; i++ ) {
            if(compareBytes(users[i].infoHash, hashInfo)){
                return true;
            }
         }
         return false;
    }       

    function getHashUserInfo(string memory phone) public view returns(bytes32) {
        bytes32 phoneHash = hashPhoneNumber(phone);
        uint i = 0;
         for (i = 0;  i <= users.length; i++ ) {
            if(compareBytes(users[i].infoHash, phoneHash)){
                return users[i].infoHash;
            }
         }
        revert("not found");
    }

    /**
     * @dev Save Document to Chain
     * 
     * Return index of Document in the document list of the owner
     */
    function saveDoc(string memory phone,bytes memory signature) public override returns (uint256) {
        bytes32 phoneHash = hashPhoneNumber(phone);
        uint index = 1000;
        for (uint i = 0;  i <= users.length; i++ ) {
           if(compareBytes(users[i].phoneHash, phoneHash)){
                index = i;
            }
        }
        require(index < 1000);
        users[index].doc[users[index].documentSize].signature = signature;
        users[index].documentSize = users[index].documentSize + 1;
        emit IndexDocument( users[index].publicKey ,users[index].documentSize - 1, users[index].doc[users[index].documentSize-1].signature);
        return  users[index].documentSize-1;
    }

    /**
     * @dev Return True if Doc wasn't change else False
     */
    function verifyDoc(string memory phone, bytes32 digest, uint indexDoc) public view override returns(bool) {
        bytes32 phoneHash = hashPhoneNumber(phone);
        uint index = 100000000000;
        for (uint i = 0;  i <= users.length; i++ ) {
           if(compareBytes(users[i].phoneHash, phoneHash)){
                index = i;
            }
        }
        require(index != 100000000000);
        return keccak256(abi.encodePacked(recoverSigner(digest,users[index].doc[indexDoc].signature))) == keccak256(abi.encodePacked(users[index].publicKey));
    }
    
    function getSignature(string memory phone, uint indexDoc) public view returns (bytes memory) {
        string memory hashPhone = bytes32ToString(keccak256(abi.encodePacked(phone)));
       uint index;
        for (uint i = 0;  i < users.length; i++ ) {
            if(keccak256(abi.encodePacked(users[i].phoneHash)) == keccak256(abi.encodePacked(hashPhone))){
                index = i;
            }
        }
        Doc memory d = users[index].doc[indexDoc];
        return d.signature;
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
     * @dev Return a hash of user information 
     */
    function hashUserInfo(string memory name, string memory cmnd, string memory dOB, string memory phone, string memory gmail, string memory publickey) 
    private pure returns (bytes32) {
        return keccak256(abi.encodePacked(name,cmnd,dOB,phone,gmail,publickey));
    }

    function hashPhoneNumber(string memory phone) 
    private pure returns (bytes32) {
        return keccak256(abi.encodePacked(phone));
    }

    function compareBytes(bytes32 a, bytes32 b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }
    function bytes32ToString(bytes32 _bytes32) public pure returns (string memory) {
        uint8 i = 0;
        while(i < 32 && _bytes32[i] != 0) {
            i++;
        }
        bytes memory bytesArray = new bytes(i);
        for (i = 0; i < 32 && _bytes32[i] != 0; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }
}

