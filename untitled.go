if input == w{
fmt.Println(input, "TrueIF")
                   } else if input == l{
                        fmt.Println("else if")

                    } else {

                    guess,_ = strconv.Atoi(input)
                    fmt.Println("Ball =", guess)
                    score=guess+score
                    fmt.Println("Score =", score)
                }