168 CHAPTER 5. DIGITAL SIGNAL PROCESSING

### 5.12 Discrete-Time Systems in the Time-Domain

A discrete-time signal s ( n ) is delayed by n samples when we write s ( n n ) , with n > 0 . Choosing n
to be negative advances the signal along the integers. As opposed to analog delays, discrete-time delays can
only be integer valued. In the frequency domain, delaying a signal corresponds to a linear phase shift of the
signal's discrete-time Fourier transform: s ( n n ) ! e S e .
Linear discrete-time systems have the superposition property.

S [ a x ( n ) + a x ( n )] = a S [ x ( n )] + a S [ x ( n )] (5.40)

A discrete-time system is called shift-invariant (analogous to time-invariant analog systems (p. 24)) if
delaying the input delays the corresponding output. If S [ x ( n )] = y ( n ) , then a shift-invariant system has
the property
S [ x ( n n )] = y ( n n ) (5.41)

We use the term shift-invariant to emphasize that delays can only have integer values in discrete-time, while
in analog signals, delays can be arbitrarily valued.
We want to concentrate on systems that are both linear and shift-invariant. It will be these that allow us
the full power of frequency-domain analysis and implementations. Because we have no physical constraints
in constructing such systems, we need only a mathematical speci cation. In analog systems, the di eren-
tial equation speci es the input-output relationship in the time-domain. The corresponding discrete-time
speci cation is the di erence equation .

y ( n ) = a y ( n 1) + + a y ( n p ) + b x ( n ) + b x ( n 1) + + b x ( n q ) (5.42)

Here, the output signal y ( n ) is related to its past values y ( n l ) , l = f 1 ; : : : ; p g , and to the current and past
values of the input signal x ( n ) . The system's characteristics are determined by the choices for the number
of coe cients p and q and the coe cients' values f a ; : : : ; a g and f b ; b ; : : : ; b g .

aside: There is an asymmetry in the coe cients: where is a ? This coe cient would multiply the
y ( n ) term in (5.42). We have essentially divided the equation by it, which does not change the
input-output relationship. We have thus created the convention that a is always one.

As opposed to di erential equations, which only provide an implicit description of a system (we must
somehow solve the di erential equation), di erence equations provide an explicit way of computing the
output for any input. We simply express the di erence equation by a program that calculates each output
from the previous output values, and the current and previous inputs.
Di erence equations are usually expressed in software with for loops. A MATLAB program that would
compute the rst 1000 values of the output has the form

for n=1:1000
y(n) = sum(a.*y(n-1:-1:n-p)) + sum(b.*x(n:-1:n-q));
end

An important detail emerges when we consider making this program work; in fact, as written it has (at
least) two bugs. What input and output values enter into the computation of y (1) ? We need values for
y (0) ; y ( 1) ; : : : values we have not yet computed. To compute them, we would need more previous values of
the output, which we have not yet computed. To compute these values, we would need even earlier values, ad
in nitum. The way out of this predicament is to specify the system's initial conditions : we must provide
the p output values that occurred before the input started. These values can be arbitrary, but the choice
does impact how the system responds to a given input. One choice gives rise to a linear system: Make the
initial conditions zero. The reason lies in the de nition of a linear system: The only way that the output to
This content is available online at http://cnx.org/content/m10251/2.23/ .
