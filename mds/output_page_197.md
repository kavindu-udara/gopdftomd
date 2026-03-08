191

### Solutions to Exercises in Chapter 5

Solution to Exercise 5.1 (p. 145)
For b -bit signed integers, the largest number is 2 1 . For b = 32 , we have 2,147,483,647 and for b = 64 ,
we have 9,223,372,036,854,775,807 or about 9 : 2 10 .
Solution to Exercise 5.2 (p. 146)
In oating point, the number of bits in the exponent determines the largest and smallest representable
numbers. For 32-bit oating point, the largest (smallest) numbers are 2 = 1 : 7 10 ( 5 : 9 10 ). For
64-bit oating point, the largest number is about 10 .
Solution to Exercise 5.3 (p. 146)
25 = 11011 and 7 = 111 . We nd that 11001 + 111 = 100000 = 32 .
Solution to Exercise 5.4 (p. 148)
The only e ect of pulse duration is to unequally weight the spectral repetitions. Because we are only
concerned with the repetition centered about the origin, the pulse duration has no signi cant e ect on
recovering a signal from its samples.
Solution to Exercise 5.5 (p. 149)

Figure 5.33

The square wave's spectrum is shown by the bolder set of lines centered about the origin. The dashed
lines correspond to the frequencies about which the spectral repetitions (due to sampling with T = 1 )
occur. As the square wave's period decreases, the negative frequency lines move to the left and the positive
frequency ones to the right.
Solution to Exercise 5.6 (p. 149)
The simplest bandlimited signal is the sine wave. At the Nyquist frequency, exactly two samples/period
would occur. Reducing the sampling rate would result in fewer samples/period, and these samples would
appear to have arisen from a lower frequency sinusoid.
Solution to Exercise 5.7 (p. 150)
The plotted temperatures were quantized to the nearest degree. Thus, the high temperature's amplitude
was quantized as a form of A/D conversion.
Solution to Exercise 5.8 (p. 151)
The signal-to-noise ratio does not depend on the signal amplitude. With an A/D range of [ A; A ] , the

quantization interval =
2 A
and the signal's rms value (again assuming it is a sinusoid) is A=
p .

Solution to Exercise 5.9 (p. 151)
Solving 2 = : 001 results in B = 10 bits.
