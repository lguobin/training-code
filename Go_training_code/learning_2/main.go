package main

import (
	File "learning_2/Package_T"
)

func Bsisc() {
	Bsisc_Slice()

	Bsisc_func()
	Bsisc_Closure()
	Bsisc_Struct()
	Bsisc_OOP()

	Bsisc_Interface()
}

func Package_T_learning() {
	File.OS_CreateFile()
	File.OS_FileOpen()
	File.OS_Write()
	File.OS_Write_String()
	File.OS_Seek()
	File.OS_Fileinfo()
	File.OS_filepath()

	File.Login_check()
	File.Count_Num_Test()

}

func main() {
	// Menu()
	// Bsisc()

	// Package_T_learning()

	File.Login_check()
	// File.Count_Num_Test()
	File.Checker_Sum()
	File.Calc()
	File.Run()

	File.Run_Error()
	File.Ptr_Swap()

	File.Time_Package()
	File.String_Demo()
}
