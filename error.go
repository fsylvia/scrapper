package util

import "log"

func HandleError(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}
