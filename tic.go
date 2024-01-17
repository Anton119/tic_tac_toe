package main

import "fmt" 

func main () {
main_arr:= [3][3]string {

{"X", "O", "O"},
{"O","X","_"},
{"X","_","O"}}
print_board(main_arr)
fmt.Println()

coordinate(main_arr)
fmt.Println()
fmt.Println()
fmt.Println("Выйграл по горизонтали - " + X_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по вертикали - " + Y_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл по диагонали - " + Z_win_or_loose(main_arr))
fmt.Println()
fmt.Println("Выйграл - " + win_or_loose(main_arr))
fmt.Println()
fmt.Println(is_full(main_arr))
fmt.Println()
print_boards(all_moves(main_arr,"X"))
fmt.Println(all_moves(main_arr, "X"))
fmt.Println()
fmt.Println(all_moves(main_arr, "O"))
}


func win_or_loose (main_arr [3][3]string) string {
   if X_win_or_loose(main_arr) == "X" || Y_win_or_loose(main_arr) == "X" || Z_win_or_loose(main_arr) == "X" { return "X"
   } else if X_win_or_loose(main_arr) == "O" || Y_win_or_loose(main_arr) == "O" || Z_win_or_loose(main_arr) == "O" { return "O"
} else { return "_" }
}


/*xxx
0x_
___


xxx
0x0
___


xxx
0x_
0__

x:=42
fmt.print(int x)

tutorial hell

*/

func all_moves (main_arr [3][3]string, sign string)[][3][3]string {
var boards[][3][3]string
for y:=0; y<3; y++ {
 for x:=0; x<3; x++ {
 if main_arr[y][x] == "_" { ma_copy:=main_arr; ma_copy[y][x] = sign; boards=append(boards, ma_copy) }
 }
 }
 return boards
 }
 
 func print_boards(boards [][3][3]string){
 for b:=0; b<len(boards); b++ {
 print_board(boards[b])
 }
 fmt.Println()
 }

// печатает поле игры
func print_board (main_arr [3][3]string) {
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
// возвращает кто выйграл по горизонтали 
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
// возвращает кто выйграл по вертикали
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
// возвращает кто выйграл по диагонали
func Z_win_or_loose (main_arr [3][3]string) string {
	for y:=0; y<3; y++ {
	if main_arr[0][0]  == "X" && main_arr[1][1]  == "X" && main_arr[2][2]  == "X" || main_arr[0][2]  == "X" && main_arr[1][1] == "X" && main_arr[2][0]  == "X" { return "X"
 } else if main_arr[0][0] == "O" && main_arr[1][1] == "O" && main_arr[2][2] == "O" || main_arr[0][2] == "O" && main_arr[1][1] == "O" && main_arr[2][0]  == "O" { return "O"
 } else { continue }
 }
 return "_"
 }

func is_full (main_arr [3][3]string) bool {
   for y:=0; y<3; y++ {
      for x:=0; x<3; x++ {
         if main_arr[y][x] == "_" { return false
         } else { continue }
      }
   }
   return true 
}

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