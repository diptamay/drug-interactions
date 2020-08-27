# Drug Interactions Problem

Given a file containing a list of drug-drug interactions, write a command line program that accepts a space-separated list of 
drugs per line and determines if there is a risk of interaction between any drugs in the list.
  
If there are multiple interactions detected for a single line of input, the program should return the most severe interaction. 
If there are multiple interactions of the same severity, the program should return the interaction that appears first in the 
interactions.json file.
The program should read its input from STDIN and write its output to STDOUT, where each line of input should generate a line 
of output in the same order.

## Examples
Note: all examples are based on the interactions.json file included with this description.

### Input (each line represents one line of input)
sildenafil tamsulosin valaciclovir

sildenafil ibuprofen

valaciclovir doxepin ticlopidine ibuprofen

### Output (order should match the input above)
MODERATE: Sildenafil may potentiate the hypotensive effect of alpha blockers, resulting in symptomatic hypotension in some patients.

No interaction

MAJOR: Valaciclovir may decrease the excretion rate of Doxepin which could result in a higher serum level.`

## Constraints
Number of medications per line between 1 and 20 Number of lines per execution between 1 and 10,000

## Assumptions/Areas of Improvement

1. `interactions.json` is not a big file and the size is more in the realm of KBs and low MBs. 
I would have to revisit how I am reading the file and building the in-memory data structure for bigger files.

2. `drugs` array in `interactions.json` is always a list of 2 drugs. If more than 2 drugs are provided then the approach 
to building the keys for the in-memory map needs to be revisited.

3. `interactions.json` is in the same folder as the source code. 

4. Code coverage is at 77.% now and can be improved.

## Run

`go run drug-interaction-store.go`

## Tests

`go test drug-interaction-store.go drug-interaction-store_test.go`

_Note: All tests are based on the `interactions.json` file included with the exercise._
