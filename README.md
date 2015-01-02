#PBJ && PB && NAAN
I first wrote Naan but beleve it will probably be the least useful. Here's the best workflow for using PB and PBJ.

###PB is used for saving things to the clipboard & for assigning it to a naanfile.

Example: This will copy "ABC" to your clipbard and save it so it can be recalled by pbj
````
        echo "ABC" | pb -n 4
````

###PBJ is used for assigning past saved things to the clipboard

Example: This will copy ABC to your clipboard
```
        pbj -n 4
```

Naan usage.
- Run it in the background. Personally, I just run it and forget about it in a tmux or screen window somewhere. 
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

Notes: Naan is always watching your clipboard. Secrets added to your clipboard may be unprotected in your filesystem for a short amount of time. I think Naan is confusing, so personally I have begun using pbj and pb.
