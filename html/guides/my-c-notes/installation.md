### macOS

On macOS the C [compilers](/guides/build-your-own-interpreter/what-are-compilers)
are typically installed as part of the Xcode Command Line Tools. You have two
options here.

#### Xcode Command Line Tools

First we need to make sure that the xcode command line tools are installed:

```bash
xcode-select --install
```

After that you should be able to run the following command to verify the
installation:

```bash
xcode-select -p
# /Library/Developer/CommandLineTools
```

#### Clang

[Clang (from the LLVM Project)](https://clang.llvm.org/) is the default C, C++, and Objective-C compiler on macOS. It is included with the Xcode Command Line Tools. You can verify that it is installed by running:

```bash
clang --version
```

In my machine I get this after running the previous in my terminal:

```bash
Apple clang version 15.0.0 (clang-1500.3.9.4)
Target: arm64-apple-darwin23.0.0
Thread model: posix
InstalledDir: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin
```

#### GCC

The other alternative is [GCC (GNU Compiler Collection)](https://gcc.gnu.org/)
which can be installed using a package manager like [Homebrew](https://brew.sh/)
like this: `brew install gcc`

To install Homebrew run the following command in your terminal:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

After the installation completes, you can now install gcc:

```bash
brew install gcc
```

Once installed you can verify the version by running on your terminal:

```bash
gcc --version
```

### Linux

#### GCC

On Linux it depends on the distribution but the most common one is [GCC (GNU Compiler Collection)](https://gcc.gnu.org/). It is the default compiler in most distros. For example, Ubuntu, Fedora,
Debian, and Arch include GCC in their default repositories or pre-install it
with developer tools.

You can install it through your favorite package manager:

```bash
#Ubuntu/Debian
sudo apt intall gcc

# Fedora:

sudo dnf install gcc

# Arch:

sudo pacman -S gcc
```

Verify the installation by running the following command:

```bash
gcc --version
```

#### Clang

An alternative would be [Clang (from the LLVM Project)](https://clang.llvm.org/). You can also install with a package manager:

```bash
# Ubuntu/Debian
sudo apt intall clang

# Fedora:
sudo dnf install clang

# Arch:
sudo pacman -S clang
```

Then, verify the installation:

```bash
clang --version
```

### Windows

On windows there a many ways to get a C, C++, compatible compiler but they are
not typically included by default.

In this guide I will just show you the way that I prefer to do it and mention
briefly other alternatives.

#### Microsoft C/C++ Compiler (MSVC)

This is the standard compiler for C and C++ on Windows. It is not
pre-installed so we need to first install it. It is included in the
installation of Microsoft Visual Studio.

**Option 1:** Download and Install [Visual Studio Community](https://visualstudio.microsoft.com/downloads/)

Make sure to also select one C/C++ workload in the installation wizard. If you
use Visual Studio as your main editor and compiler you are all set.

**Option 2:** Download and Install the [Build Tools for Visual Studio](https://visualstudio.microsoft.com/downloads/#build-tools-for-visual-studio-2022)

You will need to select the **Desktop development with C++** workload. The
installation process will install the right tools for C/C++ development, this
includes compiler, likers, assemblers, and other build tools. There will be
separate x86 and x64 compilers and tools to build code for x86, x64, ARM and
ARM64 targets.

To start using the compiler and linker, we first need to setup the command
line environment. Depending on which is the target of your game, you will find
a `.bat` program called `vcvars[target].bat`. For x64, typically found in:

```bash
\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build
```

This batch program will configure your terminal with the correct environment variables. You can alternatively run the general `vcvarsall.bat` program with the correct arguments to setup your environment. For example:

```bash
.\vcvarsall.bat x64
```

This is will configure the compiler and other tools for x64 build.

After configuring your environment, you can now verify that you have the
compiler program available in your terminal:

```bash
cl --version
```

You can read more in detail in this article:
[Use the Microsoft C++ toolset from the command line](https://learn.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-170)
