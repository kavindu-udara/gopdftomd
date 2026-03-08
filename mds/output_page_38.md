32 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.1: The generic circuit element.

Current ows through circuit elements, such as that depicted in Figure 3.1, and through conductors,
which we indicate by lines in circuit diagrams. For every circuit element we de ne a voltage and a current.
The element has a v-i relation de ned by the element's physical properties. In de ning the v-i relation, we
have the convention that positive current ows from positive to negative voltage drop. Voltage has units of
volts, and both the unit and the quantity are named for Alessandro Volta . Current has units of amperes,
and is named for the French physicist André-Marie Ampère .
Voltages and currents also carry power . Again using the convention shown in Figure 3.1 for circuit
elements, the instantaneous power at each moment of time consumed by the element is given by the
product of the voltage and current.
p ( t ) = v ( t ) i ( t )

A positive value for power indicates that at time t the circuit element is consuming power; a negative value
means it is producing power. With voltage expressed in volts and current in amperes, power de ned this
way has units of watts . Just as in all areas of physics and chemistry, power is the rate at which energy is
consumed or produced. Consequently, energy is the integral of power.

E ( t ) =

Z

p ( ) d

Again, positive energy corresponds to consumed energy and negative energy corresponds to energy produc-
tion. Note that a circuit element having a power pro le that is both positive and negative over some time
interval could consume or produce energy according to the sign of the integral of power. The units of energy
are joules since a watt equals joules/second.

Exercise 3.1 (Solution on p. 94.)
Residential energy bills typically state a home's energy usage in kilowatt-hours. Is this really a unit
of energy? If so, how many joules equals one kilowatt-hour?

### 3.2 Ideal Circuit Elements

The elementary circuit elements the resistor, capacitor, and inductor impose linear relationships between
voltage and current.

3.2.1 Resistor

Figure 3.2: Resistor. v = Ri
https://nationalmaglab.org/education/magnet-academy/history-of-electricity-magnetism/pioneers/
alessandro-volta
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Ampere.html
This content is available online at http://cnx.org/content/m0012/2.21/ .
