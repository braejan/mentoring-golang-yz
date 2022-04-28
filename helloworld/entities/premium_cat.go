package entities

const (
	SIZE_SMALL  = "small"
	SIZE_MEDIUM = "medium"
	SIZE_LARGE  = "large"
)

type PremiumCat struct {
	*Cat
	Size string
}

func NewPremiumCat(cat *Cat, size string) (premiumCat *PremiumCat) {
	premiumCat = &PremiumCat{
		Cat:  cat,
		Size: size,
	}
	return
}
