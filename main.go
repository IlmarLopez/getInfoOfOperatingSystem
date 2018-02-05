package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// OperatingSystem información del sistema operativo
type OperatingSystem struct {
	SerialNoDiskdrive string
	SerialNoBaseboard string
}

func main() {
	serialNumberDiskdrive := GetInfoOperatingSystem()

	// Print in console
	fmt.Println("Disckdrive ", serialNumberDiskdrive.SerialNoDiskdrive)
	fmt.Println("Baseboard ", serialNumberDiskdrive.SerialNoBaseboard)
}

// GetInfoOperatingSystem obtiene del sistema el serial number del disco duro
func GetInfoOperatingSystem() OperatingSystem {
	var operatingSystem OperatingSystem

	/* === LINUX === */
	if runtime.GOOS == "linux" {
		operatingSystem.SerialNoDiskdrive = _diskdriveLinux()
		operatingSystem.SerialNoBaseboard = _baseboardLinux()
	}

	/* === WINDOWS === */
	if runtime.GOOS == "windows" {
		operatingSystem.SerialNoDiskdrive = _diskdriveWindows()
		operatingSystem.SerialNoBaseboard = _baseboardWindows()
	}

	return operatingSystem
}

/* Se hace la consulta para obtener el Serial number del disco duro en windows */
func _diskdriveWindows() (serialNoDisk string) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	serialNo := strings.Replace(string(stdoutStderr), "SerialNumber", "", -1)

	// print in console
	// fmt.Printf(serialNo)

	return serialNo
}

/* Se hace la consulta para obtener el Serial number de la placa base en windows */
func _baseboardWindows() (serialNobaseboard string) {
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	serialNo := strings.Replace(string(stdoutStderr), "SerialNumber", "", -1)

	return serialNo
}

/* Se hace la consulta para obtener el Serial number del disco duro en linux*/
func _diskdriveLinux() (serialNoDisk string) {
	cmd := exec.Command("bash", "-c", `hdparm -i /dev/sda | grep SerialNo`)
	// Se obtiene el resultado de la consulta en la consola
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	// Se genera el formato de regexp que se va utiliza para detectar el tipo del valor de la variable.
	reg := regexp.MustCompile(`SerialNo=(\w+)`)

	// Se hace la coverción de stdoutStderr []byte a string.
	// Se obtiene el valor de `SerialNo=(\w+)`
	serialNoString := reg.FindString(string(stdoutStderr))

	// Se obtiene el valor que representa la variable
	serialNo := strings.Replace(serialNoString, "SerialNo=", "", -1)

	// print in console
	// fmt.Println(serialNo)

	return serialNo
}

func _baseboardLinux() (serialNobaseboard string) {
	cmd := exec.Command("bash", "-c", `dmidecode -t baseboard | grep "Serial Number"`)
	// Se obtiene el resultado de la consulta en la consola
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	serialNo := strings.Replace(string(stdoutStderr), "Serial Number:", "", -1)

	return serialNo
}
