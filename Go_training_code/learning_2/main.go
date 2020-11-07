package main

import (
	Json "learning_2/Json"
	File "learning_2/Package_T"
	_Struct "learning_2/Struct"
	_Test "learning_2/Tester"
)

func Bsisc() {
	Bsisc_Slice()

	Bsisc_func()
	Bsisc_Closure()
	Bsisc_OOP()

	Bsisc_Map()
	Bsisc_Interface()
	Bsisc_json()
}

func Bsisc_Struct() {
	_Struct.Bsisc_Struct()
	_Struct.Struct_Demo_1()
}

func Package_T_learning() {
	File.OS_CreateFile()
	File.OS_FileOpen()
	File.OS_Write()
	File.OS_Write_String()
	File.OS_Seek()
	File.OS_Fileinfo()
	File.OS_filepath()
	File.Cmd_args()

	File.Login_check()
	File.Count_Num_Test()

	File.Checker_Sum()
	File.Calc()
	File.Run_Digui()
	File.Run_Error()
	File.Ptr_Swap()

	// 使用自带包
	File.String_Demo()
	File.Time_Package()
	File.Array_Demo()
	File.Flag_Demo()
}

func Json_learning() {
	Json.Json_Demo()
	Json.Json_to_struct()
}

func Tester_learning() {
	_Test.Test_Demo()
}

func main() {
	// Menu()
	// Bsisc()
	// Package_T_learning()
	// Json_learning()

	// string_to_slice()

	// Tester_learning()

	Bsisc_Struct()
}
