# unique

Reads a file and outputs all unique lines.

The tool keeps all unique maps in memory (in a map), so the size of the file it
can process is limited by your RAM. If you give it a large file it might OOM.

The reason this exists is that I need it. If it's useful to you as well - awesome.

## Usage

When no arguments are given, `unique` reads from the standard in.  
Running ```$ echo " one\n one\n two" | unique``` will output

```
 one
 two
```

In order to read from a file, use the `-f` parameter:  
```$ unique -f file.dat```

