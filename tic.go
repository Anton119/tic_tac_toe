package main

import "fmt" 

func main () {
main_arr:= [3][3]string {

{"_", "O", "X"},
{"O","X","_"},
{"X","O","X"}}
//tick_print(main_arr)
fmt.Println()
coordinate(main_arr)
fmt.Println()
fmt.Println("Выйграл по горизонтали - " + X_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по вертикали - " + Y_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по диагонали - " + Z_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл - " + result(main_arr))
}

func result (main_arr [3][3]string) string {
   if X_win_or_loose(main_arr) == "X" || Y_win_or_loose(main_arr) == "X" || Z_win_or_loose(main_arr) == "X" { return "X"
   } else if X_win_or_loose(main_arr) == "O" || Y_win_or_loose(main_arr) == "O" || Z_win_or_loose(main_arr) == "O" { return "O"
} else { return "_" }
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

   func coordinate (main_arr [3][3]string) {
      coor:=0
      for y:=0; y<3; y++ {
         for x:=0; x<3; x++ {
            coor += 1
            switch main_arr[y][x] {
            case "X": fmt.Print("X")
            case "O": fmt.Print("O")
            case "_": fmt.Print(coor)
            }
         }
         fmt.Println()
      } 
      }
// печатает кто выйграл по горизонтали 
func X_win_or_loose (main_arr [3][3]string) string {
for y:=0; y<3; y++ {
X_check:=X_same_el(main_arr, y)
if X_check && main_arr[y][0] == "X" { return "X"
} else if X_check  && main_arr[y][0] == "O" { return "O" 
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
   for x:=0; x<3; x++ {
      Y_check:=Y_same_el(main_arr, x)
      if Y_check && main_arr[0][x] == "X" { return "X"
      } else if Y_check && main_arr[0][x] == "O" { return "O"
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
// печатает кто выйграл по диагонали
func Z_win_or_loose (main_arr [3][3]string) string {
	for y:=0; y<3; y++ {
	if main_arr[0][0]  == "X" && main_arr[1][1]  == "X" && main_arr[2][2]  == "X" || main_arr[0][2]  == "X" && main_arr[1][1] == "X" && main_arr[2][0]  == "X" { return "X"
 } else if main_arr[0][0] == "O" && main_arr[1][1] == "O" && main_arr[2][2] == "O" || main_arr[0][2] == "O" && main_arr[1][1] == "O" && main_arr[2][0]  == "O" { return "O"
 } else { continue }
 }
 return "_"
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