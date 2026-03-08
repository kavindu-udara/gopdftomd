80 CHAPTER 3. ANALOG SIGNAL PROCESSING

by the actual resistors comprising the circuit? Let's start with simple cases and build up to a complete proof.

(a) Suppose resistors R and R are connected in parallel. Show that the power dissipated by ( R k R )
equals the sum of the powers dissipated by the component resistors.
(b) Now suppose R and R are connected in series. Show the same result for this combination.
(c) Use these two results to prove the general result we seek.

Problem 3.20: Power Transmission
The network shown in Figure 3.64a represents a simple power transmission system. The generator produces

(a) Simple power transmission system

(b) Modi ed load circuit

Figure 3.64

60 Hz and is modeled by a simple Thévenin equivalent. The transmission line consists of a long length of
copper wire and can be accurately described as a 50 resistor.

(a) Determine the load current R and the average power the generator must produce so that the load
receives 1,000 watts of average power. Why does the generator need to generate more than 1,000 watts
of average power to meet this requirement?
(b) Suppose the load is changed to that shown in Figure 3.64b. Now how much power must the generator
produce to meet the same power requirement? Why is it more than it had to produce to meet the
requirement for the resistive load?
(c) The load can be compensated to have a unity power factor (see Exercise 3.13) so that the voltage
and current are in phase for maximum power e ciency. The compensation technique is to place a
circuit in parallel to the load circuit. What element works and what is its value?
(d) With this compensated circuit, how much power must the generator produce to deliver 1,000 watts
average power to the load?

Problem 3.21: Optimal Power Transmission
Figure 3.65 shows a general model for power transmission. The power generator is represented by a Thévenin
equivalent and the load by a simple impedance. In most applications, the source components are xed while
there is some latitude in choosing the load.

(a) Suppose we wanted the maximize voltage transmission: make the voltage across the load as large as
possible. What choice of load impedance creates the largest load voltage? What is the largest load
voltage?
(b) If we wanted the maximum current to pass through the load, what would we choose the load impedance
to be? What is this largest current?
(c) What choice for the load impedance maximizes the average power dissipated in the load? What is
most power the generator can deliver?
