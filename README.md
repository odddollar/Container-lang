# Container-lang

Container-lang is a lightweight interpreted scripting language based on the idea of "code containers", aiming to make it easy to create reusable code for one/many line scripts.

---

### Code containers

Code containers are the same as a line in other languages, however have a numerical id that allows them to be referenced later, meaning the same container can be used many times. Additionally, containers can be placed sequentially on one line, allowing for whole programs to be written on one line.

A container is created with the syntax `{UNIQUE_NUMERICAL_ID|CODE}`. All code is required to be placed in a container

Containers are executed from left to right, then top to bottom.

E.g.

```
{1|x <- 10}{2|y <- 11}{3|z <- 0}
{4|i <- 20}
```

is exactly the same as

```
{1|x <- 10}
{2|y <- 11}
{3|z <- 0}
{4|i <- 20}
```