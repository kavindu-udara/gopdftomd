59

Figure 3.31

In some (complicated) cases, we cannot use the simpli cation techniques such as parallel or series combi-
nation rules to solve for a circuit's input-output relation. In other modules, we wrote v-i relations and
Kirchho 's laws haphazardly, solving them more on intuition than procedure. We need a formal method
that produces a small, easy set of equations that lead directly to the input-output relation we seek. One
such technique is the node method .
The node method begins by nding all nodes places where circuit elements attach to each other in the
circuit. We call one of the nodes the reference node ; the choice of reference node is arbitrary, but it is
usually chosen to be a point of symmetry or the bottom node. For the remaining nodes, we de ne node
voltages e that represent the voltage between the node and the reference. These node voltages constitute
the only unknowns; all we need is a su cient number of equations to solve for them. In our example, we
have two node voltages. The very act of de ning node voltages is equivalent to using all the KVL
equations at your disposal . The reason for this simple, but astounding, fact is that a node voltage is
uniquely de ned regardless of what path is traced between the node and the reference. Because two paths
between a node and reference have the same voltage, the sum of voltages around the loop equals zero.
In some cases, a node voltage corresponds exactly to the voltage across a voltage source. In such cases,
the node voltage is speci ed by the source and is not an unknown. For example, in our circuit, e = v ;
thus, we need only to nd one node voltage.
The equations governing the node voltages are obtained by writing KCL equations at each node having
an unknown node voltage, using the v-i relations for each element. In our example, the only circuit equation
is
e v
+
e
+
e
= 0 (3.22)

A little re ection reveals that when writing the KCL equations for the sum of currents leaving a node, that
node's voltage will always appear with a plus sign, and all other node voltages with a minus sign. Systematic
application of this procedure makes it easy to write node equations and to check them before solving them.
Also remember to check units at this point: Every term should have units of current. In our example, solving
for the unknown node voltage is easy:

e =
R R R + R R + R R
v (3.23)

Have we really solved the circuit with the node method? Along the way, we have used KVL, KCL, and
the v-i relations. Previously, we indicated that the set of equations resulting from applying these laws is
necessary and su cient. This result guarantees that the node method can be used to solve any circuit.
One fallout of this result is that we must be able to nd any circuit variable given the node voltages and
sources. All circuit variables can be found using the v-i relations and voltage divider. For example, the

current through R equals
e
.
This content is available online at http://cnx.org/content/m0032/2.20/ .
