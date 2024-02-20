package main

import "fmt" 

func main () {
//X - AI, O - HUMAN
main_arr:= [3][3]string {

{"X", "O", "_"},
{"_","_","_"},
{"O","O","_"}}
//number_arr:= []int{6,78,32,567,1}
print_board(main_arr)
fmt.Println()
coordinate(main_arr)
fmt.Println()
//fmt.Println(max_or_min_element(number_arr, func(x int, y int)bool{return x<y}))
//fmt.Println(max_or_min_element(number_arr, func(x int, y int)bool{return x>y}))
/*fmt.Println()
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
fmt.Println()
fmt.Println(count_moves(main_arr,"X"))
*/
fmt.Println(evaluate_moves(main_arr,"X",0 ))
fmt.Println()
print_board(best_board(main_arr, "X"))
fmt.Println("???????")
print_boards(all_moves(main_arr,"X"))
fmt.Println()
//str_print_board(all_moves(main_arr, "X"))
}



func win_or_loose (main_arr [3][3]string) string {
   if X_win_or_loose(main_arr) == "X" || Y_win_or_loose(main_arr) == "X" || Z_win_or_loose(main_arr) == "X" { return "X"
   } else if X_win_or_loose(main_arr) == "O" || Y_win_or_loose(main_arr) == "O" || Z_win_or_loose(main_arr) == "O" { return "O"
} else { return "_" }
}

func another_player (player string) string {
if player == "X" {return "O"
   } else { return "X"}
}

// (тренировочная функция) считает все возможные варианты для ОСТАЛЬНЫХ ШАГОВ
func count_moves (main_arr [3][3]string, player string) int {
//main_arr:=level[i]
sum_of_levels:=0
if is_full(main_arr) == false && win_or_loose(main_arr) == "_"  {
level:= all_moves(main_arr, player)
   for i:=0;i<len(level);i++ { 
      sum_of_levels+=count_moves(level[i], another_player(player))
   }
}
// добавляем начальную доску (1) к сумме остальных досок (вариантов ходов)
return sum_of_levels+1
}

//дополненная функция count_moves , реализующая алгоритм min_max
func evaluate_moves (main_arr[3][3]string, player string, level int) int {
if win_or_loose(main_arr) == "X" /*&& is_full(main_arr) == false*/ {
    fmt.Println("final level",level)
    print_board(main_arr)
    fmt.Println("+1")
    return 1
   } else if win_or_loose(main_arr) == "O"/* && is_full(main_arr) == false*/ {
    fmt.Println("final level",level)
    print_board(main_arr)
    fmt.Println("-1")
    return -1
   } else if is_full(main_arr) == true {
    fmt.Println("final level",level)
    print_board(main_arr)
    fmt.Println(" 0")
    return 0 }
       evaluation:=all_moves(main_arr, player)
       var sum_of_evaluations []int
             for i:=0; i<len(evaluation); i++ {
             // запись результатов оценок веток в массив                  // player = another_player(player)
             sum_of_evaluations=append(sum_of_evaluations, evaluate_moves(evaluation[i], another_player(player), level+1 ))
             }
             fmt.Println("level ",level)
             str_print_board(evaluation)
                print_el(sum_of_evaluations)
                   if player == "X" { return  max_or_min_element(sum_of_evaluations,func(x int, y int)bool{return x>y})
                   } else { return max_or_min_element(sum_of_evaluations,func(x int, y int)bool{return x<y})} 
}

func print_el (arr[]int) {
  for i:=0; i<len(arr); i++ {
  if arr[i] == 1 { fmt.Print("+1","   ") 
     } else if arr[i] == -1 { fmt.Print("-1","   ") 
        } else if arr[i] == 0 { fmt.Print(" 0","   ")}
  }
  fmt.Println()
}
// 3+1
// 3+(2*4)
// 3+(x:=1)

//ищем доску для максиальной оценки 
func best_board (main_arr[3][3]string, player string) [3][3]string {
boards_tocheck:=all_moves(main_arr, player)
mboard:=boards_tocheck[0]
ev_boards:=evaluate_moves(mboard, another_player(player),0)
   for i:=0; i<len(boards_tocheck); i++ {
      var_board:=evaluate_moves(boards_tocheck[i], another_player(player),0)
      if ev_boards < var_board { ev_boards=var_board; mboard=boards_tocheck[i] }
      }
   return mboard
}

// print("xxx",end="yyy") ==  fmt.Print("xxx"); fmt.Print("yyy")
// print("xxx",end="") == fmt.Print("xxx")




//добавить доп переменную для bestboard
func max_or_min_element (arr []int, more_or_less func (int, int) bool ) int {
m:=arr[0]
   for i:=1; i<len(arr); i++ {
      if more_or_less(m, arr[i]) { continue 
      } else { m=arr[i] }
   }
return m
}

// создали массив досок , все возможные варианты одного шага
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
   fmt.Println()
    }
fmt.Println()
}

// печатает 1 доску 
func print_board (main_arr [3][3]string) {
  for x:=0; x<3; x++ {
   for y:=0; y<3; y++ {
   fmt.Print(main_arr[x][y])
   }
   fmt.Println()
   }
}


/*x_x
000
xxx

abc
def
ghi

x_x  abc
000  def
xxx  ghi
*/

//выводит первую строчку каждой доски

func str_print_board (boards [][3][3]string) {
for y:=0; y<3; y++ {
   for board:=0; board<len(boards); board++ {
      for x:=0; x<3; x++ {
         fmt.Print(boards[board][y][x])
      }
      fmt.Print("  ")
   }
   fmt.Println()
}
/*fmt.Println()
fmt.Println("без цикла")
fmt.Print(boards[0][0][0],boards[0][0][1],boards[0][0][2],"   ")
fmt.Print(boards[1][0][0],boards[1][0][1],boards[1][0][2],"   ")
fmt.Print(boards[2][0][0],boards[2][0][1],boards[2][0][2],"   ")
fmt.Print(boards[3][0][0],boards[3][0][1],boards[3][0][2],"   ")  
fmt.Print(boards[4][0][0],boards[4][0][1],boards[4][0][2],"   ")
//новая строка
fmt.Println()
fmt.Print(boards[0][1][0],boards[0][1][1],boards[0][1][2],"   ")
fmt.Print(boards[1][1][0],boards[1][1][1],boards[1][1][2],"   ")
fmt.Print(boards[2][1][0],boards[2][1][1],boards[2][1][2],"   ")
fmt.Print(boards[3][1][0],boards[3][1][1],boards[3][1][2],"   ")
fmt.Print(boards[4][1][0],boards[4][1][1],boards[4][1][2],"   ")
*/
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

