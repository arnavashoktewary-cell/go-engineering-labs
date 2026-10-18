# Go Lab 01 — Counters and Registry

**Theme:** structs, pointers, method receivers, maps, error handling, and testing  
**Difficulty:** beginner → intermediate  
**Language:** Go (standard library only)  
**Suggested duration:** two focused lab sessions; speed is not graded.

## Your mission

Build a small in-memory **counter registry**: a component that could track requests, jobs, or events in a service. It is intentionally a library first, with a tiny command-line application under `cmd/demo/` that uses it.

This is a specification, **not a solved exercise**. The starter implementation contains TODOs. The provided autograder tests will fail until you implement the required behavior. Don't edit the grader to make it pass.

## Start here — Ubuntu / WSL

From this extracted directory:

```bash
cd src
go mod init example.com/lab01
go test -v -count=1 ./...
```

You create the `go.mod` file yourself: that's part of the lab. No `go install` and no PATH edits are necessary.

Edit `counter.go` and `registry.go`; write **at least three original tests** in `student_test.go`. To see your current autograder score:

```bash
cd ..                  # back to this directory
python3 grader/grade.py
```

Other useful commands (from `src/`):

```bash
go test -run '^TestCounter_' -v -count=1
go test -run '^TestRegistry_' -v -count=1
go test ./...
go vet ./...
go run ./cmd/demo        # meaningful once the library works
```

**Note:** You should expect red test output initially. If the module has not been initialized, the grader will ask you to run `go mod init`.

## Stage 1 — Counter (50 points)

Complete the functions/methods in `src/counter.go`.

```go
type Counter struct {
    Name  string
    Value int
}

func NewCounter(name string, initial int) (*Counter, error)
func (c *Counter) Increment()
func (c *Counter) Add(delta int) error
func (c Counter) Snapshot() Counter
```

Contract:

- `NewCounter` creates a counter with the requested name and initial value. Reject the empty string as a name and any negative initial value by returning a **non-nil error**. On failure, return a nil pointer.
- `Increment` increases the stored count by 1, *modifying the original*.
- `Add` adds a nonnegative amount. If `delta < 0`, return an error and **do not modify** the counter. Zero is allowed.
- `Snapshot` returns a separate **value copy** of the counter, so modifying the returned value does not affect the original.

Assume no arithmetic overflow and non-nil `*Counter` receivers for this lab.

## Stage 2 — Registry (50 points)

Complete the functions/methods in `src/registry.go`.

```go
type Registry struct {
    counters map[string]*Counter
}

func NewRegistry() *Registry
func (r *Registry) Register(name string, initial int) error
func (r *Registry) Increment(name string) error
func (r *Registry) Get(name string) (Counter, bool)
func (r *Registry) Names() []string
func (r *Registry) Total() int
```

Contract:

- `NewRegistry` returns a usable registry with an initialized map.
- `Register` adds one counter. **Duplicate names are errors**, even if the second initial value differs. Empty names and negative initial values are also errors. Failed registrations must not change the registry.
- `Increment` modifies the registered counter with that name. Unknown names cause an error without registering anything.
- `Get` returns a **copy** of the counter and `true` if found; otherwise returns `Counter{}` and `false`. Changing the returned copy must not modify the registry.
- `Names` returns a fresh slice containing every registered name in **ascending lexicographic order**. Never rely on map iteration order.
- `Total` returns the sum of all registered counter values (0 if none).

Assume methods are called on a registry returned by `NewRegistry`, and access is single-threaded. Concurrency comes in a later lab.

## Why the API looks this way

- `Counter` is a struct holding state.
- `Increment` / `Add` have pointer receivers because they must modify that state.
- `Snapshot` has a value receiver because it intentionally produces a copy.
- `Registry` holds pointers to counters for mutable shared state *inside the registry*, but `Get` deliberately exposes a copy to callers.
- `Names` sorts explicit keys because Go maps have unspecified iteration order.

## Autograder

The grader contains **10 public tests × 10 points each = 100 automated functional points**. It reports individual results and a total. Your own tests, readability, error-message quality, and engineering design are reviewed separately when you send your solution; the 100-point local score is not the entire code review.

Public tests are visible in `src/grader_test.go` so you can learn from failures. I can apply **additional edge-case tests** when you submit your work; those are not inside the ZIP.

Do not modify:

- `src/grader_test.go`
- `grader/grade.py`

## Your test-writing requirement

Write at least three additional tests in `src/student_test.go` without copying the grader's assertions verbatim. Suggested areas: negative input should not mutate state, a returned snapshot should be independent, and registering a duplicate should preserve the original entry.

Test naming format:

```go
func TestStudentSomeBehavior(t *testing.T) {
    // Arrange, act, assert
}
```

## Lab submission

Send me a ZIP of your filled-in `src` directory or a Git bundle. I will assess correctness, your self-written tests, receiver/pointer reasoning, algorithmic choices (especially sorting and map access), and style. Do not commit generated executables.

To make a Git bundle (once you've initialized a repository and committed your work):

```bash
git bundle create lab01-submission.bundle --all
```

## Learning checkpoints

Before writing much code, answer `questions/00_pointer_trace.md` **without running it**. We will go through your reasoning together. Then tackle Stage 1 before Stage 2. You are free to inspect the grader, but attempt the spec first.
