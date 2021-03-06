# Container-lang

Container-lang is a lightweight, interpreted, esoteric scripting language based on the idea of "code containers", aiming to make it easy to create reusable code for one/many line scripts. Container-lang currently only supports numerical data-types, including integers and floats. String support is unlikely to be added in the future.

## Usage
Prebuilt interpreter files are available on the GitHub page. 

To run an Container-lang file (```.cnl``` extension), open a command prompt and type ```container-lang -f [PATH_TO_FILE]```.

Specify the ```-d``` flag to print debug information, including the token stream generated from the input file and the program completion time.

To build from source, clone the Git repo using ```git clone https://github.com/odddollar/Container-lang.git```, then run ```go build Container-lang```.

## Code containers

Code containers are the same as a line in other languages, however have a numerical ID that allows them to be referenced later, meaning the same container can be used many times. Additionally, containers can be placed sequentially on one line, allowing for whole programs to be written on one line.

A container is created with the syntax `{UNIQUE_NUMERICAL_ID|CODE}`. All code is required to be placed in a container.

Containers are executed from left to right, then top to bottom.

E.g.

```
{ 1 |x <- 10}{2|y <- 11}{ 3|z <- 56}
{4 |i<- 20}
```

is exactly the same as

```
{1|x <- 10}
{2|y <- 11}
{3|z <- 56}
{4|i <- 20}
```

## Language reference

### Comments 

Due to the way the interpreter splits the input file into code containers, any text not within a container will be treated as a comment and thus be ignored by the interpreter.

E.g.

```
{1|x <- 10}
{2|PRINTLN 5}{3| PRINTLN 10}
This is a comment
{4|mult<- x*22}
```

```
{1|x <- 10}{2|PRINTLN 5}{3| PRINTLN 10}This is also a valid comment{4|mult<- x*22}
```

### Variables

Much like other languages, variables are "boxes" that store a value for later referencing. These values can be updated at any time. However, unlike other languages that use the ```=``` symbol to assign values, Container-lang uses the arrow symbol, ```<-```.

All variables are global (E.g. those in a block can be accessed by those not in a block and vice versa).

E.g.

```
This will assign the value 1 to the variable "x"
{1|x<-1}
```

Mathematical operations can also be performed on variables using this same method. 

E.g.

```
{1|x <- 1}
{2|PRINTLN x}
{3|x <- x+1}
{4|PRINTLN x}
```

### Blocks

Blocks allow for the grouping together of other containers to be executed later. They're very similar to functions in other programming languages, however they don't accept parsing in arguments. Blocks can be placed inside other blocks.

Blocks use the syntax ```BLOCK [NUMBER_OF_CONTAINERS_TO_GROUP]```, where the number of containers argument states the next ```x``` number of containers to put in the block.

Similarly to functions, blocks won't be run until they're called. Blocks can be run/called using the ```EXECUTE``` function (continue reading) and the block container's ID, which will run all the containers contained within that block.

Blocks are incredibly useful for creating chunks of code that can be reused many times.

**NOTE**: Blocks must be defined before they can be called.

**TIP**: While not necessary, it makes reading a program so much easier if you indent the containers contained within a block. This whitespace is ignored by the interpreter.

E.g.

```
This will create a block with an ID of 1 that contains the next 2 containers
{1|BLOCK 2}
	{2|PRINTLN 10}
	{3|PRINTLN 11}
To run this block, use the EXECUTE function
{4|EXECUTE 1}
```

```
{1|PRINTLN 1}
{2|BLOCK 2}
	{3|PRINTLN 10}
	{4|PRINTLN 11}
{5|PRINTLN 2}
{6|EXECUTE 2}
This will print
"1
2
10
11"
```

Containers within a block can't be executed by containers outside the block.

E.g.

```
This will result in a "No container with specified ID" error
{1|BLOCK 2}
	{2|PRINTLN 10}
	{3|PRINTLN 11}
{4|EXECUTE 2}
```

Blocks can contain other blocks. This can lead to some interesting nested containers scenarios.

E.g.

```
This will print
"36
70
1
10
11
12"
{1|BLOCK 8}
	{2|BLOCK 2}
		{3|PRINTLN 10}
		{4|PRINTLN 11}
	{5|PRINTLN 1}
	{8|BLOCK 1}
		{9|PRINTLN 12}
	{11|EXECUTE 2}
	{12|EXECUTE 8}
{6|PRINTLN 36}
{7|PRINTLN 70}
{10|EXECUTE 1}
```

