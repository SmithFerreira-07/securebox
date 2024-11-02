# Secure App / Aplicativo Seguro

A secure password management application that allows users to store and encrypt their account credentials. The application leverages machine-specific salts, encryption techniques, and the Windows Registry for secure data handling.

Um aplicativo seguro de gerenciamento de senhas que permite aos usuários armazenar e criptografar suas credenciais de conta. O aplicativo utiliza sais específicos da máquina, técnicas de criptografia e o Registro do Windows para o manuseio seguro de dados.

## Features / Funcionalidades

- **Account Management**: Add, view, and manage your accounts and passwords securely.  
  **Gerenciamento de Contas**: Adicione, visualize e gerencie suas contas e senhas de forma segura.
  
- **Encryption**: Utilizes NaCl secret box for data encryption.  
  **Criptografia**: Utiliza a caixa secreta NaCl para criptografia de dados.

- **Machine-Specific Salts**: Generates unique salts based on the machine's ID for enhanced security.  
  **Sais Específicos da Máquina**: Gera sais exclusivos com base no ID da máquina para maior segurança.

- **Persistent Storage**: Saves vault data to a JSON file and supports backup in the Windows Registry.  
  **Armazenamento Persistente**: Salva os dados do cofre em um arquivo JSON e suporta backup no Registro do Windows.

## Requirements / Requisitos

- Go 1.16 ou posterior
- `golang.org/x/crypto` para funções de criptografia
- Windows (para operações de Registro)

- Go 1.16 or later
- `golang.org/x/crypto` for encryption functions
- Windows (for Registry operations)
