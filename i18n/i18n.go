package i18n

var translations = map[string]map[string]string{
	"en": {
		"options":                "Options: (1) Add Account (2) Show Accounts (3) Change Master Key (4) Exit",
		"choose_option":          "Choose your Option: ",
		"enter_master_key":       "Enter your master key: ",
		"account_added":          "Account added successfully.",
		"invalid_option":         "Invalid option.",
		"account_name":           "Enter account name: ",
		"account_password":       "Enter password: ",
		"show_accounts":          "Account: %s, Password: %s\n",
		"change_master_key":      "Enter new master key: ",
		"master_key_changed":     "Master key changed successfully.",
		"exiting":                "Exiting.",
		"salt_not_found":         "Stored salt not found or does not match. Storing in registry.",
		"salt_loaded":            "Salt loaded successfully from registry.",
		"failed_to_store_salt":   "Failed to store salt: %v",
		"failed_to_backup_salt":  "Failed to backup salt: %v",
		"failed_to_load_vault":   "Failed to load vault: %v",
		"failed_to_save_vault":   "Failed to save vault: %v",
		"machine_salt_fail":      "Failed to generate machine-specific salt: %v",
		"failed_to_generate_key": "Failed to generate encryption key: %v",
	},
	"pt": {
		"options":                "Opções: (1) Adicionar Conta (2) Mostrar Contas (3) Alterar Chave Mestre (4) Sair",
		"choose_option":          "Escolha sua Opção: ",
		"enter_master_key":       "Digite sua chave mestre: ",
		"account_added":          "Conta adicionada com sucesso.",
		"invalid_option":         "Opção inválida.",
		"account_name":           "Digite o nome da conta: ",
		"account_password":       "Digite a senha: ",
		"show_accounts":          "Conta: %s, Senha: %s\n",
		"change_master_key":      "Digite a nova chave mestre: ",
		"master_key_changed":     "Chave mestre alterada com sucesso.",
		"exiting":                "Saindo.",
		"salt_not_found":         "Salt armazenado não encontrado ou não corresponde. Armazenando no registro.",
		"salt_loaded":            "Salt carregado com sucesso do registro.",
		"failed_to_store_salt":   "Falha ao armazenar salt: %v",
		"failed_to_backup_salt":  "Falha ao fazer backup do salt: %v",
		"failed_to_load_vault":   "Falha ao carregar o cofre: %v",
		"failed_to_save_vault":   "Falha ao salvar o cofre: %v",
		"machine_salt_fail":      "Falha ao gerar salt específico da máquina: %v",
		"failed_to_generate_key": "Falha ao gerar a chave de criptografia: %v",
	},
}

func GetTranslation(lang, key string) string {
	if val, ok := translations[lang][key]; ok {
		return val
	}
	return key
}
