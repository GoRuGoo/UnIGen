package main

import "time"

func (o Options) generateSingleTestData()error {
	currentTime := time.Now()

	parsedPlaceholderTime,err:= time.Parse(o.format, currentTime.Format("2006-01-02-15-04-05"))
	if err!=nil{
		return err
	}

	
}
