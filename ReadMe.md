# SapApi Marketplace — Golang Implementation

A part of SapApi Marketplace service running on Go, to handle operations like Bleve search, Websockets, File uploads.

 
### 📂 File Upload & Retrieval API

This API allows uploading and retrieving files (currently supports JPEG and PNG).
It supports associating uploaded files with various entities such as brands, companies, admins, users, or resources.

---

## **1. Upload File**

### **Endpoint**

```http
POST /  
Host: localhost:8080
Content-Type: multipart/form-data
```

> **Note:** At least **one** of the following **form fields** must be provided.

### **Form Fields**

| Field       | Type | Description              |
| ----------- | ---- | ------------------------ |
| `BrandGuid` | UUID | Brand identifier         |
| `CGuid`     | UUID | Company identifier       |
| `RpAccGuid` | UUID | Admin account identifier |
| `UGuid`     | UUID | User identifier          |
| `ResGuid`   | UUID | Resource identifier      |

### **File Field**

| Field   | Type       | Description                                                                    |
| ------- | ---------- | ------------------------------------------------------------------------------ |
| `Files` | File/Image | One or more files to upload. Multiple files can be sent in the **same** field. |

**Supported formats:** `JPEG`, `PNG`

---

### **Example Request**

Multipart form-data with files:

```http
POST / HTTP/1.1
Host: localhost:8080
Content-Type: multipart/form-data; boundary=----BOUNDARY

------BOUNDARY
Content-Disposition: form-data; name="BrandGuid"
d3d10f66-a812-4a5b-8d32-b7e60184c7f5
------BOUNDARY
Content-Disposition: form-data; name="Files"; filename="download.jpeg"
Content-Type: image/jpeg

(binary content)
------BOUNDARY--
```

---

### **Example Response**

```json
{
  "status": true,
  "message": "Ok",
  "errors": null,
  "data": [
    {
      "Name": "download.jpeg",
      "Path": "./uploads/RpAccGuid/d0a51101-3bb6-4d7b-ae62-d69a6b7ab40f/images/<FSIZE>/d0a51101-3bb6-4d7b-ae62-d69a6b7ab40f.jpeg",
      "Guid": "d0a51101-3bb6-4d7b-ae62-d69a6b7ab40f",
      "Error": "",
      "TargetGuid": "f3f00f78-20b9-475d-86a1-70f376f8bb0a"
    }
  ]
}
```

---

## **2. Retrieve File**

### **Endpoint**

```http
GET /  
Host: localhost:8080
Content-Type: application/json
```

---

### **Request Body**

```json
{
  "File": "image",
  "Guid": "fd0b38ec-c409-48c5-9183-ae10d7ad8295",
  "Size": "R"
}
```

| Field  | Type   | Description                                                        |
| ------ | ------ | ------------------------------------------------------------------ |
| `File` | string | File type — `"image"` or `"doc"`                                   |
| `Guid` | UUID   | `TargetGuid` of the uploaded file                                  |
| `Size` | string | **If Image**: `"S"` = small, `"M"` = medium, `"R"` = real/original |

---

### **Example Response**

```json
{
  "status": true,
  "message": "Ok",
  "errors": null,
  "data": {
    "ImgId": 1,
    "EmpId": null,
    "BrandId": null,
    "ResId": 1,
    "CId": null,
    "UId": null,
    "RpAccId": null,
    "ImgGuid": "fd0b38ec-c409-48c5-9183-ae10d7ad8295",
    "FileName": "052288a2121369f3254fb3bfe63a.png",
    "FilePath": "uploads/commerce/Resource/1/images/R/052288a2121369f3254fb3bfe63a.png",
    "CreatedDate": "2020-10-27T17:48:16.190445Z",
    "ModifiedDate": "2020-09-21T17:12:00.43Z",
    "GCRecord": null
  }
}
```

---

## **📌 Notes**

* The API is designed for **multipart file uploads** with entity association.
* Currently supports **image storage & retrieval**, but can be extended for other file types.
* File size variations (`S`, `M`, `R`) are generated for images to optimize usage.
* Implemented in **Go (Golang)** with SapApi marketplace architecture.
