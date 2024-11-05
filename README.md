# hamming

Hamming numbers, or ugly numbers, or smooth numbers, is a sequence of numbers that they can be produced by :

```
2^x * 3^y * 5^z
```

*Where x, y, z >= 0*

For example :
- 2, is a hamming number, since : `2^1 + 3^0 + 5^0 = 2`.
- 25, is a hamming number, since : `2^0 + 3^0 + 5^2 = 25`.

It is a sequence that follows a predictable, controllable pattern of incrementation due to their specific structure. It looks like that : 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, ...

They origin from punched card readers.

In 1950, Richard W. Hamming was working on Bell Labs (home of Unix any many other great inventions), where he got frustrated of the erorrs  of the punched card reader. Soon, he  started working on several error-correcting codes,  most notably the Hamming(7,4) code which is used in many aspects of todays life. A fruit of his labour was the Hamming sequence -or hamming numbers.

Famous they became by Edsger W. Dijkstra, a honorable computer scientist (famous for Dijkstra's algorithm).
In his book "Discipline of Programming", on the 17th chapter, named "AN EXERCISE ATTRIBUTED TO R.W. HAMMING", Dijkstra travels accross the beauty of hamming numbers.
This, made hamming numbers a hot topic on academia.

The beauty of those numbers is that they consist of the multiplication of 2, 3, and 5 raised to some powers (2^x * 3^y * 5^z). This means they are very easy to factor, simple to manipulate, and portable.

Factoring a number is when you split it down to other numbers, which if you multiply them ,you get that  number. For example factoring 100 can look like this :
- `100 / 2 = 50`, no remainder, prime factors : `[2]`
- `50 / 2 = 25`. no remainder, prime factors : `[2, 2]`
- `25 / 2 = 12.5`, has remainder, we try next number
- `25 / 3 = 8.33`, same
- `25 / 4 = 6.25`, same
- `25 / 5 = 5`, no remainder, prime factors : `[2, 2, 5]`
- `5 / 2 = 2.5`, has remainder, we try next number
- `5 / 3 = 1.66`, same
- `5 / 4 = 1.25`, same
- `5 / 5 = 1`, no remainder, prime factors : `[2, 2, 5, 5]`
- We reached 1 so we stop

And indeed `2 * 2 * 5 * 5 = 100`.

If a number can't be factored, then it's a prime number. It can be divided only by 1 and it's self.

100 is a hamming number as well, since we can symbolize it as, based on the above, `2^2 * 3^0 * 5^2`.

And if factorizing a number like 100 to `2 * 2 * 5 * 5` or `2^2 * 3^0 * 5^2` seems like making our lifes harder, think that the same we can do for numbers like  `262440000000`, which is `2^9 * 3^8 * 5^7`. It's now much easier to transmit over a bad network "9, 8, 7" than this huge number.

At this point, we can call hamming numbers as "easy to factor numbers", because they consist only of 3 prime numbers : 2, 3, 5.
There are numbers that can't be constructed of the numbers 2,3 and 5 and they require other prime numbers as well, like 7 : those are not hamming numbers and are harder to factor as well.

Hamming numbers, are "easy to divide numbers" as well. Being built only from small prime factors (2, 3, and 5), are more likely to be divisible by a range of commonly needed factors than a completely random number.

Many algorithms rely on factoring to reduce complexity. Providing numbers that are easy to factor, reduces complexity even more and can play a crusual role into making things faster or even possible.

**Using Hamming numbers, doesn't do the factorization, but guarantees that the factorization or de-factorization will be easy and fast. It also provides more easily divisible numbers, compared to using random numbers.**


## Metrics

The difference is crazy. Thanks to the code provided to this Git, which provides go functions to :
- find if x is hamming number
- find next hamming number after x
- generate a list of hamming numbers of length n
- factorize a given number x (util)
- messure the time a functio took to execute (high-order, util)

We can see the difference in time it takes to factorize a hamming number and a not-hamming number.

- 51200000 is a hamming number.
- 51200001 is NOT a hamming number.
- µs is microseconds
- ms is milliseconds (1000µs)

The result for my machine (AMD A4-5300, but any CPU would behave similarily) are :
```
> go build
> ./hamming
Time taken to factorize 51200000: 38.255µs
Time taken to factorize 51200001: 1.013663ms
```
