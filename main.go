package main

import (
  /*
  #cgo CFLAGS: -I/usr/include/mysql
  #include <mysql.h>
  #include <string.h>
  #include <stdlib.h>
  #include <stdio.h>
  */
  "C"
  "unicode/utf8"
  "os"
  "os/exec"
  "fmt"
)

//export myexec
func myexec (initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, is_null *C.char, error *C.char) *C.char { 
  out, err := exec.Command(C.GoString(*args.args)).Output()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  result = C.CString(string(out))
  *length = C.ulong(utf8.RuneCountInString(C.GoString(result)))
  return result
}

//export myexec_init
func myexec_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool{
  if(int(args.arg_count) != 1) {
    return 1
  } else {
    return 0
  }
}

func main(){
}
