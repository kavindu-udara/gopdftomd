81

Figure 3.65 One way to maximize a function of a complex variable is to write the expression in terms
of the variable's real and imaginary parts, evaluate derivatives with respect to each, set both
derivatives to zero and solve the two equations simultaneously.

Problem 3.22: Big is Beautiful
Sammy wants to choose speakers that produce very loud music. He has an ampli er and notices that the
speaker terminals are labeled 8 source.

(a) What does this mean in terms of the ampli er's equivalent circuit?
(b) Any speaker Sammy attaches to the terminals can be well-modeled as a resistor. Choosing a speaker
amounts to choosing the values for the resistor. What choice would maximize the voltage across the
speakers?
(c) Sammy decides that maximizing the power delivered to the speaker might be a better choice. What
values for the speaker resistor should be chosen to maximize the power delivered to the speaker?

Problem 3.23: Sharing a Channel
Two transmitter-receiver pairs want to share the same digital communications channel. The transmitter
signals will be added together by the channel. Receiver design is greatly simpli ed if rst we remove the
unwanted transmission (as much as possible). Each transmitter signal has the form

x ( t ) = A sin (2 f t ) ; 0 t T

where the amplitude is either zero or A and each transmitter uses its own frequency f . Each frequency is
harmonically related to the bit interval duration T , where the transmitter one uses the the frequency 1 =T .
The datarate is 10 Mbps.

(a) Draw a block diagram that expresses this communication scenario.
(b) Find circuits that the receivers could employ to separate unwanted transmissions. Assume the received
signal is a voltage and the output is to be a voltage as well.
(c) Find the second transmitter's frequency so that the receivers can suppress the unwanted transmission
by at least a factor of ten.

Problem 3.24: Circuit Detective Work
In the lab, the open-circuit voltage measured across an unknown circuit's terminals equals sin ( t ) . When a

1 resistor is placed across the terminals, a voltage of
1
sin ( t + = 4) appears.

(a) What is the Thévenin equivalent circuit?
(b) What voltage will appear if we place a 1F capacitor across the terminals?

Problem 3.25: Mystery Circuit
We want to determine as much as we can about the circuit lurking in the impenetrable box shown in
Figure 3.66. A voltage source v = 2 volts has been attached to the left-hand terminals, leaving the right
terminals for tests and measurements.

(a) Sammy measures v = 10 V when a 1 resistor is attached to the terminals. Samantha says he is
wrong. Who is correct and why?
