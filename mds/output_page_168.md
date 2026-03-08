162 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Exercise 5.17 (Solution on p. 192.)
In making the complexity evaluation for the DFT, we assumed the data to be real. Three questions
emerge. First of all, the spectra of such signals have conjugate symmetry, meaning that negative

frequency components ( k =
N
+ 1 ; : : : ; N + 1 in the DFT (5.33)) can be computed from the

corresponding positive frequency components. Does this symmetry change the DFT's complexity?
Secondly, suppose the data are complex-valued; what is the DFT's complexity now? Finally, a less
important but interesting question is suppose we want K frequency values instead of N ; now what
is the complexity?

### 5.9 Fast Fourier Transform (FFT)

One wonders if the DFT can be computed faster: Does another computational procedure an algorithm
exist that can compute the same quantity, but more e ciently. We could seek methods that reduce the
constant of proportionality, but do not change the DFT's complexity O N . Here, we have something
more dramatic in mind: Can the computations be restructured so that a smaller complexity results?
In 1965, IBM researcher James Cooley and Princeton faculty member John Tukey developed what is
now known as the Fast Fourier Transform (FFT). It is an algorithm for computing that DFT that has order
O ( N log N ) for certain length inputs . Now when the length of data doubles, the spectral computational
time will not quadruple as with the DFT algorithm; instead, it approximately doubles. Later research showed
that no algorithm for computing the DFT could have a smaller complexity than the FFT. Surprisingly,
historical work has shown that Gauss in the early nineteenth century developed the same algorithm, but
did not publish it! After the FFT's rediscovery, not only was the computation of a signal's spectrum greatly
speeded, but also the added feature of algorithm meant that computations had exibility not available to
analog implementations.

Exercise 5.18 (Solution on p. 192.)
Before developing the FFT, let's try to appreciate the algorithm's impact. Suppose a short-length
transform takes 1 ms. We want to calculate a transform of a signal that is 10 times longer. Compare
how much longer a straightforward implementation of the DFT would take in comparison to an
FFT, both of which compute exactly the same quantity.

To derive the FFT, we assume that the signal's duration is a power of two: N = 2 . Consider what happens
to the even-numbered and odd-numbered elements of the sequence in the DFT calculation.

S ( k ) = s (0) + s (2) e
+ + s ( N 2) e

+ s (1) e
+ s (3) e
+ + s ( N 1) e

=

"

s (0) + s (2) e
+ + s ( N 2) e

#

+

"

s (1) + s (3) e
+ + s ( N 1) e

#

e

(5.39)

Each term in square brackets has the form of a
-length DFT. The rst one is a DFT of the even-numbered
elements, and the second of the odd-numbered elements. The rst DFT is combined with the second

multiplied by the complex exponential e
. The half-length transforms are each evaluated at frequency
indices k = 0 , : : : , N 1 . Normally, the number of frequency indices in a DFT calculation range between zero
and the transform length minus one. The computational advantage of the FFT comes from recognizing
This content is available online at http://cnx.org/content/m10250/2.17/ .
https://en.wikipedia.org/wiki/James_Cooley
https://en.wikipedia.org/wiki/John_Tukey
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Gauss.html
