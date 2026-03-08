262 CHAPTER 7. APPENDIX : 1 10 : 5 Common values for the decibel. The decibel values for all but the powers of ten are
approximate, but are accurate to a decimal place.
input amplitude at that frequency. This calculation is one reason that we plot transfer function magnitude
on a logarithmic vertical scale expressed in decibels.

### 7.2 Permutations and Combinations

The lottery game consists of picking k numbers from a pool of n . For example, you select 6 numbers out
of 60 . To win, the order in which you pick the numbers doesn't matter; you only have to choose the right set
of 6 numbers. The chances of winning equal the number of di erent length- k sequences that can be chosen.
A related, but di erent, problem is selecting the batting lineup for a baseball team. Now the order matters,
and many more choices are possible than when order does not matter.
Answering such questions occurs in many applications beyond games. In digital communications, for
example, you might ask how many possible double-bit errors can occur in a codeword. Numbering the bit
positions from 1 to N , the answer is the same as the lottery problem with k = 6 . Solving these kind of
problems amounts to understanding permutations - the number of ways of choosing things when order
matters as in baseball lineups - and combinations - the number of ways of choosing things when order does
not matter as in lotteries and bit errors.
Calculating permutations is the easiest. If we are to pick k numbers from a pool of n , we have n choices
for the rst one. For the second choice, we have n 1 . The number of length-two ordered sequences is
therefore be n ( n 1) . Continuing to choose until we make k choices means the number of permutations

is n ( n 1) ( n 2) ( n k + 1) . This result can be written in terms of factorials as
n ! n k )!
, with n ! =

n ( n 1) ( n 2) 1 . For mathematical convenience, we de ne 0! = 1 .
When order does not matter, the number of combinations equals the number of permutations divided by
the number of orderings. The number of ways a pool of k things can be ordered equals k ! . Thus, once we
choose the nine starters for our baseball game, we have 9! = 362 ; 880 di erent lineups! The symbol for the

combination of k things drawn from a pool of n is
n

k
and equals
n ! n k )! k !
.

Exercise 7.2 (Solution on p. 265.)
What are the chances of winning the lottery? Assume you pick 6 numbers from the numbers 1 - 60 .

Combinatorials occur in interesting places. For example, Newton derived that the n -th power of a sum

obeyed the formula ( x + y ) =
n

0
x +
n

1
x y +
n

2
x y + +
n

n
y .
This content is available online at http://cnx.org/content/m10262/2.13/ .
