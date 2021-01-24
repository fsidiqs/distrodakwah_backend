package productlibrary

import (
	"encoding/json"
	"strings"
)

type Variant struct {
	Name string `json:"name"`
}

func NewVariantProductVariant(variantReqJson string) ([]string, error) {
	variantReqs := []string{}
	err := json.NewDecoder(strings.NewReader(variantReqJson)).Decode(&variantReqs)

	if err != nil {
		return nil, err
	}
	return variantReqs, nil
}
