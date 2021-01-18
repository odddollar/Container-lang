# Container-lang

Container-lang is a lightweight interpreted scripting language based on the idea of "code containers", aiming to make it easy to create reusable code for one/many line scripts. Container-lang currently only supports numerical data-types, including integers and floats. String support may be added in the future.

---

### Code containers

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

---

## Language reference

### Comments 

Due to the way the interpreter splits the input file into code containers, any text not within a container will be treated as a comment and thus be ignored by the interpreter.

E.g.

```
{1|x <- 10}
{2|PRINT 5}{3| PRINT 10}
This is a comment
{4|mult<- x*22}
```

```
{1|x <- 10}{2|PRINT 5}{3| PRINT 10}This is also a valid comment{4|mult<- x*22}
```

### Variables

Much like other languages, variables are "boxes" that store a value for later referencing. These values can be updated at any time. However, unlike other languages that use the ```=``` symbol to assign values, Container-lang uses the arrow symbol, ```<-```.

E.g.

```
This will assign the value 1 to the variable "x"
{1|x<-1}
```

Mathematical operations can also be performed on variables using this same method. E.g.

```
{1|x <- 1}
{2|PRINT x}
{3|x <- x+1}
{4|PRINT x}
```

### Blocks

Blocks allow for the grouping together of other containers to be executed later. They're very similar to functions in other programming languages, however they don't accept parsing in arguments. Additionally, blocks can't be placed inside other blocks; attempting to do so will result in an ```Unrecognised request``` error.

Blocks use the syntax ```BLOCK [NUMBER_OF_CONTAINERS_TO_GROUP]```, where the number of containers argument states the next ```x``` number of containers to put in the block.

Similarly to functions, blocks won't be run until they're called. Blocks can be run/called using the ```EXECUTE``` function (continue reading) and the block container's ID, which will run all the containers contained within that block.

Blocks are incredibly useful for creating chunks of code that can be reused many times.

E.g.

```
This will create a block with an ID of 1 that contains the next 2 containers
{1|BLOCK 2}
{2|PRINT 10}
{3|PRINT 11}
To run this block, use the EXECUTE function
{4|EXECUTE 1}
```

```
{1|PRINT 1}
{2|BLOCK 2}
{3|PRINT 10}
{4|PRINT 11}
{5|PRINT 2}
{6|EXECUTE 2}
This will print
"1
2
10
11"
```

### Print function

The print function is called using the syntax ```PRINT [VALUE_TO_PRINT]``` inside of a container and is used to display text in the console.

E.g.

```
This will print "10" to the console
{1|PRINT 10}
This will print "999" to the console
{2|PRINT 999}
```

Variables can be printed by placing the variable name as the argument, as well as expressions involving variables. E.g.

```
{1|x <- 1}
{2|PRINT x}
{3|PRINT x+2}
```

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
{1|PRINT 10}
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
{1|PRINT 10}
{2|EXECUTE 1}
{3|EXECUTE 2}
```
