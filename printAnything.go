package main

import "C"
import "log"

//export Add
func Add() int {
	// use print will fine.
	print("Add called\n")
	// use log will panic .  If comment this line,  there is a probability of multiple runs.
	log.Println("Add called")
	return 1
}
