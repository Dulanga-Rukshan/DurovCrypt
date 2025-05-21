package DurovCrypt

func WelcomeMsg() string {
	welcome := `
              
            ╔═══╗               ╔═══╗             ╔╗ 
            ╚╗╔╗║               ║╔═╗║            ╔╝╚╗
            ║║║║╔╗╔╗╔══╗╔═╗╔╗╔╗║║ ╚╝╔═╗╔╗ ╔╗╔══╗╚╗╔╝
            ║║║║║║║║║╔╗║║╔╝║╚╝║║║ ╔╗║╔╝║║ ║║║╔╗║ ║║ 
            ╔╝╚╝║║╚╝║║╚╝║║║ ╚╗╔╝║╚═╝║║║ ║╚═╝║║╚╝║ ║╚╗
            ╚═══╝╚══╝╚══╝╚╝  ╚╝ ╚═══╝╚╝ ╚═╗╔╝║╔═╝ ╚═╝
                                        ╔═╝║ ║║      
                                        ╚══╝ ╚╝      
                 DUROVCRYPT VERSION: 1.0.0
  DUROVCRYPT is a go language based file Encryptor and Decryptor. DurovCrypt is Developed by Dulanga Rukshan, 
  A security enthusiast and a cryptographer who is Interested in Linux and cryptography algorithms. 

  REPO --->  https://github.com/Dulanga-Rukshan/DurovCrypt submit the bug fixes and additional feature ideas, 

  CONTACT DEVOLOPER --> Dulangarukshan@proton.me  

    USAGE:
      DurovCrypt

    COMMANDS: use arrow keys to select through the options
      Encrypt-----> Encrypt a file
      Decrypt-----> Decrypt a file
      Help   -----> Show this help message              
                    

  `
	return welcome
}

func ShowHelp() string {
	help := `
                        DUROVCRYPT VERSION: 1.0.0

DUROVCRYPT is a go language based file Encryptor and Decryptor 

USAGE:
  DurovCrypt

COMMANDS: use arrow keys to select through the options
  Encrypt-----> Encrypt a file
  Decrypt-----> Decrypt a file
  Help   -----> Show this help message

EXAMPLES:
(01). Encrypt a file:
  >>> What do you wanna perform in DurovCrypt? 
  ****You'll be prompted for options and Select Encrypt using arrow keys in your keyboard.


  Enter a FilePath/FileName to Encrypt: 
  ****Here you can enter the file name you wanna Encrypt. You can Enter file name 
      like (/home/username/Desktop/file.txt) or if file is in current directory 
      just enter filename like file.txt. 


  >> Enter password for Encrypt: 
  ****Here you have to choose a password for encrypting the file you entered. You have to consider 
      few things when choosing password, 

                ---> Password must be longer than 6 characters
                ---> Password must contain one or more Capital Letters
                ---> Password must contain one or more Simple Letters
                ---> Password must contain one or more Numbers
                ---> Password must contain one or more Symbols
                ---> And finally password cannot contain any sequential like 
                              • "1234"
                              • "abcd" 

                  Example of good password := Ehdj$@0#dn52@!.vqa7+


>> Retype password for Encrypt:
****Here you have to retype the password you entered above, if passwords are mismatched 
    operation will be fail, so type the right password, after that Encrypted file will be
    saved in unencrypted file directory as {fileOriginalname + ".drv" }. 

    **success message look like this --> "File saved as /home/username/Desktop/file.txt.drv!"

===========WARNING REMEMBER THE PASSWORD YOU ARE USING FOR ENCRYPTION, IF YOU LOST OR FORGET IT DECRYPTION IS IMPOSSIBLE!=====
                                 


(02). Decrypt a Durovcrypt encrypted file:
  >>> What do you wanna perform in DurovCrypt? Decrypt
  ****You'll be prompted for options and Select Decrypt using arrow keys in your keyboard.

  Enter a FilePath/FileName to Decrypt: 
  ****Here you have to enter the file name you wanna Decrypt. You can Enter file name 
      like (/home/username/Desktop/file.txt.drv) or if file is in current directory 
      just enter filename like file.txt.drv. Durovcrypt only decrypt files that been encrypted 
      by Durovcrypt.  


  >> Enter password for Decrypt:
  ****Here you have to enter password that used for encrypting the file. if password is wrong
      decryption will fail. after successful Decryption file will be save as {"Decrypt"+ fileOriginalname}

      **success message look like this --> File saved as /home/username/Desktop/Decryptfile.txt!


(03). Quick help:
  DurovCrypt help

SECURITY TIPS:
• REMEBER THE ENCRYPTION PASSWORD
• Use strong passwords (mix letters, numbers, symbols)
• Never share your password
• Keep backups of original files
• Delete unencrypted files after encryption

    DurovCrypt is Developed by Dulanga Rukshan, A security enthusiast and a cryptographer who is Interested in Linux and 
    cryptography algorithms. You can clone this repo --->  https://github.com/Dulanga-Rukshan/DurovCrypt submit the bug fixes and
    additional feature ideas, Also you can contact me by --> Dulangarukshan@proton.me                             
`
	return help
}

func HelpMsg() string {
	help := `COMMANDS:
    encrypt (e)    - Encrypt a file
    decrypt (d)    - Decrypt a file
    help (h)       - Show this help message`

	return help
}
