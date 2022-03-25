The challenge was solved using GOLANG

### >Instructions: 

It is necessary to have docker installed.

>> Building image

Before running the program it's necessary to build the image from our Dockerfile


```bash
# We need to be in the root directory of the project.

> docker build -t draftea-challenge .
```

>>Running container

We can run this program in two ways using docker:


>>>>Interactive mode

Writing position and instructions

```bash
> docker run -i --rm draftea-challenge
```

Example
```bash
5 3
-- New Robot --
1 1 E

RFRFRFRF
-- output --
 1 1 E 
------------

-- New Robot --
3 2 N

FRRFLLFFRRFLL
-- output --
 3 3 N LOST
------------

-- New Robot --
0 3 W

LLFFFLFLFL
-- output --
 2 3 S 
------------

```


>>>> Passing a file

Passing several operations from file.

```bash
> docker run -i --rm draftea-challenge < data.txt
```

>> Running with GO

```bash
# it's necessary to install GO and dependencies
go mod download 
go run ./cmd/cli/main.go < ./data.txt  

-- New Robot --
1 1 E

RFRFRFRF
-- last position --
    1 1 E 
-------------------

-- New Robot --
3 2 N

FRRFLLFFRRFLL
-- last position --
    3 3 N LOST
-------------------

-- New Robot --
0 3 W

LLFFFLFLFL
-- last position --
    2 3 S 
-------------------

-- Outputs --
    1 1 E 
    3 3 N LOST
    2 3 S

# Or just 
go run ./cmd/cli/main.go

5 3
1 1 E
-- New Robot --
1 1 E

RRFFRRFF
RRFFRRFF
-- last position --
    0 1 W LOST
```

>> Running tests

```bash
# it's necessary to install GO and dependencies
go mod download 
go test -v ./...
```

### Documentation

** Comments clarify the code and they are added with purpose of making the code easier to understand.

** TODO: Add more unit tests