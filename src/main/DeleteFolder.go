package main

import (
	"fmt"
	"os/exec"
)


func main() {
	fmt.Println("Start Deleting App Data Files")
	//FOR /D %%p IN ("R:\Projects\MedSuiteGit\MedExpressApplication\src\MedExpress\bin\*.*") DO rmdir "%%p" /s /q
	//DO rmdir "C:\\Users\\Antonis\\AppData\\Local\\Sigma_Solutions\\*.*" /s /q
	//filepath  := "C:\\Users\\Antonis\\AppData\\Local\\Sigma_Solutions"
	if err := exec.Command("rmdir \"c:\\Users\\Antonis\\AppData\\Local\\Sigma_Solutions\\*.*\" /s /q").Run(); err != nil {
		fmt.Println("Error:",err)
	}else {
		fmt.Println("Folder deleted successfully!")
	}
}