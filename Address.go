package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Address struct {
	apiV1URL string
	logger   func(string)
}

func NewAddress(config SDKConfig) *Address {
	return &Address{
		apiV1URL: config.BaseAPIURL,
		logger:   func(arg string) { log.Println(arg) },
	}
}

func (a *Address) GetAddressBalance(address string) (*AddressBalance, error) {
	url := fmt.Sprintf("%s/address/state/%s", a.apiV1URL, address)
	data, err := a.get(url)
	if err != nil {
		return nil, err
	}
	var balance AddressBalance
	if err := json.Unmarshal(data, &balance); err != nil {
		return nil, err
	}
	return &balance, nil
}

func (a *Address) GetAddressTransactions(address string) (*AddressTransactions, error) {
	url := fmt.Sprintf("%s/address/transactions/%s", a.apiV1URL, address)
	data, err := a.get(url)
	if err != nil {
		return nil, err
	}
	var transactions AddressTransactions
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, err
	}
	return &transactions, nil
}

func (a *Address) GetUnconfirmedAddressTransactions(address string) (*AddressTransactions, error) {
	url := fmt.Sprintf("%s/address/unconfirmed/transactions/%s", a.apiV1URL, address)
	data, err := a.get(url)
	if err != nil {
		return nil, err
	}
	var transactions AddressTransactions
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, err
	}
	return &transactions, nil
}

func (a *Address) GetAddressUTXO(address string) (*AddressUTXOs, error) {
	url := fmt.Sprintf("%s/address/utxo/%s", a.apiV1URL, address)
	data, err := a.get(url)
	if err != nil {
		return nil, err
	}
	var utxos AddressUTXOs
	if err := json.Unmarshal(data, &utxos); err != nil {
		return nil, err
	}
	return &utxos, nil
}

func (a *Address) get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		a.logger(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		a.logger(err.Error())
		return nil, err
	}

	return body, nil
}
