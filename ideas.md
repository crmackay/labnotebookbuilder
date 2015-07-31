ideas:

```go

// recursive walking function
// maybe: but found items into a channel? for a different htmlFileHandler
// maybe: create a few handlers that do the conversion and collation and output // of html
func dirHandler (dir){
    fir i, nextFile := range foundFiles {
        fileHandler(nextFile)
    }
    for i, nextDir := range foundDirs {
        dirHandler(nextDir)
    }
}

func fileHandler (file) {

    }


type fileType interface {

}


type fileTypeHandler func(file) file

mdFileHandler := func(file){
    return newFile
}

svg


tileTypeHandlers := map[string]fileTypeHandler{
    "md": mdFileHandler,
    "img": imgFileHandler,
    "svg": svgFileHandler
    "html": htmlFileHandler,
}

file fldrContents struct {
    files []file
    dirs []directory
}

```

workflow:

- scan directory -->
- put found things into channel -->
- parse each object via specific handler --> 
- put objects into a map (map["absolutepath"]fldrContents) -->
- parse each folder contents -->
- create html -->
- put into channel -->
- write to file, overwriting whats there
