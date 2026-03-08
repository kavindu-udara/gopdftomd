214 CHAPTER 6. INFORMATION COMMUNICATION

The receiver for digital communication is known as a matched lter .

Figure 6.14: The optimal receiver structure for digital communication faced with additive white noise
channels is the depicted matched lter.

This receiver multiplies the received signal by each of the possible members of the transmitter signal set,
integrates the product over the bit interval, and compares the results. Whichever path through the receiver
yields the largest value corresponds to the receiver's decision as to what bit was sent during the previous bit
interval. For the next bit interval, the multiplication and integration begins again, with the next bit decision
made at the end of the bit interval. Mathematically, the received value of b ( n ) , which we label
b
b ( n ) , is given
by

b
b ( n ) = arg max

Z

r ( t ) s ( t ) dt (6.42)

You may not have seen the arg max notation before. max f g yields the maximum value of its argument
with respect to the index i . arg max equals the value of the index that yields the maximum. Note that the
precise numerical value of the integrator's output does not matter; what does matter is its value relative to
the other integrator's output.
Let's assume a perfect channel for the moment: The received signal equals the transmitted one. If bit 0
were sent using the baseband BPSK signal set, the integrator outputs would be

Z

r ( t ) s ( t ) dt = A T

Z

r ( t ) s ( t ) dt = A T

(6.43)

If bit 1 were sent,
Z

r ( t ) s ( t ) dt = A T

Z

r ( t ) s ( t ) dt = A T

(6.44)

Exercise 6.17 (Solution on p. 256.)
Can you develop a receiver for BPSK signal sets that requires only one multiplier-integrator com-
bination?

Exercise 6.18 (Solution on p. 256.)
What is the corresponding result when the amplitude-modulated BPSK signal set is used?

Clearly, this receiver would always choose the bit correctly. Channel attenuation would not a ect this
correctness; it would only make the values smaller, but all that matters is which is largest.
