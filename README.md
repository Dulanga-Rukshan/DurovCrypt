# DurovCrypt

<p align="center">
  <img src="assests/banner.png" alt="MyTool Banner" width="500" style="border-radius:10px">
</p>

DurovCrypt is GO language based file Encryptor and Decryptor. DurovCrypt uses AES 256+GCM encryption algorithm to encrypt and decrypt files.
Why use DurovCrypt? you ask, Imagine you store passwords, financial details, or private messages in a file on your computer. Now picture that file falling 
into the wrong hands a hacker, a nosy colleague, or even malware. scary, right? 

DurovCrypt protects your sensitive files like a digital vault by Encrypting it using a password. So even someone get the file 
they cant read it. 
 
            💳 Bank details in a File? Lock them with military grade encryption.
          
            🔑 Passwords stored Locally? Make them unreadable to anyone but you.

Unlike simple password protected ZIP files (which can be cracked), DurovCrypt uses AES-256 encryption the same standard governments use combined with
Argon2, a brute force resistant algorithm that makes hacking attempts nearly impossible.

🛡️ Real-World Protection
> 📌 "My laptop got stolen with all my tax documents!" → If they were encrypted, the thief gets useless gibberish.

>  
> 📌 "I save work API keys in a text file..." → A disgruntled coworker can’t leak what they can’t read.

No more accidental leaks, encrypted files hide secrets in plain sight.

With DurovCrypt, your files stay safe even if they’re stolen. Just remember your password because without it, not even supercomputers can break in.

Okay now move on with how to install DurovCrypt. 

# 🚀 Installation Guide
 DurovCrypt is designed to be easy to install on Linux, macOS. 

## Step 01 
You have to clone this repo--> https://github.com/Dulanga-Rukshan/DurovCrypt.git using git command, Here is the example 👇
<p align="start">
  <img src="assests/Gitclone.png" alt="Gitclone.png Banner" width="500" style="border-radius:10px">
</p>

## Step 02 
after that cd in to folder DurovCrypt 👇
<p align="start">
  <img src="assests/cd.png" alt="cd.png Banner" width="500" style="border-radius:10px">
</p>

## Step 03 
After that you have to run ./install.sh command to install all dependencies for DurovCrypt, DurovCrypt is go based language. So it need go compiler 
to run in a system ./install.sh will check the system if go installed in system it will build the DurovCrypt, unless go isn't installed in a system it will 
prompt you to install. 👇
<p align="start">
  <img src="assests/DurovCryptGoask.png" alt="DurovCryptGoask.png Banner" width="500" style="border-radius:10px">
</p>
<p align="start">
  <img src="assests/DurovCryptGoask1.png" alt="DurovCryptGoask1.png Banner" width="700" style="border-radius:10px">
</p>

## Step 04 
Now DurovCrypt is Ready to use you will get a prompt like this 👇
<p align="start">
  <img src="assests/installSuccess.png" alt="installSuccess.png Banner" width="500" style="border-radius:10px">
</p>

<br>
<br>

# 🔐 Usage Guide

## Start DurovCrypt

To start DurovCrypt just type DurovCrypt in a terminal. It will prompt you to DurovCrypt main interface, Like this 👇
<p align="start">
  <img src="assests/MainInterface.png" alt="MainInterface.png Banner" width="800" style="border-radius:10px">
</p>

<br>
<br>

## Let's Start Encrypting a file

I have a file name called file.txt in my home folder it contain simple plaintext msg. I cat in to it see what's in it. 👇
<p align="start">
  <img src="assests/unencryptedFile.png" alt="unencryptedFile.png Banner" width="500" style="border-radius:10px">
</p>

To encrypt this text file i type "DurovCrypt" in a terminal. Then i have to enter the file name for encryption
And a password for it. Here DurovCrypt expect a password. I have to type and retype the password, And most IMPORTANTLY i have to
remember that password, Because without it i can't decrypt the file that i encrypted. DurovCrypt expect in a good password.
what i mean by good password is password that has 👇
<p align="start">
  <img src="assests/password.png" alt="password.png Banner" width="500" style="border-radius:10px">
</p>

Without these characteristics DurovCrypt don't accept the password for encryption, And here is how you should enter a file name  👇
<p align="start">
  <img src="assests/encryptFile.png" alt="encryptFile.png Banner" width="500" style="border-radius:10px">
</p>

Now encrypted file is save in a my unencrypted file folder as file.txt.drv 👇
<p align="start">
  <img src="assests/encryptFolder.png" alt="encryptFolder.png Banner" width="500" style="border-radius:10px">
</p>

lets look at how encrypted file look like  👇
<p align="start">
  <img src="assests/encrypted.png" alt="encrypted.png Banner" width="500" style="border-radius:10px">
</p>
Total gibberish
<br>
<br>

## Let's  Decrypt that file

To Decrypt the file that i encrypted, all i have to do is enter "DurovCrypt" in the terminal and select decrypt enter the file name 
i want to decrypt, and the password that used for encryption. It's look like this  👇
<p align="start">
  <img src="assests/decrypt.png" alt="decrypt.png Banner" width="500" style="border-radius:10px">
</p>

Now Decrypted file is saved in a encrpted file folder lets have a look in to it.  👇
<p align="start">
  <img src="assests/DecryptFolder.png" alt="DecryptFolder.png Banner" width="500" style="border-radius:10px">
</p>
<p align="start">
  <img src="assests/decryptFile.png" alt="decryptFile.png Banner" width="500" style="border-radius:10px">
</p>









