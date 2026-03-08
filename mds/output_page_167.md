161

we have K equations in N unknowns if we want to nd the signal from its sampled spectrum. This require-
ment is impossible to ful ll if K < N ; we must have K N . Our orthogonality relation essentially says that
if we have a su cient number of equations (frequency samples), the resulting set of equations can indeed be
solved.
By convention, the number of DFT frequency values K is chosen to equal the signal's duration N . The
discrete Fourier transform pair consists of

Discrete Fourier Transform Pair

S ( k ) =
X
s ( n ) e

s ( n ) =
1

X
S ( k ) e

(5.38)

Since the DFT is merely a sampled version of the DTFT, the table shown on page 159 applies to the DFT
as well, but with one notable exception: signals having even or odd symmetry. Because signals amenable to
the DFT are de ned only for n = 0 : : : ; N 1 , what does s ( n ) mean? Because the result of computing the
DFT, then the IDFT of the result, is a periodic version of the original signal, we consider the original signal
to be periodic with period N . In this case, an even signal would satisfy s ( n ) = s ( N n ) . Computing
the DFT of an even or odd signal would yield the result s ( n ) = s ( N n ) ! S ( N k ) . Also note
that using the table for real-valued signals, we would conclude that S ( k ) = S ( k ) . However, the DFT's
result is only de ned for positive k . Note that the DFT spectrum is periodic with period N , implying that
S ( k ) = S ( N k ) . Therefore, the DFT of all real-valued signals must have the property S ( k ) = S ( N k ) .

Exercise 5.16 (Solution on p. ?? )
What does this conjugate symmetry property say about S ( N= 2) when N , the transform length, is
even?

### 5.8 DFT: Computational Complexity

We now have a way of computing the spectrum for an arbitrary signal: The Discrete Fourier Transform
(DFT) (5.33) computes the spectrum at N equally spaced frequencies from a length- N sequence. An issue
that never arises in analog computation, like that performed by a circuit, is how much work it takes to
perform the signal processing operation such as ltering. In computation, this consideration translates to
the number of basic computational steps required to perform the needed processing. The number of steps,
known as the complexity , becomes equivalent to how long the computation takes (how long must we wait
for an answer). Complexity is not so much tied to speci c computers or programming languages but to how
many steps are required on any computer. Thus, a procedure's stated complexity says that the time taken
will be proportional to some function of the amount of data used in the computation and the amount
demanded.
For example, consider the formula for the discrete Fourier transform. For each frequency we choose, we
must multiply each signal value by a complex number and add together the results. For a real-valued signal,
each real-times-complex multiplication requires two real multiplications, meaning we have 2 N multiplications
to perform. To add the results together, we must keep the real and imaginary parts separate. Adding N
numbers requires N 1 additions. Consequently, each frequency requires 2 N + 2 ( N 1) = 4 N 2 basic
computational steps. As we have N frequencies, the total number of computations is N (4 N 2) .
In complexity calculations, we only worry about what happens as the data lengths increase, and take the
dominant term here the 4 N term as re ecting how much work is involved in making the computation.
As multiplicative constants don't matter since we are making a proportional to evaluation, we nd the
DFT is an O N computational procedure. This notation is read order N -squared. Thus, if we double
the length of the data, we would expect that the computation time to approximately quadruple.
This content is available online at http://cnx.org/content/m0503/2.11/ .
