package main

import (
       "io"
       "os"
       "os/signal"
       "syscall"
       "log"
       "context"
       "net/http"
       "fmt"
       "net"
       "flag"
)


const index_page =  `<!DOCTYPE html>
<html>
<head>
    <script>
    function set_yzm() {
      var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            rmp = JSON.parse(this.responseText);
            var yzm_img = document.getElementById("yzm_img");
            yzm_img.setAttribute('src', rmp.data.img);
            var yzm_2 = document.getElementById("yzm_2");
            yzm_2.setAttribute('value', rmp.data.unique_str);
        }
      };
      xhttp.open("GET", "http://sample.com");
      xhttp.send();
    };
    
    set_yzm();
    
    function loadDoc() {
      var usr = document.getElementById("usr").value;
      var psw = document.getElementById("psw").value;

      var uid = document.getElementById("user_id").value;
      var oid = document.getElementById("order_id").value;
      var yzm = document.getElementById("yzm").value;
      var yzm_2 = document.getElementById("yzm_2").value;

      var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
        if (this.readyState == 4) {
            if(this.status == 200){
                set_yzm();
                rtx = this.responseText;
                if (rtx == 'ha ha ok *~o~*'){
                    document.getElementById("usr").value = '';
                    document.getElementById("psw").value = '';
                    document.getElementById("user_id").value = '';
                    document.getElementById("order_id").value = '';
                    document.getElementById("yzm").value = '';
                };
                alert(rtx);
            } else{
                alert('服务器升级中 请稍候再试 或联系管理员');
            };
        }
      };
      xhttp.open("POST", "http://sample.com");
      // xhttp.setRequestHeader('Origin', 'fengshui.lbadvisor., com');
      xhttp.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
      xhttp.send();
    }
    </script>
</head>
<body>


<h2>Html Forms</h2>
<H1 id="demo">hulu.com</H1><br>
<div>
    user_name:
    <br>
    <input autocomplete="off" type="text" id ='usr' name="usr" value="" style="width: 180px">
    <br>
    password:
    <br>
    <input autocomplete="off" type="password" name="psw" id="psw" value="" style="width: 180px">
    <br>
    <br>
    user_id:
    <br>
    <input autocomplete="off" type="text" id ='user_id' name="user_id" value="" style="width: 180px">
    <br>
    order_id:
    <br>
    <input autocomplete="off" type="text" name="order_id" id="order_id" value="" style="width: 180px">
    <br>
    验证码:
    <br>
    <input autocomplete="off" placeholder="请输入校验码" style="width: 100px; height: 36px; border: 1px solid LightGray" type="text" rows="2" name='yzm', id='yzm'>
    <input type="hidden" id="yzm_2" name="yzm_2" value="">
    <img id='yzm_img' src="data:image/gif;base64,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" style="width: 100px; height: 37px;padding-top: 0px;vertical-align:middle" alt="" onclick="set_yzm()"/> <br>
        <br>
        <button type="button" style="font-size: 16px;width: 100px;" onclick="loadDoc()">Click Me!</button>
</div>

<p>If you click the "Submit" button, the form-data will be sent to a page called "/".</p>

</body>
</html>`


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

       var pt uint
       flag.UintVar(&pt, "port", 8080, "an listened tcp v4 port")
       flag.Parse()

       pts := fmt.Sprintf(":%d", pt)

       srv := http.Server{Addr: pts}
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, index_page)
       })

       idleConnsClosed := make(chan struct{})

       sigint := make(chan os.Signal, 1)
       signal.Notify(sigint, syscall.SIGTERM)
       signal.Notify(sigint, syscall.SIGINT)

       go func() {

              <-sigint
              if err := srv.Shutdown(context.Background()); err != nil {
                     // Error from closing listeners, or context timeout:
                     log.Printf("HTTP server Shutdown: %v", err)
                     os.Exit(0)
              }
              log.Printf("HTTP server Shutdown")
              close(idleConnsClosed)
              os.Exit(0)
       }()


       fmt.Printf("  http://%v:%d/\n\n", GetOutboundIP(), pt)

       if err := srv.ListenAndServe(); err != http.ErrServerClosed {
              // Error starting or closing listener:
              log.Printf("HTTP server ListenAndServe: %v", err)
              sigint<-syscall.SIGTERM
       }

       <-idleConnsClosed
}