package registry

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

// Salt base hardcoded
var baseSalt = []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}

// Generates a unique salt using baseSalt and machine-specific data
func GetMachineSalt() ([]byte, error) {
	machineID, err := GetMachineID()
	if err != nil {
		return nil, err
	}
	combinedSalt := append(baseSalt, machineID...)
	return combinedSalt, nil
}

// Retrieves the machine ID based on the operating system
func GetMachineID() ([]byte, error) {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("wmic", "csproduct", "get", "UUID").Output()
		if err != nil {
			return nil, err
		}
		return bytes.TrimSpace(out), nil
	}
	out, err := exec.Command("cat", "/etc/machine-id").Output()
	if err != nil {
		return nil, err
	}
	return bytes.TrimSpace(out), nil
}

// Stores the salt in the Windows Registry
func StoreSaltInRegistry(salt []byte) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\MySecureApp`, registry.WRITE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.SetBinaryValue("Salt", salt)
}

// Loads the salt from the Windows Registry
func LoadSaltFromRegistry() ([]byte, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\MySecureApp`, registry.READ)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	salt, _, err := key.GetBinaryValue("Salt")
	return salt, err
}

// Backups the salt to a hidden file
func BackupSaltToFile(salt []byte) error {
	return os.WriteFile(".salt_backup", salt, 0600)
}

// Loads the salt from the backup file
func LoadSaltFromFile() ([]byte, error) {
	return os.ReadFile(".salt_backup")
}
