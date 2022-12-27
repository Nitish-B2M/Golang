# Golang

For Linux User
=> Step-1 [ Install golang ]
  sudo apt install golang-go

=> Step-2 [ Create any workspace where you want to work ]

=> Step-3 [ open terminal of that workspace and run below command ]
  go mod init <module_name> 
  
example:- go mod init golang_learning

=> Step-4 [ Now create go file with extension .go ]
  In my case, i have created main_01.go [ Just to test is golang working or not properly ]
  
  Content inside main_01.go 
  
      package main

      import ( 
        "fmt" 
      )

      func main() {
        fmt.Println("Hello, World!")
      }

=> Step-5 [ go to same workspace terminal and run below command ]
  go run main_01.go  [ go run file_name.go ]
  
=> Step-6 
  Now you will able to print Hello, World! in terminal
  
