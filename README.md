# Tic Tac Toe
This repository contains Golang based Tic-Tac-Toe game.

### Table of Contents
**[Assignment](#assignment)**<br>
**[Installation](#installation)**<br>
**[Play](#play)**<br>
**[Design Doc, Testing info](#design-doc)**<br>
**[Future enhancement(TODO)](#future-enhancementtodo)**<br>
**[References](#references)**<br>

## Assignment
Implement a HTTP client and a HTTP server application that together play a game of tic-tac-toe
on a 3x3 board, in golang. Client starts the game and is using crosses. Server gets client moves
and responds with reasonable defence moves.

## Installation
```
git clone https://github.com/nvyas-lifesize/tic-tac-toe
```
(**Prerequisite**: Download and Install [Go](https://golang.org/doc/install))

## Play

1. Go to ```cd /{your-directory}/tic-tac-toe/server``` and run ```go run .``` which starts the server
2. Go to ```cd /{your-directory}/tic-tac-toe/client``` and run ``` go run .``` which gives you the game board screen as follows:
    ```
    . | . | . | 
    . | . | . | 
    . | . | . | 
   Enter your choice from 1 to 9:
   ```
   *Every grid is represented by the number, ```[0,0] -> 1, [0,1] - > 2 ....[2,2] -> 9``` for users convenience to choose the input
3. Select the number where you want to add `X`
4. Computer choose the `O` accordingly
5. Display the result of the game at the end(Player wins/Computer wins/Draw)

## Design Doc
[Design Document and Testing info](https://drive.google.com/file/d/1OSJR8Tm4kUeLEq87uZcClWQPxj0LbYUQ/view?usp=sharing)

## Future enhancement(TODO)
1. Add functionality to have option for user to choose `X` or `O` at the beginning
2. Add functionality for multi-player option
3. Add functionality to reset the game at any point
4. Add functionality to choose difficulty level like, Easy, Medium or Hard
5. Add the fancy User interface
6. Add the background music
7. Add functionality to host the Game on internet

## References
1. [Learn Go](https://golang.org/doc/tutorial/getting-started)
2. [Minimax Algorithm](https://en.wikipedia.org/wiki/Minimax)