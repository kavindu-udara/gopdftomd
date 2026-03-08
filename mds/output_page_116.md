110 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.8: Fourier series approximation to sq ( t ) . The number of terms in the Fourier sum is indicated
in each plot, and the square wave is shown as a dashed line over two periods. much criticism from the French Academy of Science (Laplace, Lagrange, Monge and LaCroix
comprised the review committee) for several years after its presentation on 1807. It was not resolved for
almost a century, with ultimate solution interesting and important to understand.
Let's return to the question of equality; how can the equal sign in the de nition of the Fourier series be
justi ed? The partial answer is that point wise each and every value of t equality is not guaranteed.
However, mathematicians later in the nineteenth century showed that the rms error of the Fourier series was
always zero.
lim rms ( ) = 0

What this means is that the error between a signal and its Fourier series approximation may not be zero,
but that its rms value will be zero! It is through the eyes of the rms value that we rede ne equality: The
usual de nition of equality is called point wise equality : Two signals s ( t ) , s ( t ) are said to be equal
point wise if s ( t ) = s ( t ) for all values of t . A new de nition of equality is mean-square equality : Two
signals are said to be equal in the mean square if rms ( s s ) = 0 . For Fourier series, Gibbs' phenomenon
peaks have nite height and zero width. The error di ers from zero only at isolated points whenever the
periodic signal contains discontinuities and equals about 9% of the size of the discontinuity. The value of
a function at a nite set of points does not a ect its integral. This e ect underlies the reason why de ning
the value of a discontinuous function at its discontinuity, like when we de ned the step function on page 17,
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Fourier.html

![Im162](../output/images/test_116_Im162.R10.tif)
