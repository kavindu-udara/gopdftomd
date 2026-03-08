34 CHAPTER 3. ANALOG SIGNAL PROCESSING

3.2.3 Inductor

Figure 3.4: Inductor. v = L
d
i ( t )

The inductor stores magnetic ux, with larger valued inductors capable of storing more ux. Inductance has
units of henries (H), and is named for the American physicist Joseph Henry . The di erential and integral
forms of the inductor's v-i relation are

v ( t ) = L
d
i ( t ) or i ( t ) =
1

Z

v ( ) d (3.2)

The power consumed/produced by an inductor depends on the product of the inductor current and its
derivative

p ( t ) = Li ( t )
d
i ( t )

and its total energy expenditure up to time t is given by

E ( t ) =
1
Li ( t )

3.2.4 Sources

(a)

(b)

Figure 3.5: The voltage source on the left and current source on the right are like all circuit elements
in that they have a particular relationship between the voltage and current de ned for them. For the
voltage source, v = v for any current i ; for the current source, i = i for any voltage v .

Sources of voltage and current are also circuit elements, but they are not linear in the strict sense of linear
systems. For example, the voltage source's v-i relation is v = v regardless of what the current might be.
As for the current source, i = i regardless of the voltage. Another name for a constant-valued voltage
source is a battery, and can be purchased in any supermarket. Current sources, on the other hand, are much
harder to acquire; we'll learn why later.

### 3.3 Ideal and Real-World Circuit Elements

Source and linear circuit elements are ideal circuit elements. One central notion of circuit theory is combining
the ideal elements to describe how physical elements operate in the real world. For example, the 1 k
http://siarchives.si.edu/history/exhibits/joseph-henry
This content is available online at http://cnx.org/content/m0013/2.9/ .
