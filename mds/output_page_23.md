17

Figure 2.3: The real exponential.

The unit step function (Figure 2.4) is denoted by u( t ) , and is de ned to be

u( t ) =

(
0 t < 0

1 t > 0
(2.23)

Figure 2.4: The unit step.

Origin warning: This signal is discontinuous at the origin. Its value at the origin need not be
de ned because the value doesn't matter in signal theory.

This kind of signal is used to describe signals that turn on suddenly. For example, to mathematically repre-
sent turning on an oscillator, we can write it as the product of a sinusoid and a step: s ( t ) = A sin (2 ft ) u( t ) .

2.2.5 Pulse

The unit pulse (Figure 2.5) describes turning a unit-amplitude signal on for a duration of seconds, then
turning it o .

p ( t ) =

8
>
<

>
:

0 ; t < 0

1 ; 0 < t <

0 ; t >

(2.24)

Figure 2.5: The pulse.

We will nd that this is the second most important signal in communications.
