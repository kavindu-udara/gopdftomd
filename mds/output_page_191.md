185

and we approximate the derivative by

d
y ( t )
y ( nT ) y (( n 1) T )

where T essentially amounts to a sampling interval.

(a) What is the di erence equation that must be solved to approximate the di erential equation?
(b) When x ( t ) = u( t ) , the unit step, what will be the simulated output?
(c) Assuming x ( t ) is a sinusoid, how should the sampling interval T be chosen so that the approximation
works well?

Problem 5.15: Derivatives
The derivative of a sequence makes little sense, but still, we can approximate it. The digital lter described
by the di erence equation
y ( n ) = x ( n ) x ( n 1)

resembles the derivative formula. We want to explore how well it works.

(a) What is this lter's transfer function?
(b) What is the lter's output to the depicted triangle input (Figure 5.31)?

Figure 5.31

(c) Suppose the signal x ( n ) is a sampled analog signal: x ( n ) = x ( nT ) . Under what conditions will the
lter act like a di erentiator? In other words, when will y ( n ) be proportional to
x ( t ) j ?

Problem 5.16: The DFT
Let's explore the DFT and its properties.

(a) What is the length- K DFT of length- N boxcar sequence, where N < K ?
(b) Consider the special case where K = 4 . Find the inverse DFT of the product of the DFTs of two
length-3 boxcars.
(c) If we could use DFTs to perform linear ltering, it should be true that the product of the input's DFT
and the unit-sample response's DFT equals the output's DFT. So that you can use what you just
calculated, let the input be a boxcar signal and the unit-sample response also be a boxcar. The result
of part (b) would then be the lter's output if we could implement the lter with length-4 DFTs. Does
the actual output of the boxcar- lter equal the result found in the previous part (list, p. 185)?
(d) What would you need to change so that the product of the DFTs of the input and unit-sample response
in this case equaled the DFT of the ltered output?
