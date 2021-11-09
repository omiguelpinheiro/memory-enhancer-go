# memory-enhancer
Helps exercise your memory by giving you random tokens and poems to memorize.

## Using
Every day when you first open your terminal you will be greated with a token and a piece of a poem. Your job is to remember both of them till the end of the day. But this is just level one of course, you can play with this yourself like using the history located at `$HOME/.memory-enhancer/history` to try to remember passwords and poems from past days and see if you can get more than one day right.

## Install
The installation process can be done in two ways, through a bash script or compilation from source. Both steps **MUST FOLLOW** the last steps section.

### Cloning
The easiest way to install is just cloning the repository to the default installation folder
```bash
git clone https://github.com/omiguelpinheiro/memory-enhancer.git "$HOME/.memory-enhancer"
```
### Compiling from source
You can also compile the code from source using
```bash
git clone https://github.com/omiguelpinheiro/memory-enhancer.git "$HOME/.memory-enhancer" && go build -o "$HOME/.memory-enhancer/memory" "$HOME/.memory-enhancer/main"
```
### Last steps
You **MUST** configure your `.zshrc` or `.bashrc` file in order for the program to work, run this switching `.zshrc` for `.bashrc` if you're using bash
```bash
echo "$HOME/.memory-enhancer/memory" >> "$HOME/.zshrc"
```
also, your operating system, as the good boy he is, may block you from executing the *potentially dangerous* program you just downloaded from my repository by giving you a `permission denied` error when trying to execute it. You can fix this, not at all surprisingly, by giving the binaries permission to execute. Use this command if you get the error
```bash
sudo chmod +w "$HOME/.memory-enhancer/memory"
```
And that's it, you're good to go. Also, did you know that Jupiter is twice as massive as all of the other planets of the solar system combined? So, yeah ... if you're here to memorize useless facts I think this proves you're a little step closer to success.
## Options
The script also comes with 3 parameters that you can play with and can be changed in the `memory.cfg` file located in the installation folder `.memory-enhancer` located in your home directory. Change the following options to modify:
```
TOKEN_LENGHT -> size of the token (default is 4)
LINE_MAX -> lines in the piece of poem (default is 2)
LINE_TOL -> the program will try to find verses that fit in your max number of lines with LINE_TOL lines of tolerance first, if that's not possible it will pick a random part instead.
```
## Uninstall
To uninstall just delete the .memory-enhancer folder from your home with
```bash
rm -rf "$HOME/.memory-enhancer/memory"
```
and delete this line from your `.zshrc` or `.bashrc` file
```bash
$HOME/.memory-enhancer/memory
```
but I really hope your PC crashes forever for abandoning me.