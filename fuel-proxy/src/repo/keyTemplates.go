package repo

import "fmt"

const (
	utxoHashmapKeyTemplate                   = "fuel-ee-proxy/utxo_id-coin_encoded:%s:%s:%s"
	UtxoLastProcessedBlockHashmapKeyTemplate = "fuel-ee-proxy/utxo_id-coin_encoded-block"
)

func FormUtxoHashmapKeyTemplate(ownerId string, txId string, txOutputIdx string) string {
	return fmt.Sprintf(utxoHashmapKeyTemplate, ownerId, txId, txOutputIdx)
}
