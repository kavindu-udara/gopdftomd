115

Figure 4.11: The upper plot shows the magnitude of the Fourier series spectrum for the case of T = 1
with the Fourier transform of p ( t ) shown as a dashed line. For the bottom panel, we expanded the period
to T = 5 , keeping the pulse's duration xed at 0.2, and computed its Fourier series coe cients. sinc (pronounced sink ) function, and is denoted by sinc ( t ) . Thus, the magnitude of the pulse's Fourier
transform equals j sinc ( f ) j .
The Fourier transform relates a signal's time and frequency domain representations to each other. The
direct Fourier transform (or simply the Fourier transform) calculates a signal's frequency domain repre-
sentation from its time-domain variant (4.34). The inverse Fourier transform (4.35) nds the time-domain
representation from the frequency domain. Rather than explicitly writing the required integral, we often
symbolically express these transform calculations as F ( s ) and F ( S ) , respectively.

F ( s ) = S ( f )

=

Z

s ( t ) e dt
(4.34)

F ( S ) = s ( t )

=

Z

S ( f ) e df
(4.35)

We must have s ( t ) = F ( F ( s ( t ))) and S ( f ) = F F ( S ( f )) , and these results are indeed valid with
minor exceptions.

note: Recall that the Fourier series for a square wave gives a value for the signal at the dis-
continuities equal to the average value of the jump. This value may di er from how the signal is
de ned in the time domain, but being unequal at a point is indeed minor.

Showing that you get back to where you started is di cult from an analytic viewpoint, and we won't try
here. Note that the direct and inverse transforms di er only in the sign of the exponent.

Exercise 4.12 (Solution on p. 140.)
The di ering exponent signs means that some curious results occur when we use the wrong sign.
What is F ( S ( f )) ? In other words, use the wrong exponent sign in evaluating the inverse Fourier
transform.

![Im166](../output/images/test_121_Im166.R10.tif)
