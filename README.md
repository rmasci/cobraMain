# Cobra Main.go

This is to be used in a 'Cobra' application. This allows you to set a default cmd.

To use:

1. Set the defaultCmd("<somthing>") to what you want to be your default command
2. Change the commands slice to contain all the possible commands. The ones you added with cobra add.

***Note*** Every time you 'cobra add' a new command it needs to go here. TODO: make this something we can configure on compile time in the makefile.

Use something like grep "Use:" *.go | grep -v root.go | awk '{print $3}' | sed 's/^.//;s/..$//' | tr '\n' ' '

Then set a build variable using -X to go build
