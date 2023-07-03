package main

import (
	"fmt"
	"os"

	"rochatronic.net/text-venture-service/cmd/api"
	v1 "rochatronic.net/text-venture-service/cmd/api/v1"
	"rochatronic.net/text-venture-service/cmd/model"
)

func main() {
   
welcome := `                                                          
   ______          __     _    __           __                
  /_  __/__  _  __/ /_   | |  / /__  ____  / /___  __________ 
   / / / _ \| |/_/ __/   | | / / _ \/ __ \/ __/ / / / ___/ _ \
  / / /  __/>  </ /_     | |/ /  __/ / / / /_/ /_/ / /  /  __/
 /_/  \___/_/|_|\__/     |___/\___/_/ /_/\__/\__,_/_/   \___/ 
                                                              
                                                     v%s
`
  fmt.Printf(welcome,v1.VERSION)                                 

   configFileName :=  "./config/application.yaml"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	} 
   
   txt := model.LoadTextVenture(configFileName)
   api.Start(txt)
}