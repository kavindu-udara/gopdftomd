58 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.30: Input and output waveforms for the example RL circuit when the element values are
R = 6 : 28 and L = 100mH . 3 j H (0) j =
3
. Thus, the value we

choose for the resistance will determine the scaling factor of how voltage is converted into current.
For the 60 Hz component signal, the output current is 2 j H (60) j cos (2 60 t + \ H (60)) . The total
output due to our source is

i = 2 j H (60) j cos (2 60 t + \ H (60)) + 3 H (0) (3.20)

The cuto frequency for this lter occurs when the real and imaginary parts of the transfer function's

denominator equal each other. Thus, 2 f L = R , which gives f =
R L
. We want this cuto

frequency to be much less than 60 Hz. Suppose we place it at, say, 10 Hz. This speci cation would

require the component values to be related by
R
= 20 = 62 : 8 . The transfer function at 60 Hz

would be
1 2 60 L + R
=
1

1 j + 1
=
1

1
0 : 16
1
(3.21)

which yields an attenuation (relative to the gain at zero frequency) of about 1 = 6 , and result in
an output amplitude of 0 : 3 =R relative to the constant term's amplitude of 3 =R . A factor of ten
di erence between the relative sizes of the two components seems reasonable. Having a 100 mH
inductor would require a 6.28 resistor. An easily available resistor value is 6.8 ; thus, this
choice results in cheaply and easily purchased parts. To make the resistance bigger would require a
proportionally larger inductor. Unfortunately, even a 1 H inductor is physically large; consequently
low cuto frequencies require small-valued resistors and large-valued inductors. The choice made
here represents only one compromise.
The phase of the 60 Hz component will very nearly be = 2 , leaving it to be
0 : 3 =R cos 2 60 t
= 0 : 3 =R sin (2 60 t ) . The waveforms for the input and output are shown
in Figure 3.30.

Note that the sinusoid's phase has indeed shifted; the lowpass lter not only reduced the 60 Hz signal's
amplitude, but also shifted its phase by 90 .
