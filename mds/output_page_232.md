226 CHAPTER 6. INFORMATION COMMUNICATION

transmitted again changes: So-called data bits b ( n ) emerge from the source coder at an average rate ( A )
and exit the channel at a rate 1 =E higher. We represent the fact that the bits sent through the digital
channel operate at a di erent rate by using the index l for the channel-coded bit stream c ( l ) . Note that
the blocking (framing) imposed by the channel coder does not correspond to symbol boundaries in the bit
stream b ( n ) , especially when we employ variable-length source codes.
Does any error-correcting code reduce communication errors when real-world constraints are taken into
account? The answer now is yes. To understand channel coding, we need to develop rst a general framework
for channel coding, and discover what it takes for a code to be maximally e cient: Correct as many errors
as possible using the fewest error correction bits as possible (making the e ciency K=N as large as possible).

### 6.27 Error-Correcting Codes: Hamming Distance

So-called linear codes create error-correction bits by combining the data bits linearly. The phrase linear
combination means here single-bit binary arithmetic.

0 0 = 0 1 1 = 0 0 1 = 1 1 0 = 1

0 0 = 0 1 1 = 1 0 1 = 0 1 0 = 0

For example, let's consider the speci c (3,1) error correction code described by the following coding table
and, more concisely, by the succeeding matrix expression.

c (1) = b (1)

c (2) = b (1)

c (3) = b (1)

or

c = G b

where

G =

2

6
6
4

1

1

1

3

7
7
5
c =

2

6
6
4

c (1)

c (2)

c (3)

3

7
7
5
b =
h
b (1)
i

The length- K (in this simple example K = 1 ) block of data bits is represented by the vector b , and the
length- N output block of the channel coder, known as a codeword , by c . The generator matrix G de nes
all block-oriented linear channel coders.
As we consider other block codes, the simple idea of the decoder taking a majority vote of the received bits
won't generalize easily. We need a broader view that takes into account the distance between codewords.
A length- N codeword means that the receiver must decide among the 2 possible data words to select
which of the 2 codewords was actually transmitted. As shown in Figure 6.22, we can think of the data
words geometrically. We de ne the Hamming distance between binary data words c and c , denoted by
d ( c ; c ) to be the minimum number of bits that must be ipped to go from one word to the other. For
example, the distance between codewords is 3 bits. In our table of binary arithmetic, we see that adding
a 1 corresponds to ipping a bit. Furthermore, subtraction and addition are equivalent. We can express the
Hamming distance as
d ( c ; c ) = sum (( c c )) (6.55)

Exercise 6.28 (Solution on p. 258.)
Show that adding the error vector col [1 ; 0 ; : : : ; 0] to a codeword ips the codeword's leading bit and
leaves the rest una ected.
