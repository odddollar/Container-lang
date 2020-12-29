# Container-lang

Container-lang is a lightweight interpreted scripting language based on the idea of "code containers", aiming to make it easy to create reusable code for one/many line scripts. Container-lang currently only supports numerical data-types, including integers and floats. String support may be added in the future.

---

### Code containers

Code containers are the same as a line in other languages, however have a numerical ID that allows them to be referenced later, meaning the same container can be used many times. Additionally, containers can be placed sequentially on one line, allowing for whole programs to be written on one line.

A container is created with the syntax `{UNIQUE_NUMERICAL_ID|CODE}`. All code is required to be placed in a container.

Containers are executed from left to right, then top to bottom.

E.g.

```
{ 1 |x <- 10}{2|y <- 11}{ 3|z <- 0}
{4 |i<- 20}
```

is exactly the same as

```
{1|x <- 10}
{2|y <- 11}
{3|z <- 0}
{4|i <- 20}
```

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

### Print function

The print function is called using the syntax ```PRINT [VALUE_TO_PRINT]``` inside of a container and is used to display text in the console.

E.g.

```
This will print "10" to the console
{1|PRINT 10}
This will print "999" to the console
{2|PRINT 999}
```

### Execute function

The execute function to used to execute another container by its ID. This is one of the most powerful functions in the language as it allows for the reuse of containers as many times as required. 

The syntax used is ```EXECUTE [ID_OF_CONTAINER_TO_EXECUTE]```

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

