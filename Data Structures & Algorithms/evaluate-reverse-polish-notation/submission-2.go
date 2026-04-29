func pop(stack []int) (int, []int) {
	if len(stack) == 0 {
		return -1, stack
	}
	last := len(stack) - 1
	top := stack[last]
	stack = stack[:last]
	return top, stack
}

func handleOp(op1, op2 int, token string) (int, error) {
	switch token {
	case "+":
		return op1 + op2, nil
	case "-":  // reverse the order of op since it's a stack.
		return op2 - op1, nil
	case "*":
		return op1 * op2, nil
	case "/":  // reverse the order of op since it's a stack.
		return op2 / op1, nil
	default:
		return -1, fmt.Errorf("problem while operating")
	}
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		if tokens[0] != "+" && tokens[0] != "-" && tokens[0] != "*" && tokens[0] != "/" {
			iToken, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(fmt.Sprintf("invalid token: %s", tokens[0]))
			}
			return iToken
		}
	}
	var stack []int // i'm gonna assume that this is my stack.
	var res int

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			iToken, err := strconv.Atoi(token)
			if err != nil {
				panic(fmt.Sprintf("invalid token: %s", token))
			}
			stack = append(stack, iToken)
			fmt.Println(stack)
		} else {
			var op1, op2 int
			var err error
			op1, stack = pop(stack)
			op2, stack = pop(stack)

			res, err = handleOp(op1, op2, token)
			if err != nil {
				panic(fmt.Sprintf("invalid operation: %v", err))
			}

			stack = append(stack, res)
		}
	}

	return res
}
