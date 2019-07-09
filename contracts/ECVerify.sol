pragma solidity >=0.5.0;



//TODO: extend with https://ethereum.stackexchange.com/questions/15364/ecrecover-from-geth-and-web3-eth-sign

contract ECVerify {
    // Duplicate Solidity's ecrecover, but catching the CALL return value
    function safer_ecrecover(bytes32 hash, uint8 v, bytes32 r, bytes32 s) internal returns (bool, address) {
        // We do our own memory management here. Solidity uses memory offset
        // 0x40 to store the current end of memory. We write past it (as
        // writes are memory extensions), but don't update the offset so
        // Solidity will reuse it. The memory used here is only needed for
        // this context.

        // FIXME: inline assembly can't access return values
        bool ret;
        address addr;

        assembly {
            let size := mload(0x40)
            mstore(size, hash)
            mstore(add(size, 32), v)
            mstore(add(size, 64), r)
            mstore(add(size, 96), s)

            // NOTE: we can reuse the request memory because we deal with
            //       the return code
            ret := call(3000, 1, 0, size, 128, size, 32)
            addr := mload(size)
        }

        return (ret, addr);
    }

    function ecrecovery(bytes32 hash, bytes memory sig) public returns (bool, address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if (sig.length != 65)
          return (false, address(0));

        // The signature format is a compact form of:
        //   {bytes32 r}{bytes32 s}{uint8 v}
        // Compact means, uint8 is not padded to 32 bytes.
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))

            // Here we are loading the last 32 bytes. We exploit the fact that
            // 'mload' will pad with zeroes if we overread.
            // There is no 'mload8' to do this, but that would be nicer.
            v := byte(0, mload(add(sig, 96)))

            // Alternative solution:
            // 'byte' is not working due to the Solidity parser, so lets
            // use the second best option, 'and'
            // v := and(mload(add(sig, 65)), 255)
        }

        // albeit non-transactional signatures are not specified by the YP, one would expect it
        // to match the YP range of [27, 28]
        //
        // geth uses [0, 1] and some clients have followed. This might change, see:
        //  https://github.com/ethereum/go-ethereum/issues/2053
        if (v < 27)
          v += 27;

        if (v != 27 && v != 28)
            return (false, address(0));

        return safer_ecrecover(hash, v, r, s);
    }

    function ecrecovery32(bytes32 hash, bytes32[3] memory sig) public returns (bool, address) {
        bytes32 r;
        bytes32 s;
        uint v;

        r = sig[0];
        s = sig[1];
        v = uint(sig[2]);

        if (v < 27)
            v += 27;
       
        if (v == 27)
            return safer_ecrecover(hash, 27, r, s);
        else if (v == 28)
            return safer_ecrecover(hash, 28, r, s);
        else    
            return (false, address(0));
    }

    function ecverify(bytes32 hash, bytes memory sig, address signer) public returns (bool) {
        bool ret;
        address addr;
        (ret, addr) = ecrecovery(hash, sig);
        return ret == true && addr == signer;
    }
}


