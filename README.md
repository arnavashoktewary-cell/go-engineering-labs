# go-engineering-labs
A collection of hands-on Go labs and mini-projects documenting my journey learning Go, DSA, and software engineering. Hopefully useful to fellow learners too!

Go Engineering Labs

A collection of hands-on labs and mini-projects as I learn Go, data structures and algorithms, and practical software engineering.
I'm learning by building: implementing features from specifications, writing tests, debugging failures, and improving the design as I go. I'm sharing the exercises and my implementations in case they're useful to other learners too. :)
What I'm working on
- Go fundamentals: structs, pointers, methods, maps, slices, interfaces, and errors.
- DSA through projects: choosing and implementing data structures to solve practical problems.
- Engineering habits: unit tests, Git, documentation, debugging, benchmarking, and code reviews.
- Systems and backend development: file I/O, concurrency, and HTTP services.
Labs
Lab	Project	Main topics	Status
01	Counter Registry	Structs, pointers, maps, slices, errors, testing	✅ Public tests: 100/100
02	Profile Directory	Struct initialization, pointer aliasing, maps	Planned
03–10	More Go + DSA mini-projects	Algorithms, systems, concurrency, APIs	Planned


The roadmap may evolve as I learn. Each lab is intended to be a small, self-contained project.
Running a lab
You'll need Go installed. From the repository root, for Lab 01:
cd lab01-counter-registry/src
go test -v ./...
To run the Lab 01 public autograder (if included in the lab folder):
cd ..
python3 grader/grade.py
The public autograder score reflects its provided test cases, not an exhaustive correctness guarantee. Individual labs may include their own setup instructions.
Learning approach
1. Read the requirements and plan the implementation.
2. Write the Go code and my own _test.go tests.
3. Run the tests, investigate failures, and refactor.
4. Record useful lessons and tradeoffs.
The goal is to become comfortable building and understanding real software—not just memorizing syntax or completing exercises.
For other learners
Feel free to browse the code, run the tests, and try solving the lab specifications yourself before looking at my implementations. Suggestions and constructive feedback are welcome.
