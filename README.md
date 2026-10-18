# Go Engineering Labs

Hey! This repo is basically where I'll be putting all the labs and mini-projects I'm doing while learning **Go (Golang)** and DSA.

I'm a CSE undergrad trying to get better at programming, especially the practical side of things. I've worked with C++ before, but Go is pretty new to me.

Instead of just watching tutorials or solving random problems, I thought it'd be more fun to actually build stuff and figure things out along the way.

I'm keeping everything here to track my progress, and hopefully some of these exercises might help other people learning Go too :)

## What's the plan?

I'm starting with the basics and slowly moving towards more complicated projects.

- **Go fundamentals** — structs, pointers, maps, slices, interfaces, error handling, etc.
- **DSA** — learning and implementing algorithms through actual problems and projects.
- **Testing** — writing my own tests, using autograders, and figuring out why things break.
- **Systems and backend stuff** — file handling, concurrency, networking, APIs, and whatever else I end up exploring.

The eventual goal is to get comfortable enough with Go to start contributing to open-source projects.

## Labs

| Lab | Project | Topics | Status |
|---|---|---|---|
| 01 | Counter Registry | Structs, pointers, maps, slices, errors | 100/100 ✅ |
| 02 | Profile Directory | Structs, pointers, maps (more practice!) | Up next |
| 03–10 | TBD | DSA, algorithms, backend and systems | Planned |

I'm aiming to finish around 10 labs before jumping into actual open-source repositories. The plan will probably change a bit as I go.

## Running the labs

You'll need [Go](https://go.dev/dl/) installed.

Each lab has its own folder with the source code, tests, and instructions.

For example, to run the tests for Lab 01:

```bash
cd lab01-counter-registry/src
go test -v ./...
```

To run the provided autograder:

```bash
cd ..
python3 grader/grade.py
```

Feel free to try the problems yourself before looking at my implementations!

## How I'm learning

Nothing too fancy. I read the documentation, try implementing something, inevitably get stuck somewhere, debug it, and try again.

Each lab is meant to introduce a few concepts while also making me use the stuff I've already learned.

I'll also try to document anything interesting I learn, especially things that aren't immediately obvious from the code.

## That's about it!

Also, full disclosure: the lab specifications and autograders are AI-generated. The implementations are my own. Just wanted to be transparent about that!

This is very much a learning repo, so don't expect everything to be perfect (especially the earlier labs).

If you find something useful here, great! And if you spot something stupid in my code, feel free to point it out. 

Hopefully I'll be able to look back at this repo in a few months and see how far I've come.

Happy coding! :)
