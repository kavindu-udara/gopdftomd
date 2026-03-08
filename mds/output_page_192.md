186 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Problem 5.17: DSP Tricks
Sammy is faced with computing lots of discrete Fourier transforms. He will, of course, use the FFT algorithm,
but he is behind schedule and needs to get his results as quickly as possible. He gets the idea of computing
two transforms at one time by computing the transform of s ( n ) = s ( n ) + js ( n ) , where s ( n ) and s ( n )
are two real-valued signals of which he needs to compute the spectra. The issue is whether he can retrieve
the individual DFTs from the result or not.

(a) What will be the DFT S ( k ) of this complex-valued signal in terms of S ( k ) and S ( k ) , the DFTs of
the original signals?
(b) Sammy's friend, an Aggie who knows some signal processing, says that retrieving the wanted DFTs is
easy: Just nd the real and imaginary parts of S ( k ) . Show that this approach is too simplistic.
(c) While his friend's idea is not correct, it does give him an idea. What approach will work? Hint : Use
the symmetry properties of the DFT.
(d) How does the number of computations change with this approach? Will Sammy's idea ultimately lead
to a faster computation of the required DFTs?

Problem 5.18: Discrete Cosine Transform (DCT)
The discrete cosine transform of a length- N sequence is de ned to be

S ( k ) =
X
s ( n ) cos
2 nk N

Note that the number of frequency terms is 2 N 1 : k = f 0 ; : : : ; 2 N 1 g .

(a) Find the inverse DCT.
(b) Does a Parseval's Theorem hold for the DCT?
(c) You choose to transmit information about the signal s ( n ) according to the DCT coe cients. You
could only send one, which one would you send?

Problem 5.19: A Digital Filter
A digital lter is described by the following di erence equation:

y ( n ) = ay ( n 1) + ax ( n ) x ( n 1) , a =
1

(a) What is this lter's unit sample response?
(b) What is this lter's transfer function?
(c) What is this lter's output when the input is sin
?

Problem 5.20: Another Digital Filter
A digital lter is determined by the following di erence equation.

y ( n ) = y ( n 1) + x ( n ) x ( n 4)

(a) Find this lter's unit sample response.
(b) What is the lter's transfer function? How would you characterize this lter (lowpass, highpass, special
purpose, something else)?
(c) Find the lter's output when the input is the sinusoid sin
.
(d) In another case, the input sequence is zero for n < 0 , then becomes nonzero. Sammy measures the
output to be y ( n ) = ( n ) + ( n 1) . Can his measurement be correct? In other words, is there an
input that can yield this output? If so, nd the input x ( n ) that gives rise to this output. If not, why
not?
