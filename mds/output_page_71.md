65

Because dependent sources cannot be described as impedances, and because the dependent variable
cannot disappear when you apply parallel/series combining rules, circuit simpli cations such as current
and voltage divider should not be applied in most cases. Analysis of circuits containing dependent sources
essentially requires use of formal methods, like the node method (Section 3.15). Using the node method for
such circuits is not di cult, with node voltages de ned across the source treated as if they were known (as
with independent sources). Consider the circuit shown on the top in Figure 3.39.

Figure 3.39: The top circuit depicts an op-amp in a feedback ampli er con guration. On the bottom
is the equivalent circuit, and integrates the op-amp circuit model into the circuit.

Note that the op-amp is placed in the circuit upside-down, with its inverting input at the top and
serving as the only input. As we explore op-amps in more detail in the next section, this con guration will
appear again and again, and its usefulness demonstrated. To determine how the output voltage is related to
the input voltage, we apply the node method. Only two node voltages v and v need be de ned; the
remaining nodes are across sources or serve as the reference. The node equations are

v v
+
v
+
v v
= 0 (3.24)

v ( Gv )
+
v v
+
v
= 0 (3.25)

Note that no special considerations were used in applying the node method to this dependent-source circuit.
Solving these to learn how v relates to v yields

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
v (3.26)

This expression represents the general input-output relation for this circuit, known as the standard feed-
back con guration . Once we learn more about op-amps, in particular what its typical element values
are, the expression will simplify greatly. Do note that the units check, and that the parameter G of the
dependent source is a dimensionless gain.
