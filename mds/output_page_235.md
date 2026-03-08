229

matrix. For our (7,4) code,

H =

2

6
6
6
6
6
6
6
4

1 1 1 0

0 1 1 1

1 1 0 1
|

{z

}

1 0 0

0 1 0

0 0 1
|

{z

}

3

7
7
7
7
7
7
7
5

(6.57)

The parity check matrix thus has size ( N K ) N , and the result of multiplying this matrix with a received
word is a length- ( N K ) binary vector. If no digital channel errors occur we receive a codeword so that

b c = c then H b c = 0 . For example, the rst column of G , (1 ; 0 ; 0 ; 0 ; 1 ; 0 ; 1) , is a codeword. Simple
calculations show that multiplying this vector by H results in a length- ( N K ) zero-valued vector.

Exercise 6.30 (Solution on p. 258.)
Show that H c = 0 for all the columns of G . In other words, show that HG = 0 an ( N K ) K
matrix of zeroes. Does this property guarantee that all codewords also satisfy H c = 0 ?

When the received bits b c do not form a codeword, H b c does not equal zero, indicating the presence of one or
more errors induced by the digital channel. Because the presence of an error can be mathematically written
as b c = ( c e ) , with e a vector of binary values having a 1 in those positions where a bit error occurred
(Table 6.3). The single-bit error patterns in the left column yield the decoding sequences shown in the right column.

Exercise 6.31 (Solution on p. 258.)

Show that adding the error vector (1 ; 0 ; : : : ; 0) to a codeword ips the codeword's leading bit and
leaves the rest una ected.

Consequently, H b c = H ( c e ) = H e . Because the result of the product is a length- ( N K ) vector of binary
values, we can have 2 1 non-zero values that correspond to non-zero error patterns e . To perform our
channel decoding,

1. compute (conceptually at least) H b c ;
2. if this result is zero, no detectable or correctable error occurred;
3. if non-zero, consult a table of length- ( N K ) binary vectors to associate them with the minimal error
pattern that could have resulted in the non-zero result; then
4. add the error vector thus obtained to the received vector b c to correct the error (because ( c e e ) = c ).

5. Select the data bits from the corrected word to produce the received bit sequence
b
b ( n ) .

The phrase minimal in the third item raises the point that a double (or triple or quadruple : : : ) error
occurring during the transmission/reception of one codeword can create the same received word as a single-

bit error or no error in another codeword. For example, (1 ; 0 ; 0 ; 0 ; 1 ; 0 ; 1) and (0 ; 1 ; 0 ; 0 ; 1 ; 1 ; 1) are both
