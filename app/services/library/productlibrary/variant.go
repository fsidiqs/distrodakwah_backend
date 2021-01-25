package productlibrary

import (
	"encoding/json"
	"strings"
)

type Variant struct {
	Name string `json:"name"`
}

func NewVariantProductVariant(variantReqJson string) ([]VPVariant, error) {
	variantReqs := []string{}
	err := json.NewDecoder(strings.NewReader(variantReqJson)).Decode(&variantReqs)
	if err != nil {
		return nil, err
	}
	vpVariants := make([]VPVariant, len(variantReqs))
	for i, variantReq := range variantReqs {
		vpVariants[i].Name = variantReq
	}
	return vpVariants, nil
}
