package main

import (
	"os"
	"io"
	"fmt"
)



// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error){
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error){
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func(){
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil{
		return
	}

	err = out.Sync()
	return
}



func main() {
	//Copy Test APISync
	fmt.Println("Start Copy Test APISync....")
	src, dist := "F:\\GoWork\\src\\go-tutorial\\src\\main\\apisync\\test\\mvngt-Test.bat", "R:\\Projects\\MedSuiteGit\\MedSuiteApiSync\\src\\MedSuiteApiSync\\medhistory\\mvngt.bat"
	fmt.Printf("Copying %s to %s\n", src, dist)
	//err := CopyFile(os.Args[1], os.Args[2])
	err := CopyFile(src, dist)
	if err != nil {
		fmt.Printf("Copy Test APISync failed %q\n", err)
	} else {
		fmt.Printf("Copy Test APISync succeeded\n")
	}

	//Copy Test Database
	fmt.Println("Start Copy Test Database ....")
	srcdb, distdb := "F:\\GoWork\\src\\go-tutorial\\src\\main\\apisync\\test\\MedExpress.sdf", "C:\\Users\\Antonis\\AppData\\Roaming\\MedExpress\\db\\MedExpress.sdf"
	fmt.Printf("Copying %s to %s\n", srcdb, distdb)
	//err := CopyFile(os.Args[1], os.Args[2])
	errdb := CopyFile(srcdb, distdb)
	if errdb != nil {
		fmt.Printf("Copy Test Database failed %q\n", errdb)
	} else {
		fmt.Printf("Copy Test Database succeeded\n")
	}
}