### Repeat

The repeat function is similar to a for loop in other languages, it repeats the given code a set number of times. 

It uses the syntax ```REPEAT [CONTAINER_ID_TO_REPEAT], [NUMBER_OF_TIMES_TO_REPEAT]```.

The repeat function implicitly creates a variable that keeps track of its current iteration status, named ```i[CONTAINER_ID_OF_REPEAT]```, E.g. A repeat container with an ID of ```3``` will implicitly create an iterator variable called ```i3```. 

It is not recommended to write to these variables, it can cause some weird situations, however it is possible.

E.g. 

```
This will print 
"10
10
10"
{1|BLOCK 1}
	{2|PRINTLN 10}
{3|REPEAT 1, 3}
```

In the above example, the print function is encased in a block to prevent it running by default. If the print wasn't in a block, ```10``` would be printed 4 times.

```
This will print
"1
1
1
1
1
1"
{1|BLOCK 3}
	{2|BLOCK 1}
		{3|PRINTLN 1}
	{4|REPEAT 2, 3}
{5|REPEAT 1, 2}
```

Below is an example of the implicitly created iteration variable.

```
{1|BLOCK 1}
	{2|PRINTLN i3}
{3|REPEAT 1, 5}
```

### If

If statements exist as they would in any other language, with a condition, code that runs if the condition is true and optional code that runs if the condition is false.

If statements use the syntax ```IF [CONDITION], [CONTAINER_ID_TO_RUN_IF_TRUE], [CONTAINER_ID_TO_RUN_IF_FALSE (Optional)]```.

E.g.

```
{1|BLOCK 1}
	{2|PRINTLN 10}
{3|BLOCK 1}
	{4|PRINTLN 20}
{5|x <- 1}
{6|IF x>11, 1, 3}
```

The print functions are wrapped in blocks to prevent them from running by default, only executing when called by the if statement.

### Execute function

The execute function to used to execute another container by its ID. This is one of the most powerful functions in the language as it allows for the reuse of containers as many times as required. 

The syntax used is ```EXECUTE [ID_OF_CONTAINER_TO_EXECUTE]``` inside of a container.

E.g.

```
This will execute the code in the container with ID 1
{2|EXECUTE 1}
```

```
This will print 
"10
10
10"
{1|PRINTLN 10}
{2|EXECUTE 1}
{3|EXECUTE 1}
```

Execute functions can also be used to run other execute functions.

E.g.

```
This will print 
"10
10
10"
{1|PRINTLN 10}
{2|EXECUTE 1}
{3|EXECUTE 2}
```

### Print function

The print function is called using the syntax ```PRINT [VALUE_TO_PRINT]``` or ```PRINTLN [VALUE_TO_PRINT]```inside of a container and is used to display text in the console.

```PRINT``` doesn't add a newline character, ```PRINTLN``` does.

To print a blank character (E.g. a space), provide ```BLANK``` as an argument for ```PRINT``` or ```PRINTLN```.

E.g.

```
This will print "10" to the console
{1|PRINTLN 10}
This will print "999" to the console
{2|PRINTLN 999}
```

Variables can be printed by placing the variable name as the argument, as well as expressions involving variables.

E.g.

```
{1|x <- 1}
{2|PRINTLN x}
{3|PRINTLN x+2}
```

## Examples

Basic 4 bit binary counter

```
{10|BLOCK 10}
	{7|BLOCK 8}
		{1|BLOCK 6}
			{2|BLOCK 4}
				{12|PRINT i11}
				{9|PRINT i8}
				{3|PRINT i6}
				{4|PRINTLN i5}
			{5|REPEAT 2, 2}
		{6|REPEAT 1, 2}
	{8|REPEAT 7, 2}
{11|REPEAT 10, 2}
```

Find factors of 2000000

```
{6|num <- 2000000}
{1|BLOCK 3}
    {4|BLOCK 1}
        {5|PRINTLN i2}
    {3|IF num%i2==0, 4}
{2|REPEAT 1, num+1}
```

Find prime numbers up to 1500

```
{11|BLOCK 9}
    {1|factors <- 0}
    {7|BLOCK 3}
        {4|BLOCK 1}
            {5|factors <- factors+1}
        {2|IF i12%i6==0, 4}
    {6|REPEAT 7, i12+1}
    {8|BLOCK 1}
        {10|PRINTLN i12}
    {9|IF factors==2, 8}
{12|REPEAT 11, 1500}
```
