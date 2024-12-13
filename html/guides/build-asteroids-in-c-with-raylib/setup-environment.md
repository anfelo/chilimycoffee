Before we start coding the actual game we need to make sure that we have all
the tools in place, installed in our computer.

We will be making a game in C and for that to happen, our computer must be
able to compile the code that we write. Luckily, almost every operating system
comes with a C compiler installed that we can use. Depending on which
operating system (OS) you are running on, you will use one or the other. But
don't worry I will guide you through the entire process no matter what
computer you are using.

Let us first check that you are all set.

### macOS

On macOS check that you have installed either clang or gcc:

```bash
clang --version
```

Or:

```bash
gcc --version
```

If none of those are available in your machine, check the [installation guide](/guides/my-c-notes/installation#macos).

### Windows

Check the [installation guide](/guides/my-c-notes/installation#windows) for more info

### Linux

Check if clang or gcc are already installed in your system:

```bash
clang --version
```

Or:

```bash
gcc --version
```

If none of those are installed in your Linux machine, check the [installation guide](/guides/my-c-notes/installation#linux).

This is all that you will need for now, you are almost ready to start hacking.
Now we will setup Raylib, a game library that will make our lives easier when
trying to paint some pixels into the screen, and much more.
