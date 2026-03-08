60 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.32

The presence of a current source in the circuit does not a ect the node method greatly; just include it in
writing KCL equations as a current leaving the node. The circuit has three nodes, requiring us to de ne
two node voltages. The node equations are

e
+
e e
i = 0 (Node 1)

e e
+
e
= 0 (Node 2)

Note that the node voltage corresponding to the node that we are writing KCL for enters with a positive
sign, the others with a negative sign, and that the units of each term is given in amperes. Rewrite these
equations in the standard set-of-linear-equations form.

e
1
+
1
e
1
= i

( e )
1
+ e
1
+
1
= 0

Solving these equations gives

e =
R + R
e

e =
R R + R + R
i

To nd the indicated current, we simply use i =
e
.

Example 3.6: Node Method Example
In the circuit shown in Figure 3.33, we cannot use the series/parallel combination rules: The vertical
resistor at node 1 keeps the two horizontal 1 resistors from being in series, and the 2 resistor
prevents the two 1 resistors at node 2 from being in series. We really do need the node method
to solve this circuit! Despite having six elements, we need only de ne two node voltages. The node
equations are

e v
+
e
+
e e
= 0 (Node 1)

e v
+
e
+
e e
= 0 (Node 2)

Solving these equations yields e =
v and e =
v . The output current equals
=
v .
One unfortunate consequence of using the element's numeric values from the outset is that it
becomes impossible to check units while setting up and solving equations.
