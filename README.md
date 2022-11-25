# Welcome to maths module!

a sample for creating modules in go and publish it!


# First create a module 

go mod init github.com/usesrname/modulename

## sample program

    package mathematics
    
     Sum calculates the total from a slice of numbers.
    
    func  Sum(numbers []int) int {
    
    var sum int
    
    for _, number :=  range numbers { 
    sum += number
    }
    return sum
    }

## Sample testing

   

     package mathematics
    
      
    
    import  "testing"
    
      
    
    func  TestSum(t *testing.T) {
    
      
    
    numbers := []int{1, 2, 3, 4, 5}
    
      
    
    got :=  Sum(numbers)
    
    expected :=  15
    
      
    
    if expected != got {
    
    t.Errorf("expected %d got %d", expected, got)
    
    }
    
    }

## Commands to publish:

    git add .  
    git commit -m “my first module, yay!”  
    git tag **v1.0.0**  
    git push origin **v1.0.0**

##  Let’s use our module in our project!

    package main
    
    import (
    
    "fmt"
    
    "github.com/halilylm/mathematics"
    
    )
    
    func main() {
    
    nums := []int{1, 4, 7}
    
    result := mathematics.Sum(nums)
    
    fmt.Println(result)
    
    }

##  Congratulations!
