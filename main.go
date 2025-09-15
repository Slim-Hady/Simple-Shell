package main;

import(
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)
/*
 shell is a command line u can think about it like interface 
 we write an input an take a response 
*/
func main(){

	reader := bufio.NewReader(os.Stdin);

	for {
		fmt.Print("=> ");
		input , err := reader.ReadString('\n');
		if (err != nil){
			fmt.Fprintln(os.Stderr,err);
			continue;
		}
		cleanInput := strings.TrimSpace(input);
		if (cleanInput == "exit"){
			fmt.Println("Bye :> ");
			break;
		}
		if err = execInput(input); err!=nil{
			fmt.Fprintln(os.Stderr, err);
		}
	}
}

func execInput(input string) error{
	
	if input == "" {
		return nil
	}

	cmd := exec.Command("powershell", "-Command", input)
	
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr;
	cmd.Stdout = os.Stdout;

	return cmd.Run();
}