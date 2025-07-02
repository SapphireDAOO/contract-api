package blockchain

type MarketplaceAction int

const (
	Pending MarketplaceAction = iota
	ResolveDispute
	DismissDispute
	SettleDispute
)
