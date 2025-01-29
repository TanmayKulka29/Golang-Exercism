package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
    if(avgTime==0) {
    	avgTime = 2 
    }
    return len(layers)*avgTime
}
// TODO: define the 'Quantities()' function
func Quantities(quantities []string) (int, float64) {
    var val1 int = 0
    var val2 float64 = 0.0
    for i:=0;i<len(quantities);i++ {
        if(quantities[i]=="noodles") {
        	val1+=50   
        } else if(quantities[i]=="sauce") {
            val2 += 0.2
        }
    }
    return val1, val2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList[]string, myList[]string) {
     myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities[]float64, portions int) []float64 {
    n := len(quantities)
    res := make([]float64, n)
    for i:=0;i<n;i++ {
        res[i] = quantities[i]/2
        res[i] = res[i]*float64(portions)
    }
    return res
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
