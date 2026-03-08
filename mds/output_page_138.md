132 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.26
cuto frequency to the square wave's spectrum in this case?

Problem 4.11: Mathematics with Circuits
Simple circuits can implement simple mathematical operations, such as integration and di erentiation. We
want to develop an active circuit (it contains an op-amp) having an output that is proportional to the
integral of its input. For example, you could use an integrator in a car to determine distance traveled from
the speedometer.

(a) What is the transfer function of an integrator?
(b) Find an op-amp circuit so that its voltage output is proportional to the integral of its input for all
signals.

Problem 4.12: Where is that sound coming from?
We determine where sound is coming from because we have two ears and a brain. Sound travels at a
relatively slow speed and our brain uses the fact that sound will arrive at one ear before the other. As shown
in Figure 4.26, a sound coming from the right arrives at the left ear seconds after it arrives at the right
ear.
Once the brain nds this propagation delay, it can determine the sound direction. In an attempt to
model what the brain might do, RU signal processors want to design an optimal system that delays each
ear's signal by some amount then adds them together. and are the delays applied to the left and right
signals respectively. The idea is to determine the delay values according to some criterion that is based on
what is measured by the two ears.

(a) What is the transfer function between the sound signal s ( t ) and the processor output y ( t ) ?
(b) One way of determining the delay is to choose and to maximize the power in y ( t ) . How are
these maximum-power processing delays related to ?

Problem 4.13: Arrangements of Systems
Architecting a system of modular components means arranging them in various con gurations to achieve
some overall input-output relation. For each of the con gurations shown in Figure 4.27, determine the overall
transfer function between x ( t ) and y ( t ) .
The overall transfer function for the cascade ( rst depicted system) is particularly interesting. What
does it say about the e ect of the ordering of linear, time-invariant systems in a cascade?

Problem 4.14: Filtering

Let the signal s ( t ) =
sin ( t )
be the input to a linear, time-invariant lter having the transfer function

shown in Figure 4.28. Find the expression for y ( t ) , the lter's output.
