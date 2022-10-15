# Simple CLI task list

Disclaimer: A simple task list manager written in Go. The project does not attempt to promote the best practices of Go

## Requirements

Go 1.18 (only if you want to compile it by yourself)

## Installation

There is compiled binary in `bin/` folder, you can use it, or install it by yourself:

```sh
# Clone git repo to local machine
git clone https://github.com/hinotora/go-task-list.git

# Cd to project folder
cd go-task-list

# Compile binary file with make, it will create binary `main` file in this directory, you can rename it
make

# Add execute rights to your binary
chmod +x main
```

## Usage (bash)

```sh

# 1. Get all your tasklists names and identifiers
./go-task-list --action=all

# Output
List id: bb24e8c1-1068-4638-b2e8-ccd8f3cf86bc 
List name: My todo list 
List date created: 15.10.2022 18:57




# 2. Get specific list with tasks using list identifier
./go-task-list  --action=list --list=identifier

# Output
Name: My todo list 
Date created: 15.10.2022 18:57
======================================
[3a8841cf-8aae-459e-abc0-2db45092a5a0] [-]: Go to shop 
[a405a107-65a4-4237-b922-28b1e7283d1b] [-]: Clean my room



# 3. Create new list
./go-task-list  --action=createlist --lname="List name"

# Output
Created new list [My todo list] with identifier [bb24e8c1-1068-4638-b2e8-ccd8f3cf86bc].



# 4. Create new task in list
./go-task-list  --action=createtask --lid=ListIdentifier --tname="Task name"

# Output
Added new task with identifier [3a8841cf-8aae-459e-abc0-2db45092a5a0] to list [My todo list]



# 5. Remove task from list
./go-task-list  --actoon=removetask --lid=ListIdentifier --tid=TaskIdentifier

# Output
Removed task from list [My todo list]



# 6. Mask task as done
./go-task-list  --action=done --lid=ListIdentifier --tid=TaskIdentifier

# Output
Task [Clean my room] marked as done 



# 7. Remove list
./go-task-list  --action=removelist --lid=ListIdentifier

# Output
Removed list [bb24e8c1-1068-4638-b2e8-ccd8f3cf86bc] from memory. 



```

## License

Nah, do whatever you want, but... make link to this github repo

