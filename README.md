# Character counter | ▮▮▮▮▮▮▮▮

Simple code just to try out and Binary Tree on Golang
Count characters to train openning a file and reading it, as well as printing out stuff more pretty formated.
Using a simple logic of max number of repeated chars and casting to float and then int to floor and create the repetitions to print it in kind of histogram form for a more didatic value.

## Running

git clone and than run : `go run ./cmd/characters_counter filename`

## Example

```
"testes teste bonito", receiving file content

➜  characters_counter go run ./cmd/characters_counter
Record of Repeats: 6
 't' | id:  116 | qtd:     6 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'e' | id:  101 | qtd:     4 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 ' ' | id:   32 | qtd:     2 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'b' | id:   98 | qtd:     1 | ▮▮▮▮▮▮▮▮
 's' | id:  115 | qtd:     3 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'o' | id:  111 | qtd:     2 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'n' | id:  110 | qtd:     1 | ▮▮▮▮▮▮▮▮
 'i' | id:  105 | qtd:     1 | ▮▮▮▮▮▮▮▮
```

```
"testes teste bonito
como a vida é bela, nos que fode com ela.
parabens senhor, vc é bacanal", receiving file content

➜  characters_counter go run ./cmd/characters_counter
Record of Repeats: 15
 't' | id:  116 | qtd:     6 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'e' | id:  101 | qtd:    10 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 ' ' | id:   32 | qtd:    15 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 '
' | id:   10 | qtd:     2 | ▮▮▮▮▮▮
 'b' | id:   98 | qtd:     4 | ▮▮▮▮▮▮▮▮▮▮▮▮▮
 'a' | id:   97 | qtd:     9 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 ',' | id:   44 | qtd:     2 | ▮▮▮▮▮▮
 '.' | id:   46 | qtd:     1 | ▮▮▮
 'c' | id:   99 | qtd:     4 | ▮▮▮▮▮▮▮▮▮▮▮▮▮
 'd' | id:  100 | qtd:     2 | ▮▮▮▮▮▮
 's' | id:  115 | qtd:     6 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'o' | id:  111 | qtd:     8 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'n' | id:  110 | qtd:     5 | ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮
 'i' | id:  105 | qtd:     2 | ▮▮▮▮▮▮
 'f' | id:  102 | qtd:     1 | ▮▮▮
 'h' | id:  104 | qtd:     1 | ▮▮▮
 'm' | id:  109 | qtd:     2 | ▮▮▮▮▮▮
 'l' | id:  108 | qtd:     3 | ▮▮▮▮▮▮▮▮▮▮
 'q' | id:  113 | qtd:     1 | ▮▮▮
 'p' | id:  112 | qtd:     1 | ▮▮▮
 'r' | id:  114 | qtd:     2 | ▮▮▮▮▮▮
 'v' | id:  118 | qtd:     2 | ▮▮▮▮▮▮
 'u' | id:  117 | qtd:     1 | ▮▮▮
 'Ã' | id:  195 | qtd:     2 | ▮▮▮▮▮▮
 '©' | id:  169 | qtd:     2 | ▮▮▮▮▮▮
```
