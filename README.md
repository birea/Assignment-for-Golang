                
                
                
              Volume.Finance Take-Home Programming Assignment for GoLang



Story: There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

Goal: To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Required JSON structure: 
●[["SFO", "EWR"]]        => ["SFO", "EWR"]
●[["ATL", "EWR"], ["SFO", "ATL"]]      => ["SFO", "EWR"]
●[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

Specifications: 

●Create a private GitHub repo and add https://github.com/taariq, https://github.com/verabehr, https://github.com/measure-fi  and https://github.com/Vizualni as collaborators to the project.


●Define and document the format of the API endpoint in the README.

1:Curl to test on insomnia

curl --request GET \
  --url http://127.0.0.1:8080/flight_points \
  --header 'Content-Type: application/json' \
  --data '{"paths":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] }'
2:Requests being made to '/flight_points'

{"paths":[["SFO", "EWR"]]}
output:
["SFO","EWR"]

{"paths":[["ATL", "EWR"], ["SFO", "ATL"]] }
output:
["SFO","EWR"]

{"paths":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] }
output:
["SFO","EWR"]


●Use Golang and/or any tools that you think will help you best accomplish the task at hand.

1:Markdown-for readme-rofl
2:MakeFile
3:CI , github actions
4:Lint
5:Unit Test

●When you are done with the assignment and follow up via email with an estimate of how long you spent on the task and any interesting ideas you wish to share.

1:Time spent
  3 Hours (yesterday evening)
2:I have a problem similar to this
  I have finished for 1 hour

There are 3 types of tiles, as shown in fig.1.
You are going to paste them on fig.2 to create a cube.
There are 'A' - 'F' positions (see fig. 2). You will paste tiles onto each position from 'A' to 'F'.
We can then express the cube's design as a series of numbers.
For example, We can express fig.3 as 123212 where each number corresponds to a tile.

Write a function 'trace(tiles)' that receives a series of tiles (represented as the numbers from fig. 1)
and returns the series of positions(alphabet from fig. 2) that you will pass if you trace the line from the left side of A and 
continue until looping back to the same position.

For example, in case of fig.3, the cube's design is "123212".
So if you start tracing line from the left side of A, You will pass
A => B => C => A => D => B => E => F => C => E => D => F => (LOOP HERE) A => B ...
So input is "123212" and output is "ABCADBEFCEDF"
Please refer fig.4. fig.5.

Other examples include

"123123" => "ABCADEFDBECF"
"111111" => "ABEF"
"222222" => "ACFEDF"
"333333" => "ADFECF"




