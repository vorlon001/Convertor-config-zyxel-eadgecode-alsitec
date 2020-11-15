package main

import (
	"fmt"
	"io"
	"bytes"
	//"bufio"
	//"os"
	"encoding/json"
)

type msg struct {
  Text string
}

type Context struct {
	PidFileName string
	rPipe *io.PipeReader
	wPipe *io.PipeWriter
}

func main() {
	rPipe, wPipe := io.Pipe()
	defer rPipe.Close()
	
	context := Context{PidFileName: "sdgfsdfgsdf" , rPipe: rPipe, wPipe: wPipe}
	
        bufJson := new(bytes.Buffer)
        _  = json.NewEncoder(bufJson ).Encode(context)
	//fmt.Printf( ">> %#v %#v \n", string(bufJson.Bytes()) , bufJson )

	go func(context Context , bufJson *bytes.Buffer) {
		_, _ = fmt.Fprint(context.wPipe, string(bufJson.Bytes()) )
		_, _ = fmt.Fprint(context.wPipe, "\n\n" )
		context.wPipe.Close();
	}(context, bufJson)



//      reader := bufio.NewReader(r)
//	input, _ := reader.ReadString('\n')
//	fmt.Println("1)input:",input)

	contextPipe := Context{}
	decoder := json.NewDecoder(rPipe)
	if err := decoder.Decode(&contextPipe); err != nil {
		fmt.Printf(">>> %v",err)
		return
	}
	fmt.Printf("%#v \n",contextPipe)
}
