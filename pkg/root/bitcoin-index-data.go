// +build ignore
package main

import (
	"github.com/alee792/vugu-example/pkg/bitcoin"
	"github.com/vugu/vugu"
)

// BitcoinIndexData of the DOM.
type BitcoinIndexData struct {
	priceIndex bitcoin.PriceIndex
	isLoading  bool
}

// HandleClick fetches Bitcoin's current price index.
func (data *BitcoinIndexData) HandleClick(event *vugu.DOMEvent) {
	data.priceIndex = bitcoin.PriceIndex{}

	evt := event.EventEnv()

	go func() {
		evt.Lock()
		data.isLoading = true
		evt.UnlockRender()

		pi := bitcoin.GetPriceIndex()

		evt.Lock()
		defer evt.UnlockRender()
		data.priceIndex = *pi
		data.isLoading = false
	}()
}
