package main;

import(
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)
/*
 shell is basically a commandline program that takes user input interprets it and returns a response
*/
func main(){
	// Making an object ( reader ) for taking an input in go we say Create a buffered reader to take input
	reader := bufio.NewReader(os.Stdin);
	// this infinite loop because we will not exit until user write "exit"
	for {
		// thats mark to user know where to type commands
		fmt.Print("=> ");
		/*
		we will read an input and store it in variable (input) but if occure any error will go to variable err ( it is in go return the input and result ) 
  		*/
		input , err := reader.ReadString('\n');
		// nil mean no error occuer If err is not nil print the error and continue the loop
		if (err != nil){  
			fmt.Fprintln(os.Stderr,err); // write the error
			continue; 
		}
		// we take an input and leave a space ('\n') so we remove this space and see what user write if exit so exit else execute the command
		cleanInput := strings.TrimSpace(input);
		if (cleanInput == "exit"){
			fmt.Println("Bye :> ");
			break;
		}
		// here we execute the input and return error
		if err = execInput(input); err!=nil{
			fmt.Fprintln(os.Stderr, err);
		}
	}
}

func execInput(input string) error{
	// if the input is empty just no error occurred and wait user to write a comman	
	if input == "" {
		return nil
	}
	// here (without this line u will make a shell for unix so it will not work in windows this line u can see as a built in u don't have to make args )
	// if u use unix make the args 
	cmd := exec.Command("powershell", "-Command", input)
	/*
 	if u want to make it using args split the word and execute the args but that is will not work for windows user 
	args := strings.Split(input, " ") 
    cmd := exec.Command(args[0], args[1:]...)
	you have to continue if unix .......
 	*/
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr;
	cmd.Stdout = os.Stdout;

	return cmd.Run();
}
