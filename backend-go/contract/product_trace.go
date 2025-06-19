package product

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"product_trace/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProductTrace struct {
    Contract *contracts.Contracts
    Client   *ethclient.Client
}

type Status struct {
    Status    string
    Details   string
    Timestamp int64
}

func NewProductTrace() *ProductTrace {

    // Debug: Kiểm tra biến môi trường
    infuraKey := os.Getenv("INFURA_KEY")
    contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
    privateKeyHex := os.Getenv("PRIVATE_KEY")
    log.Println("INFURA_KEY:", infuraKey)
    log.Println("CONTRACT_ADDRESS:", contractAddressHex)
    log.Println("PRIVATE_KEY:", len(privateKeyHex), "chars") // Không in trực tiếp private key để bảo mật

    if infuraKey == "" {
        log.Panic("INFURA_KEY không được thiết lập")
    }
    if contractAddressHex == "" {
        log.Panic("CONTRACT_ADDRESS không được thiết lập")
    }

    client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + infuraKey)
    if err != nil {
        log.Panic("Không thể kết nối tới Sepolia Testnet: ", err)
    }

    contractAddress := common.HexToAddress(contractAddressHex)
    if contractAddress == (common.Address{}) {
        log.Panic("CONTRACT_ADDRESS không hợp lệ")
    }

    contract, err := contracts.NewContracts(contractAddress, client)
    if err != nil {
        log.Panic("Không thể khởi tạo contract: ", err)
    }

    return &ProductTrace{
        Contract: contract,
        Client:   client,
    }
}

func (p *ProductTrace) AddProduct(name, manufacturer, initialStatus, initialDetails string) (int64, error) {
    privateKeyHex := os.Getenv("PRIVATE_KEY")
    if privateKeyHex == "" {
        log.Panic("PRIVATE_KEY không được thiết lập")
    }

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return 0, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return 0, err
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := p.Client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return 0, err
    }
    gasPrice, err := p.Client.SuggestGasPrice(context.Background())
    if err != nil {
        return 0, err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia chain ID
    if err != nil {
        return 0, err
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(500000)
    auth.GasPrice = gasPrice

    tx, err := p.Contract.AddProduct(auth, name, manufacturer, initialStatus, initialDetails)
    if err != nil {
        return 0, err
    }

    receipt, err := bind.WaitMined(context.Background(), p.Client, tx)
    if err != nil {
        return 0, err
    }
    if receipt.Status == 0 {
        return 0, err
    }

    productCount, err := p.Contract.ProductCount(nil)
    if err != nil {
        return 0, err
    }

    return productCount.Int64(), nil
}

func (p *ProductTrace) UpdateStatus(productID int64, status, details string) error {
    privateKeyHex := os.Getenv("PRIVATE_KEY")
    if privateKeyHex == "" {
        log.Panic("PRIVATE_KEY không được thiết lập")
    }

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return err
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := p.Client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return err
    }
    gasPrice, err := p.Client.SuggestGasPrice(context.Background())
    if err != nil {
        return err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia chain ID
    if err != nil {
        return err
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice = gasPrice

    tx, err := p.Contract.UpdateStatus(auth, big.NewInt(productID), status, details)
    if err != nil {
        return err
    }

    receipt, err := bind.WaitMined(context.Background(), p.Client, tx)
    if err != nil {
        return err
    }
    if receipt.Status == 0 {
        return err
    }

    return nil
}

func (p *ProductTrace) GetProduct(productID int64) (string, string, error) {
    result, err := p.Contract.GetProduct(nil, big.NewInt(productID))
    if err != nil {
        return "", "", err
    }
    return result.Name, result.Manufacturer, nil
}

func (p *ProductTrace) GetCurrentStatus(productID int64) (Status, error) {
    result, err := p.Contract.GetCurrentStatus(nil, big.NewInt(productID))
    if err != nil {
        return Status{}, err
    }
    return Status{
        Status:    result.Status,
        Details:   result.Details,
        Timestamp: result.Timestamp.Int64(),
    }, nil
}

func (p *ProductTrace) GetProductStatuses(productID int64) ([]Status, error) {
    statuses, err := p.Contract.GetProductStatuses(nil, big.NewInt(productID))
    if err != nil {
        return nil, err
    }
    result := make([]Status, len(statuses))
    for i, s := range statuses {
        result[i] = Status{
            Status:    s.Status,
            Details:   s.Details,
            Timestamp: s.Timestamp.Int64(),
        }
    }
    return result, nil
}