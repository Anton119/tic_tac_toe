package main

import "fmt" 

func main () {
main_arr:= [3][3]string {

{"_", "_", "X"},
{"O","_","X"},
{"O","X","X"}}
tick_print(main_arr)
fmt.Println()
fmt.Println("Выйграл по горизонтали - " + X_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по вертикали - " + Y_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по диагонали - " + Z_win_or_loose(main_arr))
}
// печатает поле игры
func tick_print (main_arr [3][3]string) {
for x:=0; x<3; x++ {
   for y:=0; y<3; y++ {
   fmt.Print(main_arr[x][y])
   }
   fmt.Println()
   }
}
// печатает кто выйграл по горизонтали 
func X_win_or_loose (main_arr [3][3]string) string {
for x:=0; x<3; x++ {
X_check:=X_same_el(main_arr, x)
if X_check && main_arr[x][0] == "X" { return "X"
} else if X_check  && main_arr[x][0] == "O" { return "O" 
} else { continue }
}
return "_" 
}
// ищет похожий элемент для х
func X_same_el (main_arr [3][3]string, y int) bool {
   for x:=0; x<3-1; x++ {
   if main_arr[y][x] != main_arr[y][x+1]  { return false 
   } else { continue }
   }
   return true
   }
// печатает кто выйграл по вертикали
func Y_win_or_loose (main_arr [3][3]string) string {
   for y:=0; y<3; y++ {
      Y_check:=Y_same_el(main_arr, y)
      if Y_check && main_arr[0][y] == "X" { return "X"
      } else if Y_check && main_arr[0][y] == "O" { return "O"
   } else { continue }
   
   }
   return "_"
}
// ищет похожий элемент для y
func Y_same_el (main_arr [3][3]string, x int) bool {
   for y:=0; y<3-1; y++ {
      if main_arr[y][x] != main_arr[y+1][x] { return false 
      } else { continue }
   }
   return true
}
func Z_win_or_loose (main_arr [3][3]string) string {
   for 
   Z_check:=Z_same_el(main_arr, z)

}



/*func coordinate (main_arr [3][3]string, j int) string {
for i:=0


/*func check_tick (main_arr [3][3]string) bool {
//если нет "_" , то ход невозможен (ничья)//
}
*/

/*x==true  -> x

x=true
true=true -> true
x=false
false=true -> false
*/

/*
x1o
o45
oox
4
*/

/*xo_
ox_
_ox

arr:= [] string {}
[кол-во]тип_элеметов
*/