package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) IncreaseAssetPosition(assetID string, shares int) {
	UpdateAssetPosition(i, assetID, shares)
}

func (i *Investor) DecreaseAssetPosition(assetID string, shares int) {
	UpdateAssetPosition(i, assetID, -shares)
}

func UpdateAssetPosition(i *Investor, assetID string, shares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		if shares > 0 {
			i.AddAssetPosition(NewInvestorAssetPosition(assetID, shares))
		}
	} else {
		assetPosition.Shares += shares
	}
}
