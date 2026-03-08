174 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.20: To lter a signal in the frequency domain, rst compute the DFT of the input, multiply
the result by the sampled frequency response, and nally compute the inverse DFT of the product. The
DFT's length must be at least the sum of the input's and unit-sample response's duration minus one.
We calculate these discrete Fourier transforms using the fast Fourier transform algorithm, of course.

Before detailing this procedure, let's clarify why so many new issues arose in trying to develop a frequency-
domain implementation of linear ltering. The frequency-domain relationship between a lter's input and
output is always true: Y e = H e X e . The Fourier transforms in this result are discrete-

time Fourier transforms; for example, X e =
P
x ( n ) e . Unfortunately, using this relationship
to perform ltering is restricted to the situation when we have analytic formulas for the frequency response
and the input signal. The reason why we had to invent the discrete Fourier transform (DFT) has the
same origin: The spectrum resulting from the discrete-time Fourier transform depends on the continuous
frequency variable f . That's ne for analytic calculation, but computationally we would have to make an
uncountably in nite number of computations.

note: Did you know that two kinds of in nities can be meaningfully de ned? A countably
in nite quantity means that it can be associated with a limiting process associated with integers.
An uncountably in nite quantity cannot be so associated. The number of rational numbers is
countably in nite (the numerator and denominator correspond to locating the rational by row and
column; the total number so-located can be counted, voila!); the number of irrational numbers is
uncountably in nite. Guess which is bigger?

The DFT computes the Fourier transform at a nite set of frequencies samples the true spectrum
which can lead to aliasing in the time-domain unless we sample su ciently fast. The sampling interval here

is
1
for a length- K DFT: faster sampling to avoid aliasing thus requires a longer transform calculation.

Since the longest signal among the input, unit-sample response and output is the output, it is that signal's
duration that determines the transform length. We simply extend the other two signals with zeros (zero-pad)
to compute their DFTs.

Example 5.7
Suppose we want to average daily stock prices taken over last year to yield a running weekly
average (average over ve trading sessions). The lter we want is a length-5 averager (as shown in
Figure 5.19), and the input's duration is 253 (365 calendar days minus weekend days and holidays).
The output duration will be 253 + 5 1 = 257 , and this determines the transform length we need
to use. Because we want to use the FFT, we are restricted to power-of-two transform lengths. We
need to choose any FFT length that exceeds the required DFT length. As it turns out, 256 is a
power of two ( 2 = 256 ), and this length just undershoots our required length. To use frequency
domain techniques, we must use length-512 fast Fourier transforms.
Figure 5.21 shows the input and the ltered output. The MATLAB programs that compute the
ltered output in the time and frequency domains are

Time Domain
h = [1 1 1 1 1]/5;
y = filter(h,1,[djia zeros(1,4)]);
