package request

type CategoryReq struct {
	SubdepartmentID uint   `json:"subdepartment_id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	ParentID        uint   `json:"parent_id,omitempty"`
}

type BrandReq struct {
	VendorID uint   `json:"vendor_id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}

type DepartmentReq struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type SubdepartmentReq struct {
	DepartmentID uint8  `json:"department_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}
