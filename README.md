# fwatch

fwatch is a simple tool to watch a file and execute a command when it changes.

## Build

```bash
$ go build
```

## Usage

```bash
$ ./fwatch file.txt echo 'file changed'
```

## Install

Depending on your system the installation is different.

### Linux
On Linux you can just move the compiled binary to `/usr/local/bin`:

```bash
$ sudo mv fwatch /usr/local/bin
```

### Windows
On Windows you first run the build process. After that simply move the entire project folder to `C:\Program Files` on x64 systems or `C:\Program Files (x86)` on x32 systems. 
After that is done you should see the `fwatch` folder inside the corresponding Program Files folder.

Now you need to add the fwatch folder to the path.
1. Search for "path" with the Windows search and press enter.
2. Click the button which says Environent Variables.
3. Search for the `PATH` variable in the user variables and double click on it.
4. Click on `New`
5. Enter the path to the fwatch folder. It should look like this: `C:\Program Files\fwatch` or `C:\Program Files\fwatch`

Done!
