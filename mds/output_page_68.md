62 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.35: Transfer functions of the circuits shown in Figure 3.34. Here, R = 1 , R = 1 , and C = 1 .

When R = R , as shown on the plot, the passband gain becomes half of the original, and the cuto
frequency increases by the same factor. Thus, adding R provides a 'knob' by which we can trade passband
gain for cuto frequency.

Exercise 3.16 (Solution on p. 95.)
We can change the cuto frequency without a ecting passband gain by changing the resistance in
the original circuit. Does the addition of the R resistor help in circuit design?

### 3.16 Power Conservation in Circuits

Now that we have a formal method the node method for solving circuits, we can use it to prove a powerful
result: KVL and KCL are all that are required to show that all circuits conserve power, regardless of what
elements are used to build the circuit.
First of all, de ne node voltages for all nodes in a given circuit. Any node chosen as the reference will
do. For example, in the portion of a large circuit depicted in Figure 3.36, we de ne node voltages for nodes
a, b and c. With these node voltages, we can express the voltage across any element in terms of them. For
example, the voltage across element 1 is given by v = e e . The instantaneous power for element 1
becomes
v i = ( e e ) i = e i e i

Writing the power for the other elements, we have

v i = e i e i

v i = e i e i

When we add together the element power terms, we discover that once we collect terms involving a particular
node voltage, it is multiplied by the sum of currents leaving the node minus the sum of currents entering.
For example, for node b, we have e ( i i ) . We see that the currents will obey KCL that multiply each
This content is available online at http://cnx.org/content/m17317/1.2/ .
