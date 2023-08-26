package gohelp

import (
	"errors"
	"fmt"
	"os"
)

// checks if file exists given filepath
func FileExists(pathToFile string) error{
    if _,err := os.Stat(pathToFile); err == nil {
        return nil 
    } else if errors.Is(err, os.ErrNotExist){
        return fmt.Errorf("The path: %s\nDoes not exist", pathToFile);
    }else {
        return fmt.Errorf("The following error occurred: %s",err)
    }

} 
