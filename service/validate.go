package service

func IsValidContentType(contentType string) bool {
	AllowedTypes := map[string]bool{
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	}
	return AllowedTypes[contentType]

}
func IsValidContentTypeForMail(contentType string) bool {
	AllowedTypes := map[string]bool{
		"application/pdf": true,
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	}
	return AllowedTypes[contentType]

}
