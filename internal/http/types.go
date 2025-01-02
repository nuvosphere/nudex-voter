package http

import "encoding/json"

// UNCHECKED.
type FireblocksWebhookRequest struct {
	Type      string          `json:"type"`
	TenantId  string          `json:"tenantId"`
	Timestamp int64           `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

// TRANSACTION_CREATED, TRANSACTION_STATUS_UPDATED, TRANSACTION_APPROVAL_STATUS_UPDATED.
type TransactionDetails struct {
	ID                            string                   `json:"id"`
	ExternalTxID                  string                   `json:"externalTxId"`
	Status                        string                   `json:"status"`
	SubStatus                     string                   `json:"subStatus"`
	TxHash                        string                   `json:"txHash,omitempty"`
	Operation                     string                   `json:"operation"`
	Note                          string                   `json:"note"`
	AssetID                       string                   `json:"assetId"`
	AssetType                     string                   `json:"assetType"`
	Source                        TransferPeerPathResponse `json:"source"`
	SourceAddress                 string                   `json:"sourceAddress,omitempty"`
	Destination                   TransferPeerPathResponse `json:"destination"`
	Destinations                  []DestinationsResponse   `json:"destinations,omitempty"`
	DestinationAddress            string                   `json:"destinationAddress,omitempty"`
	DestinationAddressDescription string                   `json:"destinationAddressDescription,omitempty"`
	DestinationTag                string                   `json:"destinationTag,omitempty"`
	AmountInfo                    AmountInfo               `json:"amountInfo"`
	TreatAsGrossAmount            bool                     `json:"treatAsGrossAmount"`
	FeeInfo                       FeeInfo                  `json:"feeInfo"`
	FeeCurrency                   string                   `json:"feeCurrency"`
	CreatedAt                     int64                    `json:"createdAt"`
	LastUpdated                   int64                    `json:"lastUpdated"`
	CreatedBy                     string                   `json:"createdBy"`
	SignedBy                      []string                 `json:"signedBy,omitempty"`
	RejectedBy                    string                   `json:"rejectedBy,omitempty"`
	ExchangeTxID                  string                   `json:"exchangeTxId,omitempty"`
	CustomerRefID                 string                   `json:"customerRefId,omitempty"`
	ReplacedTxHash                string                   `json:"replacedTxHash,omitempty"`
	NumOfConfirmations            int                      `json:"numOfConfirmations"`
	BlockInfo                     BlockInfo                `json:"blockInfo"`
	Index                         int                      `json:"index,omitempty"`
	SystemMessages                []SystemMessageInfo      `json:"systemMessages,omitempty"`
	AddressType                   string                   `json:"addressType,omitempty"`
	RequestedAmount               float64                  `json:"requestedAmount,omitempty"`
	Amount                        float64                  `json:"amount,omitempty"`
	NetAmount                     float64                  `json:"netAmount,omitempty"`
	AmountUSD                     float64                  `json:"amountUSD,omitempty"`
	ServiceFee                    float64                  `json:"serviceFee,omitempty"`
	NetworkFee                    float64                  `json:"networkFee,omitempty"`

	// NetworkRecords               []NetworkRecord               `json:"networkRecords,omitempty"`
	// AuthorizationInfo            `json:"authorizationInfo"`
	// AmlScreeningResult           `json:"amlScreeningResult"`
	// ExtraParameters              TransactionExtraParameters    `json:"extraParameters"`
	// SignedMessages               []SignedMessage               `json:"signedMessages,omitempty"`
	// RewardsInfo                  `json:"rewardsInfo"`
}

type AmountInfo struct {
	Amount          string `json:"amount"`
	RequestedAmount string `json:"requestedAmount"`
	NetAmount       string `json:"netAmount"`
	AmountUSD       string `json:"amountUSD"`
}

type FeeInfo struct {
	NetworkFee string `json:"networkFee"`
	ServiceFee string `json:"serviceFee"`
}

type TransferPeerPathResponse struct {
	Type    string `json:"type"`
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	SubType string `json:"subType"`
}

type DestinationsResponse struct {
	Amount                        string                   `json:"amount"`
	Destination                   TransferPeerPathResponse `json:"destination"`
	AmountUSD                     float64                  `json:"amountUSD"`
	DestinationAddress            string                   `json:"destinationAddress"`
	DestinationAddressDescription string                   `json:"destinationAddressDescription,omitempty"`
	CustomerRefID                 string                   `json:"customerRefId,omitempty"`

	// AmlScreeningResult         `json:"amlScreeningResult"`
}

type SystemMessageInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type BlockInfo struct {
	BlockHeight string `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
}
