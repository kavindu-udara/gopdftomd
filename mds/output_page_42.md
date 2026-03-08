36 CHAPTER 3. ANALOG SIGNAL PROCESSING

positive current ow is in the same direction as positive voltage drop. Once you de ne voltages and currents,
we need six non-redundant equations to solve for the six unknown voltages and currents. By specifying the
source, we have one; this amounts to providing the source's v-i relation. The v-i relations for the resistors
give us two more. We are only halfway there; where do we get the other three equations we need?
What we need to solve every circuit problem are mathematical statements that express how the circuit
elements are interconnected. Said another way, we need the laws that govern the electrical connection of
circuit elements. First of all, the places where circuit elements attach to each other are called nodes .
Two nodes are explicitly indicated in Figure 3.6; a third is at the bottom where the voltage source and
resistor R are connected. Electrical engineers tend to draw circuit diagrams schematics in a rectilinear
fashion. Thus the long line connecting the bottom of the voltage source with the bottom of the resistor
is intended to make the diagram look pretty. This line simply means that the two elements are connected
together. Kirchho 's Laws , one for voltage and one for current, determine what a connection among
circuit elements means. These laws can help us analyze this circuit.

3.4.1 Kirchho 's Current Law

At every node, the sum of all currents entering a node must equal zero. What this law means physically is
that charge cannot accumulate in a node; what goes in must come out. In our example circuit depicted in
Figure 3.6, we have a three-node circuit and thus have three KCL equations.

i i = 0

i i = 0

i + i = 0

Note that the current entering a node is the negative of the current leaving the node.
Given any two of these KCL equations, we can nd the other by adding or subtracting them. Thus, one
of them is redundant and, in mathematical terms, we can discard any one of them. The convention is to
discard the equation for the (unlabeled) node at the bottom of the circuit.
Exercise 3.2 (Solution on p. 94.)
In writing KCL equations, you will nd that in an n -node circuit, exactly one of them is always
redundant. Can you sketch a proof of why this might be true?
Hint: It has to do with the fact that charge won't accumulate in one place on its own.

3.4.2 Kirchho 's Voltage Law (KVL)

The voltage law says that the sum of voltages around every closed loop in the circuit must equal zero. A
closed loop has the obvious de nition: Starting at a node, trace a path through the circuit that returns you
to the origin node. KVL expresses the fact that electric elds are conservative: The total work performed
in moving a test charge around a closed path is zero. The KVL equation for our circuit is

v + v v = 0

In writing KVL equations, we follow the convention that an element's voltage enters with a plus sign when
traversing the closed path, we go from the positive to the negative of the voltage's de nition.
For the example circuit shown in Figure 3.6 (page 35), we have three v-i relations, two KCL equations,
and one KVL equation for solving for the circuit's six voltages and currents.

v-i: v = v KCL: i i = 0 KVL: v + v + v = 0

v = R i i i = 0

v = R i

We have exactly the right number of equations! Eventually, we will discover shortcuts for solving circuit
problems; for now, we want to eliminate all the variables but v and determine how it depends on v and
