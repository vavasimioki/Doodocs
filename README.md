# API Usage Instructions

This README provides clear instructions for interacting with the API using `curl` commands.

---

## 1.Usage
```bash
go get github.com/go-yaml/yaml v2.1.0+incompatible
go get github.com/go-yaml/yaml v2.1.0+incompatible

## 2.Usage
```bash
go run ./cmd 

## Retrieve Archive Info

Get information about an archive file.

```bash
curl -i -X POST http://localhost:8000/api/archive/info \
    -F "file=@/path/to/your/file.zip"


 my case: (curl -i -X POST   --form "files[]=@/Users/asemospanova/Downloads/IMG_5604.JPG;type=image/jpeg"   --form "files[]=@/Users/asemospanova/Downloads/IMG_5604.JPG;type=image/jpeg"   http://localhost:8000/api/archive/create)


## Create ZIP file
Get information about an archive file.
```bash
curl -X POST http://localhost:8000/api/archive/create \
    --form "files[]=@/path/to/file1.jpg;type=image/jpeg" \
    --form "files[]=@/path/to/file2.jpg;type=image/jpeg"

  my case:  (curl -X POST --form "files[]=@/Users/asemospanova/Downloads/Tengizchevroil.docx;type=application/vnd.openxmlformats-officedocument.wordprocessingml.document" http://localhost:8000/api/archive/create)

    (curl -X POST   --form "files[]=@/Users/asemospanova/Downloads/IMG_5604.JPG;type=image/jpeg"   --form "files[]=@/Users/asemospanova/Downloads/IMG_5604.JPG;type=image/jpeg"   http://localhost:8000/api/archive/create)

## Send To Mails
Get File, emails than send file to them

```bash
curl -X POST http://localhost:8000/api/archive/mail/send \
    -H "Content-Type: multipart/form-data" \
    -F "file=@/path/to/document.docx;type=application/vnd.openxmlformats-officedocument.wordprocessingml.document" \
    -F "emails=email1@example.com,email2@example.com"

my case: (curl -X POST   http://localhost:8000/api/archive/mail/send -H "Content-Type: multipart/form-data" -F "file=@/Users/asemospanova/Downloads/Tengizchevroil.docx;type=application/vnd.openxmlformats-officedocument.wordprocessingml.document" -F "emails=vavasimioki117@gmail.com,vavasimioki117@gmail.com")


