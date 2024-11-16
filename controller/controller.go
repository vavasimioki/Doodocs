package controller

import (
	"aosmanova/doodocs/models"
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
)

func ArchiveInformation(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		logger.Error(err.Error())
		return
	}
	defer file.Close()

	// fmt.Println("name", header.Filename)
	// fmt.Println("zip size", float64(header.Size))
	// z, err := zip.OpenReader(header.Filename)
	// if err != nil {
	// 	logger.Error(err.Error(), "openReading error")
	// }

	// var totalsize float64
	// for _, f := range z.File {
	// 	fmt.Printf("filename %s\n", f.Name)
	// 	// fmt.Printf("file_size %d\n", f.CompressedSize)
	// 	fmt.Printf("file_size64 %d\n", f.CompressedSize64)
	// 	totalsize += float64(f.CompressedSize64)
	// 	fmt.Printf("total_size", totalsize)
	// 	mymtipe := http.DetectContentType(f.Extra)
	// 	fmt.Printf("file_mymtipe %s\n", mymtipe)
	// }
	GetInfo(file, header, logger)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Arch Info"))
}
func GetInfo(file multipart.File, header *multipart.FileHeader, logger *slog.Logger) *models.ZipInfo {
	tempFile := fmt.Sprintf(header.Filename)
	createFile, err := os.Create(tempFile)
	if err != nil {
		logger.Error(err.Error(), "err creating file")
	}
	defer createFile.Close()

	_, err = io.Copy(createFile, file)
	if err != nil {
		logger.Error(err.Error(), "error copying file")
		return nil
	}
	filesZip, err := zip.OpenReader(tempFile)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Printf("zipSize %d\n", header.Size)
	fmt.Printf("zipName %s\n", header.Filename)
	var filesize_total float64
	name := []string{}

	for _, file := range filesZip.File {

		if file.CompressedSize64 == 0 {
			continue
		}
		f, err := file.Open()
		if err != nil {
			logger.Error(err.Error(), "error openig file")
			continue
		}
		defer f.Close()
		buf := make([]byte, 512)
		r, err := f.Read(buf)
		if err != nil {
			logger.Error(err.Error(), "error reading file")
			continue
		}
		name = append(name, file.Name[0:])
		fmt.Printf("file_name %s\n", file.Name[0:])
		fmt.Printf("file_size64 %d\n", file.CompressedSize64)
		fmt.Printf("mymtipe %s\n", http.DetectContentType([]byte(buf[:r])))
		filesize_total += float64(file.CompressedSize64)
	}
	if len(name) > 0 {
		fmt.Printf("files_count %d\n", len(name))
	}
	fmt.Printf("file_size_total %f\n", filesize_total)
	// absPath, err := filepath.Abs(tempFile)
	// if err != nil {
	// 	logger.Error(err.Error(), "error finding file path")
	// 	return nil
	// }
	// fmt.Println(absPath)
	return &models.ZipInfo{}
	// defer file.Close()
	// tmplPath := fmt.Sprintf("/tmp/%s", header.Filename)

	// c, err := os.Create(tmplPath)
	// if err != nil {
	// 	fmt.Printf("error creating %s\n", tmplPath)
	// 	return &models.ZipInfo{}
	// }
	// defer c.Close()

	// _, err = io.Copy(c, file)
	// if err != nil {
	// 	fmt.Printf("error copying %s\n", file)
	// 	return &models.ZipInfo{}
	// }

	// readfile, err := zip.OpenReader(tmplPath)
	// if err != nil {
	// 	log.Println(err)
	// 	return &models.ZipInfo{}
	// }
	// data := &models.ZipInfo{
	// 	Filename:     header.Filename,
	// 	Archive_size: float64(header.Size),
	// }
	// for _, file := range readfile.File {
	// 	f, err := file.Open()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	defer f.Close()
	// 	buf := make([]byte, 512)
	// 	r, err := f.Read(buf)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	if len(data.Files) > 0 {
	// 		// Добавляем в существующий FileInfo
	// 		data.Files[len(data.Files)-1].File_path = append(data.Files[len(data.Files)-1].File_path, file.Name)
	// 	} else {
	// 		// Создаём новую запись для FileInfo
	// 		data.Files = append(data.Files, models.FileInfo{
	// 			File_path: []string{file.Name}, // Инициализация списка путей
	// 			Size:      float64(file.CompressedSize64),
	// 			Mimetype:  http.DetectContentType(buf[:r]),
	// 		})
	// 	}

	// 	name := file.Name
	// 	fmt.Printf("file name %s\n", name)
	// 	data.Total_size += float64(file.CompressedSize64)
	// 	fmt.Printf("file total_size %f\n", data.Total_size)

	// }
	// data.Total_files = float64(len(readfile.File))
	// fmt.Printf("total_files %f\n", data.Total_files)
	// return data
}

func CreateArchive(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	w.Write([]byte("create Arch"))
}

func ArchiveSend(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	w.Write([]byte("archive send"))

}
