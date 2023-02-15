# GOven

GOven is a very simple command-line tool for generating secure passphrases. It's written in Go using only the standard library.  
GOven was inspired by [EFF Dice-Generated Passphrases](https://www.eff.org/dice).

The name *GOven* is a pun between "Go" and "güven", which means *trust* in Turkish. "Güvenli" means *secure*, but güven was preferred for brevity.

| :warning: WARNING |
|-------------------|
| I'm not a cryptographer, so it's recommended to check the source code before using GOven. |

## Usage

To generate a passphrase, simply run `goven` in your terminal. By default, 4 lowercase words will be generated, separated by a "-".  
You can use flags to customize the output:

- `help`, `h`: show help message;
- `length N`: generate `N` words instead of the default 4;
- `pascalcase`: use PascalCase instead of lowercase;
- `separator STRING`: specify a separator between words (default "-", "" for none);
- `wordlist PATH`: specify a file to use instead of the default wordlist; see [Custom wordlists](#custom-wordlists) for more info;

Note that using flags with one or two "-" is the same: `$ goven -length 5` will do the same as `$ goven --length 5`.

### Custom wordlists

By default, GOven selects words from the [EFF's Long Wordlist](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt), which contains 7776 words, but users can specify another wordlist through the `-wordlist` flag.  
Files used as wordlists must contain *exactly one word on each line*, or else unexpected behavior will occur.

For example, if `my-words.txt` has this content:

```
banana
hummus
kebab
lasagna
pasta
pizza
popcorn
sushi
tacos
tofu
```

then the command `$ goven -wordlist my-words.txt -pascalcase -length 4 -separator "#"` will output something like this:

```
Hummus#Tacos#Pizza#Popcorn

```

### Copy to clipboard

To copy a passphrase to the clipboard you will need an external tool, such as xclip:

```
# Generate a passphrase using default settings and copy it to the system's clipboard
$ goven | xclip -sel clip
```

