package blockchain

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAddressBalance(t *testing.T) {
	mockServer := createMockServer("address/state/mockAddress", `{"balance": 42}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL + "/bgl/v1/blockchain",
	}
	addressAPI := NewAddress(config)

	balance, err := addressAPI.GetAddressBalance("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if balance.Balance != 42 {
		t.Errorf("Expected balance to be decoded, got %d", balance.Balance)
	}
}

func TestGetAddressTransactions(t *testing.T) {
	mockServer := createMockServer("address/transactions/mockAddress", `{"page": 2, "list": [{"txId": "abc"}]}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL + "/bgl/v1/blockchain",
	}
	addressAPI := NewAddress(config)

	transactions, err := addressAPI.GetAddressTransactions("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if transactions.Page != 2 || len(transactions.List) != 1 || transactions.List[0].TxID != "abc" {
		t.Errorf("Expected transactions to be decoded, got %#v", transactions)
	}
}

func TestGetUnconfirmedAddressTransactions(t *testing.T) {
	mockServer := createMockServer("address/unconfirmed/transactions/mockAddress", `{"page": 3, "list": [{"txId": "def"}]}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL + "/bgl/v1/blockchain",
	}
	addressAPI := NewAddress(config)

	transactions, err := addressAPI.GetUnconfirmedAddressTransactions("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if transactions.Page != 3 || len(transactions.List) != 1 || transactions.List[0].TxID != "def" {
		t.Errorf("Expected unconfirmed transactions to be decoded, got %#v", transactions)
	}
}

func TestGetAddressUTXO(t *testing.T) {
	mockServer := createMockServer("address/unconfirmed/transactions/mockAddress", `[{"txId": "utxo", "vOut": 1}]`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL + "/bgl/v1/blockchain",
	}
	addressAPI := NewAddress(config)

	utxos, err := addressAPI.GetAddressUTXO("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(*utxos) != 1 || (*utxos)[0].TxID != "utxo" || (*utxos)[0].VOut != 1 {
		t.Errorf("Expected UTXOs to be decoded, got %#v", utxos)
	}
}

func createMockServer(path, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bgl/v1/blockchain/"+path {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
}
