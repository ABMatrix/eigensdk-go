package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type TransactionOperation string
type FeeLevel string

const (
	ContractCall TransactionOperation = "CONTRACT_CALL"
	Transfer     TransactionOperation = "TRANSFER"
	Mint         TransactionOperation = "MINT"
	Burn         TransactionOperation = "BURN"
	TypedMessage TransactionOperation = "TYPED_MESSAGE"
	Raw          TransactionOperation = "RAW"

	FeeLevelHigh   FeeLevel = "HIGH"
	FeeLevelMedium FeeLevel = "MEDIUM"
	FeeLevelLow    FeeLevel = "LOW"
)

type account struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type extraParams struct {
	Calldata string `json:"contractCallData"`
}

type ContractCallRequest struct {
	Operation       TransactionOperation `json:"operation"`
	ExternalTxID    string               `json:"externalTxId"`
	AssetID         AssetID              `json:"assetId"`
	Source          account              `json:"source"`
	Destination     account              `json:"destination"`
	Amount          string               `json:"amount"`
	ExtraParameters extraParams          `json:"extraParameters"`
	// In case a transaction is stuck, specify the hash of the stuck transaction to replace it
	// by this transaction with a higher fee, or to replace it with this transaction with a zero fee and drop it from the blockchain.
	ReplaceTxByHash string `json:"replaceTxByHash"`
	// GasPrice and GasLimit are the gas price and gas limit for the transaction.
	// If GasPrice is specified (non-1559), MaxFee and PriorityFee are not required.
	GasPrice string `json:"gasPrice"`
	GasLimit string `json:"gasLimit"`
	// MaxFee and PriorityFee are the maximum and priority fees for the transaction.
	// If the transaction is stuck, the Fireblocks platform will replace the transaction with a new one with a higher fee.
	// These fields are required if FeeLevel is not specified.
	MaxFee      string `json:"maxFee"`
	PriorityFee string `json:"priorityFee"`
	// FeeLevel is the fee level for the transaction which Fireblocks estimates based on the current network conditions.
	// The fee level can be HIGH, MEDIUM, or LOW.
	// If MaxFee and PriorityFee are not specified, the Fireblocks platform will use the default fee level MEDIUM.
	// Ref: https://developers.fireblocks.com/docs/gas-estimation#estimated-network-fee
	FeeLevel FeeLevel `json:"feeLevel"`
}

type ContractCallResponse struct {
	ID     string   `json:"id"`
	Status TxStatus `json:"status"`
}

func NewContractCallRequest(
	externalTxID string,
	assetID AssetID,
	sourceAccountID string,
	destinationAccountID string,
	amount string,
	calldata string,
	replaceTxByHash string,
	gasPrice string,
	gasLimit string,
	maxFee string,
	priorityFee string,
	feeLevel FeeLevel,
) *ContractCallRequest {
	return &ContractCallRequest{
		Operation:    ContractCall,
		ExternalTxID: externalTxID,
		AssetID:      assetID,
		Source: account{
			Type: "VAULT_ACCOUNT",
			ID:   sourceAccountID,
		},
		// https://developers.fireblocks.com/reference/transaction-sources-destinations
		Destination: account{
			Type: "EXTERNAL_WALLET",
			ID:   destinationAccountID,
		},
		Amount: amount,
		ExtraParameters: extraParams{
			Calldata: calldata,
		},
		ReplaceTxByHash: replaceTxByHash,
		GasPrice:        gasPrice,
		GasLimit:        gasLimit,
		MaxFee:          maxFee,
		PriorityFee:     priorityFee,
		FeeLevel:        feeLevel,
	}
}

func (f *client) ContractCall(ctx context.Context, req *ContractCallRequest) (*ContractCallResponse, error) {
	f.logger.Debug("Fireblocks call contract", "req", req)
	res, err := f.makeRequest(ctx, "POST", "/v1/transactions", req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	var response ContractCallResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &ContractCallResponse{
		ID:     response.ID,
		Status: response.Status,
	}, nil
}
