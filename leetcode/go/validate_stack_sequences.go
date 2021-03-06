func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    j := 0
    N := len(pushed)
    
    for _, x := range pushed {
        stack = append(stack, x)
        for (len(stack) != 0 && j <  N && stack[len(stack)-1] == popped[j]) {
            stack = stack[0:len(stack)-1]
            j++
        }
    }
    
    return j == N
}