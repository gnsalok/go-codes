package main 

import(
    "fmt"
    "sync"
)

func myFunc(wg *sync.WaitGroup){
    fmt.Println("Inside go routine!")
    wg.Done()
}

func main(){

    fmt.Println("First Line executed")
    var wg sync.WaitGroup
    wg.Add(1)

    // It will wait till go routine completed 
    go myFunc(&wg)
    wg.Wait()


    fmt.Println("Last line execute!")

}