67

Figure 3.41: The top circuit depicts an op-amp in a feedback ampli er con guration. On the bottom
is the equivalent circuit, and integrates the op-amp circuit model into the circuit.

3.19.1 Inverting Ampli er

The feedback con guration shown in Figure 3.41 is the most common op-amp circuit for obtaining what is
known as an inverting ampli er .

R R GR

1
+
1
+
1

1
+
1
+
1

1
v =
1
v (3.27)

provides the exact input-output relationship. In choosing element values with respect to op-amp character-
istics, we can simplify the expression dramatically.

Make the load resistance, R , much larger than R . This situation drops the term 1 =R from the
second factor of (3.27).
Make the resistor, R , smaller than R , which means that the 1 =R term in the third factor is negligible.

With these two design criteria, the expression (3.27) becomes

R GR

1
+
1

1
v =
1
v (3.28)

Because the gain is large and the resistance R is small, the rst term becomes 1 =G , leaving us with

1

1
+
1

1
v =
1
v (3.29)

If we select the values of R and R so that ( GR R ) , this factor will no longer depend on the

op-amp's inherent gain, and it will equal
1
.
