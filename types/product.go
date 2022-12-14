package types

type Products struct {
	ProductList []Product `json:"products,omitempty"`
}

type Product struct {
	Name           string           `json:"name,omitempty"`
	Visible        bool             `json:"visible,omitempty"`
	Description    string           `json:"description,omitempty"`
	Sku            string           `json:"sku,omitempty"`
	Stock          Stock            `json:"stock,omitempty"`
	Price          Price            `json:"price,omitempty"`
	AdditionalInfo []AdditionalInfo `json:"additionalInfoSections,omitempty"`
	ProductOptions []ProductOption  `json:"productOptions,omitempty"`
	Discount       Discount         `json:"discount,omitempty"`
}

type Stock struct {
	InventoryStatus string `json:"inventoryStatus,omitempty"`
	InStock         bool   `json:"inStock,omitempty"`
}

type Price struct {
	Currency        string         `json:"currency,omitempty"`
	Price           float64        `json:"price,omitempty"`
	DiscountedPrice float64        `json:"discountedPrice,omitempty"`
	Formatted       FormattedPrice `json:"formatted,omitempty"`
}

type FormattedPrice struct {
	Price           string `json:"price,omitempty"`
	DiscountedPrice string `json:"discounted_price,omitempty"`
}

type AdditionalInfo struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProductOption struct {
	OptionType string          `json:"optionType,omitempty"`
	Name       string          `json:"name,omitempty"`
	Choices    []ProductChoice `json:"choices,omitempty"`
}

type ProductChoice struct {
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	InStock     bool   `json:"inStock,omitempty"`
}

type Discount struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}
