# wavSteg

A command line interface based program for wav audio Steganography.<br>
it uses an algorithm that I've developed in which there are 2 steps to encoding an wav file.

:mag_right:  find closest ascii character in PCM data, and insert value to the corresponding index.<br>
:pushpin: Mark the previous index <br>


### Limitations And Encoding / Decoding times
Since the program stores the data linearly and depends on the previous encoded index, doing so makes it a whole lot more diffcult to 
make the program multi-threaded which increases the compute time.
however, since the index is cached the encoding computation isn't relooping the entire PCM data, which makes it easier to compute.



### Installation
1. <b>Install go-lang from offical website</b><br>
    ```https://go.dev/doc/install```

2. <b>build program</b>
```go build wavSteg.go```


### Running the Program
1. Place .wav file to encode inside ```input``` folder.
2. Place text file(or any other file that you can read from) inside ```input``` folder.

```
   Usage:
  -audio string
    	Audio file name
  -decode
    	Decode flag (default true)
  -encode
    	Encode flag
  -message string
    	Encode message with Command Line Interface message
  -tfile string
    	Encode message with provided text file name

```

encoding example:
```
./wavSteg -audio=input/sample-file-2.wav -tfile=input/toEncode.txt -encode=true
```
```
./wavSteg -audio=input/sample-file-2.wav -message "hello world!" -encode=true
```
decoding example:
```
./wavSteg -audio=enc_file.wav -decode=true
```
Encoded audio and Decoded file will be saved to results directory.

## Disclaimer

This program is a proof-of-concept, and should not be intended for transfring important information, and I won't be liable for any damagaes caused by this program.

## Copyright

This software is licensed under MIT. Copyright © 2022 Rani Giro
