package main
import (
    "fmt"
    "os"
    "io"
    "archive/tar"
    "compress/gzip"
)
func main() {



    // file write
    fw, err := os.Create("test4.tar.gz")
    if err != nil {
        panic(err)
    }
    defer fw.Close()
    // gzip write
    gw := gzip.NewWriter(fw)
    defer gw.Close()
    // tar write
    tw := tar.NewWriter(gw)
    defer tw.Close()
    // 打开文件夹
   
   
    // 遍历文件列表
   ///sadasda
       
        // 打开文件
        fg, err := os.Open("test4")
        fr,_:=fg.Stat()
        if err != nil {
            panic(err)
        }
        defer fg.Close()
       
        // 信息头
        h := new(tar.Header)
        h.Name = fr.Name()
         h.Size = fr.Size()
         h.Mode = int64(fr.Mode())
        h.ModTime = fr.ModTime()
        // 写信息头
        err = tw.WriteHeader(h)
        if err != nil {
            panic(err)
        }
        // 写文件
        _, err = io.Copy(tw, fg)
        if err != nil {
            panic(err)
        }
    
    fmt.Println("tar.gz ok")
}