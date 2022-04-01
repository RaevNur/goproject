# ASCII-ART
### In ascii-art you can:
- Output text as a picture Mask
- Set the color of your ascii output
- Change the text style
- Align the text as you like
- Save your result to a file and vice versa
## Instruction
### Basic output
> The first argument passed is considered as the text that you want to output
The program accepts only Ascii characters that are in the range from 32-126 ( ` ` - `~` ) inclusive.

`go run . "1 Line\n2 line lee"`
### Colorfull Output 
> Colors as well as other parameters are entered after the first argument (the text that you want to output).
You can pass 3 sub-parameters to the color parameter
(color [characters/substrings] [from-to / index of a character in a string])
Template: `"--color=RGB(255,255,255) [symbols={Characters}/substr={SubStrings}] [in={from-to/idx of element}]"`

> he color takes the name of the colors and RGB
> `--symbols={}\--substr={}` cannot accept spaces
##### Examples:
1.1 To colorize the entire text:<br>
`go run . "1 Line\n2 line lee" "--color=RGB(100, 220, 200)"`

2.1 Color only certain characters (all `le` characters):<br>
`go run . "1 Line\n2 line lee" "--color=RGB(100, 220, 200) symbols={le}"`

2.2 Color only certain characters (First 2 `l` and `e` characters):<br>
`go run . "1 Line\n2 line lee" "--color=red symbols={le} in={1-2}"`

3.1 Color only certain substrings (all `ne` substrings):<br>
`go run . "1 Line\n2 line leene" "--color=green substr={ne}"`

3.2 Color only certain substrings (First 2 `ne` substrings):<br>
`go run . "1 Line\n2 line neene" "--color=blue substr={ne} in={1-2}"`

4.1 You Can Combine Colors:<br>
`go run . "1 Line\n2 line neene" "--color=RGB(100, 220, 200) symbols={le}" "--color=blue substr={ne} in={1}"`

### Text Align
To use text align, the `--align=`parameter is used

##### Examples:
1. Default text align is `left`:<br>
`go run . "1 Line\n2 line lee" "--align=left"`

2. Text align `right` (Text moving to right wall):<br>
`go run . "1 Line\n2 line lee" "--align=right"`

3. Text align `center` (Text moving to center):<br>
`go run . "1 Line\n2 line lee" "--align=center"`

4. Text align `justify` (Removing space and The text is distributed over the entire width):<br>
`go run . "1 Line\n2 line lee" "--align=justify"`

5. For text align, the last entered parameter text align is accepted (Only the last parameter `--align=center` will be initialized)<br>
`go run . "1 Line\n2 line lee" "--align=justify" "--align=center" "--color=purple"`

### Show With Any Themes (Fonts)
Themes are how your characters will look in the output. You can create your own theme, as long as it complies with certain rules: 
- The file must be in the `txt` format
- The created file must be moved to the `themes` folder
- Each character must have a height of 8
- Look at the `standard` theme and create your own file and create your own theme

You can see all the themes in the themes folder
To use them, simply enter the file name in the themes folder without the extension

##### Examples:
1. Default theme `standard`:<br>
`go run . "1 Line\n2 line lee" standard`

2. Theme `thinkertoy`:<br>
`go run . "1 Line\n2 line lee" thinkertoy`

2. Themes like align takes last theme `thinkertoy`:<br>
`go run . "1 Line\n2 line lee" thinkertoy shadow`

> You can use custom themes


### Save Input To File
To save to a file, you need to use the `--output=` parameter and the name of the file to which you want to save the output. When saving, the program creates a file in which the result that was supposed to be output is written.
> You can also specify the theme with which you want to save the file

##### Examples:
Let's say we want to save to a file `example.txt`
```bash
$ go run . "Hello World" --output=example.txt
$ cat example.txt -e
 _    _          _   _                __          __                 _       _  $
| |  | |        | | | |               \ \        / /                | |     | | $
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | $
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | $
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | $
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| $
                                                                                $
                                                                                $
```

Let's try with a different theme
```bash
$ go run . "Hello\nWorld" --output=example1.txt shadow
$ cat example1.txt -e
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
                                             $
_|          _|                   _|       _| $
_|          _|   _|_|   _|  _|_| _|   _|_|_| $
_|    _|    _| _|    _| _|_|     _| _|    _| $
  _|  _|  _|   _|    _| _|       _| _|    _| $
    _|  _|       _|_|   _|       _|   _|_|_| $
                                             $
                                             $
```

### Read Ascii-Art In File
With the help of `--reverse=` you can read ascii drawing and the program will output it as plain text
> By default, there is a standard theme and you can read these files by setting different themes

##### Examples:
Let's say we want to save to a file `example.txt`
```bash
$ cat example.txt -e
 _    _          _   _                __          __                 _       _  $
| |  | |        | | | |               \ \        / /                | |     | | $
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | $
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | $
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | $
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| $
                                                                                $
                                                                                $
$ go run . --reverse=example.txt
Hello World
$
```
And now let's read the file with a different topic
```bash
$ cat example1.txt -e
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
                                             $
_|          _|                   _|       _| $
_|          _|   _|_|   _|  _|_| _|   _|_|_| $
_|    _|    _| _|    _| _|_|     _| _|    _| $
  _|  _|  _|   _|    _| _|       _| _|    _| $
    _|  _|       _|_|   _|       _|   _|_|_| $
                                             $
                                             $
$ go run . --reverse=example1.txt shadow
Hello
World
$
```

### The program was written:
- [Dias1c](https://github.com/Dias1c)
- [nrblzn](https://github.com/RaevNur)
