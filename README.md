# GoPrime

GoPrime is a simple, single-threaded program written in Go, that accepts an integer as the user input and performs prime factorization on it.

![Screenshot](https://i.imgur.com/H3bC4VW.png)

# Why?

I recently finished learning Go fundamentals. Prime numbers are cool, so I decided to build this tool. It might be able to help some people learning math! :)

# Technical stuff
This was my first attempt at writing a proper Go program. There's certainly more room for enhancements/optimization:
 - More tests could be written
 - factors = append(factors, i) is resource intensive, because it allocates a new underlying array every time a factor is appended
 - performPrimeFactorization() uses fmt.Print(), so if I had created a test for it, go test would output all the factorization steps during the test. Is this okay?
 - I forgot to document my code

I might update this project from time to time with enhancements and stuff. Meanwhile, let me know your feedback! :)
