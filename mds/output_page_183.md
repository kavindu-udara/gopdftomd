177

Figure 5.22: The noisy input signal is sectioned into length-48 frames, each of which is ltered using
frequency-domain techniques. Each ltered section is added to other outputs that overlap to create the
signal equivalent to having ltered the entire input. The sinusoidal component of the signal is shown as
the red dashed line.
each term.

x ( n ) =
X
x ( n mN ) ) y ( n ) =
X
y ( n mN ) (5.55)

As illustrated in Figure 5.22, note that each ltered section has a duration longer than the input. Conse-
quently, we must literally add the ltered sections together, not just butt them together.
Computational considerations reveal a substantial advantage for a frequency-domain implementation
over a time-domain one. The number of computations for a time-domain implementation essentially remains
constant whether we section the input or not. Thus, the number of computations for each output is 2 q + 1 .
In the frequency-domain approach, computation counting changes because we need only compute the lter's
frequency response H ( k ) once, which amounts to a xed overhead. We need only compute two DFTs and
multiply them to lter a section. Letting N denote a section's length, the number of computations for
a section amounts to ( N + q ) log ( N + q ) + 6 ( N + q ) . In addition, we must add the ltered outputs
together; the number of terms to add corresponds to the excess duration of the output compared with the

input ( q ). The frequency-domain approach thus requires log ( N + q )+6+
q + q
computations per output

value. For even modest lter orders, the frequency-domain approach is much faster.

Exercise 5.32 (Solution on p. 193.)
Show that as the section length increases, the frequency domain approach becomes increasingly
more e cient.

Note that the choice of section duration is arbitrary. Once the lter is chosen, we should section so that the

![Im224](../output/images/test_183_Im224.R8.tif)
