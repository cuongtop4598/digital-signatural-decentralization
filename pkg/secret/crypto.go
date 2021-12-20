package secret


// import (
// 	"github.com/ethereum/go-ethereum/crypto"
// )


// func VerifySignature(address string, nonce string, signature string) error {
// 	hash := crypto.Keccak256Hash(nonce)
// 	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
// 	if err != nil {
// 		return err
// 	}
// 	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
// }