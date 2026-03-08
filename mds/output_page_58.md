52 CHAPTER 3. ANALOG SIGNAL PROCESSING

v ( t ) i ( t ) , what is the equivalent expression using impedances? The resulting calculation reveals more about
power consumption in circuits and the introduction of the concept of average power .
When all sources produce sinusoids of frequency f , the voltage and current for any circuit element or
collection of elements are sinusoids of the same frequency.

v ( t ) = j V j cos (2 ft + )

i ( t ) = j I j cos (2 ft + )

Here, the complex amplitude of the voltage V equals j V j e and that of the current is j I j e . We can also
write the voltage and current in terms of their complex amplitudes using Euler's formula (Section 2.1.2).

v ( t ) =
1
V e + V e

i ( t ) =
1
Ie + I e

Multiplying these two expressions and simplifying gives

p ( t ) =
1
V I + V I + V Ie + V I e

=
1
Re [ V I ] +
1
Re V Ie

=
1
Re [ V I ] +
1
j V jj I j cos (4 ft + + )

We de ne
V I to be complex power . The real-part of complex power is the rst term and since it does
not change with time, it represents the power consistently consumed/produced by the circuit. The second
term varies with time at a frequency twice that of the source. Conceptually, this term details how power
sloshes back and forth in the circuit because of the sinusoidal source.
From another viewpoint, the real-part of complex power represents long-term energy consump-
tion/production. Energy is the integral of power and, as the integration interval increases, the rst term
appreciates while the time-varying term sloshes. Consequently, the most convenient de nition of the av-
erage power consumed/produced by any circuit is in terms of complex amplitudes.

P =
1
Re [ V I ] (3.14)

Exercise 3.12 (Solution on p. 94.)
Suppose the complex amplitudes of the voltage and current have xed magnitudes. What phase
relationship between voltage and current maximizes the average power? In other words, how are
and related for maximum power dissipation?

Because the complex amplitudes of the voltage and current are related by the equivalent impedance, average
power can also be written as

P =
1
Re [ Z ] j I j =
1
Re
1
j V j

These expressions generalize the results (3.3) we obtained for resistor circuits. We have derived a fundamental
result: Only the real part of impedance contributes to long-term power dissipation . Of the circuit
elements, only the resistor dissipates power. Capacitors and inductors dissipate no power in the long term.
It is important to realize that these statements apply only for sinusoidal sources. If you turn on a constant
voltage source in an RC-circuit, charging the capacitor does consume power.

Exercise 3.13 (Solution on p. 95.)
In an earlier problem (Problem 1.1), we found that the rms value of a sinusoid was its amplitude
divided by
p . What is average power expressed in terms of the rms values of the voltage and
current ( V and I respectively)?
