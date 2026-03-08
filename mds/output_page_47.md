41

Figure 3.12: The simple attenuator circuit (Figure 3.7) is attached to an oscilloscope's input. The

input-output relation for the above circuit without a load is: v =
R + R
v .

Suppose we want to pass the output signal into a voltage measurement device, such as an oscilloscope or
a voltmeter. In system-theory terms, we want to pass our circuit's output to a sink. For most applications,
we can represent these measurement devices as a resistor, with the current passing through it driving the
measurement device through some type of display. In circuits, a sink is called a load ; thus, we describe a
system-theoretic sink as a load resistance R . Thus, we have a complete system built from a cascade of
three systems: a source, a signal processing system (simple as it is), and a sink.
We must analyze afresh how this revised circuit, shown in Figure 3.12, works. Rather than de ning eight
variables and solving for the current in the load resistor, let's take a hint from other analysis (series rules
(p. 38), parallel rules (p. 40)). Resistors R and R are in a parallel con guration: The voltages across
each resistor are the same while the currents are not. Because the voltages are the same, we can nd the
current through each from their v-i relations: i =
and i =
. Considering the node where all three
resistors join, KCL says that the sum of the three currents must equal zero. Said another way, the current
entering the node through R must equal the sum of the other two currents leaving the node. Therefore,

i = i + i , which means that i = v
+
.

Let R denote the equivalent resistance of the parallel combination of R and R . Using R 's v-i
relation, the voltage across it is v =
. The KVL equation written around the leftmost loop has

v = v + v ; substituting for v , we nd

v = v
R
+ 1

or
v
=
R + R

Thus, we have the input-output relationship for our entire system having the form of voltage divider,
but it does not equal the input-output relation of the circuit without the voltage measurement device. We
can not measure voltages reliably unless the measurement device has little e ect on what we are trying to
measure. We should look more carefully to determine if any values for the load resistance would lessen its
impact on the circuit. Comparing the input-output relations before and after, what we need is R R .

As R =
1
+
1
, the approximation would apply if
1

1
or R R . This is the condition

we seek:

Voltage measurement: Voltage measurement devices must have large resistances compared with
that of the resistor across which the voltage is to be measured.

Exercise 3.7 (Solution on p. 94.)
Let's be more precise: How much larger would a load resistance need to be to a ect the input-output
relation by less than 10%? by less than 1%?
