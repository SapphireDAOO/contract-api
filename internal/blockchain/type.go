package blockchain

type MarketplaceAction int

const (
	Pending MarketplaceAction = iota
	Release
	DismissDispute
	SettleDispute
)
