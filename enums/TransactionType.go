package enums

type TransactionType int

const (
	DEBIT TransactionType = iota
	CREDIT
)

func (t TransactionType) String() string {
	return [...]string{"Debit", "Credit"}[t]
}

func (t TransactionType) EnumIndex() int {
	return int(t)
}
