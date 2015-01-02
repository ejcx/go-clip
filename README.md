#NAAN and PBJ
Naan usage.
- Run it in the background. Personally, I just run it and forget about it in a tmux or screen window somewhere
````
        go build naan.go
        ./naan.go
````
- Naan then saves your last 10 clipboard entries into files /tmp/naanfiles/naan.[1-10]
- To get the data out of those files, you can be cool and do something like 
````
        cat /tmp/naanfiles/naan.[1-10] | pbcopy
````
- or you can be really really cool and use pbj.go like so
````
        go build pbj.go
        ./pbj -n 4
````

Bugs: Copying something re-adds it to the clipboard. This is likely to change which is why it is a bug.
In the future I will probably have a program to replace pbcopy which takes an integer as an argument. This
integer would then used as the -n in pbj to recall the value saved to the file

