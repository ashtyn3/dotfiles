Vim�UnDo� �U��Q)���g�Yb��u!��ƨz���3(�&�O   5           `                       ``�Q    _�                     +   b    ����                                                                                                                                                                                                                                                                                                                                                             `_�V     �   *   ,   5      �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')5�_�                    ,   \    ����                                                                                                                                                                                                                                                                                                                                                             `_�Y     �   +   -   5      ]	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken)5�_�                    ,   f    ����                                                                                                                                                                                                                                                                                                                                                             `_�]    �               5   package user       import (   	"crypto/ecdsa"   	"fmt"   	"teth/auth"   	cr "teth/crypto"   	"teth/loggy"       1	"github.com/ethereum/go-ethereum/common/hexutil"   )	"github.com/ethereum/go-ethereum/crypto"   	"go.uber.org/zap"   )       %var logger *zap.Logger = loggy.Init()   var status int = 0       func Err(err string, e error) {   	if e != nil {   		status = 1   !		logger.Error(err, zap.Error(e))   	}   }   func (u *User) Create() bool {   	// private key encryption   (	privateKey, err := crypto.GenerateKey()   )	Err("Could not generate key pair ", err)   )	prvBytes := crypto.FromECDSA(privateKey)   '	hexPrv := hexutil.Encode(prvBytes)[2:]   *	u.PrvKey = cr.Encrypt(u.Password, hexPrv)       	// address creation   !	publicKey := privateKey.Public()   3	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)   		if !ok {   3		logger.Fatal("error casting public key to ECDSA")   	}   9	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()   	u.PubKey = address       	db := auth.Auth().DB.Init()   	query := fmt.Sprintf(`   �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')   h	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, signature)       	_, qErr := db.Exec(query)   #	Err("failed to create user", qErr)   	defer logger.Sync()   	if status == 0 {   		return true   	}   	return false   }5�_�                    ,   ^    ����                                                                                                                                                                                                                                                                                                                                                             `_�^     �   +   -   5      h	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, signature)5�_�                    ,   _    ����                                                                                                                                                                                                                                                                                                                                                             `_�_    �               5   package user       import (   	"crypto/ecdsa"   	"fmt"   	"teth/auth"   	cr "teth/crypto"   	"teth/loggy"       1	"github.com/ethereum/go-ethereum/common/hexutil"   )	"github.com/ethereum/go-ethereum/crypto"   	"go.uber.org/zap"   )       %var logger *zap.Logger = loggy.Init()   var status int = 0       func Err(err string, e error) {   	if e != nil {   		status = 1   !		logger.Error(err, zap.Error(e))   	}   }   func (u *User) Create() bool {   	// private key encryption   (	privateKey, err := crypto.GenerateKey()   )	Err("Could not generate key pair ", err)   )	prvBytes := crypto.FromECDSA(privateKey)   '	hexPrv := hexutil.Encode(prvBytes)[2:]   *	u.PrvKey = cr.Encrypt(u.Password, hexPrv)       	// address creation   !	publicKey := privateKey.Public()   3	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)   		if !ok {   3		logger.Fatal("error casting public key to ECDSA")   	}   9	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()   	u.PubKey = address       	db := auth.Auth().DB.Init()   	query := fmt.Sprintf(`   �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')   j	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, u.signature)       	_, qErr := db.Exec(query)   #	Err("failed to create user", qErr)   	defer logger.Sync()   	if status == 0 {   		return true   	}   	return false   }5�_�                    +   �    ����                                                                                                                                                                                                                                                                                                                                                             `_�d     �   *   ,   5      �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')5�_�                    +   �    ����                                                                                                                                                                                                                                                                                                                                                             `_�h     �   *   ,   5      �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s)5�_�      	              +   �    ����                                                                                                                                                                                                                                                                                                                                                             `_�j     �   *   ,   5      �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s)5�_�      
           	   +   �    ����                                                                                                                                                                                                                                                                                                                                                             `_�j    �               5   package user       import (   	"crypto/ecdsa"   	"fmt"   	"teth/auth"   	cr "teth/crypto"   	"teth/loggy"       1	"github.com/ethereum/go-ethereum/common/hexutil"   )	"github.com/ethereum/go-ethereum/crypto"   	"go.uber.org/zap"   )       %var logger *zap.Logger = loggy.Init()   var status int = 0       func Err(err string, e error) {   	if e != nil {   		status = 1   !		logger.Error(err, zap.Error(e))   	}   }   func (u *User) Create() bool {   	// private key encryption   (	privateKey, err := crypto.GenerateKey()   )	Err("Could not generate key pair ", err)   )	prvBytes := crypto.FromECDSA(privateKey)   '	hexPrv := hexutil.Encode(prvBytes)[2:]   *	u.PrvKey = cr.Encrypt(u.Password, hexPrv)       	// address creation   !	publicKey := privateKey.Public()   3	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)   		if !ok {   3		logger.Fatal("error casting public key to ECDSA")   	}   9	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()   	u.PubKey = address       	db := auth.Auth().DB.Init()   	query := fmt.Sprintf(`   �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')   j	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, u.signature)       	_, qErr := db.Exec(query)   #	Err("failed to create user", qErr)   	defer logger.Sync()   	if status == 0 {   		return true   	}   	return false   }5�_�   	              
   ,   `    ����                                                                                                                                                                                                                                                                                                                                                             ``�P     �   +   -   5      j	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, u.signature)5�_�   
                  ,   a    ����                                                                                                                                                                                                                                                                                                                                                             ``�P    �               5   package user       import (   	"crypto/ecdsa"   	"fmt"   	"teth/auth"   	cr "teth/crypto"   	"teth/loggy"       1	"github.com/ethereum/go-ethereum/common/hexutil"   )	"github.com/ethereum/go-ethereum/crypto"   	"go.uber.org/zap"   )       %var logger *zap.Logger = loggy.Init()   var status int = 0       func Err(err string, e error) {   	if e != nil {   		status = 1   !		logger.Error(err, zap.Error(e))   	}   }   func (u *User) Create() bool {   	// private key encryption   (	privateKey, err := crypto.GenerateKey()   )	Err("Could not generate key pair ", err)   )	prvBytes := crypto.FromECDSA(privateKey)   '	hexPrv := hexutil.Encode(prvBytes)[2:]   *	u.PrvKey = cr.Encrypt(u.Password, hexPrv)       	// address creation   !	publicKey := privateKey.Public()   3	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)   		if !ok {   3		logger.Fatal("error casting public key to ECDSA")   	}   9	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()   	u.PubKey = address       	db := auth.Auth().DB.Init()   	query := fmt.Sprintf(`   �		insert into users (first_name, last_name, username, email, password, prv_key, pub_key, msg_token,signature) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')   j	`, u.Firstname, u.Lastname, u.Username, u.Email, u.Password, u.PrvKey, u.PubKey, u.MsgToken, u.Signature)       	_, qErr := db.Exec(query)   #	Err("failed to create user", qErr)   	defer logger.Sync()   	if status == 0 {   		return true   	}   	return false   }5��