package main
// go build png2jpg.go
// ./png2jpg
import "os"
import "fmt"
import "log"
import "image/png"
import "image/jpeg"
import "path/filepath"


func main() {
    args_exc_prog := os.Args[1:]
    len_args := len(args_exc_prog)
    if len_args == 0 {
        log.Fatal("no file gave")
    }
    fn := args_exc_prog[0]
    
    dir, fl := filepath.Split(fn)
    ext := filepath.Ext(fl)
    bse := fl[0:len(fl)-len(ext)]
    nfn := filepath.Join(dir, bse + ".jpg")

    file, err := os.Open(fn)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }

    img, err := png.Decode(file)
    if err != nil {
     log.Fatal(err)
    }

    out, err := os.Create(nfn)
    defer out.Close()
    if err != nil {
     log.Fatal(err)
    }
    // write new image to file
    jpeg.Encode(out, img, nil)
    fmt.Println(fn, " -> ", nfn)
}
