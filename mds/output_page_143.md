137

(b) Show the block diagram of a stereo AM receiver that can yield the left and right signals as separate
outputs.
(c) What signal would be produced by a conventional coherent AM receiver that expects to receive a
standard AM signal conveying a message signal having bandwidth W ?

Problem 4.25: Novel AM Stereo Method
A clever engineer has submitted a patent for a new method for transmitting two signals simultaneously
in the same transmission bandwidth as commercial AM radio. As shown in Figure 4.33, her approach is to
modulate the positive portion of the carrier with one signal and the negative portion with a second. In detail

Figure 4.33

the two message signals m ( t ) and m ( t ) are bandlimited to W Hz and have maximal amplitudes equal to
1. The carrier has a frequency f much greater than W . The transmitted signal x ( t ) is given by

x ( t ) =

(
A (1 + am ( t )) sin (2 f t ) if sin (2 f t ) 0

A (1 + am ( t )) sin (2 f t ) if sin (2 f t ) < 0
:

In all cases, 0 < a < 1 . The plot shows the transmitted signal when the messages are sinusoids:
m ( t ) = sin (2 f t ) and m ( t ) = sin (2 2 f t ) where 2 f < W . You, as the patent examiner, must de-
termine whether the scheme meets its claims and is useful.

(a) Provide a more concise expression for the transmitted signal x ( t ) than given above.
(b) What is the receiver for this scheme? It would yield both m ( t ) and m ( t ) from x ( t ) .
(c) Find the spectrum of the positive portion of the transmitted signal.
(d) Determine whether this scheme satis es the design criteria, allowing you to grant the patent. Explain
your reasoning.

Problem 4.26: A Radical Radio Idea
An ELEC 241 student has the bright idea of using a square wave instead of a sinusoid as an AM carrier.
The transmitted signal would have the form

x ( t ) = A 1 + m ( t ) sq ( t )

where the message signal m ( t ) would be amplitude-limited: j m ( t ) j < 1
