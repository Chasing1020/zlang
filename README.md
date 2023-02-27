# Introduction

zlang is a dynamic language interpreter (aka z language) based on Go language implementation. 
It combines some syntax of JavaScript and Python and supports basic control flow of traditional languages,
including conditional statements, standard I/O, function closures, recursion, and so on. 
It is currently under active development.

## 1. Quick Start

If you have already installed the Go language environment, you can choose to download and compile it directly:

```shell
git clone https://github.com/Chasing1020/zlang.git
go mod tidy
go build
```

Alternatively, you can choose to use the compiled version directly. Try running:

```shell
./zlang run main.zjc
# Hello, world!
```

If you see "Hello, world!" outputted, then it means you have successfully installed zlang.

### 1.1. Data Types

Currently, zlang supports six basic data types: int, string, boolean, array, map, and function.

Every time you create a new variable, you need to use "let", for example:

```js
let int = 1;
let string = "a string";
let boolean = true;
let array = [1, 2, 3, 4, 5];
let map = {"k": "v"};
let add = function (a, b) {
    return a + b;
};
```

Array type: similar to the list in Python, it can store data of any type, like this:

```js
let arr = [1, 2, function (a, b) {println(a + b);}]
arr[2](arr[0], arr[1]) // 3
```

Map type: to avoid hash collisions, keys can only be int or string types. However, in one map, you can mix and match both types, like this:

```js
let map = {"chasing":1020, 1020:"chasing"}
println(map["chasing"], map[1020]) // 1020, chasing
```

### 1.2. Operators

The current version supports basic operators such as +, -, *, /, %, <, >, <=, >=, !=, ==, with the same order of precedence as in C language.

### 1.3. Control Flow

The usage of if and for loops is the same as that in C-like languages:

```js
if (true) { print("true"); } else { print("false"); }

let sum = 0;
for (let i = 0; i <= 100; i = i + 1) {
    sum = sum + i;
}
println(sum); // 5050
```

### 1.4. Built-in Functions

Evaluate a string as an expression: eval(x)
Standard output: print(x), println(x), printf(fmt, x)
Standard input: input(), which returns a string type by default
Get length: len(x), which returns the length of an array or a string
Create a new array: newArray(x), which creates a new integer array with a length of x, and is initialized to 0 by default
Type conversion: string(x), int(x), which converts between string and int types
Comparison functions: min(a, b), max(a, b), which return the minimum and maximum values of two integers

## 2. Basic Commands

To run the program, simply enter "./zlang" in the command line. To exit, press ctrl+c twice or ctrl+D once.

```shell
$ ./zlang                         
Welcome to zLang v0.0.1.
Type "help()" for more information.
> ^C
(To exit, press Ctrl+C again or Ctrl+D)
> ^C%                             
```

## 3. Todo List

Support floating point numbers and object-oriented programming.
Optimize function recursion stack copying logic.
Support finer control structures such as "break" and "continue".
Support operations such as "++", "--", "+=", and "-=".

## 4. Software License

Licensed under the Apache License 2.0.