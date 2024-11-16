package models

type ErrorResponse struct {
	// Statuscode string 	`json:"statuscode"`
	Message string `json:"message"`
}

type ZipInfo struct {
	Filename     string     `json:"filename"`
	Archive_size float64    `json:"archive_size"`
	Total_files  float64    `json:"total_files"`
	Files        []FileInfo `json:"files"`
}
type FileInfo struct {
	File_path string  `json:"file_path"`
	Size      float64 `json:"size"`
	Mimetype  string  `json:"mimetype"`
}
