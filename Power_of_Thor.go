package main

import "fmt"
//import "os"

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
    
    for {
        
        // remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
        var remainingTurns int
        fmt.Scan(&remainingTurns)
        
          if(initialTY == lightY){
                if (initialTX - lightX < 0) {
                    fmt.Println("E")
                    initialTX++
            
                } else {
                    fmt.Println("W")
                    initialTX--
                }   
            }  
        
            if(initialTX == lightX){
                if (initialTY - lightY < 0) {
                    fmt.Println("S")
                    initialTY++
            
                } else {
                    fmt.Println("N")
                    initialTY--
                }  
            }
        
        if initialTX <= 39 &&  initialTY <= 17{
    
            if  initialTX - lightX < 0{ 
                if (initialTY - lightY > 0)  {
                    fmt.Println("NE")
                    initialTY--
                    initialTX++
                }
            }
    
            if initialTX > lightX {
                if initialTY < lightY {
                    fmt.Println("SW")
                    initialTY++
                    initialTX--
                }
            }
        
            if initialTX - lightX > 0 && initialTY - lightY > 0 && initialTX <= 39 &&  initialTY <= 17{
                fmt.Println("NW")
                initialTY--
                initialTX--
            }

            if initialTX - lightX < 0 && initialTY - lightY < 0 && initialTX <= 39 &&  initialTY <= 17{
                fmt.Println("SE")
                initialTY++
                initialTX++
            }
        }  
    }
}
