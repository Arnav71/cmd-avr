# cmd-avr
The Plugin for AVR on CMD!!
# Introduction
The cmd-avr is just what it sounds like The Plugin which edits the Makefile on cmd.
# Works With Any Text Editor That Doesn't Have An AVR Plugin(Sorta)
Text Editor Requirements

1. Must Have A Console

Thats All

User Requirements

1. Patience

That's All Again

System Requirements

1. You need to install Golang

Thats All Again(2)
# Installation
Installation is a little bit tricky but that's where the user requirements come.
Steps
First Of All ,you need a common Makefile, For Windows users Who are using WinAVR or has a Makefile for avr-gcc can do this := 
1. Open up your AVR Workspace,(this can be "C:\AVRSketchbook for mine)
2. Save the Makefile into the Workspace
3. If Your Workspace isn't C:\AVRSketchbook,you need to change it,to do that read the custom workspace section
4. Now create a folder called bin.
5. cd into that bin folder by typing cd C:\bin (if you created it in the C Drive)
6. Now Build the go file using go build the\path\where\you\installed\the\go\file.go
7. Setup the %PATH% Environment Variable by typing 
   set PATH=%PATH%;C:\bin
And that's it 

# Additional Notes For Notepad++ Users
If you are a Notepad++ User then you can execute the command in Notepad++ itself(Sorry,i didn't try other text editors)!
1. Follow this guide http://vpapanik.blogspot.in/2012/08/using-winavr-with-notepad.html
2. Now create another Macro called AVR New Project the same way.
3. Inside it, type

cls

avr

4. That's it save it and enjoy

# Custom Workspace
1. Open the gofile before building it and add these changes

See these lines

op := "C:\\AVRSketchbook\\Makefile"

oe := "C:\\AVRSketchbook\\"+ ar

oa := "C:\\AVRSketchbook"

Change it into the Avr Workspace you use like

op := "C:\\AVRWork\\Makefile"

oe := "C:\\AVRWork\\" + ar 

oa := "C:\\AVRWork"

Here AVRWork is just an example
# Questions or Bugs
Email me at planetoid128@gmail.com
# Attributions
Thanks to vpapanik on creating the blog
# License
This is licensed under MIT-License
