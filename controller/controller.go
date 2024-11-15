package controller

import (
	"aosmanova/doodocs/models"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func ArchiveInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer file.Close()

	GetInfo(file, header)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Arch Info"))
}
func GetInfo(file multipart.File, header *multipart.FileHeader) *models.ZipInfo {
	defer file.Close()
	tmplPath := fmt.Sprintf("/tmp/%s", header.Filename)

	c, err := os.Create(tmplPath)
	if err != nil {
		fmt.Printf("error creating %s\n", tmplPath)
		return &models.ZipInfo{}
	}
	defer c.Close()

	_, err = io.Copy(c, file)
	if err != nil {
		fmt.Printf("error copying %s\n", file)
		return &models.ZipInfo{}
	}

	readfile, err := zip.OpenReader(tmplPath)
	if err != nil {
		log.Println(err)
		return &models.ZipInfo{}
	}
	data := &models.ZipInfo{
		Filename:     header.Filename,
		Archive_size: float64(header.Size),
	}
	for _, file := range readfile.File {
		f, err := file.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer f.Close()
		buf := make([]byte, 512)
		r, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(data.Files) > 0 {
			// Добавляем в существующий FileInfo
			data.Files[len(data.Files)-1].File_path = append(data.Files[len(data.Files)-1].File_path, file.Name)
		} else {
			// Создаём новую запись для FileInfo
			data.Files = append(data.Files, models.FileInfo{
				File_path: []string{file.Name}, // Инициализация списка путей
				Size:      float64(file.CompressedSize64),
				Mimetype:  http.DetectContentType(buf[:r]),
			})
		}

		name := file.Name
		fmt.Printf("file name %s\n", name)
		data.Total_size += float64(file.CompressedSize64)
		fmt.Printf("file total_size %f\n", data.Total_size)

	}
	data.Total_files = float64(len(readfile.File))
	fmt.Printf("total_files %f\n", data.Total_files)
	return data
}

func CreateArchive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create Arch"))
}

func ArchiveSend(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("archive send"))

}
