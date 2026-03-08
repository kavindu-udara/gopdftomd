88 CHAPTER 3. ANALOG SIGNAL PROCESSING

(b) You attach a 1 H inductor across the terminals. What voltage do you measure?

Problem 3.41: Dependent Sources
Find the voltage v in each of the circuits depicted in Figure 3.80.

(a) circuit a

(b) circuit b

Figure 3.80

Problem 3.42: Operational Ampli ers
Find the transfer function between the source voltage(s) and the indicated output voltage for the circuits
shown in Figure 3.81.

Problem 3.43: Op-Amp Circuit
The circuit shown in Figure 3.82 is claimed to serve a useful purpose.

(a) What is the transfer function relating the complex amplitude of the output signal, the current I , to
the complex amplitude of the input, the voltage V ?
(b) What equivalent circuit does the load resistor R see?
(c) Find the output current when v = V e .

Problem 3.44: Why Op-Amps are Useful
The cascade of op-amp circuits shown in Figure 3.83 illustrates the reason why op-amp realizations of transfer
functions are so useful.

(a) Find the transfer function relating the complex amplitude of the voltage v ( t ) to the source. Show
that this transfer function equals the product of each stage's transfer function.
(b) What is the load impedance appearing across the rst op-amp's output?
(c) Figure 3.84 illustrates that sometimes designs can go wrong. Find the transfer function for this
op-amp circuit, and then show that it can't work! Why can't it?

Problem 3.45: Operational Ampli ers
Consider the circuit of Figure 3.85.

(a) Find the transfer function relating the voltage v ( t ) to the source.
(b) In particular, R = 530 , C = 1 F , R = 5 : 3 k , C = 0 : 01 F , and R = R = 5 : 3 k .
Characterize the resulting transfer function and determine what use this circuit might have.

Problem 3.46: Designing a Bandpass Filter
We want to design a bandpass lter that has transfer the function

H ( f ) = 10
j 2 f
j
+ 1 j
+ 1

Here, f is the cuto frequency of the low-frequency edge of the passband and f is the cuto frequency of
the high-frequency edge. We want f = 1 kHz and f = 10 kHz .
