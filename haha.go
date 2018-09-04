package main

import "flag"
import "fmt"
import "os"
import "log"
import "net"
import "net/http"


func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}


func main() {

    var dr string
    var pt uint

    flag.StringVar(&dr, "mulu", ".", "a direction path shared")
    flag.UintVar(&pt, "port", 9999, "an listened tcp v4 port")

    // Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()
    // flag.Usage()
    // Here we'll just dump out the parsed options and
    // any trailing positional arguments. Note that we
    // need to dereference the pointers with e.g. `*wordPtr`
    // to get the actual option values.
    // fmt.Println("dir:", dr)
    // fmt.Println("port:", pt)

    fi, err := os.Stat(dr)
    if err != nil {
        fmt.Println("!!! a direction path need")
        return
    }

    mode := fi.Mode()
    if !mode.IsDir(){
        fmt.Println("!!! a direction path need")
        return
    }

    fmt.Println("dir:", dr)
    // fmt.Println("port:", pt)
    pts := fmt.Sprintf(":%d", pt)
    // fmt.Printf("%v\n", pts)
    fmt.Printf("  http://%v:%d/\n\n", GetOutboundIP(), pt)
    log.Fatal(http.ListenAndServe(pts, http.FileServer(http.Dir(dr))))

}
