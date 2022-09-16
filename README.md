# gotemplater

A templating utility that takes in a data in the form of yaml file and a go template file and render the output.

## How to use?

### docker/nerdctl
Replace the command with `docker` or `nerdctl` whichever applicable
```
<command>  run --rm -it -v ${PWD}/example:/ws -w /ws maclarensg/gotemplater gotemplater -d data.yaml -t custom.txt.gotmpl
```

### compile
```
make build 
./gotemplater -d example/data.yaml -t example/custom.txt.gotmpl
```

### Output to file path
```
<command>  run --rm -it -v ${PWD}/example:/ws -v out:/out -w /ws maclarensg/gotemplater gotemplater -d data.yaml -t custom.txt.gotmpl -o /out/output.file
```


