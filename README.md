dart2exe
--

Convert console Dart applications into single executables containing the VM and
all dependencies.


```
$ dart2exe ~/wrk/dart_test_app_2exe
2013/09/20 20:45:47 Generated: dart_test_app_2exe

$ ./dart_test_app_2exe

┏━ ┏━┃┏━┃━┏┛━━┃┏━┛┃ ┃┏━┛
┃ ┃┏━┃┏┏┛ ┃ ┏━┛┏━┛ ┛ ┏━┛
━━ ┛ ┛┛ ┛ ┛ ━━┛━━┛┛ ┛━━┛
      
I am Testing standard lib (math) and external (crypto) dependencies.
--
b92ada6b7ed827a813f6f4d45e8f577c19d9a0d10dc9e76c6e665a7c8c83368d
ed4f6d26173f8ad4d8d248368db091eeca5134148a55f49c7029d42bf009b09f
4bd39d33477621026bdf138823d2897619912b0f3b3030e98f8c34ca4e6c8ea1
13a92a7e5c72399e42829ac1f16bedc2a74097fb9f956669e100526984b54fc0
300c3ec4acc27875bb1754009f9c6b71e3ef4e076fcc1d5500d65a46ec731777
de5459c64660f26149bc599e3e66295245c3de78aaf19180fdfbd6924ac39f2a
bc155ab1be34f9c492d896494b550d6ec38a427e2e76cb5dd86b3f6238454d85
29a3c3a465b4ce4c9e489e73820337028153d5b0209d50eb57e51f64e9e054d3
--
I tested them. Pub works. Goodbye, world!
```

How does it work?
--

Yeah, it's mostly a hack.

We create a binary (using Go) that contains everything you need tar'd and
chopped into chunks. Each chunk is embedded as a literal byte array in Go
source. These arrays are rejoined and unpacked into /tmp at runtime by
a bootstrapper which just executes `dart bin/main.dart`.

Performance
--
Not entirely terrible. The whole build takes around 5 seconds on this crappy
laptop, and the binaries seem to start pretty fast.
