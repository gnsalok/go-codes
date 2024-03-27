// These init() functions can be used within a package block and 
// regardless of how many times that package is imported, the init() function will only be called once.

// This effectively allows us to set up database connections, 
// or register with various service registries, or perform any number of other tasks that you typically only want to do once.

package main

import(
	"fmt"

)


var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
    return 42
}

func init() {
    WhatIsThe = 0
}

func main() {
    if WhatIsThe == 0 {
        fmt.Println("It's all a lie.")
    }
}