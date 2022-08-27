// class FlightPathsController < ApplicationController
//   ENDPOINTS_MALFORMED_INPUT_ERROR = 
//       'Please ensure your input is in the proper format.
//        This endpoint accepts an array of ticket arrays that contain two elements, an airport source and an airport destination.
//        At least one ticket must be provided.
//        e.g. [ [source, destination], [source, destination] ]'

//   def endpoints
//     render json: FlightPath.new(tickets).airport_endpoints
//   rescue
//     render plain: ENDPOINTS_MALFORMED_INPUT_ERROR 
//   end

//   private

//   def tickets
//     params.require(:tickets)
//   end
// end


package flight

import "fmt"

func Parser(entry [][]string) ([]string, error) {
	answer := []string{}

	stack1 := []string{}
	stack2 := []string{}

	//------------------------------- push the element to the stack
	// ------------------------------we are going to pop it in both stacks
	for _, v := range entry {
		stack1 = append(stack1, v[0])
		stack2 = append(stack2, v[1])
		for i1, v1 := range stack1 {
			if v1 == v[1] {
				stack1 = append(stack1[:i1], stack1[i1+1:]...)
				stack2 = stack2[:len(stack2)-1]
				break
			}
		}
		for i2, v2 := range stack2 {
			if v2 == v[0] {
				stack2 = append(stack2[:i2], stack2[i2+1:]...)
				stack1 = stack1[:len(stack1)-1]
				break
			}
		}
	}

	// ------------------------------------ one in each stack
	if (len(stack1) > 1) || (len(stack2) > 1) {
		return answer, fmt.Errorf("invalid entry")
	}

	// ------------------------------result
	answer = append(answer, stack1[0], stack2[0])

	return answer, nil
}
