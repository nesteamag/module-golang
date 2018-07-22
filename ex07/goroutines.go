package goroutines

func Process(input chan string) chan string {
var(
	output = make(chan string)
	boolean = make(chan bool)
)
 go func() {
output <- "(" + <-input + ")"
boolean <-true
}()

 go func() {
<-boolean
close(output)
}()

return output
}

