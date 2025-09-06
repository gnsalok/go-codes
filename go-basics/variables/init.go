package main

import "fmt"

/* AnswerToLife() is guaranteed to run before init() is called, and init() is guaranteed to run before main() is called.
Keep in mind that init() is always called, regardless if there's main or not, so if you import a package that has an init function,
it will be executed.
*/
var WhatIsThe = AnswerToLife()

func AnswerToLife() int { // 1
	return 42
}

func init() { // 2
	WhatIsThe = 0
}

func main() { // 3
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}
