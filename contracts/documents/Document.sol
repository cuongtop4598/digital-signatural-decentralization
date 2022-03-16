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
        address publicKey;
        uint documentSize;
        mapping(uint => Doc) doc;
    }

    User[] users;

    struct Doc {
        bytes signature;
    }

    event Number(string);
    /**
     * @dev Return True if user information is stored successfully, otherwise return False.
     */
    function storeUser(string memory userID,string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) 
    public override returns(bytes32) {
        bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail,publicKey);
        uint256 idx = users.length;
        users.push();
        User storage u = users[idx];
        u.userID = userID;
        u.infoHash = hashInfo;
        u.publicKey = publicKey;
        u.phoneHash = hashPhoneNumber(phone);
        u.documentSize = 0;
        return u.phoneHash;
    }
    /**
     * @dev Verify user information
    */
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail, address publicKey)
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
            if(compareBytes(users[i].phoneHash, phoneHash)){
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
    function saveDoc(string memory phone,bytes memory signature) public override {
        bytes32 phoneHash = hashPhoneNumber(phone);
        uint i = 0;
        for ( i = 0;  i <= users.length; i++ ) {
           if(compareBytes(users[i].phoneHash, phoneHash)){
                 string memory numDoc = uint2str(users[i].documentSize);
                 emit Number(numDoc);
                 users[i].doc[users[i].documentSize].signature = signature;
                 users[i].documentSize = users[i].documentSize + 1;
            }
        }
       revert("not found");
    }

    function getEthSignedMessageHash(bytes32 _messageHash)
            public
            pure
            returns (bytes32)
        {
            /*
            Signature is produced by signing a keccak256 hash with the following format:
            "\x19Ethereum Signed Message\n" + len(msg) + msg
            */
            return
                keccak256(
                    abi.encodePacked("\x19Ethereum Signed Message:\n32", _messageHash)
                );
        }
    /**
     * @dev Return True if Doc wasn't change else False
     */
    function verifyDoc(string memory phone, bytes32 digest, uint indexDoc) public view override returns(bool) {
        bytes memory signature;
        address publicKey;
        (signature, publicKey) = getSignature(phone, indexDoc);
        bytes32 ethMessageHash = getEthSignedMessageHash(digest);
        address signer = recoverSigner(ethMessageHash,signature);
        return signer == publicKey;
    }


    function getSinger(string memory phone, bytes32 digest, uint indexDoc) public view  returns(bytes memory) {
        bytes memory signature;
        address publicKey;
        (signature, publicKey) = getSignature(phone, indexDoc);
        bytes32 ethMessageHash = getEthSignedMessageHash(digest);
        bytes memory signer = abi.encodePacked(recoverSigner(ethMessageHash,signature));
        return signer;
    }
    
    function getSignature(string memory phone, uint indexDoc) public view returns (bytes memory, address ) {
        string memory hashPhone = bytes32ToString(keccak256(abi.encodePacked(phone)));
       uint index;
        for (uint i = 0;  i < users.length; i++ ) {
            if(keccak256(abi.encodePacked(users[i].phoneHash)) == keccak256(abi.encodePacked(hashPhone))){
                index = i;
            }
        }
        Doc memory d = users[index].doc[indexDoc];
        return (d.signature,  users[index].publicKey);
    }

   function recoverSigner(bytes32 _ethSignedMessageHash, bytes memory _signature)
        public
        pure
        returns (address)
    {
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);

        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function splitSignature(bytes memory sig)
        public
        pure
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
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
    function hashUserInfo(string memory name, string memory cmnd, string memory dOB, string memory phone, string memory gmail, address publickey) 
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
    function compareString(string memory a, string memory b) public pure returns (bool) {
        if(bytes(a).length != bytes(b).length) {
        return false;
        } else {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
        }
    }
    function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) {
            k = k-1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
