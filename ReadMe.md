# Uplaod File
```http
POST localhost:8080/
```

**At Least One Of Form Fields Should Be Specified**

| FormFields | type | D`escription |
| :--- | :--- | :--- |`
| BrandGuid | uuid | Specify Brand UUID|
| CGuid | uuid | Specify Company UUID|
| RpAccGuid | uuid | Specify Admin UUID|
| UGuid | uuid | Specify User UUID|
| ResGuid | uuid | Specify Resource UUID|

| FileField | type | Description |
| :--- | :--- | :--- |
| Files| File/image| Put any File Or Image You can put Multiple **To the Same FileField**|

**MultipartForm Post Request With Files**

**For Now Working With Jpeg And Png Only**

## Response
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

```http
GET localhost:8080/
```

## Request
```json
{
    "File":"image",
    "Guid":"fd0b38ec-c409-48c5-9183-ae10d7ad8295",
    "Size":"R"
}
```

| FormFields | type | Description |
| :--- | :--- | :--- |
| File | string | fileType ('Image', 'doc')|
| Guid | uuid | TargetGuid|
| Size | uuid | **If Image** ('S'-small, 'M'-medium, 'R'-real)|

## Response

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
