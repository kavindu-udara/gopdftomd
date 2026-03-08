69

Figure 3.43
1
+ j 2 fC . Letting the input resistance equal R , the transfer

function of the op-amp inverting ampli er now is H ( f ) =
R =R j 2 fR C

Thus, we have the gain equal to
R
and the cuto frequency
1 C
.

Creating a speci c transfer function with op-amps does not have a unique answer. As opposed to design
with passive circuits, electronics is more exible (a cascade of circuits can be built so that each has little
e ect on the others; see Problem 3.43) and gain (increase in power and amplitude) can result. To complete
our example, let's assume we want a lowpass lter that emulates what the telephone companies do. Signals
transmitted over the telephone have an upper frequency limit of about 3 kHz. For the second design choice,
we require R C = 5 : 3 10 . Thus, many choices for resistance and capacitance values are possible. A
1 F capacitor and a 330 resistor, 10 nF and 33 k , and 10 pF and 33 M would all theoretically work.

Let's also desire a voltage gain of ten:
R
= 10 , which means R =
R
. Recall that we must have R < R .

As the op-amp's input impedance is about 1 M , we don't want R too large, and this requirement means
that the last choice for resistor/capacitor values won't work. We also need to ask for less gain than the
op-amp can provide itself. Because the feedback element is an impedance (a parallel resistor capacitor

combination), we need to examine the gain requirement more carefully. We must have
j Z j
< 10 for all

frequencies of interest. Thus,
R 1 + j 2 fR C j
=R < 10 . As this impedance decreases with frequency, the

design speci cation of
R
= 10 means that this criterion is easily met. Thus, the rst two choices for the

resistor and capacitor values (as well as many others in this range) will work well. Additional considerations
like parts cost might enter into the picture. Unless you have a high-power application (this isn't one) or ask
for high-precision components, costs don't depend heavily on component values as long as you stay close to
standard values. For resistors, having values r 10 , easily obtained values of r are 1, 1.4, 3.3, 4.7, and 6.8,
and the decades span 0 8.

Exercise 3.17 (Solution on p. 95.)
What is special about the resistor values; why these rather odd-appearing values for r ?

3.19.3 Intuitive Way of Solving Op-Amp Circuits

When we meet op-amp design speci cations, we can simplify our circuit calculations greatly, so much so that
we don't need the op-amp's circuit model to determine the transfer function. Here is our inverting ampli er.
When we take advantage of the op-amp's characteristics large input impedance, large gain, and small
output impedance we note the two following important facts.
