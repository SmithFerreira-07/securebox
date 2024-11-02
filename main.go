package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"securebox/i18n"
	"securebox/registry"

	"github.com/fatih/color"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Vault struct {
	Accounts []Account `json:"accounts"`
}

func saveVault(vault Vault, key [32]byte) error {
	data, err := json.Marshal(vault)
	if err != nil {
		return err
	}

	encryptedData, err := encrypt(data, key)
	if err != nil {
		return err
	}

	return os.WriteFile("vault.json", encryptedData, 0600)
}

func loadVault(key [32]byte) (Vault, error) {
	var vault Vault

	data, err := os.ReadFile("vault.json")
	if err != nil {
		if os.IsNotExist(err) {
			return vault, nil
		}
		return vault, err
	}

	decryptedData, err := decrypt(data, key)
	if err != nil {
		return vault, err
	}

	err = json.Unmarshal(decryptedData, &vault)
	return vault, err
}

func generateKey(masterKey string, salt []byte) ([32]byte, error) {
	derivedKey, err := scrypt.Key([]byte(masterKey), salt, 32768, 8, 1, 32)
	if err != nil {
		return [32]byte{}, err
	}
	var key32 [32]byte
	copy(key32[:], derivedKey)
	return key32, nil
}

func encrypt(data []byte, key [32]byte) ([]byte, error) {
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, err
	}

	encrypted := secretbox.Seal(nonce[:], data, &nonce, &key)
	return encrypted, nil
}

func decrypt(encrypted []byte, key [32]byte) ([]byte, error) {
	var nonce [24]byte
	copy(nonce[:], encrypted[:24])

	decrypted, ok := secretbox.Open(nil, encrypted[24:], &nonce, &key)
	if !ok {
		return nil, errors.New("decryption failed")
	}
	return decrypted, nil
}

func main() {

	red := color.New(color.FgRed).Add(color.Bold)
	asciiArt1 := `
  _________                                 __________              
 /   _____/ ____   ____  __ _________   ____\______   \ _______  ___
 \_____  \_/ __ \_/ ___\|  |  \_  __ \_/ __ \|    |  _//  _ \  \/  /																
 /        \  ___/\  \___|  |  /|  | \/\  ___/|    |   (  <_> >    < 																					
/_______  /\___  >\___  >____/ |__|    \___  >______  /\____/__/\_ \
        \/     \/     \/                   \/       \/            \/
`

	asciiArt2 := `
⠀⠀⠀⠀⣠⣾⣿⠿⢋⣡⣴⣶⣶⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⣾⣿⡟⢁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⠆⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀
⠀⢀⣾⣿⠏⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⢁⣴⣾⣿⠟⠋⣀⣤⣤⡀⠀⠀
⠀⣼⣿⡟⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⣴⣿⣿⠟⢁⣴⣿⣿⣿⣿⣿⡄⠀
⠀⣉⠛⠃⠸⢿⣿⣿⣿⣿⣿⣿⣿⣿⡟⢀⣾⣿⡿⠁⣰⣿⣿⣿⣿⣿⣿⣿⡧⠀
⠀⣿⣿⡆⢠⣤⠀⣉⠙⠛⠿⢿⣿⣿⠀⣾⣿⣿⠃⣼⣿⣿⣿⣿⠿⠛⢉⣡⣤⠀
⠀⣿⣿⡇⢸⣿⠀⣿⠿⠷⣶⠀⣈⡁⠀⠻⠿⡟⠀⠿⠟⠋⣁⣠⣴⣾⣿⣿⣿⠀
⠀⣿⣿⡇⢸⣿⠀⣿⡄⢠⣿⠀⣿⡇⠀⣶⣦⡄⠀⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⠀
⠀⣿⣿⡇⢸⣿⠀⠿⢧⣾⣿⠀⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
⠀⣿⣿⡇⢸⣿⣷⣶⣤⣄⣉⠀⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
⠀⠉⠛⠃⢸⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
⠀⠀⠀⠀⠀⠀⠉⠙⠛⠿⣿⣿⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⠿⠛⠉⠁⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠁⠀⠻⢿⡇⠀⣿⠿⠛⠉⠀⠀⠀⠀⠀⠀⠀
`

	var lang string
	red.Println(asciiArt1)
	red.Println(asciiArt2)
	fmt.Print("Select language (en/pt): ")
	fmt.Scanln(&lang)

	salt, err := registry.GetMachineSalt()
	if err != nil {
		log.Fatalf(i18n.GetTranslation(lang, "machine_salt_fail"), err)
	}

	storedSalt, err := registry.LoadSaltFromRegistry()
	if err != nil || !bytes.Equal(salt, storedSalt) {
		log.Println(i18n.GetTranslation(lang, "salt_not_found"))
		if err := registry.StoreSaltInRegistry(salt); err != nil {
			log.Fatalf(i18n.GetTranslation(lang, "failed_to_store_salt"), err)
		}
		if err := registry.BackupSaltToFile(salt); err != nil {
			log.Fatalf(i18n.GetTranslation(lang, "failed_to_backup_salt"), err)
		}
	} else {
		log.Println(i18n.GetTranslation(lang, "salt_loaded"))
	}

	var masterKey string
	fmt.Print(i18n.GetTranslation(lang, "enter_master_key"))
	fmt.Scanln(&masterKey)

	encryptionKey, err := generateKey(masterKey, salt)
	if err != nil {
		log.Fatal(i18n.GetTranslation(lang, "failed_to_generate_key"), err)
	}

	vault, err := loadVault(encryptionKey)
	if err != nil {
		log.Fatal(i18n.GetTranslation(lang, "failed_to_load_vault"), err)
	}

	for {
		fmt.Println("\n" + i18n.GetTranslation(lang, "options"))
		var choice int
		fmt.Print(i18n.GetTranslation(lang, "choose_option"))
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, password string
			fmt.Print(i18n.GetTranslation(lang, "account_name"))
			fmt.Scanln(&name)
			fmt.Print(i18n.GetTranslation(lang, "account_password"))
			fmt.Scanln(&password)

			encryptedPassword := base64.StdEncoding.EncodeToString([]byte(password))
			vault.Accounts = append(vault.Accounts, Account{Name: name, Password: encryptedPassword})

			if err := saveVault(vault, encryptionKey); err != nil {
				log.Fatal(i18n.GetTranslation(lang, "failed_to_save_vault"), err)
			}
			fmt.Println(i18n.GetTranslation(lang, "account_added"))
		case 2:
			for _, account := range vault.Accounts {
				decodedPassword, _ := base64.StdEncoding.DecodeString(account.Password)
				fmt.Printf(i18n.GetTranslation(lang, "show_accounts"), account.Name, string(decodedPassword))
			}
		case 3:
			var newMasterKey string
			fmt.Print(i18n.GetTranslation(lang, "change_master_key"))
			fmt.Scanln(&newMasterKey)

			newEncryptionKey, err := generateKey(newMasterKey, salt)
			if err != nil {
				log.Fatal(i18n.GetTranslation(lang, "failed_to_generate_key"), err)
			}

			if err := saveVault(vault, newEncryptionKey); err != nil {
				log.Fatal(i18n.GetTranslation(lang, "failed_to_save_vault"), err)
			}
			encryptionKey = newEncryptionKey
			fmt.Println(i18n.GetTranslation(lang, "master_key_changed"))
		case 4:
			fmt.Println(i18n.GetTranslation(lang, "exiting"))
			return
		default:
			fmt.Println(i18n.GetTranslation(lang, "invalid_option"))
		}
	}
}
