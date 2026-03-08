57

Figure 3.29
complex exponentials, each of which has a frequency di erent from the others. The transfer function portrays
how the circuit a ects the amplitude and phase of each component, allowing us to understand how the circuit
works on a complicated signal. Those components having a frequency less than the cuto frequency pass
through the circuit with little modi cation while those having higher frequencies are suppressed. The circuit
is said to act as a lter , ltering the source signal based on the frequency of each component complex
exponential. Because low frequencies pass through the lter, we call it a lowpass lter to express more
precisely its function.
We have also found the ease of calculating the output for sinusoidal inputs through the use of the transfer
function. Once we nd the transfer function, we can write the output directly as indicated by the output of
a circuit for a sinusoidal input (3.18).

Example 3.5
Let's apply these results to a nal example, in which the input is a voltage source and the output
is the inductor current. The source voltage equals V = 2 cos (2 60 t ) + 3 . We want the circuit to
pass constant (o set) voltage essentially unaltered (save for the fact that the output is a current
rather than a voltage) and remove the 60 Hz term. Because the input is the sum of two sinusoids a
constant is a zero-frequency cosine our approach is

1. nd the transfer function using impedances;
2. use it to nd the output due to each input component;
3. add the results;
4. nd element values that accomplish our design criteria.

Because the circuit is a series combination of elements, let's use voltage divider to nd the transfer
function between V and V , then use the v-i relation of the inductor to nd its current.

I
=
j 2 fL + j 2 fL

1 2 fL

=
1 2 fL + R

= H ( f )

(3.19)

where

voltage divider =
j 2 fL + j 2 fL

and

inductor admittance =
1 2 fL

[Do the units check?] The form of this transfer function should be familiar; it is a lowpass lter,
and it will perform our desired function once we choose element values properly.
