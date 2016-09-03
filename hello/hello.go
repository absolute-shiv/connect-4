package main

import(
	"fmt"
	"os"
   	 "os/exec"
	)

var  board[6][7] int 


func showboard (){
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
	for i := 0;i < 6; i++{
		fmt.Println("|",board[i][0],"|",board[i][1],"|",board[i][2],"|",board[i][3],"|",board[i][4],"|",board[i][5],"|",board[i][6],"|")
	}
fmt.Println( "   1   2   3   4   5   6   7   "  )
}


func main() {
	showboard()
}
