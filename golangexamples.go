package golangexamples
import "github.com/ehteshamz/greetings"
// ConcatSice is a function to return contents of slice concatenated together
func ConcatSice(sliceToConcat [] byte) string{
  var contents = ""
  for index := 0;  index < len(sliceToConcat); index++ {
     contents = contents+ string(sliceToConcat[index])
     if index != len(sliceToConcat) - 1{
       contents = contents+ "-"
     }
  }
  return  contents
}

// Encrypt is a function to return contents of slice concatenated together
func Encrypt (sliceToEncrypt *[] byte, ceaserCount int){
  for index := 0;  index < len(*sliceToEncrypt); index++ {
    (*sliceToEncrypt)[index] = ((*sliceToEncrypt)[index] + byte(ceaserCount)) %255
  }

}

// EZGreetings calls Sir Ehtesham's greetings package function
func EZGreetings(name string) string{
  return greetings.PrintGreetings(name)
}
