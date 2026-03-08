231

Figure 6.23: The probability of an error occurring in transmitted K = 4 data bits equals 1 (1 p )
as (1 p ) equals the probability that the four bits are received without error. The upper curve displays
how this probability of an error anywhere in the four-bit block varies with the signal-to-noise ratio. When
a (7,4) single-bit error correcting code is used, the transmitter reduced the energy it expends during a
single-bit transmission by 4/7, appending three extra bits for error correction. Now the probability of
any bit in the seven -bit block being in error after error correction equals 1 (1 p ) 7 p (1 p ) ,
where p is the probability of a bit error occurring in the channel when channel coding occurs. Here
(7 p ) (1 p ) equals the probability of exactly on in seven bits emerging from the channel in error; The
channel decoder corrects this type of error, and all data bits in the block are received correctly. (Solution on p. 259.)
What must the relation between N and K be for a code to correct all single- and double-bit errors
with a perfect t?

### 6.30 Noisy Channel Coding Theorem

As the block length becomes larger, more error correction will be needed. Do codes exist that can correct
all errors? Perhaps the crowning achievement of Claude Shannon's creation of information theory answers
this question. His result comes in two complementary forms: the Noisy Channel Coding Theorem and its
converse.

6.30.1 Statement of the Theorem

Let E denote the e ciency of an error-correcting code: the ratio of the number of data bits to the total
number of bits used to represent them. If the e ciency is less than the capacity of the digital channel, an
error-correcting code exists that has the property that as the length of the code increases, the probability of
an error occurring in the decoded block approaches zero.

lim Pr [ block error ] = 0 ; E < C (6.58)

6.30.2 Converse to the Theorem

If E > C , the probability of an error in a decoded block must approach one regardless of the code that might
be chosen.
lim Pr [ block error ] = 1 (6.59)
This content is available online at http://cnx.org/content/m0073/2.12/ .

![Im258](../output/images/test_237_Im258.R8.tif)
