package main

import "fmt" 

func main () {
main_arr:= [3][3]string {
{"_","O","O"},
{"X","O","_"},
{"X","_","X"}}
tick_print(main_arr)
//fmt.Println(win_or_loose_x(main_arr), win_or_loose_y(main_arr))
fmt.Println()
coordinate(main_arr)
}

// тик принт печатает поле игры
func tick_print (main_arr [3][3]string) {
	for y:=0; y<3; y++ {
		for x:=0; x<3; x++   {
			fmt.Print(main_arr[y][x])
		}
     fmt.Println()
   }
}

// печатает кто выйграл по горизонтали
func win_or_loose_x (main_arr [3][3]string) string {
	for y:=0; y<len(main_arr); y++ {
		check_x:=same_el_x(main_arr, y)
		if check_x && main_arr[y][0] == "X" { return "X"
		} else if check_x  && main_arr[y][0] == "O" { return "O" 
		} else { continue }
}
return "_" 
}

// часть win_or_loose - ищет одинаковый элемент по горизонтали
func same_el_x (main_arr [3][3]string, y int) bool {
	for x:=0; x<len(main_arr)-1; x++ {
		if main_arr[y][x] == main_arr[y][x+1] { return true 
			} else { continue }
		}
return false
}

// печатает кто выйграл по вертикали
func win_or_loose_y (main_arr [3][3]string) string {
	for y:=0; y<len(main_arr); y++ {
		check_y:=same_el_y(main_arr, y)
		if check_y && main_arr[y][0] == "X" { return "X"
		} else if check_y  && main_arr[y][0] == "O" { return "O" 
		} else { continue }
}
return "_" 
}

// часть win_or_loose - ищет одинаковый элемент по вертикали
func same_el_y (main_arr [3][3]string, y int) bool {
	for y:=0; y<len(main_arr)-1; y++ {
		if main_arr[y][0] == main_arr[y+1][0] { return true 
			} else { continue }
		}
return false
	}

	func coordinate (main_arr [3][3]string) {
		//coor:= [9]int{}
		for y:=0; y<3; y++ {
			for x:=0; x<3; x++ {
					switch main_arr[y][x] {
						case "_":fmt.Print(x)
						case "O":fmt.Print("O")
						case "X":fmt.Print("X")
					}
				}
				fmt.Println()
			}
	   }