package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)
var waitGroup sync.WaitGroup
// definir a quantidade de cpus que quero na aplicação
func init(){
	runtime.GOMAXPROCS(2)
}
func main() {
	// saber a quantidade de CPU que ta sendo utilizado para o código
	fmt.Println(runtime.NumCPU())
	// adicionar waitGroups
	waitGroup.Add(2)
	go runProcess("P1",20)
	go runProcess("P2",20)
	waitGroup.Wait()

}


func runProcess(name string, total int ){
		for i:=0; i< total; i++ {
			fmt.Println(name,"->",i)
			// tempo setado somente para ficar mais fácil a visualização
			t := time.Duration(rand.Intn(255))
			time.Sleep(time.Microsecond *t)
		}
		waitGroup.Done()
}


// PRocessos concorrentes ocorrem juntos em cores diferentes