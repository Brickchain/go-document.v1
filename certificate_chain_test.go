package document

// import (
// 	"testing"
// 	"time"

// 	"gitlab.brickchain.com/brickchain/crypto"
// )

// func TestCertificateChain(t *testing.T) {
// 	rootKey, _ := crypto.NewKey()
// 	rootPK, _ := crypto.NewPublicKey(rootKey)
// 	subKey, _ := crypto.NewKey()
// 	subPK, _ := crypto.NewPublicKey(subKey)

// 	b := NewCertificateChain(rootPK, subPK, "sign", 3600)
// 	b.DocumentTypes = []string{"mandate"}

// 	if b.HasExpired() {
// 		t.Fatal("Certificate chain has expired")
// 	}
// }

// func TestCertificateChainHasExpired(t *testing.T) {
// 	rootKey, _ := crypto.NewKey()
// 	rootPK, _ := crypto.NewPublicKey(rootKey)
// 	subKey, _ := crypto.NewKey()
// 	subPK, _ := crypto.NewPublicKey(subKey)

// 	b := NewCertificateChain(rootPK, subPK, "sign", 1)
// 	b.DocumentTypes = []string{"mandate"}

// 	time.Sleep(time.Second * 2)
// 	if !b.HasExpired() {
// 		t.Fatal("Certificate chain should have expired")
// 	}
// }

// func TestCertificateChainAllowedType(t *testing.T) {
// 	rootKey, _ := crypto.NewKey()
// 	rootPK, _ := crypto.NewPublicKey(rootKey)
// 	subKey, _ := crypto.NewKey()
// 	subPK, _ := crypto.NewPublicKey(subKey)

// 	b := NewCertificateChain(rootPK, subPK, "sign", 3600)
// 	b.DocumentTypes = []string{"*"}

// 	if !b.AllowedType("mandate") {
// 		t.Fatal("Certificate chain should have allowed the mandate document type")
// 	}
// }

// func TestCertificateChainAllowedTypeFalse(t *testing.T) {
// 	rootKey, _ := crypto.NewKey()
// 	rootPK, _ := crypto.NewPublicKey(rootKey)
// 	subKey, _ := crypto.NewKey()
// 	subPK, _ := crypto.NewPublicKey(subKey)

// 	b := NewCertificateChain(rootPK, subPK, "sign", 3600)
// 	b.DocumentTypes = []string{"none"}

// 	if b.AllowedType("mandate") {
// 		t.Fatal("Certificate chain should not have allowed the mandate document type")
// 	}
// }
