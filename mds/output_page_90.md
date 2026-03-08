84 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.70 If you were to build a circuit that was identical (from the viewpoint of the terminals) to the given one,
what would your circuit be?
(c) For your circuit, what are A and ?

Problem 3.31: Solving a Mystery Circuit
Sammy must determine as much as he can about a mystery circuit by attaching elements to the terminal
and measuring the resulting voltage. When he attaches a 1 resistor to the circuit's terminals, he measures
the voltage across the terminals to be 3 sin ( t ) . When he attaches a 1F capacitor across the terminals, the
voltage is now 3
p t = 4) .

(a) What voltage should he measure when he attaches nothing to the mystery circuit?
(b) What voltage should Sammy measure if he doubled the size of the capacitor to 2 F and attached it to
the circuit?

Problem 3.32: Find the Load Impedance
The circuit depicted in Figure 3.71 has a transfer function between the output voltage and the source equal
to

H ( f ) =
8 f
8 f + 4 + j 6 f

Figure 3.71

(a) Sketch the magnitude and phase of the transfer function.
(b) At what frequency does the phase equal = 2 ?
(c) Find a circuit that corresponds to this load impedance. Is your answer unique? If so, show it to be so;
if not, give another example.

Problem 3.33: Analog Hum Rejection
Hum refers to corruption from wall socket power that frequently sneaks into circuits. Hum gets its name
because it sounds like a persistent humming sound. We want to nd a circuit that will remove hum from any
signal. A Rice engineer suggests using a simple voltage divider circuit consisting of two series impedances
(Figure 3.72).

(a) The impedance Z is a resistor. The Rice engineer must decide between two circuits for the impedance
Z (Figure 3.73). Which of these will work?
