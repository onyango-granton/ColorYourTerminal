package main

import "fmt"

func main() {
	var Reset = "\033[0m" 
	var Red = "\033[31m" 
	var Green = "\033[32m" 
	var Yellow = "\033[33m" 
	var Blue = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan = "\033[36m" 
	var Gray = "\033[37m" 
	var White = "\033[97m"

	colorName := []string{"Reset","Red","Green","Yellow","Blue","Magenta","Cyan","Gray","White"}
	colors := []string{Reset,Red,Green, Yellow,Blue,Magenta, Cyan,Gray,White}
	colorMap := make(map[string]string)

	for i:=0; i < len(colorName);{
		for j:=i;j<len(colors);{
			colorMap[colorName[i]] = colors[j]
			fmt.Println(i,j)
			i++
			if i == len(colorMap){
				break
			}
		}
	}

	fmt.Println(colorMap)


	for k, v := range colorMap{
		fmt.Println(v, "This is",k,colorMap[Reset])
	}
}