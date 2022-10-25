# graphql-file-info
Graphql server which show upload files info

# start server (default graphql address http://localhost:8080/query)
```bash
$ go mod tidy
$ make run
```

# Queries
### Single file: _fileInfo_
```sh
query FileInfo($file: Upload!) {
    fileInfo(file: $file) {
        __typename
        file {
            mime
            name
            size
        }
    }
}

```
### Response
```json
{
  "data": {
    "result": {
      "__typename": "FileData",
      "mime": "image/png",
      "name": "color-bars-600.png",
      "size": "1.34KB"
    }
  }
}
```

### Multiple files: _filesInfo_
```sh
query FilesInfo($files: [Upload!]!) {
    filesInfo(files: $files) {
        ... on FilesInfoResult {
            files {
              mime
              name
              size
            }
        }
    }
}
```
### Response
```json
{
  "data": {
    "result": {
      "__typename": "FilesInfoResult",
      "files": [
        {
          "mime": "image/jpeg",
          "name": "Image001.jpg",
          "size": "713.88KB"
        },
        {
          "mime": "image/jpeg",
          "name": "Image002.jpg",
          "size": "725.23KB"
        }
      ]
    }
  }
}
```