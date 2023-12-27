package main

import "fmt" 

func main () {
main_arr:= [3][3]string {{"X", "O", "_"},{"O","X","_"},{"_","O","X"}}
tick_print(main_arr)
}

func tick_print (main_arr [3][3]string) {
for i:=0; i<3; i++ {
   for j:=0; j<3; j++ {
   fmt.Print(main_arr[i][j])
   }
   fmt.Println()
   }
}

func check_tick (main_arr [3][3]string) bool {
//если нет "_" , то ход невозможен (ничья)//
}

func win_or_loose (main_arr [3][3]string) string {
//кто выйграл//
}

 




/*xo_
ox_
_ox

arr:= [] string {}
[кол-во]тип_элеметов
*/