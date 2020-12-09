package productlibrary

import "mime/multipart"

type ProductImage struct {
	FileName string
	Content  multipart.File
}
