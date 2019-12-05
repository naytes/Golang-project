//imports needed for assignment below
import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
)

func errorCheck(e error) { //sends to panic if error
	if e !=nil {
		panic(e) 
	}
}

func openFile() string {
        data, err := ioutil.ReadFile(os.Args[1]) //if you type "go run go.go testFile.txt" then it  opens the file "testFile.txt"
        errorCheck(err)
       // fmt.Print(string(data))
	return string(data) //casts to string
}

func numParse(file string) string{
	var output string
	for i := 0; i < len(file); i++ { //iterates through file
		if file[i] > 47 && file[i] < 58 { //if value between ascii values of numbers do nothing
                } else { //else add to output string
			output = output + string(file[i])
		}
	}
	//fmt.Print(file)
	return output
}

func outputFile(out string){
        var outputName string //string to hold file name
        reader := bufio.NewReader(os.Stdin) //new reader
	fmt.Print("Enter desired output file name (no need to add the '.txt'): ") 
        outputName, _ = reader.ReadString('\n') //reads next string
	outputName = outputName[:(len(outputName)-1)] //for some reason i kept getting '\n' at the end of my file string so this deletes it
        outputName = outputName + ".txt" //adds '.txt' to file name
        file2, err := os.Create(outputName) //create new file for output
        errorCheck(err)
        writer := bufio.NewWriter(file2) //makes write to file call
        _, err = writer.WriteString(out) //write the string to the file
        errorCheck(err)
        fmt.Println("Your output has been saved under: " + outputName)
        file2.Sync() //syncs file
        writer.Flush() //flushes


}

func main() {
        file := openFile() //opens new file based on command line argument "go run go.go file.txt"
	//fmt.Print(file)
	file = numParse(file) //removes all numbers from file
	fmt.Print(file) //prints the file (after numbers have been removed)
	outputFile(file) //saves it in an output file
	}
