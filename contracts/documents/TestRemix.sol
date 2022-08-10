// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }
    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

library StringUtils {
    /// @dev Does a byte-by-byte lexicographical comparison of two strings.
    /// @return a negative number if `_a` is smaller, zero if they are equal
    /// and a positive numbe if `_b` is smaller.
    function compare(string memory _a, string memory _b) public pure returns (int) {
        bytes memory a = bytes(_a);
        bytes memory b = bytes(_b);
        uint minLength = a.length;
        if (b.length < minLength) minLength = b.length;
        //@todo unroll the loop into increments of 32 and do full 32 byte comparisons
        for (uint i = 0; i < minLength; i ++)
            if (a[i] < b[i])
                return -1;
            else if (a[i] > b[i])
                return 1;
        if (a.length < b.length)
            return -1;
        else if (a.length > b.length)
            return 1;
        else
            return 0;
    }
    /// @dev Compares two strings and returns true iff they are equal.
    function equal(string memory _a, string memory _b) public pure returns (bool) {
        return compare(_a, _b) == 0;
    }
    /// @dev Finds the index of the first occurrence of _needle in _haystack
    function indexOf(string memory _haystack, string memory _needle) public pure returns (int)
    {
    	bytes memory h = bytes(_haystack);
    	bytes memory n = bytes(_needle);
    	if(h.length < 1 || n.length < 1 || (n.length > h.length)) 
    		return -1;
    	else if(h.length > (2**128 -1)) // since we have to be able to return -1 (if the char isn't found or input error), this function must return an "int" type with a max length of (2^128 - 1)
    		return -1;									
    	else
    	{
    		uint subindex = 0;
    		for (uint i = 0; i < h.length; i ++)
    		{
    			if (h[i] == n[0]) // found the first char of b
    			{
    				subindex = 1;
    				while(subindex < n.length && (i + subindex) < h.length && h[i + subindex] == n[subindex]) // search until the chars don't match or until we reach the end of a or b
    				{
    					subindex++;
    				}	
    				if(subindex == n.length)
    					return int(i);
    			}
    		}
    		return -1;
    	}	
    }
}


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

contract Document is Context, IDC {
                             
    struct User {
        string userID;
        bytes32 infoHash;
        bytes32 phoneHash;
        address publicKey;
        uint documentSize;
        mapping(uint => Doc) doc;
    }
    struct Doc {
        bytes signature;
    }

    struct MultiSignerDoc {
        uint256 ID;
        uint256 startDate; // timestamp
        uint256 validDate;
        address partnerA; // the first signer 
        address partnerB; // the second signer
        bytes firstSignature;
        bytes secondSignature;
        bool valid;
    }

    User[] users;
    MultiSignerDoc[] multiSignerDocs;

    event NumberOfOwnerDocument(uint num);
    event CreateAccountSuccess(string userInfo);
    /*
     @dev The first user will sign their document and send to the other signer
    */
    function createMultipleSignerDocument(address partner, bytes memory signatureA) public override returns (uint256, uint256, bool) {
        if(partner == msg.sender) {
            return (0, 0 , false);
        }
        uint256 idx = multiSignerDocs.length;
        multiSignerDocs.push();
        MultiSignerDoc storage ms = multiSignerDocs[idx];
        ms.ID = idx;
        ms.partnerA = msg.sender;
        ms.partnerB = partner;
        ms.firstSignature = signatureA;
        ms.startDate = block.timestamp;
        ms.valid = false;
        return (ms.ID, ms.startDate , true);
    }

    /*
        @dev The second user use this function to sign the document 
    */
    function signMultipleSignerDocument(uint256 DocID, bytes memory signatureB) public override returns (uint256,bool) {
        if(DocID < multiSignerDocs.length) {
            if(multiSignerDocs[DocID].partnerB == msg.sender){
                multiSignerDocs[DocID].secondSignature = signatureB;
                multiSignerDocs[DocID].valid = true;
                multiSignerDocs[DocID].validDate = block.timestamp;
                return ( multiSignerDocs[DocID].validDate, true);
            } 
            return (0,false);
        }
        return (0,false);
    } 

    /**
     * @dev Return True if user information is stored successfully, otherwise return False.
     */
    function storeUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail,address publicKey) 
    public override returns(bytes32) {
        bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail,publicKey);
        uint256 idx = users.length;
        users.push();
        User storage u = users[idx];
        u.infoHash = hashInfo;
        u.publicKey = publicKey;
        u.phoneHash = hashPhoneNumber(phone);
        u.documentSize = 0;
        emit CreateAccountSuccess(string(abi.encodePacked(phone,'-',publicKey)));
        return u.phoneHash;
    }
    /**
     * @dev Verify user information
    */
    function verifyUser(string memory name,string memory cmnd, string memory dateOB, string memory phone,string memory gmail, address publicKey)
    public override view returns(bool) {
         bytes32 hashInfo = hashUserInfo(name,cmnd,dateOB,phone,gmail,publicKey);
         uint i = 0;
         for (i = 0;  i < users.length; i++ ) {
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
         return 0;
    }

    /**
     * @dev Save Document to Chain
     * 
     * Return index of Document in the document list of the owner
     */
    function createSingleSignerDocument(string memory phone,bytes memory signature) public override returns (bytes32) {
        bytes32 phoneHash = hashPhoneNumber(phone);
        uint i = 0;
        for ( i = 0;  i <= users.length; i++ ) {
           if(compareBytes(users[i].phoneHash, phoneHash)){
                 uint numDoc = users[i].documentSize;
                 emit NumberOfOwnerDocument(numDoc);
                 users[i].doc[users[i].documentSize].signature = signature;
                 users[i].documentSize = users[i].documentSize + 1;
                 return phoneHash;
            }
        }
        return phoneHash;
    }

    function getMessageHash(
        string memory _message
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_message));
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
    function verifyDoc(string memory phone, string memory digest, uint indexDoc) public view override returns(bool) {
        bytes memory signature;
        address publicKey;
        (signature, publicKey) = getSignature(phone, indexDoc);
        bytes32 ethMessageHash = getEthSignedMessageHash(getMessageHash(digest));
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
        for (uint i = 0;  i < users.length; i++ ) {
            if(keccak256(abi.encodePacked(users[i].phoneHash)) == keccak256(abi.encodePacked(hashPhone))){
                Doc memory d = users[i].doc[indexDoc];
                return (d.signature,  users[i].publicKey);
            }
        }
        revert('Not found');
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
}
