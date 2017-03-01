* To run the program => go run main.go

USAGE

	* You will see an interactive shell when you run the program. At the top, you will see the org hierarchy at different levels. This is so that you get an idea of what your org looks like.

	* You can add people to the organization by selecting "1". Once you add a person, you can see them in the org chart.

	* You can check the nearest reporting manager for any two employees by selecting "2". You will be prompted for the names of the two employees.

	* You can exit the shell by typing "e" in the main menu.

	* If at any time you are stuck, press enter and you will be redirected to something sensible.


TESTS

	* Have a look at /models/org_test.go and /models/queue_test.go. In the first one, I have provided test cases for different scenarios with a sample org hierarchy. In the second one, I'm testing the queue data structure that I am using for my level order traversal.

	* To execute the tests, go to the root of the project and execute => go test ./...

	* Feel free to experiment with the various test cases.

QUESTIONS/QUERIES

	* I've provided comments for the major functions. Please have a look at those for clarifications on what algorithm(s) I'm using.

	* You can reach me at bala.theripper@gmail.com

