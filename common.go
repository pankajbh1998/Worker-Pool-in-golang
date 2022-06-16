package main

import "time"


// GetRandTask gives random tasks
func GetRandTask()interface{}{
	return time.Now().Unix()
}