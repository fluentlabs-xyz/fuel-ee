package utxoRepo

type utxoInternalStatus int

const (
	TxCommitInternalStatusNew = iota
	TxCommitInternalStatusCompleting
	TxCommitInternalStatusCompleted
	TxCommitInternalStatusFailed
)

func (w utxoInternalStatus) ToString() string {
	return [...]string{"new", "completing", "completed", "failed"}[w]
}
