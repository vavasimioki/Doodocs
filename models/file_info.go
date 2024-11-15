package models

type ZipInfo struct {
	Filename     string
	Archive_size float64
	Total_size   float64
	Total_files  float64
	Files        []FileInfo
}
type FileInfo struct {
	File_path []string
	Size      float64
	Mimetype  string
}
